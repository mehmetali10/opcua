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
func (s *AttributeService) ReadRequest(sc *uasc.SecureChannel, r ua.Request) (ua.Response, error) {
	debug.Printf("Handling %T\n", r)

	req, ok := r.(*ua.ReadRequest)
	if !ok {
		debug.Printf("handleReadRequest: Expected *ua.ReadRequest, got %T", r)
		return nil, ua.StatusBadRequestTypeInvalid
	}

	results := make([]*ua.DataValue, len(req.NodesToRead))
	for i, n := range req.NodesToRead {
		debug.Printf("read: node=%s attr=%s", n.NodeID, n.AttributeID)

		dv := &ua.DataValue{
			EncodingMask:    ua.DataValueServerTimestamp,
			ServerTimestamp: time.Now(),
		}

		v, err := s.srv.AddressSpace().Attribute(n.NodeID, n.AttributeID)
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
func (s *AttributeService) HistoryReadRequest(sc *uasc.SecureChannel, r ua.Request) (ua.Response, error) {
	debug.Printf("Handling %T\n", r)

	req, ok := r.(*ua.HistoryReadRequest)
	if !ok {
		debug.Printf("handleHistoryReadRequest: Expected *ua.HistoryReadRequest, got %T", r)
		return nil, ua.StatusBadRequestTypeInvalid
	}

	//TODO: Replace with proper response once implemented
	response := &ua.ServiceFault{ResponseHeader: responseHeader(req.RequestHeader.RequestHandle, ua.StatusBadServiceUnsupported)}
	// response := &ua.HistoryReadResponse{
	//	ResponseHeader: responseHeader(req.RequestHeader.RequestHandle, ua.StatusOK),
	//  ... remaining fields
	//}

	return response, nil
}

// https://reference.opcfoundation.org/Core/Part4/v105/docs/5.10.4
func (s *AttributeService) WriteRequest(sc *uasc.SecureChannel, r ua.Request) (ua.Response, error) {
	debug.Printf("Handling %T\n", r)

	req, ok := r.(*ua.WriteRequest)
	if !ok {
		debug.Printf("handleWriteRequest: Expected *ua.WriteRequest, got %T", r)
		return nil, ua.StatusBadRequestTypeInvalid
	}

	//TODO: Replace with proper response once implemented
	response := &ua.ServiceFault{ResponseHeader: responseHeader(req.RequestHeader.RequestHandle, ua.StatusBadServiceUnsupported)}
	// response := &ua.WriteResponse{
	//	ResponseHeader: responseHeader(req.RequestHeader.RequestHandle, ua.StatusOK),
	//  ... remaining fields
	//}

	return response, nil
}

// https://reference.opcfoundation.org/Core/Part4/v105/docs/5.10.5
func (s *AttributeService) HistoryUpdateRequest(sc *uasc.SecureChannel, r ua.Request) (ua.Response, error) {
	debug.Printf("Handling %T\n", r)

	req, ok := r.(*ua.HistoryUpdateRequest)
	if !ok {
		debug.Printf("handleHistoryUpdateRequest: Expected *ua.HistoryUpdateRequest, got %T", r)
		return nil, ua.StatusBadRequestTypeInvalid
	}

	//TODO: Replace with proper response once implemented
	response := &ua.ServiceFault{ResponseHeader: responseHeader(req.RequestHeader.RequestHandle, ua.StatusBadServiceUnsupported)}
	// response := &ua.HistoryUpdateResponse{
	//	ResponseHeader: responseHeader(req.RequestHeader.RequestHandle, ua.StatusOK),
	//  ... remaining fields
	//}

	return response, nil
}
