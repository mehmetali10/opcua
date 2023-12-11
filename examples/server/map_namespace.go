package main

import (
	"errors"
	"log"
	"sync"
	"time"

	"github.com/gopcua/opcua/debug"
	"github.com/gopcua/opcua/id"
	"github.com/gopcua/opcua/ua"
	"github.com/gopcua/opcua/uasc"
)

type MapNamespace struct {
	name            string
	Mu              sync.RWMutex
	Data            map[string]any
	Subs            map[uint32]*MapReadWriterSub
	PublishRequests chan PubReq

	id uint16
}

func NewMapNamespace(name string) *MapNamespace {
	mrw := MapNamespace{}
	mrw.name = name
	mrw.Subs = make(map[uint32]*MapReadWriterSub)
	mrw.PublishRequests = make(chan PubReq, 100)
	mrw.Data = make(map[string]any)
	return &mrw
}

func (s *MapNamespace) CreateSubscription(sc *uasc.SecureChannel, r ua.Request, reqID uint32) (ua.Response, error) {
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
		SeqNums:                   make(map[uint32]struct{}),
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

func (s *MapNamespace) Publish(sc *uasc.SecureChannel, r ua.Request, reqID uint32) (ua.Response, error) {
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

	return nil, nil
}

func (s *MapNamespace) CreateMonitoredItems(sc *uasc.SecureChannel, r ua.Request, reqID uint32) (ua.Response, error) {
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
	s.Subs[subID].Tags = make(map[uint32]string)
	for i := range req.ItemsToCreate {
		item := req.ItemsToCreate[i]
		monitored_key := item.ItemToMonitor.NodeID.String()
		monitored_key = strip_crap(monitored_key)
		log.Printf("Adding monitored item '%s' to sub #%d as %d", monitored_key, subID, item.RequestedParameters.ClientHandle)
		s.Subs[subID].Tags[item.RequestedParameters.ClientHandle] = monitored_key
		res[i] = &ua.MonitoredItemCreateResult{
			StatusCode:              ua.StatusOK,
			MonitoredItemID:         uint32(i),
			RevisedSamplingInterval: s.Subs[subID].RevisedPublishingInterval,
			RevisedQueueSize:        1,
			FilterResult:            ua.NewExtensionObject(nil),
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

func (s *MapNamespace) ID() uint16 {
	return s.id
}
func (ns *MapNamespace) SetID(id uint16) {
	ns.id = id
}

func (ns *MapNamespace) Browse(bd *ua.BrowseDescription) *ua.BrowseResult {
	ns.Mu.RLock()
	defer ns.Mu.RUnlock()

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

		return &ua.BrowseResult{StatusCode: ua.StatusBadNodeIDUnknown}
	}

	if bd.NodeID.IntID() == id.RootFolder {

		refs := make([]*ua.ReferenceDescription, 1)
		newid := ua.NewNumericNodeID(ns.id, id.ObjectsFolder)
		expnewid := ua.NewNumericExpandedNodeID(ns.id, id.ObjectsFolder)
		refs[0] = &ua.ReferenceDescription{
			ReferenceTypeID: newid,
			NodeID:          expnewid,
			BrowseName:      &ua.QualifiedName{NamespaceIndex: ns.id, Name: "objects"},
			DisplayName:     &ua.LocalizedText{Text: "objects"},
			TypeDefinition:  expnewid,
		}

		return &ua.BrowseResult{
			StatusCode: ua.StatusGood,
			References: refs,
		}

	}

	refs := make([]*ua.ReferenceDescription, len(ns.Data))

	keyid := 0
	for k := range ns.Data {
		key := k
		newid := ua.NewStringNodeID(ns.id, key)
		expnewid := ua.NewStringExpandedNodeID(ns.id, key)

		refs[keyid] = &ua.ReferenceDescription{
			ReferenceTypeID: newid,
			IsForward:       true,
			NodeID:          expnewid,
			BrowseName:      &ua.QualifiedName{NamespaceIndex: ns.ID(), Name: key},
			DisplayName:     &ua.LocalizedText{EncodingMask: ua.LocalizedTextText, Text: key},
			NodeClass:       ua.NodeClassVariable, // when support is added for nested maps, this will be NodeClassObject
			TypeDefinition:  expnewid,
		}
		keyid++
	}

	return &ua.BrowseResult{
		StatusCode: ua.StatusGood,
		References: refs,
	}

}

func (ns *MapNamespace) Attribute(n *ua.NodeID, a ua.AttributeID) *ua.DataValue {
	log.Printf("read: node=%s attr=%s", n.String(), a)

	if n.IntID() != 0 {
		// this is not one of our normal tags.
		if n.IntID() != id.ObjectsFolder {
			return &ua.DataValue{
				EncodingMask:    ua.DataValueServerTimestamp | ua.DataValueStatusCode,
				ServerTimestamp: time.Now(),
				Status:          ua.StatusBadNodeIDInvalid,
			}
		}

		attrval, err := ns.Objects().Attribute(a)
		if err != nil {
			return &ua.DataValue{
				EncodingMask:    ua.DataValueServerTimestamp | ua.DataValueStatusCode,
				ServerTimestamp: time.Now(),
				Status:          ua.StatusBadAttributeIDInvalid,
			}
		}

		return &ua.DataValue{
			EncodingMask:    ua.DataValueServerTimestamp | ua.DataValueStatusCode | ua.DataValueValue,
			ServerTimestamp: time.Now(),
			Status:          ua.StatusOK,
			Value:           attrval.Value,
		}

	}

	dv := &ua.DataValue{
		EncodingMask:    ua.DataValueServerTimestamp | ua.DataValueStatusCode,
		ServerTimestamp: time.Now(),
		Status:          ua.StatusBad,
	}

	key := strip_crap(n.String())
	log.Printf("Read req for %s", key)
	var err error

	// we are going to use the node id directly to look it up from our data map.
	if a == ua.AttributeIDValue {
		dv.Status = ua.StatusOK
		dv.EncodingMask |= ua.DataValueValue
		v, ok := ns.Data[key]
		if !ok {
			return &ua.DataValue{
				EncodingMask:    ua.DataValueServerTimestamp | ua.DataValueStatusCode,
				ServerTimestamp: time.Now(),
				Status:          ua.StatusBadNodeIDUnknown,
			}
		}
		switch tv := v.(type) {
		case string:
			dv.Value = ua.MustVariant(tv)
		case int:
			// we can't use an int because it is of unspecified length.  I'm going to use int64 so that we don't
			// have to worry about cutting data off. probably.
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
	// nothing in this namespace has an ID Description
	if a == ua.AttributeIDDescription {
		dv.Status = ua.StatusOK
		dv.EncodingMask |= ua.DataValueValue
		dv.Value = ua.MustVariant("")
	}
	// nothing in this namespace has event notifiers
	if a == ua.AttributeIDEventNotifier {
		dv.Status = ua.StatusOK
		dv.EncodingMask |= ua.DataValueValue
		dv.Value = ua.MustVariant(int16(0))
	}

	// values are in section 5.1.2 of the standard.
	// https://reference.opcfoundation.org/Core/Part6/v104/docs/5.1.2
	if a == ua.AttributeIDDataType {
		dv.Status = ua.StatusOK
		dv.EncodingMask |= ua.DataValueValue
		v := ns.Data[key]
		switch v.(type) {
		case string:
			dv.Value, err = ua.NewVariant(ua.NewNumericNodeID(0, 12))
			if err != nil {
				log.Printf("problem creating variant: %v", err)
			}
		case int:
			// we can't use an int because it is of unspecified length.  I'm going to use int64 so that we don't
			// have to worry about cutting data off.
			dv.Value, err = ua.NewVariant(ua.NewNumericNodeID(0, 6))
			if err != nil {
				log.Printf("problem creating variant: %v", err)
			}
		case int32:
			dv.Value, err = ua.NewVariant(ua.NewNumericNodeID(0, 6))
			if err != nil {
				log.Printf("problem creating variant: %v", err)
			}
		case float32:
			dv.Value, err = ua.NewVariant(ua.NewNumericNodeID(0, 10))
			if err != nil {
				log.Printf("problem creating variant: %v", err)
			}
		case float64:
			dv.Value, err = ua.NewVariant(ua.NewNumericNodeID(0, 11))
			if err != nil {
				log.Printf("problem creating variant: %v", err)
			}
		case bool:
			dv.Value, err = ua.NewVariant(ua.NewNumericNodeID(0, 1))
			if err != nil {
				log.Printf("problem creating variant: %v", err)
			}
		default:
			dv.Value, err = ua.NewVariant(ua.NewNumericNodeID(0, 24))
			if err != nil {
				log.Printf("problem creating variant: %v", err)
			}
		}
	}

	// when we support arrays this will have to change.
	if a == ua.AttributeIDValueRank {
		dv.Status = ua.StatusOK
		dv.EncodingMask |= ua.DataValueValue
		dv.Value = ua.MustVariant(int32(-1))
	}

	// when we support arrays this will have to change.
	if a == ua.AttributeIDArrayDimensions {
		dv.Status = ua.StatusOK
		dv.EncodingMask |= ua.DataValueValue
		dv.Value = ua.MustVariant([]uint32{})
	}

	if dv.Value == nil {
		log.Printf("bad dv.Value!")
	}
	debug.Printf("Read '%s' = '%v' (%v)", key, dv.Value, dv.Value.Value())

	return dv
}
