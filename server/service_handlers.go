// Copyright 2018-2019 opcua authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package server

import (
	"context"
	"time"

	"github.com/gopcua/opcua/debug"
	"github.com/gopcua/opcua/id"
	"github.com/gopcua/opcua/ua"
	"github.com/gopcua/opcua/uasc"
)

type Handler func(*uasc.SecureChannel, ua.Request) (ua.Response, error)

func (s *Server) initHandlers() {
	// s.registerHandlerFunc(id.ServiceFault_Encoding_DefaultBinary, handleServiceFault)

	discovery := &DiscoveryService{s}
	s.registerHandler(id.FindServersRequest_Encoding_DefaultBinary, discovery.FindServersRequest)
	s.registerHandler(id.FindServersOnNetworkRequest_Encoding_DefaultBinary, discovery.FindServersOnNetworkRequest)
	s.registerHandler(id.GetEndpointsRequest_Encoding_DefaultBinary, discovery.GetEndpointsRequest)
	s.registerHandler(id.RegisterServerRequest_Encoding_DefaultBinary, discovery.RegisterServerRequest)
	s.registerHandler(id.RegisterServer2Request_Encoding_DefaultBinary, discovery.RegisterServer2Request)

	// SecureChannel service (handled in the uasc stack)
	// s.registerHandlerFunc(id.OpenSecureChannelRequest_Encoding_DefaultBinary, handleOpenSecureChannelRequest)
	// s.registerHandlerFunc(id.CloseSecureChannelRequest_Encoding_DefaultBinary, handleCloseSecureChannelRequest)

	session := &SessionService{s}
	s.registerHandler(id.CreateSessionRequest_Encoding_DefaultBinary, session.CreateSessionRequest)
	s.registerHandler(id.ActivateSessionRequest_Encoding_DefaultBinary, session.ActivateSessionRequest)
	s.registerHandler(id.CloseSessionRequest_Encoding_DefaultBinary, session.CloseSessionRequest)
	s.registerHandler(id.CancelRequest_Encoding_DefaultBinary, session.CancelRequest)

	node := &NodeManagementService{s}
	s.registerHandler(id.AddNodesRequest_Encoding_DefaultBinary, node.AddNodesRequest)
	s.registerHandler(id.AddReferencesRequest_Encoding_DefaultBinary, node.AddReferencesRequest)
	s.registerHandler(id.DeleteNodesRequest_Encoding_DefaultBinary, node.DeleteNodesRequest)
	s.registerHandler(id.DeleteReferencesRequest_Encoding_DefaultBinary, node.DeleteReferencesRequest)

	view := &ViewService{s}
	s.registerHandler(id.BrowseRequest_Encoding_DefaultBinary, view.BrowseRequest)
	s.registerHandler(id.BrowseNextRequest_Encoding_DefaultBinary, view.BrowseNextRequest)
	s.registerHandler(id.TranslateBrowsePathsToNodeIDsRequest_Encoding_DefaultBinary, view.TranslateBrowsePathsToNodeIDsRequest)
	s.registerHandler(id.RegisterNodesRequest_Encoding_DefaultBinary, view.RegisterNodesRequest)
	s.registerHandler(id.UnregisterNodesRequest_Encoding_DefaultBinary, view.UnregisterNodesRequest)

	query := &QueryService{s}
	s.registerHandler(id.QueryFirstRequest_Encoding_DefaultBinary, query.QueryFirstRequest)
	s.registerHandler(id.QueryNextRequest_Encoding_DefaultBinary, query.QueryNextRequest)

	attr := &AttributeService{s}
	s.registerHandler(id.ReadRequest_Encoding_DefaultBinary, attr.ReadRequest)
	s.registerHandler(id.HistoryReadRequest_Encoding_DefaultBinary, attr.HistoryReadRequest)
	s.registerHandler(id.WriteRequest_Encoding_DefaultBinary, attr.WriteRequest)
	s.registerHandler(id.HistoryUpdateRequest_Encoding_DefaultBinary, attr.HistoryUpdateRequest)

	method := &MethodService{s}
	// s.registerHandler(id.CallMethodRequest_Encoding_DefaultBinary, method.CallMethodRequest) // todo(fs): I think this is bogus
	s.registerHandler(id.CallRequest_Encoding_DefaultBinary, method.CallRequest)

	item := &MonitoredItemService{s}
	// s.registerHandler(id.MonitoredItemCreateRequest_Encoding_DefaultBinary, item.MonitoredItemCreateRequest)
	s.registerHandler(id.CreateMonitoredItemsRequest_Encoding_DefaultBinary, item.CreateMonitoredItemsRequest)
	// s.registerHandler(id.MonitoredItemModifyRequest_Encoding_DefaultBinary, item.MonitoredItemModifyRequest)
	s.registerHandler(id.ModifyMonitoredItemsRequest_Encoding_DefaultBinary, item.ModifyMonitoredItemsRequest)
	s.registerHandler(id.SetMonitoringModeRequest_Encoding_DefaultBinary, item.SetMonitoringModeRequest)
	s.registerHandler(id.SetTriggeringRequest_Encoding_DefaultBinary, item.SetTriggeringRequest)
	s.registerHandler(id.DeleteMonitoredItemsRequest_Encoding_DefaultBinary, item.DeleteMonitoredItemsRequest)

	sub := &SubscriptionService{s}
	s.registerHandler(id.CreateSubscriptionRequest_Encoding_DefaultBinary, sub.CreateSubscriptionRequest)
	s.registerHandler(id.ModifySubscriptionRequest_Encoding_DefaultBinary, sub.ModifySubscriptionRequest)
	s.registerHandler(id.SetPublishingModeRequest_Encoding_DefaultBinary, sub.SetPublishingModeRequest)
	s.registerHandler(id.PublishRequest_Encoding_DefaultBinary, sub.PublishRequest)
	s.registerHandler(id.RepublishRequest_Encoding_DefaultBinary, sub.RepublishRequest)
	s.registerHandler(id.TransferSubscriptionsRequest_Encoding_DefaultBinary, sub.TransferSubscriptionsRequest)
	s.registerHandler(id.DeleteSubscriptionsRequest_Encoding_DefaultBinary, sub.DeleteSubscriptionsRequest)
}

func (s *Server) registerHandler(typeID uint16, h Handler) {
	s.handlers[typeID] = h
}

func (s *Server) handleService(ctx context.Context, sc *uasc.SecureChannel, reqID uint32, req ua.Request) {
	debug.Printf("handleService: Got: %T\n", req)

	var resp ua.Response
	var err error

	typeID := ua.ServiceTypeID(req)
	h, ok := s.handlers[typeID]
	if ok {
		resp, err = h(sc, req)
	} else {
		if typeID == 0 {
			debug.Printf("unknown service %T. Did you call register?", req)
		}
		err = ua.StatusBadServiceUnsupported
	}

	if err != nil {
		if statusCode, ok := err.(ua.StatusCode); ok {
			resp = &ua.ServiceFault{ResponseHeader: responseHeader(0, statusCode)}
		} else {
			resp = &ua.ServiceFault{ResponseHeader: responseHeader(0, ua.StatusBadUnexpectedError)}
		}
	}

	if resp == nil {
		return
	}

	err = sc.SendResponseWithContext(ctx, reqID, resp)
	if err != nil {
		debug.Printf("Error sending response: %s\n", err)
	}
}

func responseHeader(reqID uint32, statusCode ua.StatusCode) *ua.ResponseHeader {
	return &ua.ResponseHeader{
		Timestamp:          time.Now(),
		RequestHandle:      reqID,
		ServiceResult:      statusCode,
		ServiceDiagnostics: &ua.DiagnosticInfo{},
		StringTable:        []string{},
		AdditionalHeader:   ua.NewExtensionObject(nil),
	}
}

// func handleServiceFault(s *Server, sc *uasc.SecureChannel, r ua.Request) (ua.Response, error) {
// 	debug.Printf("Handling %T\n", r)

// 	req, ok := r.(*ua.ServiceFault)
// 	if !ok {
// 		debug.Printf("handleServiceFault: Expected *ua.ServiceFault, got %T", r)
// 		return nil, ua.StatusBadRequestTypeInvalid
// 	}
// 	debug.Printf("Got ServiceFault: %s", req.ResponseHeader.ServiceResult)

// 	// No response required
// 	return nil, nil
// }
