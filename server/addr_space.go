// Copyright 2018-2020 opcua authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package server

import (
	"fmt"
	"sync"
	"time"

	"github.com/gopcua/opcua/ua"
)

type Node interface {
	ID() *ua.NodeID
	Attribute(ua.AttributeID) (*AttrValue, error)
}

type AttrValue struct {
	Value           *ua.Variant
	SourceTimestamp time.Time
}

func (v *AttrValue) String() string {
	return fmt.Sprintf("v=%#v ts=%s", v.Value.Value(), v.SourceTimestamp)
}

func NewAttrValue(v *ua.Variant) *AttrValue {
	return &AttrValue{Value: v}
}

type AddressSpace interface {
	Attribute(id *ua.NodeID, attr ua.AttributeID) (*AttrValue, error)
}

// MergeAS merges multiple address spaces in ascending order.
type MergeAS struct {
	as []AddressSpace
}

func NewMergeAS(as ...AddressSpace) *MergeAS {
	return &MergeAS{as}
}

func (s *MergeAS) Attribute(id *ua.NodeID, attr ua.AttributeID) (*AttrValue, error) {
	for _, as := range s.as {
		v, err := as.Attribute(id, attr)
		if err == ua.StatusBadNodeIDUnknown {
			continue
		}
		return v, err
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
