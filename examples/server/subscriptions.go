package main

import (
	"context"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/gopcua/opcua/debug"
	"github.com/gopcua/opcua/id"
	"github.com/gopcua/opcua/server"
	"github.com/gopcua/opcua/server/attrs"
	"github.com/gopcua/opcua/ua"
	"github.com/gopcua/opcua/uasc"
)

type PubReq struct {
	Req *ua.PublishRequest
	ID  uint32
}

type MapReadWriterSub struct {
	MapReadWriter             *MapNamespace
	ID                        uint32
	RevisedPublishingInterval float64
	RevisedLifetimeCount      uint32
	RevisedMaxKeepAliveCount  uint32
	Tags                      map[uint32]string
	TagListLock               sync.Mutex
	Channel                   *uasc.SecureChannel
	SequenceID                uint32
	SeqNums                   map[uint32]struct{}
}

func NewMapReadWriterSub() *MapReadWriterSub {
	return &MapReadWriterSub{
		Tags:    map[uint32]string{},
		SeqNums: map[uint32]struct{}{},
	}
}

func (s *MapReadWriterSub) Start() {
	go s.run()

}

// this function should be run as a go-routine and will handle sending data out
// to the client at the correct rate assuming there are publish requests queued up.
func (s *MapReadWriterSub) run() {
	for {
		// wait for the publish interval
		time.Sleep(time.Millisecond * time.Duration(s.RevisedPublishingInterval))
		s.TagListLock.Lock()
		l := len(s.Tags)
		s.TagListLock.Unlock()
		if l == 0 {
			log.Print("No tags in sub list")
			continue
		}
		log.Printf("Waiting for publish req on sub #%d", s.ID)
		pubreq := <-s.MapReadWriter.PublishRequests
		// see if we have any tags
		s.SequenceID++
		log.Printf("Got publish req on sub #%d.  Sequence %d", s.ID, s.SequenceID)
		// then get all the tags and send them back to the client

		vals := make([]*ua.MonitoredItemNotification, len(s.Tags))
		s.TagListLock.Lock()
		log.Printf("Responding with %d elements: %v", len(s.Tags), s.Tags)
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
			dv.EncodingMask |= ua.DataValueStatusCode
			dv.Status = ua.StatusOK
			dv.EncodingMask |= ua.DataValueSourceTimestamp
			dv.SourceTimestamp = time.Now()
			dv.EncodingMask |= ua.DataValueValue
			vals[i].ClientHandle = i
			vals[i].Value = dv
			log.Printf("Value: %v", dv.Value.Value())
		}
		for x := range pubreq.Req.SubscriptionAcknowledgements {
			a := pubreq.Req.SubscriptionAcknowledgements[x]
			delete(s.SeqNums, a.SequenceNumber)
		}
		log.Printf("values: %v", *vals[0])
		s.TagListLock.Unlock()
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

func safeReq[T ua.Request](r ua.Request) (T, error) {
	var t T
	req, ok := r.(T)
	if !ok {
		debug.Printf("expected %T, got %T", t, r)
		return t, ua.StatusBadRequestTypeInvalid
	}
	return req, nil
}

func (s *MapNamespace) CustomWrite(sc *uasc.SecureChannel, r ua.Request, reqID uint32) (ua.Response, error) {

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

func strip_crap(s string) string {
	seq_pos := strings.LastIndex(s, "s=")
	if seq_pos < 0 {
		return s
	}
	return s[seq_pos+2:]
}

func (ns *MapNamespace) Name() string {
	return ns.name
}
func (ns *MapNamespace) AddNode(n *server.Node) *server.Node {
	return n
}
func (ns *MapNamespace) Node(id *ua.NodeID) *server.Node {
	return nil

}
func (ns *MapNamespace) Objects() *server.Node {
	oid := ua.NewNumericNodeID(ns.ID(), id.ObjectsFolder)
	eoid := ua.NewNumericExpandedNodeID(ns.ID(), id.ObjectsFolder)
	reftype := ua.NewNumericNodeID(ns.ID(), id.Organizes) // folder
	n := server.NewNode(
		oid,
		map[ua.AttributeID]*ua.Variant{
			ua.AttributeIDNodeClass:   ua.MustVariant(uint32(ua.NodeClassObject)),
			ua.AttributeIDBrowseName:  ua.MustVariant(attrs.BrowseName(ns.name)),
			ua.AttributeIDDisplayName: ua.MustVariant(attrs.DisplayName(ns.name, ns.name)),
		},
		[]*ua.ReferenceDescription{{
			ReferenceTypeID: reftype,
			IsForward:       true,
			NodeID:          eoid,
			BrowseName:      &ua.QualifiedName{NamespaceIndex: ns.ID(), Name: ns.name},
			DisplayName:     &ua.LocalizedText{EncodingMask: ua.LocalizedTextText, Text: ns.name},
			NodeClass:       ua.NodeClassObject,
			TypeDefinition:  eoid,
		}},
		nil,
	)
	return n

}
func (ns *MapNamespace) Root() *server.Node {
	n := server.NewNode(
		ua.NewNumericNodeID(ns.ID(), id.RootFolder),
		map[ua.AttributeID]*ua.Variant{
			ua.AttributeIDNodeClass:   ua.MustVariant(uint32(ua.NodeClassObject)),
			ua.AttributeIDBrowseName:  ua.MustVariant(attrs.BrowseName("Root")),
			ua.AttributeIDDisplayName: ua.MustVariant(attrs.DisplayName("Root", "")),
		},
		[]*ua.ReferenceDescription{},
		nil,
	)
	return n

}
