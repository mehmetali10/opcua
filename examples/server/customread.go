package main

import (
	"strings"
	"sync"
	"time"

	"github.com/gopcua/opcua/debug"
	"github.com/gopcua/opcua/ua"
	"github.com/gopcua/opcua/uasc"
)

type MapReadWriter struct {
	Mu   sync.RWMutex
	Data map[string]any
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

		key := strip_sequals(n.NodeID.String())

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
		key := strip_sequals(n.NodeID.String())

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
		if bd.NodeID.IntID() != 84 {

			resp.Results[i] = &ua.BrowseResult{StatusCode: ua.StatusBadNodeIDUnknown}
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

func strip_sequals(s string) string {
	return strings.TrimPrefix(s, "s=")
}
