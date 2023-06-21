package server

import (
	"time"

	"github.com/gopcua/opcua/id"
	"github.com/gopcua/opcua/server/attrs"
	"github.com/gopcua/opcua/ua"
)

func CurrentTimeNode() *Node {
	return NewNode(
		ua.NewNumericNodeID(0, 2258),
		map[ua.AttributeID]*ua.Variant{
			ua.AttributeIDBrowseName: ua.MustVariant(attrs.BrowseName("CurrentTime")),
		},
		nil,
		func() *ua.Variant { return ua.MustVariant(time.Now()) },
	)
}

func NamespacesNode(s *Server) *Node {
	return NewNode(
		ua.NewNumericNodeID(0, 2255),
		map[ua.AttributeID]*ua.Variant{
			ua.AttributeIDBrowseName: ua.MustVariant(attrs.BrowseName("Namespaces")),
		},
		nil,
		func() *ua.Variant { return ua.MustVariant(s.Namespaces()) },
	)
}

func ServerStatusNode(s *Server) *Node {
	return NewNode(
		ua.NewNumericNodeID(0, 2256),
		map[ua.AttributeID]*ua.Variant{
			ua.AttributeIDBrowseName: ua.MustVariant(attrs.BrowseName("ServerStatus")),
		},
		nil,
		func() *ua.Variant { return ua.MustVariant(ua.NewExtensionObject(s.Status())) },
	)
}

func ServerCapabilitiesNodes(s *Server) []*Node {
	var nodes []*Node
	nodes = append(nodes, NewNode(
		ua.NewNumericNodeID(0, id.Server_ServerCapabilities_OperationLimits_MaxNodesPerRead),
		map[ua.AttributeID]*ua.Variant{
			ua.AttributeIDBrowseName: ua.MustVariant(attrs.BrowseName("MaxNodesPerRead")),
		},
		nil,
		func() *ua.Variant { return ua.MustVariant(s.cfg.cap.OperationalLimits.MaxNodesPerRead) },
	))
	return nodes
}

func RootNode() *Node {
	return NewNode(
		ua.NewNumericNodeID(0, id.RootFolder),
		map[ua.AttributeID]*ua.Variant{
			ua.AttributeIDNodeClass:  ua.MustVariant(attrs.NodeClass(ua.NodeClassObject)),
			ua.AttributeIDBrowseName: ua.MustVariant(attrs.BrowseName("Root")),
			ua.AttributeIDDataType:   ua.MustVariant(ua.NewNumericExpandedNodeID(0, id.DataTypesFolder)),
		},
		nil,
		nil,
	)
}
