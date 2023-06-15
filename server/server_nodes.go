package server

import (
	"time"

	"github.com/gopcua/opcua/id"
	"github.com/gopcua/opcua/ua"
)

func CurrentTimeNode() *node {
	return &node{
		id:    ua.NewNumericNodeID(0, 2258),
		value: func() *ua.Variant { return ua.MustVariant(time.Now()) },
		attr: map[ua.AttributeID]*ua.Variant{
			ua.AttributeIDBrowseName: ua.MustVariant(&ua.QualifiedName{Name: "CurrentTime"}),
		},
	}
}

func NamespacesNode(s *Server) *node {
	return &node{
		id:    ua.NewNumericNodeID(0, 2255),
		value: func() *ua.Variant { return ua.MustVariant(s.Namespaces()) },
		attr: map[ua.AttributeID]*ua.Variant{
			ua.AttributeIDBrowseName: ua.MustVariant(&ua.QualifiedName{Name: "Namespaces"}),
		},
	}
}

func ServerStatusNode(s *Server) *node {
	return &node{
		id:    ua.NewNumericNodeID(0, 2256),
		value: func() *ua.Variant { return ua.MustVariant(ua.NewExtensionObject(s.Status())) },
		attr: map[ua.AttributeID]*ua.Variant{
			ua.AttributeIDBrowseName: ua.MustVariant(&ua.QualifiedName{Name: "ServerStatus"}),
		},
	}
}

func RootNode(s *Server) *node {
	return &node{
		id: ua.NewNumericNodeID(0, id.RootFolder),
		attr: map[ua.AttributeID]*ua.Variant{
			ua.AttributeIDNodeClass:  ua.MustVariant(ua.NodeClassObject),
			ua.AttributeIDDataType:   ua.MustVariant(ua.NewNumericExpandedNodeID(0, id.DataTypesFolder)),
			ua.AttributeIDBrowseName: ua.MustVariant(&ua.QualifiedName{Name: "Root"}),
		},
	}
}
