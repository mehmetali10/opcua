// Copyright 2018-2020 opcua authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package server

import (
	"sync"
	"time"

	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"

	"github.com/gopcua/opcua/debug"
	"github.com/gopcua/opcua/id"
	"github.com/gopcua/opcua/ua"
)

var (
	ObjectsFolder = ua.NewNumericNodeID(0, id.ObjectsFolder)
	RootFolder    = ua.NewNumericNodeID(0, id.RootFolder)
)

type NameSpace interface {
	Name() string
	AddNode(n *Node) *Node
	Node(id *ua.NodeID) *Node
	Objects() *Node
	Root() *Node
	Browse(req *ua.BrowseDescription) *ua.BrowseResult
	ID() uint16
	SetID(uint16)
	Attribute(*ua.NodeID, ua.AttributeID) *ua.DataValue
}

// the base "node-centric" namespace
type NodeNameSpace struct {
	name  string
	mu    sync.RWMutex
	nodes []*Node
	m     map[string]*Node
	id    uint16
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
	return as.Node(ObjectsFolder)
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
