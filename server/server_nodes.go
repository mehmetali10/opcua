package server

import (
	"time"

	"github.com/gopcua/opcua/ua"
)

type ValueNode struct {
	NodeID     *ua.NodeID
	BrowseName string
	Value      func() interface{}
}

func (n *ValueNode) ID() *ua.NodeID {
	return n.NodeID
}

func (n *ValueNode) Attribute(attr ua.AttributeID) (*AttrValue, error) {
	switch attr {
	case ua.AttributeIDBrowseName:
		return NewAttrValue(ua.MustVariant(n.BrowseName)), nil
	case ua.AttributeIDValue:
		v, err := ua.NewVariant(n.Value())
		if err != nil {
			return NewAttrValue(ua.MustVariant(nil)), nil
		}
		return NewAttrValue(v), nil
	default:
		return nil, ua.StatusBadAttributeIDInvalid
	}
}

func CurrentTimeNode() *ValueNode {
	return &ValueNode{
		NodeID:     ua.NewNumericNodeID(0, 2258),
		BrowseName: "CurrentTime",
		Value: func() interface{} {
			return time.Now()
		},
	}
}

func NamespacesNode(s *Server) *ValueNode {
	return &ValueNode{
		NodeID:     ua.NewNumericNodeID(0, 2255),
		BrowseName: "Namespaces",
		Value: func() interface{} {
			return s.Namespaces()
		},
	}
}

func ServerStatusNode(s *Server) *ValueNode {
	return &ValueNode{
		NodeID:     ua.NewNumericNodeID(0, 2256),
		BrowseName: "ServerStatus",
		Value: func() interface{} {
			return ua.NewExtensionObject(s.Status())
		},
	}
}
