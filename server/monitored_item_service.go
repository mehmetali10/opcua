package server

import (
	"github.com/gopcua/opcua/debug"
	"github.com/gopcua/opcua/ua"
	"github.com/gopcua/opcua/uasc"
)

type MonitoredItemService struct {
	srv *Server
}

func (s *MonitoredItemService) MonitoredItemCreateRequest(sc *uasc.SecureChannel, r ua.Request) (ua.Response, error) {
	debug.Printf("Handling %T\n", r)

	//req, ok := r.(*ua.MonitoredItemCreateRequest)
	//if !ok {
	//	debug.Printf("handleMonitoredItemCreateRequest: Expected *ua.MonitoredItemCreateRequest, got %T", r)
	//	return nil, ua.StatusBadRequestTypeInvalid
	//}

	//TODO: Replace with proper response once implemented
	response := &ua.ServiceFault{ResponseHeader: responseHeader(0, ua.StatusBadServiceUnsupported)}
	// response := &ua.MonitoredItemCreateResponse{
	//	ResponseHeader: responseHeader(req.RequestHeader.RequestHandle, ua.StatusOK),
	//  ... remaining fields
	//}

	return response, nil
}

func (s *MonitoredItemService) CreateMonitoredItemsRequest(sc *uasc.SecureChannel, r ua.Request) (ua.Response, error) {
	debug.Printf("Handling %T\n", r)

	req, ok := r.(*ua.CreateMonitoredItemsRequest)
	if !ok {
		debug.Printf("handleCreateMonitoredItemsRequest: Expected *ua.CreateMonitoredItemsRequest, got %T", r)
		return nil, ua.StatusBadRequestTypeInvalid
	}

	//TODO: Replace with proper response once implemented
	response := &ua.ServiceFault{ResponseHeader: responseHeader(req.RequestHeader.RequestHandle, ua.StatusBadServiceUnsupported)}
	// response := &ua.CreateMonitoredItemsResponse{
	//	ResponseHeader: responseHeader(req.RequestHeader.RequestHandle, ua.StatusOK),
	//  ... remaining fields
	//}

	return response, nil
}

func (s *MonitoredItemService) MonitoredItemModifyRequest(sc *uasc.SecureChannel, r ua.Request) (ua.Response, error) {
	debug.Printf("Handling %T\n", r)

	//req, ok := r.(*ua.MonitoredItemModifyRequest)
	//if !ok {
	//	debug.Printf("handleMonitoredItemModifyRequest: Expected *ua.MonitoredItemModifyRequest, got %T", r)
	//	return nil, ua.StatusBadRequestTypeInvalid
	//}

	//TODO: Replace with proper response once implemented
	response := &ua.ServiceFault{ResponseHeader: responseHeader(0, ua.StatusBadServiceUnsupported)}
	// response := &ua.MonitoredItemModifyResponse{
	//	ResponseHeader: responseHeader(req.RequestHeader.RequestHandle, ua.StatusOK),
	//  ... remaining fields
	//}

	return response, nil
}

func (s *MonitoredItemService) ModifyMonitoredItemsRequest(sc *uasc.SecureChannel, r ua.Request) (ua.Response, error) {
	debug.Printf("Handling %T\n", r)

	req, ok := r.(*ua.ModifyMonitoredItemsRequest)
	if !ok {
		debug.Printf("handleModifyMonitoredItemsRequest: Expected *ua.ModifyMonitoredItemsRequest, got %T", r)
		return nil, ua.StatusBadRequestTypeInvalid
	}

	//TODO: Replace with proper response once implemented
	response := &ua.ServiceFault{ResponseHeader: responseHeader(req.RequestHeader.RequestHandle, ua.StatusBadServiceUnsupported)}
	// response := &ua.ModifyMonitoredItemsResponse{
	//	ResponseHeader: responseHeader(req.RequestHeader.RequestHandle, ua.StatusOK),
	//  ... remaining fields
	//}

	return response, nil
}

func (s *MonitoredItemService) SetMonitoringModeRequest(sc *uasc.SecureChannel, r ua.Request) (ua.Response, error) {
	debug.Printf("Handling %T\n", r)

	req, ok := r.(*ua.SetMonitoringModeRequest)
	if !ok {
		debug.Printf("handleSetMonitoringModeRequest: Expected *ua.SetMonitoringModeRequest, got %T", r)
		return nil, ua.StatusBadRequestTypeInvalid
	}

	//TODO: Replace with proper response once implemented
	response := &ua.ServiceFault{ResponseHeader: responseHeader(req.RequestHeader.RequestHandle, ua.StatusBadServiceUnsupported)}
	// response := &ua.SetMonitoringModeResponse{
	//	ResponseHeader: responseHeader(req.RequestHeader.RequestHandle, ua.StatusOK),
	//  ... remaining fields
	//}

	return response, nil
}

func (s *MonitoredItemService) SetTriggeringRequest(sc *uasc.SecureChannel, r ua.Request) (ua.Response, error) {
	debug.Printf("Handling %T\n", r)

	req, ok := r.(*ua.SetTriggeringRequest)
	if !ok {
		debug.Printf("handleSetTriggeringRequest: Expected *ua.SetTriggeringRequest, got %T", r)
		return nil, ua.StatusBadRequestTypeInvalid
	}

	//TODO: Replace with proper response once implemented
	response := &ua.ServiceFault{ResponseHeader: responseHeader(req.RequestHeader.RequestHandle, ua.StatusBadServiceUnsupported)}
	// response := &ua.SetTriggeringResponse{
	//	ResponseHeader: responseHeader(req.RequestHeader.RequestHandle, ua.StatusOK),
	//  ... remaining fields
	//}

	return response, nil
}

func (s *MonitoredItemService) DeleteMonitoredItemsRequest(sc *uasc.SecureChannel, r ua.Request) (ua.Response, error) {
	debug.Printf("Handling %T\n", r)

	req, ok := r.(*ua.DeleteMonitoredItemsRequest)
	if !ok {
		debug.Printf("handleDeleteMonitoredItemsRequest: Expected *ua.DeleteMonitoredItemsRequest, got %T", r)
		return nil, ua.StatusBadRequestTypeInvalid
	}

	//TODO: Replace with proper response once implemented
	response := &ua.ServiceFault{ResponseHeader: responseHeader(req.RequestHeader.RequestHandle, ua.StatusBadServiceUnsupported)}
	// response := &ua.DeleteMonitoredItemsResponse{
	//	ResponseHeader: responseHeader(req.RequestHeader.RequestHandle, ua.StatusOK),
	//  ... remaining fields
	//}

	return response, nil
}
