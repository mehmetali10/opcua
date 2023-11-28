// Copyright 2018-2020 opcua authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package server

import (
	"sync"

	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"

	"github.com/gopcua/opcua/id"
	"github.com/gopcua/opcua/ua"
)

var (
	ObjectsFolder = ua.NewNumericNodeID(0, id.ObjectsFolder)
	RootFolder    = ua.NewNumericNodeID(0, id.RootFolder)
)

type AddressSpace struct {
	mu    sync.Mutex
	nodes []*Node
	m     map[string]*Node
}

func NewAddressSpace() *AddressSpace {
	return &AddressSpace{m: map[string]*Node{}}
}

func (as *AddressSpace) AddNode(n *Node) *Node {
	as.mu.Lock()
	defer as.mu.Unlock()

	nn := &Node{
		id:   n.id,
		attr: maps.Clone(n.attr),
		refs: slices.Clone(n.refs),
		val:  n.val,
		as:   as,
	}

	// todo(fs): this is wrong since this leaves the old node in the list.
	as.nodes = append(as.nodes, nn)
	k := nn.ID().String()
	as.m[k] = nn
	return nn
}

func (as *AddressSpace) Attribute(id *ua.NodeID, attr ua.AttributeID) (*AttrValue, error) {
	n := as.Node(id)
	if n == nil {
		return nil, ua.StatusBadNodeIDUnknown
	}
	return n.Attribute(attr)
}

func (as *AddressSpace) Node(id *ua.NodeID) *Node {
	as.mu.Lock()
	defer as.mu.Unlock()
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

func (as *AddressSpace) Objects() *Node {
	return as.Node(ObjectsFolder)
}

func (as *AddressSpace) Root() *Node {
	return as.Node(RootFolder)
}
