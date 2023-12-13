package server

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/gopcua/opcua/ua"
	"github.com/gopcua/opcua/uasc"
)

type PubReq struct {
	// The data of the publish request
	Req *ua.PublishRequest

	// The request ID (from the header) of the publish request.  This has to be used when replying.
	ID uint32
}

type MonitoredItem struct {
}

type Subscription struct {
	srv                       Server
	ID                        uint32
	RevisedPublishingInterval float64
	RevisedLifetimeCount      uint32
	RevisedMaxKeepAliveCount  uint32
	ReadValueList             map[uint32]ua.ReadValueID
	ReadValueLock             sync.Mutex
	Channel                   *uasc.SecureChannel
	SequenceID                uint32
	SeqNums                   map[uint32]struct{}
}

func NewSubscription() *Subscription {
	return &Subscription{
		ReadValueList: map[uint32]ua.ReadValueID{},
		SeqNums:       map[uint32]struct{}{},
	}
}

func (s *Subscription) Start() {
	go s.run()

}

// this function should be run as a go-routine and will handle sending data out
// to the client at the correct rate assuming there are publish requests queued up.
// This is the simplest possible subscription logic where we don't check if a value was changed
// or rely on some other event notification method.  We just send the values at the subscription rate
func (s *Subscription) run() {
	for {
		// wait for the publish interval
		time.Sleep(time.Millisecond * time.Duration(s.RevisedPublishingInterval))
		s.ReadValueLock.Lock()
		l := len(s.ReadValueList)
		s.ReadValueLock.Unlock()
		if l == 0 {
			log.Print("No tags in sub list")
			continue
		}
		log.Printf("Waiting for publish req on sub #%d", s.ID)
		pubreq := <-s.srv.PublishRequests
		// see if we have any tags
		s.SequenceID++
		log.Printf("Got publish req on sub #%d.  Sequence %d", s.ID, s.SequenceID)
		// then get all the tags and send them back to the client

		vals := make([]*ua.MonitoredItemNotification, len(s.ReadValueList))
		s.ReadValueLock.Lock()
		log.Printf("Responding with %d elements: %v", len(s.ReadValueList), s.ReadValueList)
		for i := range s.ReadValueList {
			rv := s.ReadValueList[i]
			ns, err := s.srv.Namespace(int(rv.NodeID.Namespace()))
			vals[i] = new(ua.MonitoredItemNotification)
			if err != nil {
				log.Printf("error getting namespace %d: %v", rv.NodeID.Namespace(), err)
				vals[i].Value = &ua.DataValue{}
				vals[i].Value.Status = ua.StatusBad
				vals[i].Value.EncodingMask |= ua.DataValueStatusCode
			}
			dv := ns.Attribute(rv.NodeID, rv.AttributeID)
			vals[i].ClientHandle = i
			vals[i].Value = dv
		}
		for x := range pubreq.Req.SubscriptionAcknowledgements {
			a := pubreq.Req.SubscriptionAcknowledgements[x]
			delete(s.SeqNums, a.SequenceNumber)
		}
		log.Printf("values: %v", *vals[0])
		s.ReadValueLock.Unlock()
		dcn := ua.DataChangeNotification{
			MonitoredItems:  vals,
			DiagnosticInfos: []*ua.DiagnosticInfo{},
		}
		eo := make([]*ua.ExtensionObject, 1)
		eo[0] = ua.NewExtensionObject(&dcn)
		eo[0].UpdateMask()

		msg := ua.NotificationMessage{
			SequenceNumber:   s.SequenceID,
			PublishTime:      time.Now(),
			NotificationData: eo,
		}
		s.SeqNums[s.SequenceID] = struct{}{}

		seqnums := make([]uint32, len(s.SeqNums))
		idx := 0
		for i := range s.SeqNums {
			seqnums[idx] = i
			idx++
		}
		log.Printf("sequence numbers: %d", seqnums)

		response := &ua.PublishResponse{
			ResponseHeader: &ua.ResponseHeader{
				Timestamp:          time.Now(),
				RequestHandle:      pubreq.Req.RequestHeader.RequestHandle,
				ServiceResult:      ua.StatusOK,
				ServiceDiagnostics: &ua.DiagnosticInfo{},
				StringTable:        []string{},
				AdditionalHeader:   ua.NewExtensionObject(nil),
			},
			SubscriptionID:           s.ID,
			MoreNotifications:        false,
			NotificationMessage:      &msg,
			AvailableSequenceNumbers: seqnums,
			Results:                  []ua.StatusCode{},
			DiagnosticInfos:          []*ua.DiagnosticInfo{},
		}
		err := s.Channel.SendResponseWithContext(context.Background(), pubreq.ID, response)
		if err != nil {
			log.Printf("problem sending channel response: %v", err)
			log.Printf("Killing subscription %d", s.ID)
			return
		}
		log.Printf("Publish OK for %d", s.ID)
		// wait till we've got a publish request.
	}
}
