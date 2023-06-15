// Copyright 2018-2020 opcua authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package server

import (
	"sync"
	"time"

	"github.com/gopcua/opcua/ua"
)

type Node interface {
	ID() *ua.NodeID
	Value() *ua.Variant
	Attribute(ua.AttributeID) (*AttrValue, error)
}

type node struct {
	id    *ua.NodeID
	value func() *ua.Variant
	attr  map[ua.AttributeID]*ua.Variant

	superTypeID *ua.NodeID
}

func NewNode(id *ua.NodeID, v func() *ua.Variant, attr map[ua.AttributeID]*ua.Variant) *node {
	return &node{id, v, attr, ua.NewTwoByteNodeID(0)}
}

func (n *node) ID() *ua.NodeID {
	return n.id
}

func (n *node) Value() *ua.Variant {
	return n.value()
}

func (n *node) Attribute(id ua.AttributeID) (*AttrValue, error) {
	if n.attr == nil {
		return nil, ua.StatusBadAttributeIDInvalid
	}
	if v := n.attr[id]; v != nil {
		return NewAttrValue(v), nil
	}
	return nil, ua.StatusBadAttributeIDInvalid
}

type AttrValue struct {
	Value           *ua.Variant
	SourceTimestamp time.Time
}

func NewAttrValue(v *ua.Variant) *AttrValue {
	return &AttrValue{Value: v}
}

type AddressSpace interface {
	Attribute(*ua.NodeID, ua.AttributeID) (*AttrValue, error)
	Node(*ua.NodeID) (Node, error)
}

// MergeAS merges multiple address spaces in ascending order.
type MergeAS struct {
	as []AddressSpace
}

func NewMergeAS(as ...AddressSpace) *MergeAS {
	return &MergeAS{as}
}

func (s *MergeAS) Attribute(id *ua.NodeID, attr ua.AttributeID) (*AttrValue, error) {
	n, err := s.Node(id)
	if err != nil {
		return nil, err
	}
	return n.Attribute(attr)
}

func (s *MergeAS) Node(id *ua.NodeID) (Node, error) {
	for _, as := range s.as {
		n, err := as.Node(id)
		if err == ua.StatusBadNodeIDUnknown {
			continue
		}
		return n, err
	}
	return nil, ua.StatusBadNodeIDUnknown
}

// SyncAS implements a mutable address space which is safe for concurrent access.
type SyncAS struct {
	mu    sync.Mutex
	nodes map[string]Node
}

func NewSyncAS() *SyncAS {
	return &SyncAS{nodes: make(map[string]Node)}
}

func (a *SyncAS) AddNodes(nodes ...Node) error {
	a.mu.Lock()
	defer a.mu.Unlock()
	for _, n := range nodes {
		a.nodes[n.ID().String()] = n
	}
	return nil
}

func (a *SyncAS) Attribute(id *ua.NodeID, attr ua.AttributeID) (*AttrValue, error) {
	a.mu.Lock()
	defer a.mu.Unlock()
	n := a.nodes[id.String()]
	if n == nil {
		return nil, ua.StatusBadNodeIDUnknown
	}
	return n.Attribute(attr)
}

func (a *SyncAS) Node(id *ua.NodeID) (Node, error) {
	a.mu.Lock()
	defer a.mu.Unlock()
	n := a.nodes[id.String()]
	if n == nil {
		return nil, ua.StatusBadNodeIDUnknown
	}
	return n, nil
}

// interface compliance
var _ Node = (*node)(nil)
var _ AddressSpace = (*MergeAS)(nil)
var _ AddressSpace = (*SyncAS)(nil)
