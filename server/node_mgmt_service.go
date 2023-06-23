package server

import (
	"github.com/gopcua/opcua/debug"
	"github.com/gopcua/opcua/ua"
	"github.com/gopcua/opcua/uasc"
)

// NodeManagementService implements the Node Management Service Set.
//
// https://reference.opcfoundation.org/Core/Part4/v105/docs/5.7
type NodeManagementService struct {
	srv *Server
}

// https://reference.opcfoundation.org/Core/Part4/v105/docs/5.7.2
func (s *NodeManagementService) AddNodesRequest(sc *uasc.SecureChannel, r ua.Request) (ua.Response, error) {
	debug.Printf("Handling %T\n", r)

	req, ok := r.(*ua.AddNodesRequest)
	if !ok {
		debug.Printf("handleAddNodesRequest: Expected *ua.AddNodesRequest, got %T", r)
		return nil, ua.StatusBadRequestTypeInvalid
	}

	//TODO: Replace with proper response once implemented
	response := &ua.ServiceFault{ResponseHeader: responseHeader(req.RequestHeader.RequestHandle, ua.StatusBadServiceUnsupported)}
	// response := &ua.AddNodesResponse{
	//	ResponseHeader: responseHeader(req.RequestHeader.RequestHandle, ua.StatusOK),
	//  ... remaining fields
	//}

	return response, nil
}

// https://reference.opcfoundation.org/Core/Part4/v105/docs/5.7.3
func (s *NodeManagementService) AddReferencesRequest(sc *uasc.SecureChannel, r ua.Request) (ua.Response, error) {
	debug.Printf("Handling %T\n", r)

	req, ok := r.(*ua.AddReferencesRequest)
	if !ok {
		debug.Printf("handleAddReferencesRequest: Expected *ua.AddReferencesRequest, got %T", r)
		return nil, ua.StatusBadRequestTypeInvalid
	}

	//TODO: Replace with proper response once implemented
	response := &ua.ServiceFault{ResponseHeader: responseHeader(req.RequestHeader.RequestHandle, ua.StatusBadServiceUnsupported)}
	// response := &ua.AddReferencesResponse{
	//	ResponseHeader: responseHeader(req.RequestHeader.RequestHandle, ua.StatusOK),
	//  ... remaining fields
	//}

	return response, nil
}

// https://reference.opcfoundation.org/Core/Part4/v105/docs/5.7.4
func (s *NodeManagementService) DeleteNodesRequest(sc *uasc.SecureChannel, r ua.Request) (ua.Response, error) {
	debug.Printf("Handling %T\n", r)

	req, ok := r.(*ua.DeleteNodesRequest)
	if !ok {
		debug.Printf("handleDeleteNodesRequest: Expected *ua.DeleteNodesRequest, got %T", r)
		return nil, ua.StatusBadRequestTypeInvalid
	}

	//TODO: Replace with proper response once implemented
	response := &ua.ServiceFault{ResponseHeader: responseHeader(req.RequestHeader.RequestHandle, ua.StatusBadServiceUnsupported)}
	// response := &ua.DeleteNodesResponse{
	//	ResponseHeader: responseHeader(req.RequestHeader.RequestHandle, ua.StatusOK),
	//  ... remaining fields
	//}

	return response, nil
}

// https://reference.opcfoundation.org/Core/Part4/v105/docs/5.7.5
func (s *NodeManagementService) DeleteReferencesRequest(sc *uasc.SecureChannel, r ua.Request) (ua.Response, error) {
	debug.Printf("Handling %T\n", r)

	req, ok := r.(*ua.DeleteReferencesRequest)
	if !ok {
		debug.Printf("handleDeleteReferencesRequest: Expected *ua.DeleteReferencesRequest, got %T", r)
		return nil, ua.StatusBadRequestTypeInvalid
	}

	//TODO: Replace with proper response once implemented
	response := &ua.ServiceFault{ResponseHeader: responseHeader(req.RequestHeader.RequestHandle, ua.StatusBadServiceUnsupported)}
	// response := &ua.DeleteReferencesResponse{
	//	ResponseHeader: responseHeader(req.RequestHeader.RequestHandle, ua.StatusOK),
	//  ... remaining fields
	//}

	return response, nil
}
