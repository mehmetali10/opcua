// Copyright 2018-2020 opcua authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package server

import (
	"sync"
	"sync/atomic"
	"time"

	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"

	"github.com/gopcua/opcua/debug"
	"github.com/gopcua/opcua/id"
	"github.com/gopcua/opcua/server/attrs"
	"github.com/gopcua/opcua/ua"
)

var (
	ObjectsFolder = ua.NewNumericNodeID(0, id.ObjectsFolder)
	RootFolder    = ua.NewNumericNodeID(0, id.RootFolder)
)

// These are all the functions a namespace needs in order to provide nodes into the server
type NameSpace interface {
	// Name of the namespace.  Per the standard it should be an URI.
	Name() string

	// This function should create a new node
	AddNode(n *Node) *Node

	// This function should lookup and return the node indicated by the Node ID
	Node(id *ua.NodeID) *Node

	// This function should return the base Objects node that contains other nodes
	Objects() *Node

	// This function should return the root node
	Root() *Node

	// This is the function to list all available nodes to the client that is browsing.
	// The BrowseDescription has the root node of the browse and what kind of nodes the
	// client is looking for.  The Browse Result should have the list of matching nodes.
	Browse(req *ua.BrowseDescription) *ua.BrowseResult

	// ID and SetID are the namespace ID number of this namespace.  When you add it to the server
	// with srv.AddNamespace(xxx) it will set these for you.
	ID() uint16
	SetID(uint16)

	// These are the functions for reading and writing arbitrary attributes.  The most common
	// is the value attribute, but many clients also read the datatype and description attributes.
	// as well as attributes related to array bounds
	Attribute(*ua.NodeID, ua.AttributeID) *ua.DataValue
	SetAttribute(*ua.NodeID, ua.AttributeID, *ua.DataValue) ua.StatusCode
}

// the base "node-centric" namespace
type NodeNameSpace struct {
	name            string
	mu              sync.RWMutex
	nodes           []*Node
	m               map[string]*Node
	id              uint16
	nodeid_sequence uint32
}

func (ns *NodeNameSpace) GetNextNodeID() uint32 {
	if ns.nodeid_sequence < 100 {
		ns.nodeid_sequence = 100
	}
	return atomic.AddUint32(&(ns.nodeid_sequence), 1)
}

func NewNodeNameSpace(srv *Server, name string) *NodeNameSpace {
	ns := &NodeNameSpace{
		name:  name,
		nodes: make([]*Node, 0),
		m:     make(map[string]*Node),
	}
	srv.AddNamespace(ns, false, false)

	//objectsNode := NewFolderNode(ua.NewNumericNodeID(ns.id, id.ObjectsFolder), ns.name)
	oid := ua.NewNumericNodeID(ns.ID(), id.ObjectsFolder)
	//eoid := ua.NewNumericExpandedNodeID(ns.ID(), id.ObjectsFolder)
	//typedef := ua.NewNumericExpandedNodeID(0, id.ObjectsFolder)
	//reftype := ua.NewTwoByteNodeID(uint8(id.HasComponent)) // folder
	objectsNode := NewNode(
		oid,
		map[ua.AttributeID]*ua.Variant{
			ua.AttributeIDNodeClass:     ua.MustVariant(uint32(ua.NodeClassObject)),
			ua.AttributeIDBrowseName:    ua.MustVariant(attrs.BrowseName(ns.name)),
			ua.AttributeIDDisplayName:   ua.MustVariant(attrs.DisplayName(ns.name, ns.name)),
			ua.AttributeIDDescription:   ua.MustVariant(""),
			ua.AttributeIDEventNotifier: ua.MustVariant(int16(0)),
		},
		[]*ua.ReferenceDescription{},
		nil,
	)

	ns.AddNode(objectsNode)

	return ns

}

func (ns *NodeNameSpace) Name() string {
	return ns.name
}

func NewNameSpace(name string) *NodeNameSpace {
	return &NodeNameSpace{name: name, m: map[string]*Node{}}
}

func (as *NodeNameSpace) AddNode(n *Node) *Node {
	as.mu.Lock()
	defer as.mu.Unlock()

	nn := &Node{
		id:   n.id,
		attr: maps.Clone(n.attr),
		refs: slices.Clone(n.refs),
		val:  n.val,
		ns:   as,
	}

	// todo(fs): this is wrong since this leaves the old node in the list.
	as.nodes = append(as.nodes, nn)
	k := nn.ID().String()
	as.m[k] = nn
	return nn
}

func (as *NodeNameSpace) AddNewVariableNode(name string, value any) *Node {
	n := NewVariableNode(ua.NewNumericNodeID(as.id, as.GetNextNodeID()), name, value)
	as.AddNode(n)
	return n
}

func (as *NodeNameSpace) Attribute(id *ua.NodeID, attr ua.AttributeID) *ua.DataValue {
	n := as.Node(id)
	if n == nil {
		return &ua.DataValue{
			EncodingMask:    ua.DataValueServerTimestamp | ua.DataValueStatusCode,
			ServerTimestamp: time.Now(),
			Status:          ua.StatusBadNodeIDUnknown,
		}
	}

	a, err := n.Attribute(attr)
	if err != nil {
		return &ua.DataValue{
			EncodingMask:    ua.DataValueServerTimestamp | ua.DataValueStatusCode,
			ServerTimestamp: time.Now(),
			Status:          ua.StatusBadAttributeIDInvalid,
		}
	}
	return &ua.DataValue{
		EncodingMask:    ua.DataValueServerTimestamp | ua.DataValueStatusCode | ua.DataValueValue,
		ServerTimestamp: time.Now(),
		Status:          ua.StatusOK,
		Value:           a.Value,
	}
}

func (as *NodeNameSpace) Node(id *ua.NodeID) *Node {
	as.mu.RLock()
	defer as.mu.RUnlock()
	if id == nil {
		return nil
	}
	k := id.String()
	n := as.m[k]
	if n == nil {
		return nil
	}
	return n
}

func (as *NodeNameSpace) Objects() *Node {
	of := ua.NewNumericNodeID(as.id, id.ObjectsFolder)
	return as.Node(of)
}

func (as *NodeNameSpace) Root() *Node {
	return as.Node(RootFolder)
}

func (ns *NodeNameSpace) Browse(bd *ua.BrowseDescription) *ua.BrowseResult {
	ns.mu.RLock()
	defer ns.mu.RUnlock()

	debug.Printf("BrowseRequest: id=%s mask=%08b\n", bd.NodeID, bd.ResultMask)

	n := ns.Node(bd.NodeID)
	if n == nil {
		return &ua.BrowseResult{StatusCode: ua.StatusBadNodeIDUnknown}
	}

	refs := make([]*ua.ReferenceDescription, 0, len(n.refs))

	for i := range n.refs {
		r := n.refs[i]
		// we can't have nils in these or the encoder will fail.
		if r.NodeID == nil || r.BrowseName == nil || r.DisplayName == nil || r.TypeDefinition == nil {
			continue
		}
		rf := &ua.ReferenceDescription{
			ReferenceTypeID: r.ReferenceTypeID,
			IsForward:       r.IsForward,
			NodeID:          r.NodeID,
			BrowseName:      r.BrowseName,
			DisplayName:     r.DisplayName,
			NodeClass:       r.NodeClass,
			TypeDefinition:  r.TypeDefinition,
		}

		refs = append(refs, rf)
	}

	return &ua.BrowseResult{
		StatusCode: ua.StatusGood,
		References: refs,
	}

}

func (ns *NodeNameSpace) ID() uint16 {
	return ns.id
}

func (ns *NodeNameSpace) SetID(id uint16) {
	ns.id = id
}
func (as *NodeNameSpace) SetAttribute(id *ua.NodeID, attr ua.AttributeID, val *ua.DataValue) ua.StatusCode {
	n := as.Node(id)
	if n == nil {
		return ua.StatusBadNodeIDUnknown
	}

	err := n.SetAttribute(attr, *val)
	if err != nil {
		return ua.StatusBadAttributeIDInvalid
	}

	return ua.StatusOK
}
