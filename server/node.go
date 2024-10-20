package server

import (
	"time"

	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"

	"github.com/gopcua/opcua/id"
	"github.com/gopcua/opcua/server/attrs"
	"github.com/gopcua/opcua/server/refs"
	"github.com/gopcua/opcua/ua"
)

type Attributes map[ua.AttributeID]*ua.Variant

type References []*ua.ReferenceDescription

type ValueFunc func() *ua.Variant

type AttrValue struct {
	Value           *ua.Variant
	SourceTimestamp time.Time
}

func NewAttrValue(v *ua.Variant) *AttrValue {
	return &AttrValue{Value: v}
}

type Node struct {
	id   *ua.NodeID
	attr Attributes
	refs References
	val  ValueFunc

	ns NameSpace
}

func NewNode(id *ua.NodeID, attr Attributes, refs References, val ValueFunc) *Node {
	n := &Node{id, attr, refs, val, nil}
	n.sanitize()
	return n
}

func NewFolderNode(nodeID *ua.NodeID, name string) *Node {
	reftype := ua.NewTwoByteNodeID(uint8(id.HasComponent)) // folder
	eoid := ua.NewNumericExpandedNodeID(nodeID.Namespace(), id.ObjectsFolder)
	typedef := ua.NewNumericExpandedNodeID(0, id.ObjectsFolder)
	n := NewNode(
		nodeID,
		map[ua.AttributeID]*ua.Variant{
			ua.AttributeIDNodeClass:     ua.MustVariant(uint32(ua.NodeClassObject)),
			ua.AttributeIDBrowseName:    ua.MustVariant(attrs.BrowseName(name)),
			ua.AttributeIDDisplayName:   ua.MustVariant(attrs.DisplayName(name, name)),
			ua.AttributeIDDescription:   ua.MustVariant(ua.NodeClassObject),
			ua.AttributeIDEventNotifier: ua.MustVariant(int16(0)),
		},
		[]*ua.ReferenceDescription{{
			ReferenceTypeID: reftype,
			IsForward:       true,
			NodeID:          eoid,
			BrowseName:      &ua.QualifiedName{NamespaceIndex: nodeID.Namespace(), Name: name},
			DisplayName:     &ua.LocalizedText{EncodingMask: ua.LocalizedTextText, Text: name},
			NodeClass:       ua.NodeClassObject,
			TypeDefinition:  typedef,
		}},
		nil,
	)
	return n
}

func NewVariableNode(nodeID *ua.NodeID, name string, value any) *Node {
	//eoid := ua.NewNumericExpandedNodeID(nodeID.Namespace(), nodeID.IntID())
	typedef := ua.NewNumericExpandedNodeID(0, id.VariableNode)
	n := NewNode(
		nodeID,
		map[ua.AttributeID]*ua.Variant{
			ua.AttributeIDNodeClass:     ua.MustVariant(uint32(ua.NodeClassVariable)),
			ua.AttributeIDBrowseName:    ua.MustVariant(attrs.BrowseName(name)),
			ua.AttributeIDDisplayName:   ua.MustVariant(attrs.DisplayName(name, name)),
			ua.AttributeIDDescription:   ua.MustVariant(uint32(ua.NodeClassVariable)),
			ua.AttributeIDDataType:      ua.MustVariant(typedef),
			ua.AttributeIDEventNotifier: ua.MustVariant(int16(0)),
		},
		[]*ua.ReferenceDescription{},
		func() *ua.Variant {
			return ua.MustVariant(value)
		},
	)
	return n
}

func (n *Node) sanitize() {
	if n.attr == nil {
		n.attr = Attributes{}
	}
	if n.attr[ua.AttributeIDBrowseName] == nil {
		n.SetBrowseName("")
	}
	if n.attr[ua.AttributeIDDisplayName] == nil {
		n.SetDisplayName("", "")
	}
	if n.DisplayName().Text == "" {
		n.SetDisplayName(n.BrowseName().Name, "")
	}
	if n.attr[ua.AttributeIDDescription] == nil {
		n.SetDescription("", "")
	}
	if n.attr[ua.AttributeIDDataType] == nil {
		n.attr[ua.AttributeIDDataType] = ua.MustVariant(ua.NewTwoByteExpandedNodeID(0))
	}
}

func (n *Node) ID() *ua.NodeID {
	return n.id
}

func (n *Node) Value() *ua.Variant {
	if n.val == nil {
		return nil
	}
	return n.val()
}

func (n *Node) Attribute(id ua.AttributeID) (*AttrValue, error) {
	switch {
	case id == ua.AttributeIDValue:
		if n.val != nil {
			return NewAttrValue(n.val()), nil
		}
		return nil, ua.StatusBadAttributeIDInvalid
	case n.attr == nil:
		return nil, ua.StatusBadAttributeIDInvalid
	default:
		if v := n.attr[id]; v != nil {
			return NewAttrValue(v), nil
		}
		return nil, ua.StatusBadAttributeIDInvalid
	}
}
func (n *Node) SetAttribute(id ua.AttributeID, val ua.DataValue) error {
	switch {
	case id == ua.AttributeIDValue:

		// TODO: probably need to do some type checking here.
		// And some permissions tests
		n.val = func() *ua.Variant {
			return ua.MustVariant(val.Value.Value())
		}

		return nil
	}
	return ua.StatusBadNodeAttributesInvalid
}

func (n *Node) BrowseName() *ua.QualifiedName {
	v := n.attr[ua.AttributeIDBrowseName]
	if v == nil || v.Value() == nil {
		return &ua.QualifiedName{}
	}
	return v.Value().(*ua.QualifiedName)
}

func (n *Node) SetBrowseName(s string) {
	n.attr[ua.AttributeIDBrowseName] = ua.MustVariant(&ua.QualifiedName{Name: s})
}

func (n *Node) DisplayName() *ua.LocalizedText {
	v := n.attr[ua.AttributeIDDisplayName]
	if v == nil || v.Value() == nil {
		return &ua.LocalizedText{}
	}
	val := v.Value().(*ua.LocalizedText)
	val.UpdateMask()
	return val
}

func (n *Node) SetDisplayName(text, locale string) {
	n.attr[ua.AttributeIDDisplayName] = ua.MustVariant(&ua.LocalizedText{Text: text, Locale: locale})
}

func (n *Node) Description() *ua.LocalizedText {
	v := n.attr[ua.AttributeIDDescription]
	if v == nil || v.Value() == nil {
		return &ua.LocalizedText{}
	}
	return v.Value().(*ua.LocalizedText)
}

func (n *Node) SetDescription(text, locale string) {
	n.attr[ua.AttributeIDDescription] = ua.MustVariant(&ua.LocalizedText{Text: text, Locale: locale})
}

func (n *Node) DataType() *ua.ExpandedNodeID {
	v := n.attr[ua.AttributeIDDataType]
	if v == nil || v.Value() == nil {
		return ua.NewTwoByteExpandedNodeID(0)
	}
	return v.Value().(*ua.ExpandedNodeID)
}

func (n *Node) SetNodeClass(nc ua.NodeClass) {
	n.attr[ua.AttributeIDNodeClass] = ua.MustVariant(uint32(nc))
}

func (n *Node) NodeClass() ua.NodeClass {
	v := n.attr[ua.AttributeIDDescription]
	if v == nil || v.Value() == nil {
		return ua.NodeClassUnspecified
	}
	return ua.NodeClass(v.Value().(uint32))

}

func (n *Node) AddObject(o *Node) *Node {
	nn := &Node{
		id:   o.id,
		attr: maps.Clone(o.attr),
		refs: slices.Clone(o.refs),
	}
	if n.attr == nil {
		n.attr = Attributes{}
	}
	nn.SetNodeClass(ua.NodeClassObject)
	n.refs = append(n.refs, refs.Organizes(nn.id, nn.BrowseName().Name, nn.DisplayName().Text, nn.DataType()))
	return n.ns.AddNode(nn)
}

func (n *Node) AddVariable(o *Node) *Node {
	nn := &Node{
		id:   o.id,
		attr: maps.Clone(o.attr),
		refs: slices.Clone(o.refs),
		val:  o.val,
	}
	if n.attr == nil {
		n.attr = Attributes{}
	}
	nn.SetNodeClass(ua.NodeClassVariable)
	n.refs = append(n.refs, refs.Organizes(nn.id, nn.BrowseName().Name, nn.DisplayName().Text, nn.DataType()))
	return nn
}

func (n *Node) AddRef(o *Node) {
	//eoid := ua.NewNumericExpandedNodeID(o.ns.ID(), o.)
	eoid := ua.NewExpandedNodeID(o.ID(), "", 0)

	ref := ua.ReferenceDescription{
		ReferenceTypeID: &ua.NodeID{}, //o.refs[0].ReferenceTypeID,
		IsForward:       true,
		NodeID:          eoid,
		BrowseName:      o.BrowseName(),
		DisplayName:     o.DisplayName(),
		NodeClass:       o.NodeClass(),
		TypeDefinition:  o.DataType(),
	}
	n.refs = append(n.refs, &ref)
}
