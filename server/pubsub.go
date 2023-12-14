package server

import (
	"context"
	"log"
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

type Subscription struct {
	srv                       *SubscriptionService
	ID                        uint32
	RevisedPublishingInterval float64
	RevisedLifetimeCount      uint32
	RevisedMaxKeepAliveCount  uint32
	Channel                   *uasc.SecureChannel
	SequenceID                uint32
	SeqNums                   map[uint32]struct{}
	T                         *time.Ticker

	NotifyChannel chan *ua.MonitoredItemNotification
	publishQueue  []*ua.MonitoredItemNotification
}

func NewSubscription() *Subscription {
	return &Subscription{
		SeqNums: map[uint32]struct{}{},
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
	s.T = time.NewTicker(time.Millisecond * time.Duration(s.RevisedPublishingInterval))
	for {
		// we don't need to do anything if we don't have at least one thing to publish so lets get that first
		s.publishQueue = make([]*ua.MonitoredItemNotification, 1, 32)
		s.publishQueue[0] = <-s.NotifyChannel
		// collect monitored item notifications
		var pubreq PubReq
	L1:
		for {
			// TODO: this select should also monitor <-s.T.C and if it comes up we should send a keepalive?
			select {
			case newNotification := <-s.NotifyChannel:
				s.publishQueue = append(s.publishQueue, newNotification)
			case pubreq = <-s.srv.PublishRequests:
				// once we get a publish request, we should move on to publish them back
				log.Printf("Got publish request.")
				break L1
			}
		}
		// now we need to continue to collect notifications until we've reached our publishing interval
	L2:
		for {
			select {
			case newNotification := <-s.NotifyChannel:
				s.publishQueue = append(s.publishQueue, newNotification)
			case <-s.T.C:
				// we've gone long enough that we can send
				break L2
			}
		}

		// see if we have any tags
		s.SequenceID++
		log.Printf("Got publish req on sub #%d.  Sequence %d", s.ID, s.SequenceID)
		// then get all the tags and send them back to the client

		for x := range pubreq.Req.SubscriptionAcknowledgements {
			a := pubreq.Req.SubscriptionAcknowledgements[x]
			delete(s.SeqNums, a.SequenceNumber)
		}
		dcn := ua.DataChangeNotification{
			MonitoredItems:  s.publishQueue,
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
		log.Printf("Published %d items OK for %d", len(s.publishQueue), s.ID)
		// wait till we've got a publish request.
	}
}

//PublishRequest_Encoding_DefaultBinary
