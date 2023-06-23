package server

import (
	"github.com/gopcua/opcua/debug"
	"github.com/gopcua/opcua/ua"
	"github.com/gopcua/opcua/uasc"
)

// QueryService implements the Query Service Set.
//
// https://reference.opcfoundation.org/Core/Part4/v105/docs/5.9
type QueryService struct {
	srv *Server
}

// https://reference.opcfoundation.org/Core/Part4/v105/docs/5.9.3
func (s *QueryService) QueryFirstRequest(sc *uasc.SecureChannel, r ua.Request) (ua.Response, error) {
	debug.Printf("Handling %T\n", r)

	req, ok := r.(*ua.QueryFirstRequest)
	if !ok {
		debug.Printf("handleQueryFirstRequest: Expected *ua.QueryFirstRequest, got %T", r)
		return nil, ua.StatusBadRequestTypeInvalid
	}

	//TODO: Replace with proper response once implemented
	response := &ua.ServiceFault{ResponseHeader: responseHeader(req.RequestHeader.RequestHandle, ua.StatusBadServiceUnsupported)}
	// response := &ua.QueryFirstResponse{
	//	ResponseHeader: responseHeader(req.RequestHeader.RequestHandle, ua.StatusOK),
	//  ... remaining fields
	//}

	return response, nil
}

// https://reference.opcfoundation.org/Core/Part4/v105/docs/5.9.4
func (s *QueryService) QueryNextRequest(sc *uasc.SecureChannel, r ua.Request) (ua.Response, error) {
	debug.Printf("Handling %T\n", r)

	req, ok := r.(*ua.QueryNextRequest)
	if !ok {
		debug.Printf("handleQueryNextRequest: Expected *ua.QueryNextRequest, got %T", r)
		return nil, ua.StatusBadRequestTypeInvalid
	}

	//TODO: Replace with proper response once implemented
	response := &ua.ServiceFault{ResponseHeader: responseHeader(req.RequestHeader.RequestHandle, ua.StatusBadServiceUnsupported)}
	// response := &ua.QueryNextResponse{
	//	ResponseHeader: responseHeader(req.RequestHeader.RequestHandle, ua.StatusOK),
	//  ... remaining fields
	//}

	return response, nil
}
