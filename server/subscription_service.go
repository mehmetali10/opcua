package server

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/gopcua/opcua/debug"
	"github.com/gopcua/opcua/ua"
	"github.com/gopcua/opcua/uasc"
)

// SubscriptionService implements the Subscription Service Set.
//
// https://reference.opcfoundation.org/Core/Part4/v105/docs/5.13
type SubscriptionService struct {
	srv *Server
	// pub sub stuff
	PublishRequests chan PubReq
	Mu              sync.Mutex
	Subs            map[uint32]*Subscription
}

// get rid of all references to a subscription and all monitored items that are pointed at this subscription.
func (s *SubscriptionService) DeleteSubscription(id uint32) {
	s.Mu.Lock()
	defer s.Mu.Unlock()

	delete(s.Subs, id)

	// ask the monitored item service to purge out any items that use this subscription
	s.srv.MonitoredItemService.DeleteSub(id)

}

// https://reference.opcfoundation.org/Core/Part4/v105/docs/5.13.2
func (s *SubscriptionService) CreateSubscription(sc *uasc.SecureChannel, r ua.Request, reqID uint32) (ua.Response, error) {
	debug.Printf("Handling %T", r)

	req, err := safeReq[*ua.CreateSubscriptionRequest](r)
	if err != nil {
		return nil, err
	}

	s.Mu.Lock()
	defer s.Mu.Unlock()

	newsubid := uint32(len(s.Subs))

	log.Printf("Create Sub %d", newsubid)

	sub := Subscription{
		srv:                       s,
		Channel:                   sc,
		ID:                        newsubid,
		RevisedPublishingInterval: req.RequestedPublishingInterval,
		RevisedLifetimeCount:      req.RequestedLifetimeCount,
		RevisedMaxKeepAliveCount:  req.RequestedMaxKeepAliveCount,
		SeqNums:                   make(map[uint32]struct{}),
		NotifyChannel:             make(chan *ua.MonitoredItemNotification),
	}
	s.Subs[newsubid] = &sub
	sub.Start()

	resp := &ua.CreateSubscriptionResponse{
		ResponseHeader: &ua.ResponseHeader{
			Timestamp:          time.Now(),
			RequestHandle:      req.RequestHeader.RequestHandle,
			ServiceResult:      ua.StatusOK,
			ServiceDiagnostics: &ua.DiagnosticInfo{},
			StringTable:        []string{},
			AdditionalHeader:   ua.NewExtensionObject(nil),
		},
		SubscriptionID:            uint32(newsubid),
		RevisedPublishingInterval: req.RequestedPublishingInterval,
		RevisedLifetimeCount:      req.RequestedLifetimeCount,
		RevisedMaxKeepAliveCount:  req.RequestedMaxKeepAliveCount,
	}
	return resp, nil
}

// https://reference.opcfoundation.org/Core/Part4/v105/docs/5.13.3
func (s *SubscriptionService) ModifySubscription(sc *uasc.SecureChannel, r ua.Request, reqID uint32) (ua.Response, error) {
	debug.Printf("Handling %T", r)

	req, err := safeReq[*ua.ModifySubscriptionRequest](r)
	if err != nil {
		return nil, err
	}
	return serviceUnsupported(req.RequestHeader), nil
}

// https://reference.opcfoundation.org/Core/Part4/v105/docs/5.13.4
func (s *SubscriptionService) SetPublishingMode(sc *uasc.SecureChannel, r ua.Request, reqID uint32) (ua.Response, error) {
	debug.Printf("Handling %T", r)

	req, err := safeReq[*ua.SetPublishingModeRequest](r)
	if err != nil {
		return nil, err
	}
	return serviceUnsupported(req.RequestHeader), nil
}

// https://reference.opcfoundation.org/Core/Part4/v105/docs/5.13.5
func (s *SubscriptionService) Publish(sc *uasc.SecureChannel, r ua.Request, reqID uint32) (ua.Response, error) {
	log.Printf("Raw Publish req")

	req, err := safeReq[*ua.PublishRequest](r)
	if err != nil {
		log.Printf("ERROR: bad PublishRequest Struct")
		return nil, err
	}

	select {
	case s.PublishRequests <- PubReq{Req: req, ID: reqID}:
	default:
		log.Printf("Too many publish reqs.")
	}

	// per opcua spec, we don't respond now.  When data is available on the subscription,
	// the Subscription will respond in the background.
	return nil, nil
}

// https://reference.opcfoundation.org/Core/Part4/v105/docs/5.13.6
func (s *SubscriptionService) Republish(sc *uasc.SecureChannel, r ua.Request, reqID uint32) (ua.Response, error) {
	debug.Printf("Handling %T", r)

	req, err := safeReq[*ua.RepublishRequest](r)
	if err != nil {
		return nil, err
	}
	return serviceUnsupported(req.RequestHeader), nil
}

// https://reference.opcfoundation.org/Core/Part4/v105/docs/5.13.7
func (s *SubscriptionService) TransferSubscriptions(sc *uasc.SecureChannel, r ua.Request, reqID uint32) (ua.Response, error) {
	debug.Printf("Handling %T", r)

	req, err := safeReq[*ua.TransferSubscriptionsRequest](r)
	if err != nil {
		return nil, err
	}
	return serviceUnsupported(req.RequestHeader), nil
}

// https://reference.opcfoundation.org/Core/Part4/v105/docs/5.13.8
func (s *SubscriptionService) DeleteSubscriptions(sc *uasc.SecureChannel, r ua.Request, reqID uint32) (ua.Response, error) {
	debug.Printf("Handling %T", r)

	req, err := safeReq[*ua.DeleteSubscriptionsRequest](r)
	if err != nil {
		return nil, err
	}

	results := make([]ua.StatusCode, len(req.SubscriptionIDs))
	for i := range req.SubscriptionIDs {
		s.DeleteSubscription(req.SubscriptionIDs[i])
		results[i] = ua.StatusOK
	}
	return &ua.DeleteSubscriptionsResponse{
		ResponseHeader: &ua.ResponseHeader{
			Timestamp:          time.Now(),
			RequestHandle:      req.RequestHeader.RequestHandle,
			ServiceResult:      ua.StatusOK,
			ServiceDiagnostics: &ua.DiagnosticInfo{},
			StringTable:        []string{},
			AdditionalHeader:   ua.NewExtensionObject(nil),
		},
		Results:         results,                //                  []StatusCode
		DiagnosticInfos: []*ua.DiagnosticInfo{}, //          []*DiagnosticInfo
	}, nil
}

type PubReq struct {
	// The data of the publish request
	Req *ua.PublishRequest

	// The request ID (from the header) of the publish request.  This has to be used when replying.
	ID uint32
}

// This is the type that with its run() function will work in the bakground fullfilling subscription
// publishes.
//
// MonitoredItems will send updates on the NotifyChannel to let the background task know that
// an event has occured that needs to be published.
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
}

func NewSubscription() *Subscription {
	return &Subscription{
		SeqNums:       map[uint32]struct{}{},
		NotifyChannel: make(chan *ua.MonitoredItemNotification, 100),
	}
}

func (s *Subscription) Start() {
	go s.run()

}

// this function should be run as a go-routine and will handle sending data out
// to the client at the correct rate assuming there are publish requests queued up.
func (s *Subscription) run() {
	// if this go routine dies, we need to delete ourselves.
	defer s.srv.DeleteSubscription(s.ID)

	s.T = time.NewTicker(time.Millisecond * time.Duration(s.RevisedPublishingInterval))
	for {
		// we don't need to do anything if we don't have at least one thing to publish so lets get that first
		publishQueue := make(map[uint32]*ua.MonitoredItemNotification)
		// TODO: this should also monitor <-s.T.C and if it comes up we should send a keepalive?
		publishQueue[0] = <-s.NotifyChannel
		// collect monitored item notifications
		var pubreq PubReq
	L1:
		for {
			select {
			case newNotification := <-s.NotifyChannel:
				publishQueue[newNotification.ClientHandle] = newNotification
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
				publishQueue[newNotification.ClientHandle] = newNotification
			case <-s.T.C:
				// we've gone long enough that we can send
				break L2
			}
		}

		s.SequenceID++
		log.Printf("Got publish req on sub #%d.  Sequence %d", s.ID, s.SequenceID)
		// then get all the tags and send them back to the client

		for x := range pubreq.Req.SubscriptionAcknowledgements {
			a := pubreq.Req.SubscriptionAcknowledgements[x]
			delete(s.SeqNums, a.SequenceNumber)
		}

		final_items := make([]*ua.MonitoredItemNotification, len(publishQueue))
		i := 0
		for k := range publishQueue {
			final_items[i] = publishQueue[k]
			i++
		}

		dcn := ua.DataChangeNotification{
			MonitoredItems:  final_items,
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
		log.Printf("Published %d items OK for %d", len(publishQueue), s.ID)
		// wait till we've got a publish request.
	}
}

//PublishRequest_Encoding_DefaultBinary
