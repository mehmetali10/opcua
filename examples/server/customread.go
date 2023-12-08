package main

import (
	"context"
	"errors"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/gopcua/opcua/debug"
	"github.com/gopcua/opcua/id"
	"github.com/gopcua/opcua/ua"
	"github.com/gopcua/opcua/uasc"
)

type MapReadWriter struct {
	Mu              sync.RWMutex
	Data            map[string]any
	Subs            map[uint32]*MapReadWriterSub
	PublishRequests chan struct{}
}

type MapReadWriterSub struct {
	MapReadWriter             *MapReadWriter
	ID                        uint32
	RevisedPublishingInterval float64
	RevisedLifetimeCount      uint32
	RevisedMaxKeepAliveCount  uint32
	Tags                      []string
	TagListLock               sync.Mutex
	Channel                   *uasc.SecureChannel
	SequenceID                uint32
}

func NewMapReadWriterSub() *MapReadWriterSub {
	return &MapReadWriterSub{
		Tags: []string{},
	}
}

func (s *MapReadWriterSub) Start() {
	go s.run()

}

// this function should be run as a go-routine and will handle sending data out
// to the client at the correct rate assuming there are publish requests queued up.
func (s *MapReadWriterSub) run() {
	for {
		s.SequenceID++
		log.Printf("Got publish req on sub #%d.  Sequence %d", s.ID, s.SequenceID)
		// then get all the tags and send them back to the client

		vals := make([]*ua.MonitoredItemNotification, len(s.Tags))
		s.TagListLock.Lock()
		log.Printf("Responding with %d elements", len(s.Tags))
		for i := range s.Tags {
			vals[i] = new(ua.MonitoredItemNotification)
			key := s.Tags[i]
			v := s.MapReadWriter.Data[key]
			dv := new(ua.DataValue)
			switch tv := v.(type) {
			case string:
				dv.Value = ua.MustVariant(tv)
			case int:
				// we can't use an int because it is of unspecified length.  I'm going to use int64 so that we don't
				// have to worry about cutting data off.
				dv.Value = ua.MustVariant(int64(tv))
			case int32:
				dv.Value = ua.MustVariant(tv)
			case float32:
				dv.Value = ua.MustVariant(tv)
			case float64:
				dv.Value = ua.MustVariant(tv)
			case bool:
				dv.Value = ua.MustVariant(tv)
			default:
				dv.Value = ua.MustVariant(tv)
			}
			vals[i].ClientHandle = uint32(i)
			vals[i].Value = dv
		}
		s.TagListLock.Unlock()
		dcn := ua.DataChangeNotification{
			MonitoredItems:  vals,
			DiagnosticInfos: []*ua.DiagnosticInfo{},
		}
		eo := make([]*ua.ExtensionObject, 1)
		eo[0] = ua.NewExtensionObject(dcn)

		msg := ua.NotificationMessage{
			SequenceNumber:   s.SequenceID,
			PublishTime:      time.Now(),
			NotificationData: eo,
		}

		response := &ua.PublishResponse{
			ResponseHeader: &ua.ResponseHeader{
				Timestamp:          time.Now(),
				RequestHandle:      s.SequenceID,
				ServiceResult:      ua.StatusOK,
				ServiceDiagnostics: &ua.DiagnosticInfo{},
				StringTable:        []string{},
				AdditionalHeader:   ua.NewExtensionObject(nil),
			},
			SubscriptionID:      s.ID,
			MoreNotifications:   false,
			NotificationMessage: &msg,
		}
		err := s.Channel.SendResponseWithContext(context.Background(), s.SequenceID, response)
		if err != nil {
			log.Printf("problem sending channel response: %v", err)
			log.Printf("Killing subscription %d", s.ID)
			return
		}
		log.Printf("Publish OK for %d", s.ID)
		// then wait for the publish interval
		time.Sleep(time.Millisecond * time.Duration(s.RevisedPublishingInterval))
		// wait till we've got a publish request.
		log.Printf("Waiting for publish req on sub #%d", s.ID)
		<-s.MapReadWriter.PublishRequests
	}
}

func (s *MapReadWriter) CustomRead(sc *uasc.SecureChannel, r ua.Request) (ua.Response, error) {
	debug.Printf("Handling %T", r)

	req, err := safeReq[*ua.ReadRequest](r)
	if err != nil {
		return nil, err
	}
	s.Mu.RLock()
	defer s.Mu.RUnlock()

	results := make([]*ua.DataValue, len(req.NodesToRead))
	for i, n := range req.NodesToRead {
		debug.Printf("read: node=%s attr=%s", n.NodeID, n.AttributeID)

		dv := &ua.DataValue{
			EncodingMask:    ua.DataValueServerTimestamp | ua.DataValueStatusCode,
			ServerTimestamp: time.Now(),
			Status:          ua.StatusBad,
		}

		key := strip_crap(n.NodeID.String())
		log.Printf("Read req for %s", key)

		// we would normally look up the node in our actual address space, but since that's dumb, we're just
		// going to use the node id directly to look it up from our data map.
		if n.AttributeID == ua.AttributeIDValue {
			dv.Status = ua.StatusOK
			dv.EncodingMask |= ua.DataValueValue
			v := s.Data[key]
			switch tv := v.(type) {
			case string:
				dv.Value = ua.MustVariant(tv)
			case int:
				// we can't use an int because it is of unspecified length.  I'm going to use int64 so that we don't
				// have to worry about cutting data off.
				dv.Value = ua.MustVariant(int64(tv))
			case int32:
				dv.Value = ua.MustVariant(tv)
			case float32:
				dv.Value = ua.MustVariant(tv)
			case float64:
				dv.Value = ua.MustVariant(tv)
			case bool:
				dv.Value = ua.MustVariant(tv)
			default:
				dv.Value = ua.MustVariant(tv)
			}
		}

		debug.Printf("Read '%s' = '%v' (%v)", key, dv.Value, dv.Value.Value())

		results[i] = dv
	}

	response := &ua.ReadResponse{
		ResponseHeader: &ua.ResponseHeader{
			Timestamp:          time.Now(),
			RequestHandle:      req.RequestHeader.RequestHandle,
			ServiceResult:      ua.StatusOK,
			ServiceDiagnostics: &ua.DiagnosticInfo{},
			StringTable:        []string{},
			AdditionalHeader:   ua.NewExtensionObject(nil),
		},
		Results: results,
	}

	return response, nil
}

func safeReq[T ua.Request](r ua.Request) (T, error) {
	var t T
	req, ok := r.(T)
	if !ok {
		debug.Printf("expected %T, got %T", t, r)
		return t, ua.StatusBadRequestTypeInvalid
	}
	return req, nil
}

func (s *MapReadWriter) CustomWrite(sc *uasc.SecureChannel, r ua.Request) (ua.Response, error) {

	req, err := safeReq[*ua.WriteRequest](r)
	if err != nil {
		return nil, err
	}

	s.Mu.Lock()
	defer s.Mu.Unlock()

	status := ua.StatusBad

	for _, n := range req.NodesToWrite {
		debug.Printf("write: node=%s attr=%s", n.NodeID, n.AttributeID)
		key := strip_crap(n.NodeID.String())

		// we would normally look up the node in our actual address space, but since that's dumb, we're just
		// going to use the node id directly to look it up from our data map.
		if n.AttributeID == ua.AttributeIDValue {
			v := n.Value.Value.Value()
			s.Data[key] = v
			status = ua.StatusOK
			debug.Printf("write: key=%s value=%s", key, v)
		}

	}
	response := &ua.ReadResponse{
		ResponseHeader: &ua.ResponseHeader{
			Timestamp:          time.Now(),
			RequestHandle:      req.RequestHeader.RequestHandle,
			ServiceResult:      status,
			ServiceDiagnostics: &ua.DiagnosticInfo{},
			StringTable:        []string{},
			AdditionalHeader:   ua.NewExtensionObject(nil),
		},
	}

	return response, nil
}

func (s *MapReadWriter) CustomBrowse(sc *uasc.SecureChannel, r ua.Request) (ua.Response, error) {
	req, err := safeReq[*ua.BrowseRequest](r)
	if err != nil {
		return nil, err
	}
	s.Mu.RLock()
	defer s.Mu.RUnlock()

	resp := &ua.BrowseResponse{
		ResponseHeader: &ua.ResponseHeader{
			Timestamp:          time.Now(),
			RequestHandle:      req.RequestHeader.RequestHandle,
			ServiceResult:      ua.StatusOK,
			ServiceDiagnostics: &ua.DiagnosticInfo{},
			StringTable:        []string{},
			AdditionalHeader:   ua.NewExtensionObject(nil),
		},

		Results: make([]*ua.BrowseResult, len(req.NodesToBrowse)),
	}

	/*keys := make([]string, len(s.Data))
	i := 0
	for k := range s.Data {
		keys[i] = k
		i++
	}

	slices.Sort(keys)
	*/

	count := len(req.NodesToBrowse)

	debug.Printf("BrowseRequest: len(req.NodesToBrowse)=%d", count)

	for i, bd := range req.NodesToBrowse {
		debug.Printf("BrowseRequest: id=%s mask=%08b\n", bd.NodeID, bd.ResultMask)
		// nodes igniton tried to browse:
		//
		//OPCBinarySchema_TypeSystem
		//BaseDataType
		//ObjectsFolder
		//Server_ServerStatus_State id# 2259 - this appears to be a keepalive / status kind of thing.  doesn't seem to actually matter
		//
		log.Printf("Browse req for %s", bd.NodeID.String())
		if bd.NodeID.IntID() != id.RootFolder && bd.NodeID.IntID() != id.ObjectsFolder {

			resp.Results[i] = &ua.BrowseResult{StatusCode: ua.StatusBadNodeIDUnknown}
			continue
		}

		if bd.NodeID.IntID() == id.RootFolder {

			refs := make([]*ua.ReferenceDescription, 1)
			newid := ua.NewNumericNodeID(0, id.ObjectsFolder)
			expnewid := ua.NewNumericExpandedNodeID(0, id.ObjectsFolder)
			refs[0] = &ua.ReferenceDescription{
				ReferenceTypeID: newid,
				NodeID:          expnewid,
				BrowseName:      &ua.QualifiedName{NamespaceIndex: 0, Name: "objects"},
				DisplayName:     &ua.LocalizedText{Text: "objects"},
				TypeDefinition:  expnewid,
			}

			resp.Results[i] = &ua.BrowseResult{
				StatusCode: ua.StatusGood,
				References: refs,
			}

			continue
		}

		refs := make([]*ua.ReferenceDescription, len(s.Data))

		keyid := 0
		for k := range s.Data {
			key := k
			newid := ua.NewStringNodeID(0, key)
			expnewid := ua.NewStringExpandedNodeID(0, key)

			refs[keyid] = &ua.ReferenceDescription{
				ReferenceTypeID: newid,
				NodeID:          expnewid,
				BrowseName:      &ua.QualifiedName{NamespaceIndex: 0, Name: key},
				DisplayName:     &ua.LocalizedText{Text: key},
				TypeDefinition:  expnewid,
			}
			keyid++
		}

		resp.Results[i] = &ua.BrowseResult{
			StatusCode: ua.StatusGood,
			References: refs,
		}
	}

	return resp, nil
}

func strip_crap(s string) string {
	seq_pos := strings.LastIndex(s, "s=")
	if seq_pos < 0 {
		return s
	}
	return s[seq_pos+2:]
}

func (s *MapReadWriter) CreateSubscription(sc *uasc.SecureChannel, r ua.Request) (ua.Response, error) {
	debug.Printf("Handling %T", r)

	req, err := safeReq[*ua.CreateSubscriptionRequest](r)
	if err != nil {
		return nil, err
	}

	newsubid := uint32(len(s.Subs) + 5)

	log.Printf("Create Sub %d", newsubid)

	sub := MapReadWriterSub{
		MapReadWriter:             s,
		Channel:                   sc,
		ID:                        newsubid,
		RevisedPublishingInterval: req.RequestedPublishingInterval,
		RevisedLifetimeCount:      req.RequestedLifetimeCount,
		RevisedMaxKeepAliveCount:  req.RequestedMaxKeepAliveCount,
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

//PublishRequest_Encoding_DefaultBinary

func (s *MapReadWriter) Publish(sc *uasc.SecureChannel, r ua.Request) (ua.Response, error) {
	log.Printf("Raw Publish req")

	req, err := safeReq[*ua.PublishRequest](r)
	if err != nil {
		log.Printf("ERROR: bad PublishRequest Struct")
		return nil, err
	}
	res := make([]ua.StatusCode, len(req.SubscriptionAcknowledgements))

	for i := range req.SubscriptionAcknowledgements {
		subID := req.SubscriptionAcknowledgements[i].SubscriptionID
		log.Printf("Raw sub req for #%d", subID)
		_, ok := s.Subs[subID]
		if !ok {
			// bad sub ID.
			res[i] = ua.StatusBadSubscriptionIDInvalid
			continue
		}
		// if the sub request buffer isn't full we're OK otherwise send that status.
		select {
		case s.PublishRequests <- struct{}{}:
			res[i] = ua.StatusOK
			log.Printf("Sub Queued for %d", subID)
		default:
			res[i] = ua.StatusBadTooManyPublishRequests
			log.Printf("Sub Queue full for %d", subID)
		}
	}

	select {
	case s.PublishRequests <- struct{}{}:
	default:
		log.Printf("Too many publish reqs.")
	}
	/*
		resp := &ua.PublishResponse{
			ResponseHeader: &ua.ResponseHeader{
				Timestamp:          time.Now(),
				RequestHandle:      req.RequestHeader.RequestHandle,
				ServiceResult:      ua.StatusOK,
				ServiceDiagnostics: &ua.DiagnosticInfo{},
				StringTable:        []string{},
				AdditionalHeader:   ua.NewExtensionObject(nil),
			},
			SubscriptionID:           1,
			AvailableSequenceNumbers: []uint32{},                // []uint32
			MoreNotifications:        false,                     //        bool
			NotificationMessage:      &ua.NotificationMessage{}, //      *NotificationMessage
			Results:                  res,                       //                  []StatusCode
			DiagnosticInfos:          []*ua.DiagnosticInfo{},    //          []*DiagnosticInfo
		}

		return resp, nil
	*/
	return nil, nil
}

func (s *MapReadWriter) CreateMonitoredItems(sc *uasc.SecureChannel, r ua.Request) (ua.Response, error) {
	debug.Printf("Handling %T", r)

	req, err := safeReq[*ua.CreateMonitoredItemsRequest](r)
	if err != nil {
		return nil, err
	}

	count := len(req.ItemsToCreate)

	res := make([]*ua.MonitoredItemCreateResult, count)

	subID := req.SubscriptionID
	log.Printf("Creating monitored items for sub #%d", subID)
	_, ok := s.Subs[subID]
	if !ok {
		return nil, errors.New("sub doesn't exist")
	}
	s.Subs[subID].TagListLock.Lock()
	s.Subs[subID].Tags = make([]string, 0)
	for i := range req.ItemsToCreate {
		item := req.ItemsToCreate[i]
		monitored_key := item.ItemToMonitor.NodeID.String()
		monitored_key = strip_crap(monitored_key)
		log.Printf("Adding monitored item '%s' to sub #%d", monitored_key, subID)
		s.Subs[subID].Tags = append(s.Subs[subID].Tags, monitored_key)
		res[i] = &ua.MonitoredItemCreateResult{
			StatusCode:              ua.StatusOK,
			MonitoredItemID:         uint32(i),
			RevisedSamplingInterval: s.Subs[subID].RevisedPublishingInterval,
			RevisedQueueSize:        1,
			FilterResult:            &ua.ExtensionObject{},
		}

	}
	s.Subs[subID].TagListLock.Unlock()

	resp := &ua.CreateMonitoredItemsResponse{
		ResponseHeader: &ua.ResponseHeader{
			Timestamp:          time.Now(),
			RequestHandle:      req.RequestHeader.RequestHandle,
			ServiceResult:      ua.StatusOK,
			ServiceDiagnostics: &ua.DiagnosticInfo{},
			StringTable:        []string{},
			AdditionalHeader:   ua.NewExtensionObject(nil),
		},
		Results:         res,                    //                  []StatusCode
		DiagnosticInfos: []*ua.DiagnosticInfo{}, //          []*DiagnosticInfo
	}

	return resp, nil

}
