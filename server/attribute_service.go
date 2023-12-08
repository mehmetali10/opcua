package server

import (
	"time"

	"github.com/gopcua/opcua/debug"
	"github.com/gopcua/opcua/ua"
	"github.com/gopcua/opcua/uasc"
)

// AttributeService implements the Attribute Service Set.
//
// https://reference.opcfoundation.org/Core/Part4/v105/docs/5.10
type AttributeService struct {
	srv *Server
}

// https://reference.opcfoundation.org/Core/Part4/v105/docs/5.10.2
func (s *AttributeService) Read(sc *uasc.SecureChannel, r ua.Request, reqID uint32) (ua.Response, error) {
	debug.Printf("Handling %T", r)

	req, err := safeReq[*ua.ReadRequest](r)
	if err != nil {
		return nil, err
	}

	results := make([]*ua.DataValue, len(req.NodesToRead))
	for i, n := range req.NodesToRead {
		debug.Printf("read: node=%s attr=%s", n.NodeID, n.AttributeID)

		dv := &ua.DataValue{
			EncodingMask:    ua.DataValueServerTimestamp,
			ServerTimestamp: time.Now(),
		}

		a := s.srv.AddressSpace()
		v, err := a.Attribute(n.NodeID, n.AttributeID)
		switch x := err.(type) {
		case nil:
			dv.EncodingMask |= ua.DataValueStatusCode | ua.DataValueValue
			dv.Status = ua.StatusOK
			dv.Value = v.Value

		case ua.StatusCode:
			dv.EncodingMask |= ua.DataValueStatusCode
			dv.Status = x

		default:
			debug.Printf("read: node=%s attr=%s err=%s", n.NodeID, n.AttributeID, err)
			dv.EncodingMask |= ua.DataValueStatusCode
			dv.Status = ua.StatusBadInternalError
		}

		if v != nil && !v.SourceTimestamp.IsZero() {
			dv.EncodingMask |= ua.DataValueSourceTimestamp
			dv.SourceTimestamp = v.SourceTimestamp
		}

		results[i] = dv
	}

	response := &ua.ReadResponse{
		ResponseHeader: responseHeader(req.RequestHeader.RequestHandle, ua.StatusOK),
		Results:        results,
	}

	return response, nil
}

// https://reference.opcfoundation.org/Core/Part4/v105/docs/5.10.3
func (s *AttributeService) HistoryRead(sc *uasc.SecureChannel, r ua.Request, reqID uint32) (ua.Response, error) {
	debug.Printf("Handling %T", r)

	req, err := safeReq[*ua.HistoryReadRequest](r)
	if err != nil {
		return nil, err
	}
	return serviceUnsupported(req.RequestHeader), nil
}

// https://reference.opcfoundation.org/Core/Part4/v105/docs/5.10.4
func (s *AttributeService) Write(sc *uasc.SecureChannel, r ua.Request, reqID uint32) (ua.Response, error) {
	debug.Printf("Handling %T", r)

	req, err := safeReq[*ua.WriteRequest](r)
	if err != nil {
		return nil, err
	}
	return serviceUnsupported(req.RequestHeader), nil
}

// https://reference.opcfoundation.org/Core/Part4/v105/docs/5.10.5
func (s *AttributeService) HistoryUpdate(sc *uasc.SecureChannel, r ua.Request, reqID uint32) (ua.Response, error) {
	debug.Printf("Handling %T", r)

	req, err := safeReq[*ua.HistoryUpdateRequest](r)
	if err != nil {
		return nil, err
	}
	return serviceUnsupported(req.RequestHeader), nil
}
