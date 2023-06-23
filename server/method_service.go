package server

import (
	"github.com/gopcua/opcua/debug"
	"github.com/gopcua/opcua/ua"
	"github.com/gopcua/opcua/uasc"
)

// MethodService implements the Method Service Set.
//
// https://reference.opcfoundation.org/Core/Part4/v105/docs/5.11
type MethodService struct {
	srv *Server
}

// https://reference.opcfoundation.org/Core/Part4/v105/docs/5.11.2
func (s *MethodService) Call(sc *uasc.SecureChannel, r ua.Request) (ua.Response, error) {
	debug.Printf("Handling %T\n", r)

	req, ok := r.(*ua.CallRequest)
	if !ok {
		debug.Printf("handleCallRequest: Expected *ua.CallRequest, got %T", r)
		return nil, ua.StatusBadRequestTypeInvalid
	}

	//TODO: Replace with proper response once implemented
	response := &ua.ServiceFault{ResponseHeader: responseHeader(req.RequestHeader.RequestHandle, ua.StatusBadServiceUnsupported)}
	// response := &ua.CallResponse{
	//	ResponseHeader: responseHeader(req.RequestHeader.RequestHandle, ua.StatusOK),
	//  ... remaining fields
	//}

	return response, nil
}
