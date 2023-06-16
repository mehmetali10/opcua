// Generated code. DO NOT EDIT
// Copyright 2018-2023 opcua authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.
package server

import (
	"github.com/gopcua/opcua/id"
	"github.com/gopcua/opcua/ua"
)

func PredefinedNodes() []Node {
	return []Node{
		&node{
			id: ua.NewNumericNodeID(0, 3062),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default Binary"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default Binary"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 3063),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default XML"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default XML"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 24),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "BaseDataType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "BaseDataType"}),
			},
			refs: []*ua.ReferenceDescription{
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 26),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 17),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 19),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 18),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 29),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 22),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 16),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 1),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 23),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 12),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 13),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 15),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 20),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 25),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 14),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 21),
					IsForward:       true,
				},
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 26),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Number"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Number"}),
			},
			refs: []*ua.ReferenceDescription{
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 50),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 27),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 11),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 10),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 28),
					IsForward:       true,
				},
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 27),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Integer"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Integer"}),
			},
			refs: []*ua.ReferenceDescription{
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 6),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 8),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 4),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 2),
					IsForward:       true,
				},
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 28),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "UInteger"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "UInteger"}),
			},
			refs: []*ua.ReferenceDescription{
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 9),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 3),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 7),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 5),
					IsForward:       true,
				},
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 29),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Enumeration"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Enumeration"}),
			},
			refs: []*ua.ReferenceDescription{
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 24216),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 12552),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 11939),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 14647),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 12077),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 19730),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 24222),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 15874),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 302),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 15008),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 98),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 24224),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 852),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 120),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 24210),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 307),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 24218),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 15632),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 24212),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 11234),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 24220),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 257),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 24214),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 576),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 11293),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 19723),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 851),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 315),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 348),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 256),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 20408),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 303),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 890),
					IsForward:       true,
				},
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 1),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Boolean"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Boolean"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 2),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "SByte"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "SByte"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 3),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Byte"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Byte"}),
			},
			refs: []*ua.ReferenceDescription{
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 15033),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 15031),
					IsForward:       true,
				},
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 4),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Int16"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Int16"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 5),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "UInt16"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "UInt16"}),
			},
			refs: []*ua.ReferenceDescription{
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 15904),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 32251),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 95),
					IsForward:       true,
				},
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 6),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Int32"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Int32"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 7),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "UInt32"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "UInt32"}),
			},
			refs: []*ua.ReferenceDescription{
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 24279),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 24277),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 23564),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 15406),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 15646),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 17588),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 25517),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 15583),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 31917),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 15654),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 20998),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 94),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 15658),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 347),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 288),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 15642),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 289),
					IsForward:       true,
				},
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 8),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Int64"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Int64"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 9),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "UInt64"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "UInt64"}),
			},
			refs: []*ua.ReferenceDescription{
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 11737),
					IsForward:       true,
				},
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 10),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Float"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Float"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 11),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Double"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Double"}),
			},
			refs: []*ua.ReferenceDescription{
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 290),
					IsForward:       true,
				},
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 12),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "String"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "String"}),
			},
			refs: []*ua.ReferenceDescription{
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 25726),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 12881),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 23751),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 12880),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 295),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 291),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 12879),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 24263),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 31918),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 12878),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 12877),
					IsForward:       true,
				},
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 13),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "DateTime"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "DateTime"}),
			},
			refs: []*ua.ReferenceDescription{
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 294),
					IsForward:       true,
				},
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 14),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Guid"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Guid"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "ByteString"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "ByteString"}),
			},
			refs: []*ua.ReferenceDescription{
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 521),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 30),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 16307),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 311),
					IsForward:       true,
				},
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 16),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "XmlElement"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "XmlElement"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 17),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "NodeId"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "NodeId"}),
			},
			refs: []*ua.ReferenceDescription{
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 388),
					IsForward:       true,
				},
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 18),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "ExpandedNodeId"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "ExpandedNodeId"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 19),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "StatusCode"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "StatusCode"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 20),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "QualifiedName"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "QualifiedName"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 21),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "LocalizedText"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "LocalizedText"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 22),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Structure"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Structure"}),
			},
			refs: []*ua.ReferenceDescription{
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 432),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 12755),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 12079),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 15605),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 15629),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 859),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 15580),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 12890),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 15621),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 23468),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 18813),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 11944),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 18806),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 874),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 32285),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 15634),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 23603),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 15598),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 32660),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 14525),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 12756),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 11943),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 96),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 891),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 338),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 15630),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 583),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 97),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 101),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 344),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 296),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 15578),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 25220),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 304),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 948),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 14524),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 308),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 887),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 7594),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 32659),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 14273),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 540),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 18809),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 12554),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 537),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 376),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 14593),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 15534),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 14533),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 871),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 12080),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 24033),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 920),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 17548),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 897),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 589),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 316),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 379),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 14744),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 8912),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 385),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 15628),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 12172),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 659),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 15611),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 15622),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 312),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 25520),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 12189),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 853),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 15616),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 894),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 15528),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 24105),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 18807),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 15597),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 719),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 15617),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 868),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 382),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 299),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 24281),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 884),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 15618),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 24106),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 331),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 23498),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 16313),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 15530),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 18811),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 23601),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 12171),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 856),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 24107),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 15623),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 25270),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 25519),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 862),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 15609),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 877),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 865),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 15502),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 586),
					IsForward:       true,
				},
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 23),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "DataValue"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "DataValue"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 25),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "DiagnosticInfo"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "DiagnosticInfo"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 30),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Image"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Image"}),
			},
			refs: []*ua.ReferenceDescription{
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 2000),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 2002),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 2003),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 2001),
					IsForward:       true,
				},
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 50),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Decimal"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Decimal"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 31),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassReferenceType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "References"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "References"}),
			},
			refs: []*ua.ReferenceDescription{
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 32),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 33),
					IsForward:       true,
				},
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 32),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassReferenceType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "NonHierarchicalReferences"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "NonHierarchicalReferences"}),
			},
			refs: []*ua.ReferenceDescription{
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 39),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 23562),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 51),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 37),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 23469),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 25257),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 41),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 24137),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 25255),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 117),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 9005),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 25237),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 40),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 17603),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 9004),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 53),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 25253),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 38),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 9006),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 17597),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 25258),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 54),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 52),
					IsForward:       true,
				},
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 33),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassReferenceType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "HierarchicalReferences"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "HierarchicalReferences"}),
				ua.AttributeIDInverseName: ua.MustVariant(&ua.LocalizedText{Locale: "", Text: "InverseHierarchicalReferences"}),
			},
			refs: []*ua.ReferenceDescription{
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 36),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 25254),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 35),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 25345),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 14936),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 34),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 25256),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 25238),
					IsForward:       true,
				},
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 34),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassReferenceType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "HasChild"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "HasChild"}),
				ua.AttributeIDInverseName: ua.MustVariant(&ua.LocalizedText{Locale: "", Text: "ChildOf"}),
			},
			refs: []*ua.ReferenceDescription{
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 45),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 44),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 32679),
					IsForward:       true,
				},
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 35),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassReferenceType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Organizes"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Organizes"}),
				ua.AttributeIDInverseName: ua.MustVariant(&ua.LocalizedText{Locale: "", Text: "OrganizedBy"}),
			},
			refs: []*ua.ReferenceDescription{
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 16362),
					IsForward:       true,
				},
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 36),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassReferenceType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "HasEventSource"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "HasEventSource"}),
				ua.AttributeIDInverseName: ua.MustVariant(&ua.LocalizedText{Locale: "", Text: "EventSourceOf"}),
			},
			refs: []*ua.ReferenceDescription{
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 48),
					IsForward:       true,
				},
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 37),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassReferenceType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "HasModellingRule"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "HasModellingRule"}),
				ua.AttributeIDInverseName: ua.MustVariant(&ua.LocalizedText{Locale: "", Text: "ModellingRuleOf"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 38),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassReferenceType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "HasEncoding"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "HasEncoding"}),
				ua.AttributeIDInverseName: ua.MustVariant(&ua.LocalizedText{Locale: "", Text: "EncodingOf"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 39),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassReferenceType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "HasDescription"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "HasDescription"}),
				ua.AttributeIDInverseName: ua.MustVariant(&ua.LocalizedText{Locale: "", Text: "DescriptionOf"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 40),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassReferenceType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "HasTypeDefinition"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "HasTypeDefinition"}),
				ua.AttributeIDInverseName: ua.MustVariant(&ua.LocalizedText{Locale: "", Text: "TypeDefinitionOf"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 41),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassReferenceType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "GeneratesEvent"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "GeneratesEvent"}),
				ua.AttributeIDInverseName: ua.MustVariant(&ua.LocalizedText{Locale: "", Text: "GeneratedBy"}),
			},
			refs: []*ua.ReferenceDescription{
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 3065),
					IsForward:       true,
				},
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 3065),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassReferenceType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "AlwaysGeneratesEvent"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "AlwaysGeneratesEvent"}),
				ua.AttributeIDInverseName: ua.MustVariant(&ua.LocalizedText{Locale: "", Text: "AlwaysGeneratedBy"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 44),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassReferenceType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Aggregates"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Aggregates"}),
				ua.AttributeIDInverseName: ua.MustVariant(&ua.LocalizedText{Locale: "", Text: "AggregatedBy"}),
			},
			refs: []*ua.ReferenceDescription{
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 56),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 46),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 47),
					IsForward:       true,
				},
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 45),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassReferenceType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "HasSubtype"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "HasSubtype"}),
				ua.AttributeIDInverseName: ua.MustVariant(&ua.LocalizedText{Locale: "", Text: "SubtypeOf"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 46),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassReferenceType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "HasProperty"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "HasProperty"}),
				ua.AttributeIDInverseName: ua.MustVariant(&ua.LocalizedText{Locale: "", Text: "PropertyOf"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 47),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassReferenceType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "HasComponent"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "HasComponent"}),
				ua.AttributeIDInverseName: ua.MustVariant(&ua.LocalizedText{Locale: "", Text: "ComponentOf"}),
			},
			refs: []*ua.ReferenceDescription{
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 15112),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 14476),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 17604),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 18805),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 16361),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 24136),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 15296),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 18804),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 25262),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 15297),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 129),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 49),
					IsForward:       true,
				},
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 48),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassReferenceType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "HasNotifier"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "HasNotifier"}),
				ua.AttributeIDInverseName: ua.MustVariant(&ua.LocalizedText{Locale: "", Text: "NotifierOf"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 49),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassReferenceType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "HasOrderedComponent"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "HasOrderedComponent"}),
				ua.AttributeIDInverseName: ua.MustVariant(&ua.LocalizedText{Locale: "", Text: "OrderedComponentOf"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 51),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassReferenceType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "FromState"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "FromState"}),
				ua.AttributeIDInverseName: ua.MustVariant(&ua.LocalizedText{Locale: "", Text: "ToTransition"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 52),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassReferenceType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "ToState"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "ToState"}),
				ua.AttributeIDInverseName: ua.MustVariant(&ua.LocalizedText{Locale: "", Text: "FromTransition"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 53),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassReferenceType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "HasCause"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "HasCause"}),
				ua.AttributeIDInverseName: ua.MustVariant(&ua.LocalizedText{Locale: "", Text: "MayBeCausedBy"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 54),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassReferenceType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "HasEffect"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "HasEffect"}),
				ua.AttributeIDInverseName: ua.MustVariant(&ua.LocalizedText{Locale: "", Text: "MayBeEffectedBy"}),
			},
			refs: []*ua.ReferenceDescription{
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 17985),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 17984),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 17983),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 17276),
					IsForward:       true,
				},
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 117),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassReferenceType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "HasSubStateMachine"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "HasSubStateMachine"}),
				ua.AttributeIDInverseName: ua.MustVariant(&ua.LocalizedText{Locale: "", Text: "SubStateMachineOf"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 56),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassReferenceType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "HasHistoricalConfiguration"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "HasHistoricalConfiguration"}),
				ua.AttributeIDInverseName: ua.MustVariant(&ua.LocalizedText{Locale: "", Text: "HistoricalConfigurationOf"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 24136),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassReferenceType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "HasStructuredComponent"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "HasStructuredComponent"}),
				ua.AttributeIDInverseName: ua.MustVariant(&ua.LocalizedText{Locale: "", Text: "IsStructuredComponentOf"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 24137),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassReferenceType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "AssociatedWith"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "AssociatedWith"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 58),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "BaseObjectType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "BaseObjectType"}),
			},
			refs: []*ua.ReferenceDescription{
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 2029),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 15906),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 76),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 18001),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 11575),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 19677),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 21145),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 15607),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 2340),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 15489),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 24264),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 15471),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 17602),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 75),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 25227),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 2318),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 2033),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 2310),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 2026),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 15306),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 26871),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 23832),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 17997),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 12555),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 21104),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 14232),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 2041),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 25221),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 23518),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 11616),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 32286),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 14209),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 15620),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 15108),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 2034),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 12581),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 21090),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 17998),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 17852),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 11645),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 2330),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 11187),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 2299),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 17721),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 12556),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 23455),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 77),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 25337),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 61),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 14643),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 15305),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 15319),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 14509),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 21096),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 21091),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 2307),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 2013),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 2004),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 17589),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 15298),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 23828),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 11163),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 15744),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 17279),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 2020),
					IsForward:       true,
				},
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 61),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "FolderType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "FolderType"}),
			},
			refs: []*ua.ReferenceDescription{
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 13353),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 17591),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 16405),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 14477),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 23556),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 15452),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 23795),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 25346),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 11564),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 17496),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 23456),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 13813),
					IsForward:       true,
				},
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 62),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassVariableType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "BaseVariableType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "BaseVariableType"}),
			},
			refs: []*ua.ReferenceDescription{
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 68),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 63),
					IsForward:       true,
				},
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 63),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassVariableType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "BaseDataVariableType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "BaseDataVariableType"}),
			},
			refs: []*ua.ReferenceDescription{
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 2138),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 32657),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 69),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 17714),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 2172),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 2755),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 16309),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 2762),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 11487),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 2244),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 17709),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 72),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 2137),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 18779),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 17986),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 2197),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 32244),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 2243),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 2171),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 15383),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 18772),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 17277),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 9002),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 2365),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 2380),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 2165),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 18786),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 19725),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 15113),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 2164),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 2196),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 2150),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 3051),
					IsForward:       true,
				},
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 68),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassVariableType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "PropertyType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "PropertyType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 69),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassVariableType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "DataTypeDescriptionType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "DataTypeDescriptionType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 72),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassVariableType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "DataTypeDictionaryType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "DataTypeDictionaryType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 75),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "DataTypeSystemType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "DataTypeSystemType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 76),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "DataTypeEncodingType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "DataTypeEncodingType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 120),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "NamingRuleType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "NamingRuleType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 77),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "ModellingRuleType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "ModellingRuleType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 78),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Mandatory"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Mandatory"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 80),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Optional"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Optional"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 83),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "ExposesItsArray"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "ExposesItsArray"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 11508),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "OptionalPlaceholder"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "OptionalPlaceholder"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 11510),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "MandatoryPlaceholder"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "MandatoryPlaceholder"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 84),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Root"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Root"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 85),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Objects"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Objects"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 86),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Types"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Types"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 87),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Views"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Views"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 88),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "ObjectTypes"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "ObjectTypes"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 89),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "VariableTypes"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "VariableTypes"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 90),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "DataTypes"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "DataTypes"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 91),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "ReferenceTypes"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "ReferenceTypes"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 92),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "XML Schema"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "XML Schema"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 93),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "OPC Binary"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "OPC Binary"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 129),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassReferenceType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "HasArgumentDescription"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "HasArgumentDescription"}),
				ua.AttributeIDInverseName: ua.MustVariant(&ua.LocalizedText{Locale: "", Text: "ArgumentDescriptionOf"}),
			},
			refs: []*ua.ReferenceDescription{
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 131),
					IsForward:       true,
				},
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 131),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassReferenceType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "HasOptionalInputArgumentDescription"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "HasOptionalInputArgumentDescription"}),
				ua.AttributeIDInverseName: ua.MustVariant(&ua.LocalizedText{Locale: "", Text: "OptionalInputArgumentDescriptionOf"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15957),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "http://opcfoundation.org/UA/"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "http://opcfoundation.org/UA/"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 3068),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassVariable),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "NodeVersion"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "NodeVersion"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 12170),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassVariable),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "ViewVersion"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "ViewVersion"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 3067),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassVariable),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Icon"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Icon"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 3069),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassVariable),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "LocalTime"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "LocalTime"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 3070),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassVariable),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "AllowNulls"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "AllowNulls"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 11433),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassVariable),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "ValueAsText"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "ValueAsText"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 11498),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassVariable),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "MaxStringLength"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "MaxStringLength"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15002),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassVariable),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "MaxCharacters"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "MaxCharacters"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 12908),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassVariable),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "MaxByteStringLength"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "MaxByteStringLength"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 11512),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassVariable),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "MaxArrayLength"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "MaxArrayLength"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 11513),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassVariable),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "EngineeringUnits"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "EngineeringUnits"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 11432),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassVariable),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "EnumStrings"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "EnumStrings"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 3071),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassVariable),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "EnumValues"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "EnumValues"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 12745),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassVariable),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "OptionSetValues"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "OptionSetValues"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 32750),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassVariable),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "OptionSetLength"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "OptionSetLength"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 3072),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassVariable),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "InputArguments"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "InputArguments"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 3073),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassVariable),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "OutputArguments"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "OutputArguments"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 17605),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassVariable),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "DefaultInstanceBrowseName"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "DefaultInstanceBrowseName"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 2000),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "ImageBMP"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "ImageBMP"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 2001),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "ImageGIF"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "ImageGIF"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 2002),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "ImageJPG"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "ImageJPG"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 2003),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "ImagePNG"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "ImagePNG"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 16307),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "AudioDataType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "AudioDataType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 12756),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Union"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Union"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 23751),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "UriString"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "UriString"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 2004),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "ServerType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "ServerType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 2013),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "ServerCapabilitiesType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "ServerCapabilitiesType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 2020),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "ServerDiagnosticsType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "ServerDiagnosticsType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 2026),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "SessionsDiagnosticsSummaryType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "SessionsDiagnosticsSummaryType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 2029),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "SessionDiagnosticsObjectType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "SessionDiagnosticsObjectType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 2033),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "VendorServerInfoType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "VendorServerInfoType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 2034),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "ServerRedundancyType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "ServerRedundancyType"}),
			},
			refs: []*ua.ReferenceDescription{
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 2036),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 2039),
					IsForward:       true,
				},
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 2036),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "TransparentRedundancyType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "TransparentRedundancyType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 2039),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "NonTransparentRedundancyType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "NonTransparentRedundancyType"}),
			},
			refs: []*ua.ReferenceDescription{
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 11945),
					IsForward:       true,
				},
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 11945),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "NonTransparentNetworkRedundancyType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "NonTransparentNetworkRedundancyType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 11564),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "OperationLimitsType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "OperationLimitsType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 11575),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "FileType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "FileType"}),
			},
			refs: []*ua.ReferenceDescription{
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 12522),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 25482),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 11595),
					IsForward:       true,
				},
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 11595),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "AddressSpaceFileType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "AddressSpaceFileType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 11616),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "NamespaceMetadataType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "NamespaceMetadataType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 11645),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "NamespacesType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "NamespacesType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 2041),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "BaseEventType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "BaseEventType"}),
			},
			refs: []*ua.ReferenceDescription{
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 2130),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 2052),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 2782),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 11436),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 2738),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 2311),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 3035),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 2132),
					IsForward:       true,
				},
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 2052),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "AuditEventType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "AuditEventType"}),
			},
			refs: []*ua.ReferenceDescription{
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 2090),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 2058),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 23606),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 12561),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 12620),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 2127),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 2099),
					IsForward:       true,
				},
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 2058),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "AuditSecurityEventType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "AuditSecurityEventType"}),
			},
			refs: []*ua.ReferenceDescription{
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 2059),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 2080),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 2069),
					IsForward:       true,
				},
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 2059),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "AuditChannelEventType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "AuditChannelEventType"}),
			},
			refs: []*ua.ReferenceDescription{
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 2060),
					IsForward:       true,
				},
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 2060),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "AuditOpenSecureChannelEventType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "AuditOpenSecureChannelEventType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 2069),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "AuditSessionEventType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "AuditSessionEventType"}),
			},
			refs: []*ua.ReferenceDescription{
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 2075),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 2071),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 2078),
					IsForward:       true,
				},
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 2071),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "AuditCreateSessionEventType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "AuditCreateSessionEventType"}),
			},
			refs: []*ua.ReferenceDescription{
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 2748),
					IsForward:       true,
				},
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 2748),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "AuditUrlMismatchEventType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "AuditUrlMismatchEventType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 2075),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "AuditActivateSessionEventType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "AuditActivateSessionEventType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 2078),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "AuditCancelEventType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "AuditCancelEventType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 2080),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "AuditCertificateEventType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "AuditCertificateEventType"}),
			},
			refs: []*ua.ReferenceDescription{
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 2088),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 2086),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 2087),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 2089),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 2082),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 2085),
					IsForward:       true,
				},
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 2082),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "AuditCertificateDataMismatchEventType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "AuditCertificateDataMismatchEventType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 2085),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "AuditCertificateExpiredEventType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "AuditCertificateExpiredEventType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 2086),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "AuditCertificateInvalidEventType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "AuditCertificateInvalidEventType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 2087),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "AuditCertificateUntrustedEventType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "AuditCertificateUntrustedEventType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 2088),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "AuditCertificateRevokedEventType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "AuditCertificateRevokedEventType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 2089),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "AuditCertificateMismatchEventType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "AuditCertificateMismatchEventType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 2090),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "AuditNodeManagementEventType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "AuditNodeManagementEventType"}),
			},
			refs: []*ua.ReferenceDescription{
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 2095),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 2097),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 2091),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 2093),
					IsForward:       true,
				},
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 2091),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "AuditAddNodesEventType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "AuditAddNodesEventType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 2093),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "AuditDeleteNodesEventType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "AuditDeleteNodesEventType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 2095),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "AuditAddReferencesEventType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "AuditAddReferencesEventType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 2097),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "AuditDeleteReferencesEventType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "AuditDeleteReferencesEventType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 2099),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "AuditUpdateEventType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "AuditUpdateEventType"}),
			},
			refs: []*ua.ReferenceDescription{
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 2100),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 2104),
					IsForward:       true,
				},
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 2100),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "AuditWriteUpdateEventType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "AuditWriteUpdateEventType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 2104),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "AuditHistoryUpdateEventType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "AuditHistoryUpdateEventType"}),
			},
			refs: []*ua.ReferenceDescription{
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 3006),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 3012),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 2999),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 19095),
					IsForward:       true,
				},
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 2127),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "AuditUpdateMethodEventType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "AuditUpdateMethodEventType"}),
			},
			refs: []*ua.ReferenceDescription{
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 32306),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 17641),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 2315),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 32260),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 2790),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 18011),
					IsForward:       true,
				},
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 2130),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "SystemEventType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "SystemEventType"}),
			},
			refs: []*ua.ReferenceDescription{
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 2788),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 2789),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 2787),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 2131),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 11446),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 15535),
					IsForward:       true,
				},
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 2131),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "DeviceFailureEventType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "DeviceFailureEventType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 11446),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "SystemStatusChangeEventType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "SystemStatusChangeEventType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 2132),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "BaseModelChangeEventType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "BaseModelChangeEventType"}),
			},
			refs: []*ua.ReferenceDescription{
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 2133),
					IsForward:       true,
				},
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 2133),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "GeneralModelChangeEventType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "GeneralModelChangeEventType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 2738),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "SemanticChangeEventType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "SemanticChangeEventType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 3035),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "EventQueueOverflowEventType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "EventQueueOverflowEventType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 11436),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "ProgressEventType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "ProgressEventType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 23606),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "AuditClientEventType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "AuditClientEventType"}),
			},
			refs: []*ua.ReferenceDescription{
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 23926),
					IsForward:       true,
				},
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 23926),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "AuditClientUpdateMethodResultEventType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "AuditClientUpdateMethodResultEventType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 2340),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "AggregateFunctionType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "AggregateFunctionType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 2137),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassVariableType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "ServerVendorCapabilityType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "ServerVendorCapabilityType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 2138),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassVariableType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "ServerStatusType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "ServerStatusType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 3051),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassVariableType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "BuildInfoType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "BuildInfoType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 2150),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassVariableType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "ServerDiagnosticsSummaryType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "ServerDiagnosticsSummaryType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 2164),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassVariableType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "SamplingIntervalDiagnosticsArrayType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "SamplingIntervalDiagnosticsArrayType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 2165),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassVariableType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "SamplingIntervalDiagnosticsType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "SamplingIntervalDiagnosticsType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 2171),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassVariableType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "SubscriptionDiagnosticsArrayType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "SubscriptionDiagnosticsArrayType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 2172),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassVariableType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "SubscriptionDiagnosticsType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "SubscriptionDiagnosticsType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 2196),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassVariableType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "SessionDiagnosticsArrayType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "SessionDiagnosticsArrayType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 2197),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassVariableType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "SessionDiagnosticsVariableType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "SessionDiagnosticsVariableType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 2243),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassVariableType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "SessionSecurityDiagnosticsArrayType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "SessionSecurityDiagnosticsArrayType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 2244),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassVariableType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "SessionSecurityDiagnosticsType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "SessionSecurityDiagnosticsType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 11487),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassVariableType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "OptionSetType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "OptionSetType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 16309),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassVariableType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "SelectionListType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "SelectionListType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 17986),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassVariableType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "AudioVariableType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "AudioVariableType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 3048),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "EventTypes"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "EventTypes"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 31915),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Locations"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Locations"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 2253),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Server"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Server"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 11312),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassVariable),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "CurrentServerId"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "CurrentServerId"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 11313),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassVariable),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "RedundantServerArray"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "RedundantServerArray"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 11314),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassVariable),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "ServerUriArray"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "ServerUriArray"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 14415),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassVariable),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "ServerNetworkGroups"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "ServerNetworkGroups"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 11192),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "HistoryServerCapabilities"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "HistoryServerCapabilities"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 23562),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassReferenceType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "IsDeprecated"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "IsDeprecated"}),
				ua.AttributeIDInverseName: ua.MustVariant(&ua.LocalizedText{Locale: "", Text: "Deprecates"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 11737),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "BitFieldMaskDataType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "BitFieldMaskDataType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 24263),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "SemanticVersionString"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "SemanticVersionString"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 14533),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "KeyValuePair"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "KeyValuePair"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 16313),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "AdditionalParametersType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "AdditionalParametersType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 17548),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "EphemeralKeyType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "EphemeralKeyType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15528),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "EndpointType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "EndpointType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 31917),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Handle"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Handle"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 31918),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "TrimmedString"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "TrimmedString"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 2299),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "StateMachineType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "StateMachineType"}),
			},
			refs: []*ua.ReferenceDescription{
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 2771),
					IsForward:       true,
				},
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 2755),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassVariableType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "StateVariableType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "StateVariableType"}),
			},
			refs: []*ua.ReferenceDescription{
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 2760),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 8995),
					IsForward:       true,
				},
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 2762),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassVariableType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "TransitionVariableType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "TransitionVariableType"}),
			},
			refs: []*ua.ReferenceDescription{
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 2767),
					IsForward:       true,
				},
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 2771),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "FiniteStateMachineType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "FiniteStateMachineType"}),
			},
			refs: []*ua.ReferenceDescription{
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 9318),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 2929),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 2391),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 15803),
					IsForward:       true,
				},
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 2760),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassVariableType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "FiniteStateVariableType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "FiniteStateVariableType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 2767),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassVariableType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "FiniteTransitionVariableType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "FiniteTransitionVariableType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 2307),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "StateType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "StateType"}),
			},
			refs: []*ua.ReferenceDescription{
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 15109),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 2309),
					IsForward:       true,
				},
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 2309),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "InitialStateType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "InitialStateType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 2310),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "TransitionType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "TransitionType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15109),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "ChoiceStateType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "ChoiceStateType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15112),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassReferenceType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "HasGuard"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "HasGuard"}),
				ua.AttributeIDInverseName: ua.MustVariant(&ua.LocalizedText{Locale: "", Text: "GuardOf"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15113),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassVariableType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "GuardVariableType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "GuardVariableType"}),
			},
			refs: []*ua.ReferenceDescription{
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 15128),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 15317),
					IsForward:       true,
				},
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15128),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassVariableType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "ExpressionGuardVariableType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "ExpressionGuardVariableType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15317),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassVariableType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "ElseGuardVariableType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "ElseGuardVariableType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 17709),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassVariableType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "RationalNumberType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "RationalNumberType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 17714),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassVariableType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "VectorType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "VectorType"}),
			},
			refs: []*ua.ReferenceDescription{
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 17716),
					IsForward:       true,
				},
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 17716),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassVariableType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "3DVectorType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "3DVectorType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 18772),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassVariableType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "CartesianCoordinatesType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "CartesianCoordinatesType"}),
			},
			refs: []*ua.ReferenceDescription{
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 18774),
					IsForward:       true,
				},
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 18774),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassVariableType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "3DCartesianCoordinatesType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "3DCartesianCoordinatesType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 18779),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassVariableType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "OrientationType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "OrientationType"}),
			},
			refs: []*ua.ReferenceDescription{
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 18781),
					IsForward:       true,
				},
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 18781),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassVariableType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "3DOrientationType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "3DOrientationType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 18786),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassVariableType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "FrameType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "FrameType"}),
			},
			refs: []*ua.ReferenceDescription{
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 18791),
					IsForward:       true,
				},
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 18791),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassVariableType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "3DFrameType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "3DFrameType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 18806),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "RationalNumber"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "RationalNumber"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 18807),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Vector"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Vector"}),
			},
			refs: []*ua.ReferenceDescription{
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 18808),
					IsForward:       true,
				},
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 18808),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "3DVector"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "3DVector"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 18809),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "CartesianCoordinates"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "CartesianCoordinates"}),
			},
			refs: []*ua.ReferenceDescription{
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 18810),
					IsForward:       true,
				},
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 18810),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "3DCartesianCoordinates"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "3DCartesianCoordinates"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 18811),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Orientation"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Orientation"}),
			},
			refs: []*ua.ReferenceDescription{
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 18812),
					IsForward:       true,
				},
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 18812),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "3DOrientation"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "3DOrientation"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 18813),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Frame"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Frame"}),
			},
			refs: []*ua.ReferenceDescription{
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 18814),
					IsForward:       true,
				},
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 18814),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "3DFrame"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "3DFrame"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 2311),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "TransitionEventType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "TransitionEventType"}),
			},
			refs: []*ua.ReferenceDescription{
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 2378),
					IsForward:       true,
				},
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 2315),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "AuditUpdateStateEventType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "AuditUpdateStateEventType"}),
			},
			refs: []*ua.ReferenceDescription{
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 11856),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 3806),
					IsForward:       true,
				},
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 11939),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "OpenFileMode"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "OpenFileMode"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 13353),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "FileDirectoryType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "FileDirectoryType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 16314),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "FileSystem"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "FileSystem"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15744),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "TemporaryFileTransferType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "TemporaryFileTransferType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15803),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "FileTransferStateMachineType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "FileTransferStateMachineType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15607),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "RoleSetType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "RoleSetType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15620),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "RoleType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "RoleType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15632),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "IdentityCriteriaType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "IdentityCriteriaType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15634),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "IdentityMappingRuleType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "IdentityMappingRuleType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 17641),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "RoleMappingRuleChangedAuditEventType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "RoleMappingRuleChangedAuditEventType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15644),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Anonymous"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Anonymous"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15656),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "AuthenticatedUser"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "AuthenticatedUser"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15668),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Observer"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Observer"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15680),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Operator"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Operator"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 16036),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Engineer"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Engineer"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15692),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Supervisor"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Supervisor"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15716),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "ConfigureAdmin"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "ConfigureAdmin"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15704),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "SecurityAdmin"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "SecurityAdmin"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 25565),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "SecurityKeyServerAdmin"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "SecurityKeyServerAdmin"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 25603),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "SecurityKeyServerAccess"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "SecurityKeyServerAccess"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 25584),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "SecurityKeyServerPush"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "SecurityKeyServerPush"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 17589),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "DictionaryEntryType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "DictionaryEntryType"}),
			},
			refs: []*ua.ReferenceDescription{
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 17598),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 17600),
					IsForward:       true,
				},
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 17591),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "DictionaryFolderType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "DictionaryFolderType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 17594),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Dictionaries"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Dictionaries"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 17597),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassReferenceType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "HasDictionaryEntry"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "HasDictionaryEntry"}),
				ua.AttributeIDInverseName: ua.MustVariant(&ua.LocalizedText{Locale: "", Text: "DictionaryEntryOf"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 17598),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "IrdiDictionaryEntryType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "IrdiDictionaryEntryType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 17600),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "UriDictionaryEntryType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "UriDictionaryEntryType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 17602),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "BaseInterfaceType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "BaseInterfaceType"}),
			},
			refs: []*ua.ReferenceDescription{
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 24158),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 25218),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 24173),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 23513),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 24183),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 24148),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 24199),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 24188),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 24167),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 24205),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 24202),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 24233),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 24179),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 24169),
					IsForward:       true,
				},
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 17708),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "InterfaceTypes"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "InterfaceTypes"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 17603),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassReferenceType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "HasInterface"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "HasInterface"}),
				ua.AttributeIDInverseName: ua.MustVariant(&ua.LocalizedText{Locale: "", Text: "InterfaceOf"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 17604),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassReferenceType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "HasAddIn"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "HasAddIn"}),
				ua.AttributeIDInverseName: ua.MustVariant(&ua.LocalizedText{Locale: "", Text: "AddInOf"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 23498),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "CurrencyUnitType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "CurrencyUnitType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 23501),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassVariable),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "CurrencyUnit"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "CurrencyUnit"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 23513),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "IOrderedObjectType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "IOrderedObjectType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 23518),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "OrderedListType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "OrderedListType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 2365),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassVariableType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "DataItemType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "DataItemType"}),
			},
			refs: []*ua.ReferenceDescription{
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 2372),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 15318),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 12021),
					IsForward:       true,
				},
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15318),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassVariableType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "BaseAnalogType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "BaseAnalogType"}),
			},
			refs: []*ua.ReferenceDescription{
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 17497),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 2368),
					IsForward:       true,
				},
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 2368),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassVariableType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "AnalogItemType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "AnalogItemType"}),
			},
			refs: []*ua.ReferenceDescription{
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 17570),
					IsForward:       true,
				},
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 17497),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassVariableType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "AnalogUnitType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "AnalogUnitType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 17570),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassVariableType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "AnalogUnitRangeType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "AnalogUnitRangeType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 2372),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassVariableType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "DiscreteItemType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "DiscreteItemType"}),
			},
			refs: []*ua.ReferenceDescription{
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 2376),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 11238),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 2373),
					IsForward:       true,
				},
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 2373),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassVariableType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "TwoStateDiscreteType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "TwoStateDiscreteType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 2376),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassVariableType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "MultiStateDiscreteType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "MultiStateDiscreteType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 11238),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassVariableType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "MultiStateValueDiscreteType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "MultiStateValueDiscreteType"}),
			},
			refs: []*ua.ReferenceDescription{
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 19077),
					IsForward:       true,
				},
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 12021),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassVariableType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "ArrayItemType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "ArrayItemType"}),
			},
			refs: []*ua.ReferenceDescription{
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 12047),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 12057),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 12029),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 12068),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 12038),
					IsForward:       true,
				},
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 12029),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassVariableType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "YArrayItemType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "YArrayItemType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 12038),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassVariableType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "XYArrayItemType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "XYArrayItemType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 12047),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassVariableType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "ImageItemType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "ImageItemType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 12057),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassVariableType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "CubeItemType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "CubeItemType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 12068),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassVariableType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "NDimensionArrayItemType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "NDimensionArrayItemType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 8995),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassVariableType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "TwoStateVariableType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "TwoStateVariableType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 9002),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassVariableType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "ConditionVariableType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "ConditionVariableType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 9004),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassReferenceType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "HasTrueSubState"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "HasTrueSubState"}),
				ua.AttributeIDInverseName: ua.MustVariant(&ua.LocalizedText{Locale: "", Text: "IsTrueSubStateOf"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 9005),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassReferenceType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "HasFalseSubState"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "HasFalseSubState"}),
				ua.AttributeIDInverseName: ua.MustVariant(&ua.LocalizedText{Locale: "", Text: "IsFalseSubStateOf"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 16361),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassReferenceType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "HasAlarmSuppressionGroup"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "HasAlarmSuppressionGroup"}),
				ua.AttributeIDInverseName: ua.MustVariant(&ua.LocalizedText{Locale: "", Text: "IsAlarmSuppressionGroupOf"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 16362),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassReferenceType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "AlarmGroupMember"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "AlarmGroupMember"}),
				ua.AttributeIDInverseName: ua.MustVariant(&ua.LocalizedText{Locale: "", Text: "MemberOfAlarmGroup"}),
			},
			refs: []*ua.ReferenceDescription{
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 32059),
					IsForward:       true,
				},
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 32059),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassReferenceType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "AlarmSuppressionGroupMember"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "AlarmSuppressionGroupMember"}),
				ua.AttributeIDInverseName: ua.MustVariant(&ua.LocalizedText{Locale: "", Text: "MemberOfAlarmSuppressionGroup"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 2782),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "ConditionType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "ConditionType"}),
			},
			refs: []*ua.ReferenceDescription{
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 2881),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 2830),
					IsForward:       true,
				},
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 2830),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "DialogConditionType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "DialogConditionType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 2881),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "AcknowledgeableConditionType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "AcknowledgeableConditionType"}),
			},
			refs: []*ua.ReferenceDescription{
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 2915),
					IsForward:       true,
				},
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 2915),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "AlarmConditionType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "AlarmConditionType"}),
			},
			refs: []*ua.ReferenceDescription{
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 17080),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 10523),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 2955),
					IsForward:       true,
				},
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 16405),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "AlarmGroupType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "AlarmGroupType"}),
			},
			refs: []*ua.ReferenceDescription{
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 32064),
					IsForward:       true,
				},
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 32064),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "AlarmSuppressionGroupType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "AlarmSuppressionGroupType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 2929),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "ShelvedStateMachineType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "ShelvedStateMachineType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 2955),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "LimitAlarmType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "LimitAlarmType"}),
			},
			refs: []*ua.ReferenceDescription{
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 9341),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 9906),
					IsForward:       true,
				},
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 9318),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "ExclusiveLimitStateMachineType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "ExclusiveLimitStateMachineType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 9341),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "ExclusiveLimitAlarmType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "ExclusiveLimitAlarmType"}),
			},
			refs: []*ua.ReferenceDescription{
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 9623),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 9764),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 9482),
					IsForward:       true,
				},
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 9906),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "NonExclusiveLimitAlarmType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "NonExclusiveLimitAlarmType"}),
			},
			refs: []*ua.ReferenceDescription{
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 10060),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 10368),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 10214),
					IsForward:       true,
				},
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 10060),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "NonExclusiveLevelAlarmType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "NonExclusiveLevelAlarmType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 9482),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "ExclusiveLevelAlarmType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "ExclusiveLevelAlarmType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 10368),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "NonExclusiveDeviationAlarmType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "NonExclusiveDeviationAlarmType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 10214),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "NonExclusiveRateOfChangeAlarmType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "NonExclusiveRateOfChangeAlarmType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 9764),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "ExclusiveDeviationAlarmType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "ExclusiveDeviationAlarmType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 9623),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "ExclusiveRateOfChangeAlarmType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "ExclusiveRateOfChangeAlarmType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 10523),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "DiscreteAlarmType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "DiscreteAlarmType"}),
			},
			refs: []*ua.ReferenceDescription{
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 10637),
					IsForward:       true,
				},
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 10637),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "OffNormalAlarmType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "OffNormalAlarmType"}),
			},
			refs: []*ua.ReferenceDescription{
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 11753),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 18496),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 18347),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 10751),
					IsForward:       true,
				},
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 11753),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "SystemOffNormalAlarmType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "SystemOffNormalAlarmType"}),
			},
			refs: []*ua.ReferenceDescription{
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 13225),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 19297),
					IsForward:       true,
				},
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 10751),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "TripAlarmType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "TripAlarmType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 18347),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "InstrumentDiagnosticAlarmType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "InstrumentDiagnosticAlarmType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 18496),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "SystemDiagnosticAlarmType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "SystemDiagnosticAlarmType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 13225),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "CertificateExpirationAlarmType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "CertificateExpirationAlarmType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 17080),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "DiscrepancyAlarmType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "DiscrepancyAlarmType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 11163),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "BaseConditionClassType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "BaseConditionClassType"}),
			},
			refs: []*ua.ReferenceDescription{
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 18665),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 17220),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 11165),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 11164),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 17219),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 17221),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 11166),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 17218),
					IsForward:       true,
				},
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 11164),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "ProcessConditionClassType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "ProcessConditionClassType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 11165),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "MaintenanceConditionClassType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "MaintenanceConditionClassType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 11166),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "SystemConditionClassType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "SystemConditionClassType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 17218),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "SafetyConditionClassType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "SafetyConditionClassType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 17219),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "HighlyManagedAlarmConditionClassType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "HighlyManagedAlarmConditionClassType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 17220),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "TrainingConditionClassType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "TrainingConditionClassType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 18665),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "StatisticalConditionClassType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "StatisticalConditionClassType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 17221),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "TestingConditionClassType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "TestingConditionClassType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 2790),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "AuditConditionEventType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "AuditConditionEventType"}),
			},
			refs: []*ua.ReferenceDescription{
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 17259),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 2829),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 8944),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 11093),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 2803),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 8927),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 8961),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 15013),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 17225),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 17242),
					IsForward:       true,
				},
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 2803),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "AuditConditionEnableEventType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "AuditConditionEnableEventType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 2829),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "AuditConditionCommentEventType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "AuditConditionCommentEventType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 8927),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "AuditConditionRespondEventType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "AuditConditionRespondEventType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 8944),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "AuditConditionAcknowledgeEventType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "AuditConditionAcknowledgeEventType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 8961),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "AuditConditionConfirmEventType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "AuditConditionConfirmEventType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 11093),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "AuditConditionShelvingEventType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "AuditConditionShelvingEventType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 17225),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "AuditConditionSuppressionEventType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "AuditConditionSuppressionEventType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 17242),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "AuditConditionSilenceEventType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "AuditConditionSilenceEventType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15013),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "AuditConditionResetEventType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "AuditConditionResetEventType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 17259),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "AuditConditionOutOfServiceEventType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "AuditConditionOutOfServiceEventType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 2787),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "RefreshStartEventType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "RefreshStartEventType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 2788),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "RefreshEndEventType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "RefreshEndEventType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 2789),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "RefreshRequiredEventType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "RefreshRequiredEventType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 9006),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassReferenceType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "HasCondition"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "HasCondition"}),
				ua.AttributeIDInverseName: ua.MustVariant(&ua.LocalizedText{Locale: "", Text: "IsConditionOf"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 17276),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassReferenceType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "HasEffectDisable"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "HasEffectDisable"}),
				ua.AttributeIDInverseName: ua.MustVariant(&ua.LocalizedText{Locale: "", Text: "MayBeDisabledBy"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 17983),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassReferenceType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "HasEffectEnable"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "HasEffectEnable"}),
				ua.AttributeIDInverseName: ua.MustVariant(&ua.LocalizedText{Locale: "", Text: "MayBeEnabledBy"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 17984),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassReferenceType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "HasEffectSuppressed"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "HasEffectSuppressed"}),
				ua.AttributeIDInverseName: ua.MustVariant(&ua.LocalizedText{Locale: "", Text: "MayBeSuppressedBy"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 17985),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassReferenceType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "HasEffectUnsuppressed"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "HasEffectUnsuppressed"}),
				ua.AttributeIDInverseName: ua.MustVariant(&ua.LocalizedText{Locale: "", Text: "MayBeUnsuppressedBy"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 17279),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "AlarmMetricsType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "AlarmMetricsType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 17277),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassVariableType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "AlarmRateVariableType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "AlarmRateVariableType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 32244),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassVariableType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "AlarmStateVariableType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "AlarmStateVariableType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 32251),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "AlarmMask"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "AlarmMask"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 2391),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "ProgramStateMachineType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "ProgramStateMachineType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 2378),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "ProgramTransitionEventType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "ProgramTransitionEventType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 11856),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "AuditProgramTransitionEventType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "AuditProgramTransitionEventType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 3806),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "ProgramTransitionAuditEventType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "ProgramTransitionAuditEventType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 2380),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassVariableType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "ProgramDiagnosticType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "ProgramDiagnosticType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15383),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassVariableType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "ProgramDiagnostic2Type"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "ProgramDiagnostic2Type"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 11214),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassVariable),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Annotations"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Annotations"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 2318),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "HistoricalDataConfigurationType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "HistoricalDataConfigurationType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 11202),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "HA Configuration"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "HA Configuration"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 11215),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassVariable),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "HistoricalEventFilter"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "HistoricalEventFilter"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 2330),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "HistoryServerCapabilitiesType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "HistoryServerCapabilitiesType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 2999),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "AuditHistoryEventUpdateEventType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "AuditHistoryEventUpdateEventType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 3006),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "AuditHistoryValueUpdateEventType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "AuditHistoryValueUpdateEventType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 19095),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "AuditHistoryAnnotationUpdateEventType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "AuditHistoryAnnotationUpdateEventType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 3012),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "AuditHistoryDeleteEventType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "AuditHistoryDeleteEventType"}),
			},
			refs: []*ua.ReferenceDescription{
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 3014),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 3022),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 3019),
					IsForward:       true,
				},
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 3014),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "AuditHistoryRawModifyDeleteEventType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "AuditHistoryRawModifyDeleteEventType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 3019),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "AuditHistoryAtTimeDeleteEventType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "AuditHistoryAtTimeDeleteEventType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 3022),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "AuditHistoryEventDeleteEventType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "AuditHistoryEventDeleteEventType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 12522),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "TrustListType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "TrustListType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 23564),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "TrustListValidationOptions"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "TrustListValidationOptions"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 12552),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "TrustListMasks"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "TrustListMasks"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 12554),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "TrustListDataType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "TrustListDataType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 19297),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "TrustListOutOfDateAlarmType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "TrustListOutOfDateAlarmType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 12555),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "CertificateGroupType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "CertificateGroupType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 13813),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "CertificateGroupFolderType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "CertificateGroupFolderType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 12556),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "CertificateType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "CertificateType"}),
			},
			refs: []*ua.ReferenceDescription{
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 12558),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 15181),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 12557),
					IsForward:       true,
				},
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 12557),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "ApplicationCertificateType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "ApplicationCertificateType"}),
			},
			refs: []*ua.ReferenceDescription{
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 23537),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 12560),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 12559),
					IsForward:       true,
				},
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 12558),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "HttpsCertificateType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "HttpsCertificateType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15181),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "UserCredentialCertificateType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "UserCredentialCertificateType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 12559),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "RsaMinApplicationCertificateType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "RsaMinApplicationCertificateType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 12560),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "RsaSha256ApplicationCertificateType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "RsaSha256ApplicationCertificateType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 23537),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "EccApplicationCertificateType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "EccApplicationCertificateType"}),
			},
			refs: []*ua.ReferenceDescription{
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 23538),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 23542),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 23541),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 23540),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 23539),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 23543),
					IsForward:       true,
				},
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 23538),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "EccNistP256ApplicationCertificateType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "EccNistP256ApplicationCertificateType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 23539),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "EccNistP384ApplicationCertificateType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "EccNistP384ApplicationCertificateType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 23540),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "EccBrainpoolP256r1ApplicationCertificateType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "EccBrainpoolP256r1ApplicationCertificateType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 23541),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "EccBrainpoolP384r1ApplicationCertificateType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "EccBrainpoolP384r1ApplicationCertificateType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 23542),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "EccCurve25519ApplicationCertificateType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "EccCurve25519ApplicationCertificateType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 23543),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "EccCurve448ApplicationCertificateType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "EccCurve448ApplicationCertificateType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 32260),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "TrustListUpdateRequestedAuditEventType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "TrustListUpdateRequestedAuditEventType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 12561),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "TrustListUpdatedAuditEventType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "TrustListUpdatedAuditEventType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 32285),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "TransactionErrorType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "TransactionErrorType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 32286),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "TransactionDiagnosticsType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "TransactionDiagnosticsType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 12581),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "ServerConfigurationType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "ServerConfigurationType"}),
			},
			refs: []*ua.ReferenceDescription{
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 25731),
					IsForward:       true,
				},
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 32306),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "CertificateUpdateRequestedAuditEventType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "CertificateUpdateRequestedAuditEventType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 12620),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "CertificateUpdatedAuditEventType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "CertificateUpdatedAuditEventType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 12637),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "ServerConfiguration"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "ServerConfiguration"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 17496),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "KeyCredentialConfigurationFolderType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "KeyCredentialConfigurationFolderType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 18155),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "KeyCredentialConfiguration"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "KeyCredentialConfiguration"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 18001),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "KeyCredentialConfigurationType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "KeyCredentialConfigurationType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 18011),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "KeyCredentialAuditEventType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "KeyCredentialAuditEventType"}),
			},
			refs: []*ua.ReferenceDescription{
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 18047),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 18029),
					IsForward:       true,
				},
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 18029),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "KeyCredentialUpdatedAuditEventType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "KeyCredentialUpdatedAuditEventType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 18047),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "KeyCredentialDeletedAuditEventType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "KeyCredentialDeletedAuditEventType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 23556),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "AuthorizationServicesConfigurationFolderType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "AuthorizationServicesConfigurationFolderType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 17732),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "AuthorizationServices"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "AuthorizationServices"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 17852),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "AuthorizationServiceConfigurationType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "AuthorizationServiceConfigurationType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 11187),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "AggregateConfigurationType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "AggregateConfigurationType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 2341),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Interpolative"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Interpolative"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 2342),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Average"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Average"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 2343),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "TimeAverage"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "TimeAverage"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 11285),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "TimeAverage2"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "TimeAverage2"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 2344),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Total"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Total"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 11304),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Total2"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Total2"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 2346),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Minimum"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Minimum"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 2347),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Maximum"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Maximum"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 2348),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "MinimumActualTime"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "MinimumActualTime"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 2349),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "MaximumActualTime"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "MaximumActualTime"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 2350),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Range"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Range"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 11286),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Minimum2"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Minimum2"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 11287),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Maximum2"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Maximum2"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 11305),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "MinimumActualTime2"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "MinimumActualTime2"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 11306),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "MaximumActualTime2"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "MaximumActualTime2"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 11288),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Range2"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Range2"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 2351),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "AnnotationCount"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "AnnotationCount"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 2352),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Count"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Count"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 11307),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "DurationInStateZero"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "DurationInStateZero"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 11308),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "DurationInStateNonZero"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "DurationInStateNonZero"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 2355),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "NumberOfTransitions"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "NumberOfTransitions"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 2357),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Start"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Start"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 2358),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "End"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "End"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 2359),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Delta"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Delta"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 11505),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "StartBound"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "StartBound"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 11506),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "EndBound"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "EndBound"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 11507),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "DeltaBounds"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "DeltaBounds"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 2360),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "DurationGood"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "DurationGood"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 2361),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "DurationBad"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "DurationBad"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 2362),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "PercentGood"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "PercentGood"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 2363),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "PercentBad"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "PercentBad"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 2364),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "WorstQuality"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "WorstQuality"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 11292),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "WorstQuality2"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "WorstQuality2"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 11426),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "StandardDeviationSample"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "StandardDeviationSample"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 11427),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "StandardDeviationPopulation"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "StandardDeviationPopulation"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 11428),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "VarianceSample"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "VarianceSample"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 11429),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "VariancePopulation"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "VariancePopulation"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15534),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "DataTypeSchemaHeader"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "DataTypeSchemaHeader"}),
			},
			refs: []*ua.ReferenceDescription{
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 15006),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 14523),
					IsForward:       true,
				},
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 14525),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "DataTypeDescription"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "DataTypeDescription"}),
			},
			refs: []*ua.ReferenceDescription{
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 15005),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 15487),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 15488),
					IsForward:       true,
				},
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15487),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "StructureDescription"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "StructureDescription"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15488),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "EnumDescription"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "EnumDescription"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15005),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "SimpleTypeDescription"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "SimpleTypeDescription"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15006),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "UABinaryFileDataType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "UABinaryFileDataType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 24105),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "PortableQualifiedName"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "PortableQualifiedName"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 24106),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "PortableNodeId"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "PortableNodeId"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 24107),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "UnsignedRationalNumber"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "UnsignedRationalNumber"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 14647),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "PubSubState"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "PubSubState"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 14523),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "DataSetMetaDataType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "DataSetMetaDataType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 14524),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "FieldMetaData"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "FieldMetaData"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15904),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "DataSetFieldFlags"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "DataSetFieldFlags"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 14593),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "ConfigurationVersionDataType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "ConfigurationVersionDataType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15578),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "PublishedDataSetDataType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "PublishedDataSetDataType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15580),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "PublishedDataSetSourceDataType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "PublishedDataSetSourceDataType"}),
			},
			refs: []*ua.ReferenceDescription{
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 15582),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 25269),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 15581),
					IsForward:       true,
				},
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 14273),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "PublishedVariableDataType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "PublishedVariableDataType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15581),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "PublishedDataItemsDataType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "PublishedDataItemsDataType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15582),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "PublishedEventsDataType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "PublishedEventsDataType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 25269),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "PublishedDataSetCustomSourceDataType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "PublishedDataSetCustomSourceDataType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15583),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "DataSetFieldContentMask"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "DataSetFieldContentMask"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15597),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "DataSetWriterDataType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "DataSetWriterDataType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15598),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "DataSetWriterTransportDataType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "DataSetWriterTransportDataType"}),
			},
			refs: []*ua.ReferenceDescription{
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 15669),
					IsForward:       true,
				},
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15605),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "DataSetWriterMessageDataType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "DataSetWriterMessageDataType"}),
			},
			refs: []*ua.ReferenceDescription{
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 15652),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 15664),
					IsForward:       true,
				},
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15609),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "PubSubGroupDataType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "PubSubGroupDataType"}),
			},
			refs: []*ua.ReferenceDescription{
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 15480),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 15520),
					IsForward:       true,
				},
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15480),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "WriterGroupDataType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "WriterGroupDataType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15611),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "WriterGroupTransportDataType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "WriterGroupTransportDataType"}),
			},
			refs: []*ua.ReferenceDescription{
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 15667),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 15532),
					IsForward:       true,
				},
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15616),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "WriterGroupMessageDataType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "WriterGroupMessageDataType"}),
			},
			refs: []*ua.ReferenceDescription{
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 15645),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 15657),
					IsForward:       true,
				},
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15617),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "PubSubConnectionDataType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "PubSubConnectionDataType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15618),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "ConnectionTransportDataType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "ConnectionTransportDataType"}),
			},
			refs: []*ua.ReferenceDescription{
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 15007),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 17467),
					IsForward:       true,
				},
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15502),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "NetworkAddressDataType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "NetworkAddressDataType"}),
			},
			refs: []*ua.ReferenceDescription{
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 15510),
					IsForward:       true,
				},
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15510),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "NetworkAddressUrlDataType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "NetworkAddressUrlDataType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15520),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "ReaderGroupDataType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "ReaderGroupDataType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15621),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "ReaderGroupTransportDataType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "ReaderGroupTransportDataType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15622),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "ReaderGroupMessageDataType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "ReaderGroupMessageDataType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15623),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "DataSetReaderDataType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "DataSetReaderDataType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15628),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "DataSetReaderTransportDataType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "DataSetReaderTransportDataType"}),
			},
			refs: []*ua.ReferenceDescription{
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 15670),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 23614),
					IsForward:       true,
				},
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15629),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "DataSetReaderMessageDataType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "DataSetReaderMessageDataType"}),
			},
			refs: []*ua.ReferenceDescription{
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 15665),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 15653),
					IsForward:       true,
				},
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15630),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "SubscribedDataSetDataType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "SubscribedDataSetDataType"}),
			},
			refs: []*ua.ReferenceDescription{
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 15631),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 15635),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 23599),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 23600),
					IsForward:       true,
				},
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15631),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "TargetVariablesDataType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "TargetVariablesDataType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 14744),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "FieldTargetDataType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "FieldTargetDataType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15874),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "OverrideValueHandling"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "OverrideValueHandling"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15635),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "SubscribedDataSetMirrorDataType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "SubscribedDataSetMirrorDataType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15530),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "PubSubConfigurationDataType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "PubSubConfigurationDataType"}),
			},
			refs: []*ua.ReferenceDescription{
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 23602),
					IsForward:       true,
				},
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 23599),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "StandaloneSubscribedDataSetRefDataType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "StandaloneSubscribedDataSetRefDataType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 23600),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "StandaloneSubscribedDataSetDataType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "StandaloneSubscribedDataSetDataType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 23601),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "SecurityGroupDataType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "SecurityGroupDataType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 25270),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "PubSubKeyPushTargetDataType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "PubSubKeyPushTargetDataType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 23602),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "PubSubConfiguration2DataType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "PubSubConfiguration2DataType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 20408),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "DataSetOrderingType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "DataSetOrderingType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15642),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "UadpNetworkMessageContentMask"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "UadpNetworkMessageContentMask"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15645),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "UadpWriterGroupMessageDataType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "UadpWriterGroupMessageDataType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15646),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "UadpDataSetMessageContentMask"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "UadpDataSetMessageContentMask"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15652),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "UadpDataSetWriterMessageDataType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "UadpDataSetWriterMessageDataType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15653),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "UadpDataSetReaderMessageDataType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "UadpDataSetReaderMessageDataType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15654),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "JsonNetworkMessageContentMask"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "JsonNetworkMessageContentMask"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15657),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "JsonWriterGroupMessageDataType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "JsonWriterGroupMessageDataType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15658),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "JsonDataSetMessageContentMask"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "JsonDataSetMessageContentMask"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15664),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "JsonDataSetWriterMessageDataType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "JsonDataSetWriterMessageDataType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15665),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "JsonDataSetReaderMessageDataType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "JsonDataSetReaderMessageDataType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 23603),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "QosDataType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "QosDataType"}),
			},
			refs: []*ua.ReferenceDescription{
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 23608),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 23604),
					IsForward:       true,
				},
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 23604),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "TransmitQosDataType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "TransmitQosDataType"}),
			},
			refs: []*ua.ReferenceDescription{
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 23605),
					IsForward:       true,
				},
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 23605),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "TransmitQosPriorityDataType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "TransmitQosPriorityDataType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 23608),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "ReceiveQosDataType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "ReceiveQosDataType"}),
			},
			refs: []*ua.ReferenceDescription{
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 23609),
					IsForward:       true,
				},
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 23609),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "ReceiveQosPriorityDataType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "ReceiveQosPriorityDataType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 17467),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "DatagramConnectionTransportDataType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "DatagramConnectionTransportDataType"}),
			},
			refs: []*ua.ReferenceDescription{
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 23612),
					IsForward:       true,
				},
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 23612),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "DatagramConnectionTransport2DataType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "DatagramConnectionTransport2DataType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15532),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "DatagramWriterGroupTransportDataType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "DatagramWriterGroupTransportDataType"}),
			},
			refs: []*ua.ReferenceDescription{
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 23613),
					IsForward:       true,
				},
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 23613),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "DatagramWriterGroupTransport2DataType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "DatagramWriterGroupTransport2DataType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 23614),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "DatagramDataSetReaderTransportDataType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "DatagramDataSetReaderTransportDataType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15007),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "BrokerConnectionTransportDataType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "BrokerConnectionTransportDataType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15008),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "BrokerTransportQualityOfService"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "BrokerTransportQualityOfService"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15667),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "BrokerWriterGroupTransportDataType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "BrokerWriterGroupTransportDataType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15669),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "BrokerDataSetWriterTransportDataType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "BrokerDataSetWriterTransportDataType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15670),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "BrokerDataSetReaderTransportDataType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "BrokerDataSetReaderTransportDataType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15906),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "PubSubKeyServiceType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "PubSubKeyServiceType"}),
			},
			refs: []*ua.ReferenceDescription{
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 14416),
					IsForward:       true,
				},
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15452),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "SecurityGroupFolderType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "SecurityGroupFolderType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15471),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "SecurityGroupType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "SecurityGroupType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 25345),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassReferenceType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "HasPushedSecurityGroup"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "HasPushedSecurityGroup"}),
				ua.AttributeIDInverseName: ua.MustVariant(&ua.LocalizedText{Locale: "", Text: "HasPushTarget"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 25337),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "PubSubKeyPushTargetType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "PubSubKeyPushTargetType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 25346),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "PubSubKeyPushTargetFolderType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "PubSubKeyPushTargetFolderType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 14416),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "PublishSubscribeType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "PublishSubscribeType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 14443),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "PublishSubscribe"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "PublishSubscribe"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 32405),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "DataSetClasses"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "DataSetClasses"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 14476),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassReferenceType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "HasPubSubConnection"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "HasPubSubConnection"}),
				ua.AttributeIDInverseName: ua.MustVariant(&ua.LocalizedText{Locale: "", Text: "PubSubConnectionOf"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 25482),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "PubSubConfigurationType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "PubSubConfigurationType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 25517),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "PubSubConfigurationRefMask"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "PubSubConfigurationRefMask"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 25519),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "PubSubConfigurationRefDataType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "PubSubConfigurationRefDataType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 25520),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "PubSubConfigurationValueDataType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "PubSubConfigurationValueDataType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 14509),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "PublishedDataSetType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "PublishedDataSetType"}),
			},
			refs: []*ua.ReferenceDescription{
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 14572),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 14534),
					IsForward:       true,
				},
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15489),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "ExtensionFieldsType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "ExtensionFieldsType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 14936),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassReferenceType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "DataSetToWriter"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "DataSetToWriter"}),
				ua.AttributeIDInverseName: ua.MustVariant(&ua.LocalizedText{Locale: "", Text: "WriterToDataSet"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 14534),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "PublishedDataItemsType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "PublishedDataItemsType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 14572),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "PublishedEventsType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "PublishedEventsType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 14477),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "DataSetFolderType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "DataSetFolderType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 14209),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "PubSubConnectionType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "PubSubConnectionType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 17721),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "ConnectionTransportType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "ConnectionTransportType"}),
			},
			refs: []*ua.ReferenceDescription{
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 15155),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 15064),
					IsForward:       true,
				},
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 14232),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "PubSubGroupType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "PubSubGroupType"}),
			},
			refs: []*ua.ReferenceDescription{
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 17725),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 17999),
					IsForward:       true,
				},
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 17725),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "WriterGroupType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "WriterGroupType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15296),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassReferenceType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "HasDataSetWriter"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "HasDataSetWriter"}),
				ua.AttributeIDInverseName: ua.MustVariant(&ua.LocalizedText{Locale: "", Text: "IsWriterInGroup"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 18804),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassReferenceType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "HasWriterGroup"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "HasWriterGroup"}),
				ua.AttributeIDInverseName: ua.MustVariant(&ua.LocalizedText{Locale: "", Text: "IsWriterGroupOf"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 17997),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "WriterGroupTransportType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "WriterGroupTransportType"}),
			},
			refs: []*ua.ReferenceDescription{
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 21136),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 21133),
					IsForward:       true,
				},
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 17998),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "WriterGroupMessageType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "WriterGroupMessageType"}),
			},
			refs: []*ua.ReferenceDescription{
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 21105),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 21126),
					IsForward:       true,
				},
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 17999),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "ReaderGroupType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "ReaderGroupType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15297),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassReferenceType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "HasDataSetReader"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "HasDataSetReader"}),
				ua.AttributeIDInverseName: ua.MustVariant(&ua.LocalizedText{Locale: "", Text: "IsReaderInGroup"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 18805),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassReferenceType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "HasReaderGroup"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "HasReaderGroup"}),
				ua.AttributeIDInverseName: ua.MustVariant(&ua.LocalizedText{Locale: "", Text: "IsReaderGroupOf"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 21090),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "ReaderGroupTransportType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "ReaderGroupTransportType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 21091),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "ReaderGroupMessageType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "ReaderGroupMessageType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15298),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "DataSetWriterType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "DataSetWriterType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15305),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "DataSetWriterTransportType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "DataSetWriterTransportType"}),
			},
			refs: []*ua.ReferenceDescription{
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 21138),
					IsForward:       true,
				},
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 21096),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "DataSetWriterMessageType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "DataSetWriterMessageType"}),
			},
			refs: []*ua.ReferenceDescription{
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 21111),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 21128),
					IsForward:       true,
				},
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15306),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "DataSetReaderType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "DataSetReaderType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15319),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "DataSetReaderTransportType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "DataSetReaderTransportType"}),
			},
			refs: []*ua.ReferenceDescription{
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 24016),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 21142),
					IsForward:       true,
				},
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 21104),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "DataSetReaderMessageType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "DataSetReaderMessageType"}),
			},
			refs: []*ua.ReferenceDescription{
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 21116),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 21130),
					IsForward:       true,
				},
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15108),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "SubscribedDataSetType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "SubscribedDataSetType"}),
			},
			refs: []*ua.ReferenceDescription{
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 15127),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 15111),
					IsForward:       true,
				},
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15111),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "TargetVariablesType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "TargetVariablesType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15127),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "SubscribedDataSetMirrorType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "SubscribedDataSetMirrorType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 23795),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "SubscribedDataSetFolderType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "SubscribedDataSetFolderType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 23828),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "StandaloneSubscribedDataSetType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "StandaloneSubscribedDataSetType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 14643),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "PubSubStatusType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "PubSubStatusType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 19677),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "PubSubDiagnosticsType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "PubSubDiagnosticsType"}),
			},
			refs: []*ua.ReferenceDescription{
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 19968),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 19786),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 20027),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 19834),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 19903),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 19732),
					IsForward:       true,
				},
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 19723),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "DiagnosticsLevel"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "DiagnosticsLevel"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 19725),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassVariableType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "PubSubDiagnosticsCounterType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "PubSubDiagnosticsCounterType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 19730),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "PubSubDiagnosticsCounterClassification"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "PubSubDiagnosticsCounterClassification"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 19732),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "PubSubDiagnosticsRootType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "PubSubDiagnosticsRootType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 19786),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "PubSubDiagnosticsConnectionType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "PubSubDiagnosticsConnectionType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 19834),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "PubSubDiagnosticsWriterGroupType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "PubSubDiagnosticsWriterGroupType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 19903),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "PubSubDiagnosticsReaderGroupType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "PubSubDiagnosticsReaderGroupType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 19968),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "PubSubDiagnosticsDataSetWriterType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "PubSubDiagnosticsDataSetWriterType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 20027),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "PubSubDiagnosticsDataSetReaderType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "PubSubDiagnosticsDataSetReaderType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 23832),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "PubSubCapabilitiesType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "PubSubCapabilitiesType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15535),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "PubSubStatusEventType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "PubSubStatusEventType"}),
			},
			refs: []*ua.ReferenceDescription{
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 15548),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 15563),
					IsForward:       true,
				},
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15548),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "PubSubTransportLimitsExceedEventType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "PubSubTransportLimitsExceedEventType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15563),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "PubSubCommunicationFailureEventType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "PubSubCommunicationFailureEventType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 21105),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "UadpWriterGroupMessageType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "UadpWriterGroupMessageType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 21111),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "UadpDataSetWriterMessageType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "UadpDataSetWriterMessageType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 21116),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "UadpDataSetReaderMessageType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "UadpDataSetReaderMessageType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 21126),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "JsonWriterGroupMessageType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "JsonWriterGroupMessageType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 21128),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "JsonDataSetWriterMessageType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "JsonDataSetWriterMessageType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 21130),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "JsonDataSetReaderMessageType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "JsonDataSetReaderMessageType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15064),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "DatagramConnectionTransportType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "DatagramConnectionTransportType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 21133),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "DatagramWriterGroupTransportType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "DatagramWriterGroupTransportType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 24016),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "DatagramDataSetReaderTransportType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "DatagramDataSetReaderTransportType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15155),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "BrokerConnectionTransportType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "BrokerConnectionTransportType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 21136),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "BrokerWriterGroupTransportType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "BrokerWriterGroupTransportType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 21138),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "BrokerDataSetWriterTransportType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "BrokerDataSetWriterTransportType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 21142),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "BrokerDataSetReaderTransportType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "BrokerDataSetReaderTransportType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 21145),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "NetworkAddressType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "NetworkAddressType"}),
			},
			refs: []*ua.ReferenceDescription{
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 21147),
					IsForward:       true,
				},
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 21147),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "NetworkAddressUrlType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "NetworkAddressUrlType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 23455),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "AliasNameType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "AliasNameType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 23456),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "AliasNameCategoryType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "AliasNameCategoryType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 23468),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "AliasNameDataType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "AliasNameDataType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 23469),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassReferenceType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "AliasFor"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "AliasFor"}),
				ua.AttributeIDInverseName: ua.MustVariant(&ua.LocalizedText{Locale: "", Text: "HasAlias"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 23470),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Aliases"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Aliases"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 23479),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "TagVariables"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "TagVariables"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 23488),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Topics"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Topics"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 24264),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "UserManagementType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "UserManagementType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 24277),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "PasswordOptionsMask"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "PasswordOptionsMask"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 24279),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "UserConfigurationMask"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "UserConfigurationMask"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 24281),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "UserManagementDataType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "UserManagementDataType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 24290),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "UserManagement"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "UserManagement"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 19077),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassVariableType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "MultiStateDictionaryEntryDiscreteBaseType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "MultiStateDictionaryEntryDiscreteBaseType"}),
			},
			refs: []*ua.ReferenceDescription{
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 19084),
					IsForward:       true,
				},
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 19084),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassVariableType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "MultiStateDictionaryEntryDiscreteType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "MultiStateDictionaryEntryDiscreteType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 25726),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "EncodedTicket"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "EncodedTicket"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 25731),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "ApplicationConfigurationType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "ApplicationConfigurationType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 26871),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "ProvisionableDeviceType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "ProvisionableDeviceType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 29878),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "ProvisionableDevice"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "ProvisionableDevice"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 24148),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "IIetfBaseNetworkInterfaceType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "IIetfBaseNetworkInterfaceType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 24158),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "IIeeeBaseEthernetPortType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "IIeeeBaseEthernetPortType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 24233),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "IIeeeAutoNegotiationStatusType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "IIeeeAutoNegotiationStatusType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 24167),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "IBaseEthernetCapabilitiesType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "IBaseEthernetCapabilitiesType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 25218),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "IVlanIdType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "IVlanIdType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 24169),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "ISrClassType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "ISrClassType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 24173),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "IIeeeBaseTsnStreamType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "IIeeeBaseTsnStreamType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 24179),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "IIeeeBaseTsnTrafficSpecificationType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "IIeeeBaseTsnTrafficSpecificationType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 24183),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "IIeeeBaseTsnStatusStreamType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "IIeeeBaseTsnStatusStreamType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 24188),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "IIeeeTsnInterfaceConfigurationType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "IIeeeTsnInterfaceConfigurationType"}),
			},
			refs: []*ua.ReferenceDescription{
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 24191),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 24195),
					IsForward:       true,
				},
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 24191),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "IIeeeTsnInterfaceConfigurationTalkerType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "IIeeeTsnInterfaceConfigurationTalkerType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 24195),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "IIeeeTsnInterfaceConfigurationListenerType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "IIeeeTsnInterfaceConfigurationListenerType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 24199),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "IIeeeTsnMacAddressType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "IIeeeTsnMacAddressType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 24202),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "IIeeeTsnVlanTagType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "IIeeeTsnVlanTagType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 24205),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "IPriorityMappingEntryType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "IPriorityMappingEntryType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 24210),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Duplex"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Duplex"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 24212),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "InterfaceAdminStatus"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "InterfaceAdminStatus"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 24214),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "InterfaceOperStatus"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "InterfaceOperStatus"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 24216),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "NegotiationStatus"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "NegotiationStatus"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 24218),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "TsnFailureCode"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "TsnFailureCode"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 24220),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "TsnStreamState"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "TsnStreamState"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 24222),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "TsnTalkerStatus"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "TsnTalkerStatus"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 24224),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "TsnListenerStatus"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "TsnListenerStatus"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 25220),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "PriorityMappingEntryType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "PriorityMappingEntryType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 24226),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Resources"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Resources"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 24227),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Communication"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Communication"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 24228),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "MappingTables"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "MappingTables"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 24229),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "NetworkInterfaces"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "NetworkInterfaces"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 24230),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Streams"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Streams"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 24231),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "TalkerStreams"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "TalkerStreams"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 24232),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "ListenerStreams"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "ListenerStreams"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 25221),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "IetfBaseNetworkInterfaceType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "IetfBaseNetworkInterfaceType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 25227),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObjectType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "PriorityMappingTableType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "PriorityMappingTableType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 25237),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassReferenceType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "UsesPriorityMappingTable"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "UsesPriorityMappingTable"}),
				ua.AttributeIDInverseName: ua.MustVariant(&ua.LocalizedText{Locale: "", Text: "UsedByNetworkInterface"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 25238),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassReferenceType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "HasLowerLayerInterface"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "HasLowerLayerInterface"}),
				ua.AttributeIDInverseName: ua.MustVariant(&ua.LocalizedText{Locale: "", Text: "HasHigherLayerInterface"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 25253),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassReferenceType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "IsExecutableOn"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "IsExecutableOn"}),
				ua.AttributeIDInverseName: ua.MustVariant(&ua.LocalizedText{Locale: "", Text: "CanExecute"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 25254),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassReferenceType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Controls"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Controls"}),
				ua.AttributeIDInverseName: ua.MustVariant(&ua.LocalizedText{Locale: "", Text: "IsControlledBy"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 25255),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassReferenceType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Utilizes"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Utilizes"}),
				ua.AttributeIDInverseName: ua.MustVariant(&ua.LocalizedText{Locale: "", Text: "IsUtilizedBy"}),
			},
			refs: []*ua.ReferenceDescription{
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 25265),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 25261),
					IsForward:       true,
				},
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 25265),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassReferenceType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "IsExecutingOn"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "IsExecutingOn"}),
				ua.AttributeIDInverseName: ua.MustVariant(&ua.LocalizedText{Locale: "", Text: "Executes"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 25256),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassReferenceType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Requires"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Requires"}),
				ua.AttributeIDInverseName: ua.MustVariant(&ua.LocalizedText{Locale: "", Text: "IsRequiredBy"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 25257),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassReferenceType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "IsPhysicallyConnectedTo"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "IsPhysicallyConnectedTo"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 25258),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassReferenceType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "RepresentsSameEntityAs"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "RepresentsSameEntityAs"}),
			},
			refs: []*ua.ReferenceDescription{
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 25260),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 25259),
					IsForward:       true,
				},
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 25259),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassReferenceType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "RepresentsSameHardwareAs"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "RepresentsSameHardwareAs"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 25260),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassReferenceType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "RepresentsSameFunctionalityAs"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "RepresentsSameFunctionalityAs"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 25261),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassReferenceType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "IsHostedBy"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "IsHostedBy"}),
				ua.AttributeIDInverseName: ua.MustVariant(&ua.LocalizedText{Locale: "", Text: "Hosts"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 25262),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassReferenceType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "HasPhysicalComponent"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "HasPhysicalComponent"}),
				ua.AttributeIDInverseName: ua.MustVariant(&ua.LocalizedText{Locale: "", Text: "PhysicalComponentOf"}),
			},
			refs: []*ua.ReferenceDescription{
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 25264),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 25263),
					IsForward:       true,
				},
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 25263),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassReferenceType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "HasContainedComponent"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "HasContainedComponent"}),
				ua.AttributeIDInverseName: ua.MustVariant(&ua.LocalizedText{Locale: "", Text: "ContainedComponentOf"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 25264),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassReferenceType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "HasAttachedComponent"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "HasAttachedComponent"}),
				ua.AttributeIDInverseName: ua.MustVariant(&ua.LocalizedText{Locale: "", Text: "AttachedComponentOf"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 32679),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassReferenceType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "HasReferenceDescription"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "HasReferenceDescription"}),
				ua.AttributeIDInverseName: ua.MustVariant(&ua.LocalizedText{Locale: "", Text: "ReferenceDescriptionOf"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 32657),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassVariableType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "ReferenceDescriptionVariableType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "ReferenceDescriptionVariableType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 32659),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "ReferenceDescriptionDataType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "ReferenceDescriptionDataType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 32660),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "ReferenceListEntryDataType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "ReferenceListEntryDataType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 256),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "IdType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "IdType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 257),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "NodeClass"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "NodeClass"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 94),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "PermissionType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "PermissionType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15031),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "AccessLevelType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "AccessLevelType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15406),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "AccessLevelExType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "AccessLevelExType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15033),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "EventNotifierType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "EventNotifierType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 95),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "AccessRestrictionType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "AccessRestrictionType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 96),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "RolePermissionType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "RolePermissionType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 97),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "DataTypeDefinition"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "DataTypeDefinition"}),
			},
			refs: []*ua.ReferenceDescription{
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 100),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 99),
					IsForward:       true,
				},
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 98),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "StructureType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "StructureType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 101),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "StructureField"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "StructureField"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 99),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "StructureDefinition"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "StructureDefinition"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 100),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "EnumDefinition"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "EnumDefinition"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 296),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Argument"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Argument"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 7594),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "EnumValueType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "EnumValueType"}),
			},
			refs: []*ua.ReferenceDescription{
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 102),
					IsForward:       true,
				},
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 102),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "EnumField"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "EnumField"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 12755),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "OptionSet"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "OptionSet"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 12877),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "NormalizedString"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "NormalizedString"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 12878),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "DecimalString"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "DecimalString"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 12879),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "DurationString"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "DurationString"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 12880),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "TimeString"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "TimeString"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 12881),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "DateString"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "DateString"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 290),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Duration"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Duration"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 294),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "UtcTime"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "UtcTime"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 295),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "LocaleId"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "LocaleId"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 8912),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "TimeZoneDataType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "TimeZoneDataType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 17588),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Index"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Index"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 288),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "IntegerId"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "IntegerId"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 307),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "ApplicationType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "ApplicationType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 308),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "ApplicationDescription"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "ApplicationDescription"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 20998),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "VersionTime"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "VersionTime"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 12189),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "ServerOnNetwork"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "ServerOnNetwork"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 311),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "ApplicationInstanceCertificate"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "ApplicationInstanceCertificate"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 302),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "MessageSecurityMode"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "MessageSecurityMode"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 303),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "UserTokenType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "UserTokenType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 304),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "UserTokenPolicy"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "UserTokenPolicy"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 312),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "EndpointDescription"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "EndpointDescription"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 432),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "RegisteredServer"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "RegisteredServer"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 12890),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "DiscoveryConfiguration"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "DiscoveryConfiguration"}),
			},
			refs: []*ua.ReferenceDescription{
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 12891),
					IsForward:       true,
				},
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 12891),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "MdnsDiscoveryConfiguration"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "MdnsDiscoveryConfiguration"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 315),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "SecurityTokenRequestType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "SecurityTokenRequestType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 344),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "SignedSoftwareCertificate"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "SignedSoftwareCertificate"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 388),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "SessionAuthenticationToken"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "SessionAuthenticationToken"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 316),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "UserIdentityToken"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "UserIdentityToken"}),
			},
			refs: []*ua.ReferenceDescription{
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 322),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 938),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 319),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 325),
					IsForward:       true,
				},
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 319),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "AnonymousIdentityToken"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "AnonymousIdentityToken"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 322),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "UserNameIdentityToken"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "UserNameIdentityToken"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 325),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "X509IdentityToken"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "X509IdentityToken"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 938),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "IssuedIdentityToken"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "IssuedIdentityToken"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 348),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "NodeAttributesMask"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "NodeAttributesMask"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 376),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "AddNodesItem"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "AddNodesItem"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 379),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "AddReferencesItem"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "AddReferencesItem"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 382),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "DeleteNodesItem"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "DeleteNodesItem"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 385),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "DeleteReferencesItem"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "DeleteReferencesItem"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 347),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "AttributeWriteMask"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "AttributeWriteMask"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 521),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "ContinuationPoint"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "ContinuationPoint"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 537),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "RelativePathElement"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "RelativePathElement"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 540),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "RelativePath"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "RelativePath"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 289),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Counter"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Counter"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 291),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "NumericRange"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "NumericRange"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 331),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "EndpointConfiguration"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "EndpointConfiguration"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 576),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "FilterOperator"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "FilterOperator"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 583),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "ContentFilterElement"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "ContentFilterElement"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 586),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "ContentFilter"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "ContentFilter"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 589),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "FilterOperand"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "FilterOperand"}),
			},
			refs: []*ua.ReferenceDescription{
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 598),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 595),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 601),
					IsForward:       true,
				},
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 592),
					IsForward:       true,
				},
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 592),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "ElementOperand"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "ElementOperand"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 595),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "LiteralOperand"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "LiteralOperand"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 598),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "AttributeOperand"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "AttributeOperand"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 601),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "SimpleAttributeOperand"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "SimpleAttributeOperand"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 659),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "HistoryEvent"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "HistoryEvent"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 11234),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "HistoryUpdateType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "HistoryUpdateType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 11293),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "PerformUpdateType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "PerformUpdateType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 719),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "MonitoringFilter"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "MonitoringFilter"}),
			},
			refs: []*ua.ReferenceDescription{
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
					TypeDefinition:  ua.NewNumericExpandedNodeID(0, 725),
					IsForward:       true,
				},
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 725),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "EventFilter"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "EventFilter"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 948),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "AggregateConfiguration"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "AggregateConfiguration"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 920),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "HistoryEventFieldList"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "HistoryEventFieldList"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 338),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "BuildInfo"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "BuildInfo"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 851),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "RedundancySupport"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "RedundancySupport"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 852),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "ServerState"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "ServerState"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 853),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "RedundantServerDataType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "RedundantServerDataType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 11943),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "EndpointUrlListDataType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "EndpointUrlListDataType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 11944),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "NetworkGroupDataType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "NetworkGroupDataType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 856),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "SamplingIntervalDiagnosticsDataType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "SamplingIntervalDiagnosticsDataType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 859),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "ServerDiagnosticsSummaryDataType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "ServerDiagnosticsSummaryDataType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 862),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "ServerStatusDataType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "ServerStatusDataType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 865),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "SessionDiagnosticsDataType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "SessionDiagnosticsDataType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 868),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "SessionSecurityDiagnosticsDataType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "SessionSecurityDiagnosticsDataType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 871),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "ServiceCounterDataType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "ServiceCounterDataType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 299),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "StatusResult"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "StatusResult"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 874),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "SubscriptionDiagnosticsDataType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "SubscriptionDiagnosticsDataType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 877),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "ModelChangeStructureDataType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "ModelChangeStructureDataType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 897),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "SemanticChangeStructureDataType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "SemanticChangeStructureDataType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 884),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Range"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Range"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 887),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "EUInformation"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "EUInformation"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 12077),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "AxisScaleEnumeration"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "AxisScaleEnumeration"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 12171),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "ComplexNumberType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "ComplexNumberType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 12172),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "DoubleComplexNumberType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "DoubleComplexNumberType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 12079),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "AxisInformation"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "AxisInformation"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 12080),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "XVType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "XVType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 894),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "ProgramDiagnosticDataType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "ProgramDiagnosticDataType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 24033),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "ProgramDiagnostic2DataType"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "ProgramDiagnostic2DataType"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 891),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Annotation"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Annotation"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 890),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassDataType),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "ExceptionDeviationFormat"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "ExceptionDeviationFormat"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 12766),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default Binary"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default Binary"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 14846),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default Binary"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default Binary"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 17537),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default Binary"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default Binary"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 17549),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default Binary"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default Binary"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15671),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default Binary"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default Binary"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 18815),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default Binary"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default Binary"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 18816),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default Binary"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default Binary"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 18817),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default Binary"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default Binary"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 18818),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default Binary"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default Binary"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 18819),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default Binary"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default Binary"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 18820),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default Binary"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default Binary"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 18821),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default Binary"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default Binary"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 18822),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default Binary"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default Binary"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 18823),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default Binary"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default Binary"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15736),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default Binary"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default Binary"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 23507),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default Binary"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default Binary"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 12680),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default Binary"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default Binary"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 32382),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default Binary"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default Binary"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15676),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default Binary"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default Binary"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 125),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default Binary"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default Binary"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 126),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default Binary"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default Binary"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 127),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default Binary"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default Binary"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15421),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default Binary"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default Binary"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15422),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default Binary"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default Binary"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 24108),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default Binary"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default Binary"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 24109),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default Binary"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default Binary"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 24110),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default Binary"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default Binary"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 124),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default Binary"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default Binary"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 14839),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default Binary"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default Binary"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 14847),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default Binary"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default Binary"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15677),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default Binary"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default Binary"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15678),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default Binary"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default Binary"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 14323),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default Binary"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default Binary"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15679),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default Binary"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default Binary"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15681),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default Binary"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default Binary"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 25529),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default Binary"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default Binary"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15682),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default Binary"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default Binary"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15683),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default Binary"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default Binary"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15688),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default Binary"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default Binary"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15689),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default Binary"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default Binary"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 21150),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default Binary"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default Binary"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15691),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default Binary"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default Binary"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15693),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default Binary"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default Binary"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15694),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default Binary"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default Binary"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15695),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default Binary"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default Binary"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 21151),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default Binary"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default Binary"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 21152),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default Binary"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default Binary"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 21153),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default Binary"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default Binary"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15701),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default Binary"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default Binary"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15702),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default Binary"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default Binary"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15703),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default Binary"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default Binary"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15705),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default Binary"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default Binary"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15706),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default Binary"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default Binary"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15707),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default Binary"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default Binary"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15712),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default Binary"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default Binary"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 14848),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default Binary"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default Binary"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15713),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default Binary"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default Binary"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 21154),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default Binary"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default Binary"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 23851),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default Binary"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default Binary"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 23852),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default Binary"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default Binary"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 23853),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default Binary"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default Binary"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 25530),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default Binary"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default Binary"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 23854),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default Binary"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default Binary"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15715),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default Binary"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default Binary"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15717),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default Binary"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default Binary"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15718),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default Binary"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default Binary"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15719),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default Binary"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default Binary"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15724),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default Binary"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default Binary"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15725),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default Binary"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default Binary"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 23855),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default Binary"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default Binary"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 23856),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default Binary"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default Binary"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 23857),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default Binary"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default Binary"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 23860),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default Binary"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default Binary"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 23861),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default Binary"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default Binary"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 17468),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default Binary"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default Binary"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 23864),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default Binary"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default Binary"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 21155),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default Binary"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default Binary"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 23865),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default Binary"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default Binary"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 23866),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default Binary"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default Binary"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15479),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default Binary"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default Binary"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15727),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default Binary"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default Binary"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15729),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default Binary"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default Binary"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15733),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default Binary"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default Binary"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 25531),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default Binary"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default Binary"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 25532),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default Binary"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default Binary"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 23499),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default Binary"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default Binary"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 24292),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default Binary"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default Binary"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 25239),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default Binary"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default Binary"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 32661),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default Binary"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default Binary"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 32662),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default Binary"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default Binary"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 128),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default Binary"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default Binary"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 121),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default Binary"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default Binary"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 14844),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default Binary"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default Binary"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 122),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default Binary"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default Binary"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 123),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default Binary"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default Binary"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 298),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default Binary"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default Binary"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 8251),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default Binary"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default Binary"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 14845),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default Binary"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default Binary"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 12765),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default Binary"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default Binary"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 8917),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default Binary"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default Binary"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 310),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default Binary"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default Binary"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 12207),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default Binary"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default Binary"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 306),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default Binary"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default Binary"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 314),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default Binary"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default Binary"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 434),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default Binary"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default Binary"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 12900),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default Binary"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default Binary"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 12901),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default Binary"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default Binary"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 346),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default Binary"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default Binary"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 318),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default Binary"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default Binary"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 321),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default Binary"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default Binary"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 324),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default Binary"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default Binary"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 327),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default Binary"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default Binary"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 940),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default Binary"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default Binary"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 378),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default Binary"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default Binary"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 381),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default Binary"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default Binary"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 384),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default Binary"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default Binary"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 387),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default Binary"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default Binary"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 539),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default Binary"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default Binary"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 542),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default Binary"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default Binary"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 333),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default Binary"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default Binary"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 585),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default Binary"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default Binary"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 588),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default Binary"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default Binary"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 591),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default Binary"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default Binary"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 594),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default Binary"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default Binary"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 597),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default Binary"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default Binary"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 600),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default Binary"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default Binary"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 603),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default Binary"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default Binary"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 661),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default Binary"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default Binary"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 721),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default Binary"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default Binary"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 727),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default Binary"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default Binary"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 950),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default Binary"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default Binary"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 922),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default Binary"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default Binary"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 340),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default Binary"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default Binary"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 855),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default Binary"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default Binary"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 11957),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default Binary"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default Binary"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 11958),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default Binary"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default Binary"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 858),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default Binary"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default Binary"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 861),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default Binary"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default Binary"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 864),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default Binary"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default Binary"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 867),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default Binary"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default Binary"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 870),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default Binary"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default Binary"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 873),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default Binary"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default Binary"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 301),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default Binary"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default Binary"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 876),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default Binary"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default Binary"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 879),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default Binary"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default Binary"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 899),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default Binary"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default Binary"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 886),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default Binary"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default Binary"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 889),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default Binary"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default Binary"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 12181),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default Binary"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default Binary"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 12182),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default Binary"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default Binary"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 12089),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default Binary"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default Binary"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 12090),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default Binary"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default Binary"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 896),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default Binary"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default Binary"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 24034),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default Binary"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default Binary"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 893),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default Binary"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default Binary"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 7617),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassVariable),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Opc.Ua"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Opc.Ua"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 12758),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default XML"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default XML"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 14802),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default XML"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default XML"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 17541),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default XML"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default XML"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 17553),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default XML"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default XML"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15949),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default XML"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default XML"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 18851),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default XML"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default XML"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 18852),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default XML"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default XML"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 18853),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default XML"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default XML"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 18854),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default XML"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default XML"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 18855),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default XML"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default XML"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 18856),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default XML"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default XML"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 18857),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default XML"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default XML"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 18858),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default XML"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default XML"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 18859),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default XML"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default XML"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15728),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default XML"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default XML"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 23520),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default XML"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default XML"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 12676),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default XML"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default XML"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 32386),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default XML"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default XML"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15950),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default XML"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default XML"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 14796),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default XML"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default XML"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15589),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default XML"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default XML"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15590),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default XML"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default XML"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15529),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default XML"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default XML"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15531),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default XML"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default XML"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 24120),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default XML"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default XML"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 24121),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default XML"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default XML"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 24122),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default XML"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default XML"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 14794),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default XML"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default XML"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 14795),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default XML"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default XML"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 14803),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default XML"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default XML"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15951),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default XML"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default XML"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15952),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default XML"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default XML"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 14319),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default XML"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default XML"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15953),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default XML"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default XML"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15954),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default XML"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default XML"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 25545),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default XML"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default XML"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15955),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default XML"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default XML"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15956),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default XML"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default XML"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15987),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default XML"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default XML"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15988),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default XML"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default XML"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 21174),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default XML"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default XML"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15990),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default XML"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default XML"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15991),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default XML"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default XML"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15992),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default XML"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default XML"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15993),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default XML"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default XML"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 21175),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default XML"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default XML"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 21176),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default XML"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default XML"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 21177),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default XML"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default XML"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15995),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default XML"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default XML"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15996),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default XML"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default XML"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 16007),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default XML"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default XML"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 16008),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default XML"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default XML"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 16009),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default XML"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default XML"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 16010),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default XML"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default XML"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 16011),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default XML"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default XML"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 14804),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default XML"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default XML"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 16012),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default XML"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default XML"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 21178),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default XML"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default XML"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 23919),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default XML"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default XML"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 23920),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default XML"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default XML"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 23921),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default XML"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default XML"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 25546),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default XML"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default XML"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 23922),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default XML"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default XML"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 16014),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default XML"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default XML"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 16015),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default XML"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default XML"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 16016),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default XML"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default XML"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 16017),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default XML"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default XML"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 16018),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default XML"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default XML"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 16019),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default XML"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default XML"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 23923),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default XML"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default XML"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 23924),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default XML"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default XML"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 23925),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default XML"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default XML"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 23928),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default XML"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default XML"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 23929),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default XML"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default XML"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 17472),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default XML"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default XML"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 23932),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default XML"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default XML"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 21179),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default XML"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default XML"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 23933),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default XML"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default XML"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 23934),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default XML"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default XML"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15579),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default XML"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default XML"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 16021),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default XML"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default XML"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 16022),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default XML"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default XML"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 16023),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default XML"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default XML"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 25547),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default XML"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default XML"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 25548),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default XML"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default XML"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 23505),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default XML"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default XML"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 24296),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default XML"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default XML"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 25243),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default XML"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default XML"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 32669),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default XML"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default XML"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 32670),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default XML"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default XML"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 16126),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default XML"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default XML"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 14797),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default XML"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default XML"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 14800),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default XML"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default XML"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 14798),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default XML"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default XML"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 14799),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default XML"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default XML"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 297),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default XML"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default XML"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 7616),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default XML"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default XML"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 14801),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default XML"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default XML"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 12757),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default XML"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default XML"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 8913),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default XML"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default XML"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 309),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default XML"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default XML"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 12195),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default XML"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default XML"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 305),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default XML"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default XML"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 313),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default XML"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default XML"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 433),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default XML"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default XML"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 12892),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default XML"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default XML"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 12893),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default XML"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default XML"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 345),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default XML"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default XML"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 317),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default XML"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default XML"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 320),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default XML"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default XML"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 323),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default XML"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default XML"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 326),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default XML"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default XML"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 939),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default XML"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default XML"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 377),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default XML"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default XML"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 380),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default XML"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default XML"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 383),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default XML"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default XML"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 386),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default XML"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default XML"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 538),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default XML"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default XML"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 541),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default XML"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default XML"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 332),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default XML"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default XML"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 584),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default XML"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default XML"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 587),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default XML"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default XML"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 590),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default XML"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default XML"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 593),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default XML"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default XML"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 596),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default XML"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default XML"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 599),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default XML"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default XML"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 602),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default XML"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default XML"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 660),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default XML"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default XML"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 720),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default XML"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default XML"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 726),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default XML"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default XML"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 949),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default XML"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default XML"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 921),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default XML"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default XML"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 339),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default XML"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default XML"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 854),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default XML"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default XML"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 11949),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default XML"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default XML"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 11950),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default XML"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default XML"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 857),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default XML"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default XML"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 860),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default XML"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default XML"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 863),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default XML"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default XML"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 866),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default XML"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default XML"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 869),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default XML"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default XML"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 872),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default XML"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default XML"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 300),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default XML"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default XML"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 875),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default XML"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default XML"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 878),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default XML"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default XML"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 898),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default XML"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default XML"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 885),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default XML"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default XML"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 888),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default XML"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default XML"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 12173),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default XML"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default XML"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 12174),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default XML"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default XML"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 12081),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default XML"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default XML"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 12082),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default XML"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default XML"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 895),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default XML"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default XML"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 24038),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default XML"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default XML"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 892),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default XML"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default XML"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 8252),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassVariable),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Opc.Ua"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Opc.Ua"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15085),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default JSON"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default JSON"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15041),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default JSON"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default JSON"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 17547),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default JSON"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default JSON"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 17557),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default JSON"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default JSON"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 16150),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default JSON"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default JSON"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 19064),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default JSON"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default JSON"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 19065),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default JSON"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default JSON"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 19066),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default JSON"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default JSON"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 19067),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default JSON"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default JSON"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 19068),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default JSON"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default JSON"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 19069),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default JSON"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default JSON"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 19070),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default JSON"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default JSON"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 19071),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default JSON"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default JSON"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 19072),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default JSON"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default JSON"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15042),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default JSON"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default JSON"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 23528),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default JSON"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default JSON"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15044),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default JSON"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default JSON"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 32390),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default JSON"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default JSON"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 16151),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default JSON"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default JSON"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15057),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default JSON"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default JSON"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15058),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default JSON"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default JSON"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15059),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default JSON"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default JSON"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15700),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default JSON"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default JSON"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15714),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default JSON"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default JSON"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 24132),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default JSON"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default JSON"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 24133),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default JSON"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default JSON"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 24134),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default JSON"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default JSON"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15050),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default JSON"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default JSON"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15051),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default JSON"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default JSON"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15049),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default JSON"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default JSON"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 16152),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default JSON"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default JSON"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 16153),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default JSON"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default JSON"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15060),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default JSON"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default JSON"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 16154),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default JSON"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default JSON"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 16155),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default JSON"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default JSON"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 25561),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default JSON"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default JSON"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 16156),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default JSON"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default JSON"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 16157),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default JSON"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default JSON"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 16158),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default JSON"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default JSON"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 16159),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default JSON"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default JSON"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 21198),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default JSON"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default JSON"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 16161),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default JSON"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default JSON"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 16280),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default JSON"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default JSON"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 16281),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default JSON"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default JSON"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 16282),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default JSON"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default JSON"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 21199),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default JSON"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default JSON"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 21200),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default JSON"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default JSON"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 21201),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default JSON"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default JSON"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 16284),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default JSON"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default JSON"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 16285),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default JSON"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default JSON"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 16286),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default JSON"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default JSON"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 16287),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default JSON"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default JSON"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 16288),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default JSON"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default JSON"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 16308),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default JSON"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default JSON"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 16310),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default JSON"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default JSON"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15061),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default JSON"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default JSON"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 16311),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default JSON"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default JSON"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 21202),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default JSON"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default JSON"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 23987),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default JSON"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default JSON"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 23988),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default JSON"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default JSON"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 23989),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default JSON"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default JSON"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 25562),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default JSON"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default JSON"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 23990),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default JSON"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default JSON"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 16323),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default JSON"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default JSON"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 16391),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default JSON"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default JSON"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 16392),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default JSON"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default JSON"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 16393),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default JSON"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default JSON"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 16394),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default JSON"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default JSON"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 16404),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default JSON"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default JSON"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 23991),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default JSON"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default JSON"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 23992),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default JSON"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default JSON"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 23993),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default JSON"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default JSON"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 23996),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default JSON"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default JSON"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 23997),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default JSON"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default JSON"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 17476),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default JSON"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default JSON"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 24000),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default JSON"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default JSON"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 21203),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default JSON"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default JSON"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 24001),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default JSON"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default JSON"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 24002),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default JSON"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default JSON"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15726),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default JSON"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default JSON"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 16524),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default JSON"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default JSON"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 16525),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default JSON"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default JSON"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 16526),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default JSON"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default JSON"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 25563),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default JSON"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default JSON"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 25564),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default JSON"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default JSON"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 23511),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default JSON"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default JSON"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 24300),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default JSON"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default JSON"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 25247),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default JSON"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default JSON"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 32677),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default JSON"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default JSON"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 32678),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default JSON"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default JSON"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15062),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default JSON"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default JSON"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15063),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default JSON"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default JSON"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15065),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default JSON"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default JSON"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15066),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default JSON"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default JSON"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15067),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default JSON"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default JSON"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15081),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default JSON"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default JSON"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15082),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default JSON"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default JSON"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15083),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default JSON"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default JSON"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15084),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default JSON"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default JSON"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15086),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default JSON"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default JSON"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15087),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default JSON"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default JSON"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15095),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default JSON"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default JSON"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15098),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default JSON"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default JSON"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15099),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default JSON"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default JSON"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15102),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default JSON"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default JSON"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15105),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default JSON"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default JSON"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15106),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default JSON"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default JSON"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15136),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default JSON"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default JSON"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15140),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default JSON"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default JSON"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15141),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default JSON"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default JSON"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15142),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default JSON"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default JSON"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15143),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default JSON"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default JSON"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15144),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default JSON"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default JSON"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15165),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default JSON"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default JSON"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15169),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default JSON"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default JSON"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15172),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default JSON"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default JSON"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15175),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default JSON"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default JSON"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15188),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default JSON"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default JSON"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15189),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default JSON"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default JSON"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15199),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default JSON"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default JSON"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15204),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default JSON"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default JSON"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15205),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default JSON"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default JSON"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15206),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default JSON"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default JSON"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15207),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default JSON"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default JSON"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15208),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default JSON"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default JSON"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15209),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default JSON"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default JSON"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15210),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default JSON"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default JSON"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15273),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default JSON"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default JSON"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15293),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default JSON"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default JSON"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15295),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default JSON"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default JSON"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15304),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default JSON"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default JSON"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15349),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default JSON"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default JSON"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15361),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default JSON"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default JSON"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15362),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default JSON"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default JSON"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15363),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default JSON"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default JSON"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15364),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default JSON"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default JSON"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15365),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default JSON"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default JSON"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15366),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default JSON"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default JSON"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15367),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default JSON"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default JSON"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15368),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default JSON"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default JSON"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15369),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default JSON"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default JSON"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15370),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default JSON"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default JSON"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15371),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default JSON"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default JSON"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15372),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default JSON"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default JSON"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15373),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default JSON"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default JSON"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15374),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default JSON"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default JSON"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15375),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default JSON"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default JSON"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15376),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default JSON"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default JSON"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15377),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default JSON"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default JSON"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15378),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default JSON"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default JSON"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15379),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default JSON"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default JSON"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15380),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default JSON"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default JSON"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15381),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default JSON"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default JSON"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 24042),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default JSON"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default JSON"}),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15382),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant(ua.NodeClassObject),
				ua.AttributeIDBrowseName:  ua.MustVariant(&ua.QualifiedName{Name: "Default JSON"}),
				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text: "Default JSON"}),
			},
		},
	}
}
