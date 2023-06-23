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
	s.registerHandler(id.FindServersRequest_Encoding_DefaultBinary, discovery.FindServers)
	s.registerHandler(id.FindServersOnNetworkRequest_Encoding_DefaultBinary, discovery.FindServersOnNetwork)
	s.registerHandler(id.GetEndpointsRequest_Encoding_DefaultBinary, discovery.GetEndpoints)
	s.registerHandler(id.RegisterServerRequest_Encoding_DefaultBinary, discovery.RegisterServer)
	s.registerHandler(id.RegisterServer2Request_Encoding_DefaultBinary, discovery.RegisterServer2)

	// SecureChannel service (handled in the uasc stack)
	// s.registerHandlerFunc(id.OpenSecureChannelRequest_Encoding_DefaultBinary, handleOpenSecureChannel)
	// s.registerHandlerFunc(id.CloseSecureChannelRequest_Encoding_DefaultBinary, handleCloseSecureChannel)

	session := &SessionService{s}
	s.registerHandler(id.CreateSessionRequest_Encoding_DefaultBinary, session.CreateSession)
	s.registerHandler(id.ActivateSessionRequest_Encoding_DefaultBinary, session.ActivateSession)
	s.registerHandler(id.CloseSessionRequest_Encoding_DefaultBinary, session.CloseSession)
	s.registerHandler(id.CancelRequest_Encoding_DefaultBinary, session.Cancel)

	node := &NodeManagementService{s}
	s.registerHandler(id.AddNodesRequest_Encoding_DefaultBinary, node.AddNodes)
	s.registerHandler(id.AddReferencesRequest_Encoding_DefaultBinary, node.AddReferences)
	s.registerHandler(id.DeleteNodesRequest_Encoding_DefaultBinary, node.DeleteNodes)
	s.registerHandler(id.DeleteReferencesRequest_Encoding_DefaultBinary, node.DeleteReferences)

	view := &ViewService{s}
	s.registerHandler(id.BrowseRequest_Encoding_DefaultBinary, view.Browse)
	s.registerHandler(id.BrowseNextRequest_Encoding_DefaultBinary, view.BrowseNext)
	s.registerHandler(id.TranslateBrowsePathsToNodeIDsRequest_Encoding_DefaultBinary, view.TranslateBrowsePathsToNodeIDs)
	s.registerHandler(id.RegisterNodesRequest_Encoding_DefaultBinary, view.RegisterNodes)
	s.registerHandler(id.UnregisterNodesRequest_Encoding_DefaultBinary, view.UnregisterNodes)

	query := &QueryService{s}
	s.registerHandler(id.QueryFirstRequest_Encoding_DefaultBinary, query.QueryFirst)
	s.registerHandler(id.QueryNextRequest_Encoding_DefaultBinary, query.QueryNext)

	attr := &AttributeService{s}
	s.registerHandler(id.ReadRequest_Encoding_DefaultBinary, attr.Read)
	s.registerHandler(id.HistoryReadRequest_Encoding_DefaultBinary, attr.HistoryRead)
	s.registerHandler(id.WriteRequest_Encoding_DefaultBinary, attr.Write)
	s.registerHandler(id.HistoryUpdateRequest_Encoding_DefaultBinary, attr.HistoryUpdate)

	method := &MethodService{s}
	// s.registerHandler(id.CallMethodRequest_Encoding_DefaultBinary, method.CallMethod) // todo(fs): I think this is bogus
	s.registerHandler(id.CallRequest_Encoding_DefaultBinary, method.Call)

	item := &MonitoredItemService{s}
	// s.registerHandler(id.MonitoredItemCreateRequest_Encoding_DefaultBinary, item.MonitoredItemCreate)
	s.registerHandler(id.CreateMonitoredItemsRequest_Encoding_DefaultBinary, item.CreateMonitoredItems)
	// s.registerHandler(id.MonitoredItemModifyRequest_Encoding_DefaultBinary, item.MonitoredItemModify)
	s.registerHandler(id.ModifyMonitoredItemsRequest_Encoding_DefaultBinary, item.ModifyMonitoredItems)
	s.registerHandler(id.SetMonitoringModeRequest_Encoding_DefaultBinary, item.SetMonitoringMode)
	s.registerHandler(id.SetTriggeringRequest_Encoding_DefaultBinary, item.SetTriggering)
	s.registerHandler(id.DeleteMonitoredItemsRequest_Encoding_DefaultBinary, item.DeleteMonitoredItems)

	sub := &SubscriptionService{s}
	s.registerHandler(id.CreateSubscriptionRequest_Encoding_DefaultBinary, sub.CreateSubscription)
	s.registerHandler(id.ModifySubscriptionRequest_Encoding_DefaultBinary, sub.ModifySubscription)
	s.registerHandler(id.SetPublishingModeRequest_Encoding_DefaultBinary, sub.SetPublishingMode)
	s.registerHandler(id.PublishRequest_Encoding_DefaultBinary, sub.Publish)
	s.registerHandler(id.RepublishRequest_Encoding_DefaultBinary, sub.Republish)
	s.registerHandler(id.TransferSubscriptionsRequest_Encoding_DefaultBinary, sub.TransferSubscriptions)
	s.registerHandler(id.DeleteSubscriptionsRequest_Encoding_DefaultBinary, sub.DeleteSubscriptions)
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
