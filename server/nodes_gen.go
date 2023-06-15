// Generated code. DO NOT EDIT
// Copyright 2018-2023 opcua authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.
package server

import "github.com/gopcua/opcua/ua"

func PredefinedNodes() []Node {
	return []Node{
		&node{
			id: ua.NewNumericNodeID(0, 3062),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default Binary"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default Binary"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 3063),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default XML"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default XML"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 1),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Boolean"),
				ua.AttributeIDDisplayName: ua.MustVariant("Boolean"),
			},
			superTypeID: ua.NewNumericNodeID(0, 24),
		},
		&node{
			id: ua.NewNumericNodeID(0, 2),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("SByte"),
				ua.AttributeIDDisplayName: ua.MustVariant("SByte"),
			},
			superTypeID: ua.NewNumericNodeID(0, 27),
		},
		&node{
			id: ua.NewNumericNodeID(0, 3),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Byte"),
				ua.AttributeIDDisplayName: ua.MustVariant("Byte"),
			},
			superTypeID: ua.NewNumericNodeID(0, 28),
		},
		&node{
			id: ua.NewNumericNodeID(0, 4),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Int16"),
				ua.AttributeIDDisplayName: ua.MustVariant("Int16"),
			},
			superTypeID: ua.NewNumericNodeID(0, 27),
		},
		&node{
			id: ua.NewNumericNodeID(0, 5),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("UInt16"),
				ua.AttributeIDDisplayName: ua.MustVariant("UInt16"),
			},
			superTypeID: ua.NewNumericNodeID(0, 28),
		},
		&node{
			id: ua.NewNumericNodeID(0, 6),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Int32"),
				ua.AttributeIDDisplayName: ua.MustVariant("Int32"),
			},
			superTypeID: ua.NewNumericNodeID(0, 27),
		},
		&node{
			id: ua.NewNumericNodeID(0, 7),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("UInt32"),
				ua.AttributeIDDisplayName: ua.MustVariant("UInt32"),
			},
			superTypeID: ua.NewNumericNodeID(0, 28),
		},
		&node{
			id: ua.NewNumericNodeID(0, 8),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Int64"),
				ua.AttributeIDDisplayName: ua.MustVariant("Int64"),
			},
			superTypeID: ua.NewNumericNodeID(0, 27),
		},
		&node{
			id: ua.NewNumericNodeID(0, 9),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("UInt64"),
				ua.AttributeIDDisplayName: ua.MustVariant("UInt64"),
			},
			superTypeID: ua.NewNumericNodeID(0, 28),
		},
		&node{
			id: ua.NewNumericNodeID(0, 10),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Float"),
				ua.AttributeIDDisplayName: ua.MustVariant("Float"),
			},
			superTypeID: ua.NewNumericNodeID(0, 26),
		},
		&node{
			id: ua.NewNumericNodeID(0, 11),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Double"),
				ua.AttributeIDDisplayName: ua.MustVariant("Double"),
			},
			superTypeID: ua.NewNumericNodeID(0, 26),
		},
		&node{
			id: ua.NewNumericNodeID(0, 12),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("String"),
				ua.AttributeIDDisplayName: ua.MustVariant("String"),
			},
			superTypeID: ua.NewNumericNodeID(0, 24),
		},
		&node{
			id: ua.NewNumericNodeID(0, 13),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("DateTime"),
				ua.AttributeIDDisplayName: ua.MustVariant("DateTime"),
			},
			superTypeID: ua.NewNumericNodeID(0, 24),
		},
		&node{
			id: ua.NewNumericNodeID(0, 14),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Guid"),
				ua.AttributeIDDisplayName: ua.MustVariant("Guid"),
			},
			superTypeID: ua.NewNumericNodeID(0, 24),
		},
		&node{
			id: ua.NewNumericNodeID(0, 15),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("ByteString"),
				ua.AttributeIDDisplayName: ua.MustVariant("ByteString"),
			},
			superTypeID: ua.NewNumericNodeID(0, 24),
		},
		&node{
			id: ua.NewNumericNodeID(0, 16),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("XmlElement"),
				ua.AttributeIDDisplayName: ua.MustVariant("XmlElement"),
			},
			superTypeID: ua.NewNumericNodeID(0, 24),
		},
		&node{
			id: ua.NewNumericNodeID(0, 17),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("NodeId"),
				ua.AttributeIDDisplayName: ua.MustVariant("NodeId"),
			},
			superTypeID: ua.NewNumericNodeID(0, 24),
		},
		&node{
			id: ua.NewNumericNodeID(0, 18),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("ExpandedNodeId"),
				ua.AttributeIDDisplayName: ua.MustVariant("ExpandedNodeId"),
			},
			superTypeID: ua.NewNumericNodeID(0, 24),
		},
		&node{
			id: ua.NewNumericNodeID(0, 19),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("StatusCode"),
				ua.AttributeIDDisplayName: ua.MustVariant("StatusCode"),
			},
			superTypeID: ua.NewNumericNodeID(0, 24),
		},
		&node{
			id: ua.NewNumericNodeID(0, 20),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("QualifiedName"),
				ua.AttributeIDDisplayName: ua.MustVariant("QualifiedName"),
			},
			superTypeID: ua.NewNumericNodeID(0, 24),
		},
		&node{
			id: ua.NewNumericNodeID(0, 21),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("LocalizedText"),
				ua.AttributeIDDisplayName: ua.MustVariant("LocalizedText"),
			},
			superTypeID: ua.NewNumericNodeID(0, 24),
		},
		&node{
			id: ua.NewNumericNodeID(0, 23),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("DataValue"),
				ua.AttributeIDDisplayName: ua.MustVariant("DataValue"),
			},
			superTypeID: ua.NewNumericNodeID(0, 24),
		},
		&node{
			id: ua.NewNumericNodeID(0, 25),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("DiagnosticInfo"),
				ua.AttributeIDDisplayName: ua.MustVariant("DiagnosticInfo"),
			},
			superTypeID: ua.NewNumericNodeID(0, 24),
		},
		&node{
			id: ua.NewNumericNodeID(0, 50),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Decimal"),
				ua.AttributeIDDisplayName: ua.MustVariant("Decimal"),
			},
			superTypeID: ua.NewNumericNodeID(0, 26),
		},
		&node{
			id: ua.NewNumericNodeID(0, 35),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ReferenceType_32"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Organizes"),
				ua.AttributeIDDisplayName: ua.MustVariant("Organizes"),
				ua.AttributeIDInverseName: ua.MustVariant(&ua.LocalizedText{Locale: "", Text: "OrganizedBy"}),
			},
			superTypeID: ua.NewNumericNodeID(0, 33),
		},
		&node{
			id: ua.NewNumericNodeID(0, 36),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ReferenceType_32"),
				ua.AttributeIDBrowseName:  ua.MustVariant("HasEventSource"),
				ua.AttributeIDDisplayName: ua.MustVariant("HasEventSource"),
				ua.AttributeIDInverseName: ua.MustVariant(&ua.LocalizedText{Locale: "", Text: "EventSourceOf"}),
			},
			superTypeID: ua.NewNumericNodeID(0, 33),
		},
		&node{
			id: ua.NewNumericNodeID(0, 37),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ReferenceType_32"),
				ua.AttributeIDBrowseName:  ua.MustVariant("HasModellingRule"),
				ua.AttributeIDDisplayName: ua.MustVariant("HasModellingRule"),
				ua.AttributeIDInverseName: ua.MustVariant(&ua.LocalizedText{Locale: "", Text: "ModellingRuleOf"}),
			},
			superTypeID: ua.NewNumericNodeID(0, 32),
		},
		&node{
			id: ua.NewNumericNodeID(0, 38),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ReferenceType_32"),
				ua.AttributeIDBrowseName:  ua.MustVariant("HasEncoding"),
				ua.AttributeIDDisplayName: ua.MustVariant("HasEncoding"),
				ua.AttributeIDInverseName: ua.MustVariant(&ua.LocalizedText{Locale: "", Text: "EncodingOf"}),
			},
			superTypeID: ua.NewNumericNodeID(0, 32),
		},
		&node{
			id: ua.NewNumericNodeID(0, 39),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ReferenceType_32"),
				ua.AttributeIDBrowseName:  ua.MustVariant("HasDescription"),
				ua.AttributeIDDisplayName: ua.MustVariant("HasDescription"),
				ua.AttributeIDInverseName: ua.MustVariant(&ua.LocalizedText{Locale: "", Text: "DescriptionOf"}),
			},
			superTypeID: ua.NewNumericNodeID(0, 32),
		},
		&node{
			id: ua.NewNumericNodeID(0, 40),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ReferenceType_32"),
				ua.AttributeIDBrowseName:  ua.MustVariant("HasTypeDefinition"),
				ua.AttributeIDDisplayName: ua.MustVariant("HasTypeDefinition"),
				ua.AttributeIDInverseName: ua.MustVariant(&ua.LocalizedText{Locale: "", Text: "TypeDefinitionOf"}),
			},
			superTypeID: ua.NewNumericNodeID(0, 32),
		},
		&node{
			id: ua.NewNumericNodeID(0, 41),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ReferenceType_32"),
				ua.AttributeIDBrowseName:  ua.MustVariant("GeneratesEvent"),
				ua.AttributeIDDisplayName: ua.MustVariant("GeneratesEvent"),
				ua.AttributeIDInverseName: ua.MustVariant(&ua.LocalizedText{Locale: "", Text: "GeneratedBy"}),
			},
			superTypeID: ua.NewNumericNodeID(0, 32),
		},
		&node{
			id: ua.NewNumericNodeID(0, 3065),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ReferenceType_32"),
				ua.AttributeIDBrowseName:  ua.MustVariant("AlwaysGeneratesEvent"),
				ua.AttributeIDDisplayName: ua.MustVariant("AlwaysGeneratesEvent"),
				ua.AttributeIDInverseName: ua.MustVariant(&ua.LocalizedText{Locale: "", Text: "AlwaysGeneratedBy"}),
			},
			superTypeID: ua.NewNumericNodeID(0, 41),
		},
		&node{
			id: ua.NewNumericNodeID(0, 45),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ReferenceType_32"),
				ua.AttributeIDBrowseName:  ua.MustVariant("HasSubtype"),
				ua.AttributeIDDisplayName: ua.MustVariant("HasSubtype"),
				ua.AttributeIDInverseName: ua.MustVariant(&ua.LocalizedText{Locale: "", Text: "SubtypeOf"}),
			},
			superTypeID: ua.NewNumericNodeID(0, 34),
		},
		&node{
			id: ua.NewNumericNodeID(0, 46),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ReferenceType_32"),
				ua.AttributeIDBrowseName:  ua.MustVariant("HasProperty"),
				ua.AttributeIDDisplayName: ua.MustVariant("HasProperty"),
				ua.AttributeIDInverseName: ua.MustVariant(&ua.LocalizedText{Locale: "", Text: "PropertyOf"}),
			},
			superTypeID: ua.NewNumericNodeID(0, 44),
		},
		&node{
			id: ua.NewNumericNodeID(0, 47),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ReferenceType_32"),
				ua.AttributeIDBrowseName:  ua.MustVariant("HasComponent"),
				ua.AttributeIDDisplayName: ua.MustVariant("HasComponent"),
				ua.AttributeIDInverseName: ua.MustVariant(&ua.LocalizedText{Locale: "", Text: "ComponentOf"}),
			},
			superTypeID: ua.NewNumericNodeID(0, 44),
		},
		&node{
			id: ua.NewNumericNodeID(0, 48),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ReferenceType_32"),
				ua.AttributeIDBrowseName:  ua.MustVariant("HasNotifier"),
				ua.AttributeIDDisplayName: ua.MustVariant("HasNotifier"),
				ua.AttributeIDInverseName: ua.MustVariant(&ua.LocalizedText{Locale: "", Text: "NotifierOf"}),
			},
			superTypeID: ua.NewNumericNodeID(0, 36),
		},
		&node{
			id: ua.NewNumericNodeID(0, 49),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ReferenceType_32"),
				ua.AttributeIDBrowseName:  ua.MustVariant("HasOrderedComponent"),
				ua.AttributeIDDisplayName: ua.MustVariant("HasOrderedComponent"),
				ua.AttributeIDInverseName: ua.MustVariant(&ua.LocalizedText{Locale: "", Text: "OrderedComponentOf"}),
			},
			superTypeID: ua.NewNumericNodeID(0, 47),
		},
		&node{
			id: ua.NewNumericNodeID(0, 51),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ReferenceType_32"),
				ua.AttributeIDBrowseName:  ua.MustVariant("FromState"),
				ua.AttributeIDDisplayName: ua.MustVariant("FromState"),
				ua.AttributeIDInverseName: ua.MustVariant(&ua.LocalizedText{Locale: "", Text: "ToTransition"}),
			},
			superTypeID: ua.NewNumericNodeID(0, 32),
		},
		&node{
			id: ua.NewNumericNodeID(0, 52),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ReferenceType_32"),
				ua.AttributeIDBrowseName:  ua.MustVariant("ToState"),
				ua.AttributeIDDisplayName: ua.MustVariant("ToState"),
				ua.AttributeIDInverseName: ua.MustVariant(&ua.LocalizedText{Locale: "", Text: "FromTransition"}),
			},
			superTypeID: ua.NewNumericNodeID(0, 32),
		},
		&node{
			id: ua.NewNumericNodeID(0, 53),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ReferenceType_32"),
				ua.AttributeIDBrowseName:  ua.MustVariant("HasCause"),
				ua.AttributeIDDisplayName: ua.MustVariant("HasCause"),
				ua.AttributeIDInverseName: ua.MustVariant(&ua.LocalizedText{Locale: "", Text: "MayBeCausedBy"}),
			},
			superTypeID: ua.NewNumericNodeID(0, 32),
		},
		&node{
			id: ua.NewNumericNodeID(0, 54),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ReferenceType_32"),
				ua.AttributeIDBrowseName:  ua.MustVariant("HasEffect"),
				ua.AttributeIDDisplayName: ua.MustVariant("HasEffect"),
				ua.AttributeIDInverseName: ua.MustVariant(&ua.LocalizedText{Locale: "", Text: "MayBeEffectedBy"}),
			},
			superTypeID: ua.NewNumericNodeID(0, 32),
		},
		&node{
			id: ua.NewNumericNodeID(0, 117),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ReferenceType_32"),
				ua.AttributeIDBrowseName:  ua.MustVariant("HasSubStateMachine"),
				ua.AttributeIDDisplayName: ua.MustVariant("HasSubStateMachine"),
				ua.AttributeIDInverseName: ua.MustVariant(&ua.LocalizedText{Locale: "", Text: "SubStateMachineOf"}),
			},
			superTypeID: ua.NewNumericNodeID(0, 32),
		},
		&node{
			id: ua.NewNumericNodeID(0, 56),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ReferenceType_32"),
				ua.AttributeIDBrowseName:  ua.MustVariant("HasHistoricalConfiguration"),
				ua.AttributeIDDisplayName: ua.MustVariant("HasHistoricalConfiguration"),
				ua.AttributeIDInverseName: ua.MustVariant(&ua.LocalizedText{Locale: "", Text: "HistoricalConfigurationOf"}),
			},
			superTypeID: ua.NewNumericNodeID(0, 44),
		},
		&node{
			id: ua.NewNumericNodeID(0, 24136),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ReferenceType_32"),
				ua.AttributeIDBrowseName:  ua.MustVariant("HasStructuredComponent"),
				ua.AttributeIDDisplayName: ua.MustVariant("HasStructuredComponent"),
				ua.AttributeIDInverseName: ua.MustVariant(&ua.LocalizedText{Locale: "", Text: "IsStructuredComponentOf"}),
			},
			superTypeID: ua.NewNumericNodeID(0, 47),
		},
		&node{
			id: ua.NewNumericNodeID(0, 24137),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ReferenceType_32"),
				ua.AttributeIDBrowseName:  ua.MustVariant("AssociatedWith"),
				ua.AttributeIDDisplayName: ua.MustVariant("AssociatedWith"),
			},
			superTypeID: ua.NewNumericNodeID(0, 32),
		},
		&node{
			id: ua.NewNumericNodeID(0, 58),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ObjectType_8"),
				ua.AttributeIDBrowseName:  ua.MustVariant("BaseObjectType"),
				ua.AttributeIDDisplayName: ua.MustVariant("BaseObjectType"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 61),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ObjectType_8"),
				ua.AttributeIDBrowseName:  ua.MustVariant("FolderType"),
				ua.AttributeIDDisplayName: ua.MustVariant("FolderType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 58),
		},
		&node{
			id: ua.NewNumericNodeID(0, 63),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("VariableType_16"),
				ua.AttributeIDBrowseName:  ua.MustVariant("BaseDataVariableType"),
				ua.AttributeIDDisplayName: ua.MustVariant("BaseDataVariableType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 62),
		},
		&node{
			id: ua.NewNumericNodeID(0, 68),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("VariableType_16"),
				ua.AttributeIDBrowseName:  ua.MustVariant("PropertyType"),
				ua.AttributeIDDisplayName: ua.MustVariant("PropertyType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 62),
		},
		&node{
			id: ua.NewNumericNodeID(0, 69),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("VariableType_16"),
				ua.AttributeIDBrowseName:  ua.MustVariant("DataTypeDescriptionType"),
				ua.AttributeIDDisplayName: ua.MustVariant("DataTypeDescriptionType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 63),
		},
		&node{
			id: ua.NewNumericNodeID(0, 72),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("VariableType_16"),
				ua.AttributeIDBrowseName:  ua.MustVariant("DataTypeDictionaryType"),
				ua.AttributeIDDisplayName: ua.MustVariant("DataTypeDictionaryType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 63),
		},
		&node{
			id: ua.NewNumericNodeID(0, 75),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ObjectType_8"),
				ua.AttributeIDBrowseName:  ua.MustVariant("DataTypeSystemType"),
				ua.AttributeIDDisplayName: ua.MustVariant("DataTypeSystemType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 58),
		},
		&node{
			id: ua.NewNumericNodeID(0, 76),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ObjectType_8"),
				ua.AttributeIDBrowseName:  ua.MustVariant("DataTypeEncodingType"),
				ua.AttributeIDDisplayName: ua.MustVariant("DataTypeEncodingType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 58),
		},
		&node{
			id: ua.NewNumericNodeID(0, 120),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("NamingRuleType"),
				ua.AttributeIDDisplayName: ua.MustVariant("NamingRuleType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 29),
		},
		&node{
			id: ua.NewNumericNodeID(0, 77),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ObjectType_8"),
				ua.AttributeIDBrowseName:  ua.MustVariant("ModellingRuleType"),
				ua.AttributeIDDisplayName: ua.MustVariant("ModellingRuleType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 58),
		},
		&node{
			id: ua.NewNumericNodeID(0, 78),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Mandatory"),
				ua.AttributeIDDisplayName: ua.MustVariant("Mandatory"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 80),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Optional"),
				ua.AttributeIDDisplayName: ua.MustVariant("Optional"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 83),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("ExposesItsArray"),
				ua.AttributeIDDisplayName: ua.MustVariant("ExposesItsArray"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 11508),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("OptionalPlaceholder"),
				ua.AttributeIDDisplayName: ua.MustVariant("OptionalPlaceholder"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 11510),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("MandatoryPlaceholder"),
				ua.AttributeIDDisplayName: ua.MustVariant("MandatoryPlaceholder"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 84),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Root"),
				ua.AttributeIDDisplayName: ua.MustVariant("Root"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 85),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Objects"),
				ua.AttributeIDDisplayName: ua.MustVariant("Objects"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 86),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Types"),
				ua.AttributeIDDisplayName: ua.MustVariant("Types"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 87),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Views"),
				ua.AttributeIDDisplayName: ua.MustVariant("Views"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 88),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("ObjectTypes"),
				ua.AttributeIDDisplayName: ua.MustVariant("ObjectTypes"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 89),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("VariableTypes"),
				ua.AttributeIDDisplayName: ua.MustVariant("VariableTypes"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 90),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("DataTypes"),
				ua.AttributeIDDisplayName: ua.MustVariant("DataTypes"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 91),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("ReferenceTypes"),
				ua.AttributeIDDisplayName: ua.MustVariant("ReferenceTypes"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 92),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("XML Schema"),
				ua.AttributeIDDisplayName: ua.MustVariant("XML Schema"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 93),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("OPC Binary"),
				ua.AttributeIDDisplayName: ua.MustVariant("OPC Binary"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 129),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ReferenceType_32"),
				ua.AttributeIDBrowseName:  ua.MustVariant("HasArgumentDescription"),
				ua.AttributeIDDisplayName: ua.MustVariant("HasArgumentDescription"),
				ua.AttributeIDInverseName: ua.MustVariant(&ua.LocalizedText{Locale: "", Text: "ArgumentDescriptionOf"}),
			},
			superTypeID: ua.NewNumericNodeID(0, 47),
		},
		&node{
			id: ua.NewNumericNodeID(0, 131),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ReferenceType_32"),
				ua.AttributeIDBrowseName:  ua.MustVariant("HasOptionalInputArgumentDescription"),
				ua.AttributeIDDisplayName: ua.MustVariant("HasOptionalInputArgumentDescription"),
				ua.AttributeIDInverseName: ua.MustVariant(&ua.LocalizedText{Locale: "", Text: "OptionalInputArgumentDescriptionOf"}),
			},
			superTypeID: ua.NewNumericNodeID(0, 129),
		},
		&node{
			id: ua.NewNumericNodeID(0, 15957),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("http://opcfoundation.org/UA/"),
				ua.AttributeIDDisplayName: ua.MustVariant("http://opcfoundation.org/UA/"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 3068),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Variable_2"),
				ua.AttributeIDBrowseName:  ua.MustVariant("NodeVersion"),
				ua.AttributeIDDisplayName: ua.MustVariant("NodeVersion"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 12170),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Variable_2"),
				ua.AttributeIDBrowseName:  ua.MustVariant("ViewVersion"),
				ua.AttributeIDDisplayName: ua.MustVariant("ViewVersion"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 3067),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Variable_2"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Icon"),
				ua.AttributeIDDisplayName: ua.MustVariant("Icon"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 3069),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Variable_2"),
				ua.AttributeIDBrowseName:  ua.MustVariant("LocalTime"),
				ua.AttributeIDDisplayName: ua.MustVariant("LocalTime"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 3070),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Variable_2"),
				ua.AttributeIDBrowseName:  ua.MustVariant("AllowNulls"),
				ua.AttributeIDDisplayName: ua.MustVariant("AllowNulls"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 11433),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Variable_2"),
				ua.AttributeIDBrowseName:  ua.MustVariant("ValueAsText"),
				ua.AttributeIDDisplayName: ua.MustVariant("ValueAsText"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 11498),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Variable_2"),
				ua.AttributeIDBrowseName:  ua.MustVariant("MaxStringLength"),
				ua.AttributeIDDisplayName: ua.MustVariant("MaxStringLength"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15002),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Variable_2"),
				ua.AttributeIDBrowseName:  ua.MustVariant("MaxCharacters"),
				ua.AttributeIDDisplayName: ua.MustVariant("MaxCharacters"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 12908),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Variable_2"),
				ua.AttributeIDBrowseName:  ua.MustVariant("MaxByteStringLength"),
				ua.AttributeIDDisplayName: ua.MustVariant("MaxByteStringLength"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 11512),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Variable_2"),
				ua.AttributeIDBrowseName:  ua.MustVariant("MaxArrayLength"),
				ua.AttributeIDDisplayName: ua.MustVariant("MaxArrayLength"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 11513),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Variable_2"),
				ua.AttributeIDBrowseName:  ua.MustVariant("EngineeringUnits"),
				ua.AttributeIDDisplayName: ua.MustVariant("EngineeringUnits"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 11432),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Variable_2"),
				ua.AttributeIDBrowseName:  ua.MustVariant("EnumStrings"),
				ua.AttributeIDDisplayName: ua.MustVariant("EnumStrings"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 3071),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Variable_2"),
				ua.AttributeIDBrowseName:  ua.MustVariant("EnumValues"),
				ua.AttributeIDDisplayName: ua.MustVariant("EnumValues"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 12745),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Variable_2"),
				ua.AttributeIDBrowseName:  ua.MustVariant("OptionSetValues"),
				ua.AttributeIDDisplayName: ua.MustVariant("OptionSetValues"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 32750),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Variable_2"),
				ua.AttributeIDBrowseName:  ua.MustVariant("OptionSetLength"),
				ua.AttributeIDDisplayName: ua.MustVariant("OptionSetLength"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 3072),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Variable_2"),
				ua.AttributeIDBrowseName:  ua.MustVariant("InputArguments"),
				ua.AttributeIDDisplayName: ua.MustVariant("InputArguments"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 3073),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Variable_2"),
				ua.AttributeIDBrowseName:  ua.MustVariant("OutputArguments"),
				ua.AttributeIDDisplayName: ua.MustVariant("OutputArguments"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 17605),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Variable_2"),
				ua.AttributeIDBrowseName:  ua.MustVariant("DefaultInstanceBrowseName"),
				ua.AttributeIDDisplayName: ua.MustVariant("DefaultInstanceBrowseName"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 2000),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("ImageBMP"),
				ua.AttributeIDDisplayName: ua.MustVariant("ImageBMP"),
			},
			superTypeID: ua.NewNumericNodeID(0, 30),
		},
		&node{
			id: ua.NewNumericNodeID(0, 2001),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("ImageGIF"),
				ua.AttributeIDDisplayName: ua.MustVariant("ImageGIF"),
			},
			superTypeID: ua.NewNumericNodeID(0, 30),
		},
		&node{
			id: ua.NewNumericNodeID(0, 2002),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("ImageJPG"),
				ua.AttributeIDDisplayName: ua.MustVariant("ImageJPG"),
			},
			superTypeID: ua.NewNumericNodeID(0, 30),
		},
		&node{
			id: ua.NewNumericNodeID(0, 2003),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("ImagePNG"),
				ua.AttributeIDDisplayName: ua.MustVariant("ImagePNG"),
			},
			superTypeID: ua.NewNumericNodeID(0, 30),
		},
		&node{
			id: ua.NewNumericNodeID(0, 16307),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("AudioDataType"),
				ua.AttributeIDDisplayName: ua.MustVariant("AudioDataType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 15),
		},
		&node{
			id: ua.NewNumericNodeID(0, 23751),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("UriString"),
				ua.AttributeIDDisplayName: ua.MustVariant("UriString"),
			},
			superTypeID: ua.NewNumericNodeID(0, 12),
		},
		&node{
			id: ua.NewNumericNodeID(0, 2004),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ObjectType_8"),
				ua.AttributeIDBrowseName:  ua.MustVariant("ServerType"),
				ua.AttributeIDDisplayName: ua.MustVariant("ServerType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 58),
		},
		&node{
			id: ua.NewNumericNodeID(0, 2013),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ObjectType_8"),
				ua.AttributeIDBrowseName:  ua.MustVariant("ServerCapabilitiesType"),
				ua.AttributeIDDisplayName: ua.MustVariant("ServerCapabilitiesType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 58),
		},
		&node{
			id: ua.NewNumericNodeID(0, 2020),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ObjectType_8"),
				ua.AttributeIDBrowseName:  ua.MustVariant("ServerDiagnosticsType"),
				ua.AttributeIDDisplayName: ua.MustVariant("ServerDiagnosticsType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 58),
		},
		&node{
			id: ua.NewNumericNodeID(0, 2026),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ObjectType_8"),
				ua.AttributeIDBrowseName:  ua.MustVariant("SessionsDiagnosticsSummaryType"),
				ua.AttributeIDDisplayName: ua.MustVariant("SessionsDiagnosticsSummaryType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 58),
		},
		&node{
			id: ua.NewNumericNodeID(0, 2029),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ObjectType_8"),
				ua.AttributeIDBrowseName:  ua.MustVariant("SessionDiagnosticsObjectType"),
				ua.AttributeIDDisplayName: ua.MustVariant("SessionDiagnosticsObjectType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 58),
		},
		&node{
			id: ua.NewNumericNodeID(0, 2033),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ObjectType_8"),
				ua.AttributeIDBrowseName:  ua.MustVariant("VendorServerInfoType"),
				ua.AttributeIDDisplayName: ua.MustVariant("VendorServerInfoType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 58),
		},
		&node{
			id: ua.NewNumericNodeID(0, 2034),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ObjectType_8"),
				ua.AttributeIDBrowseName:  ua.MustVariant("ServerRedundancyType"),
				ua.AttributeIDDisplayName: ua.MustVariant("ServerRedundancyType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 58),
		},
		&node{
			id: ua.NewNumericNodeID(0, 2036),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ObjectType_8"),
				ua.AttributeIDBrowseName:  ua.MustVariant("TransparentRedundancyType"),
				ua.AttributeIDDisplayName: ua.MustVariant("TransparentRedundancyType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 2034),
		},
		&node{
			id: ua.NewNumericNodeID(0, 2039),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ObjectType_8"),
				ua.AttributeIDBrowseName:  ua.MustVariant("NonTransparentRedundancyType"),
				ua.AttributeIDDisplayName: ua.MustVariant("NonTransparentRedundancyType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 2034),
		},
		&node{
			id: ua.NewNumericNodeID(0, 11945),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ObjectType_8"),
				ua.AttributeIDBrowseName:  ua.MustVariant("NonTransparentNetworkRedundancyType"),
				ua.AttributeIDDisplayName: ua.MustVariant("NonTransparentNetworkRedundancyType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 2039),
		},
		&node{
			id: ua.NewNumericNodeID(0, 11564),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ObjectType_8"),
				ua.AttributeIDBrowseName:  ua.MustVariant("OperationLimitsType"),
				ua.AttributeIDDisplayName: ua.MustVariant("OperationLimitsType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 61),
		},
		&node{
			id: ua.NewNumericNodeID(0, 11575),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ObjectType_8"),
				ua.AttributeIDBrowseName:  ua.MustVariant("FileType"),
				ua.AttributeIDDisplayName: ua.MustVariant("FileType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 58),
		},
		&node{
			id: ua.NewNumericNodeID(0, 11595),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ObjectType_8"),
				ua.AttributeIDBrowseName:  ua.MustVariant("AddressSpaceFileType"),
				ua.AttributeIDDisplayName: ua.MustVariant("AddressSpaceFileType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 11575),
		},
		&node{
			id: ua.NewNumericNodeID(0, 11616),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ObjectType_8"),
				ua.AttributeIDBrowseName:  ua.MustVariant("NamespaceMetadataType"),
				ua.AttributeIDDisplayName: ua.MustVariant("NamespaceMetadataType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 58),
		},
		&node{
			id: ua.NewNumericNodeID(0, 11645),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ObjectType_8"),
				ua.AttributeIDBrowseName:  ua.MustVariant("NamespacesType"),
				ua.AttributeIDDisplayName: ua.MustVariant("NamespacesType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 58),
		},
		&node{
			id: ua.NewNumericNodeID(0, 23926),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ObjectType_8"),
				ua.AttributeIDBrowseName:  ua.MustVariant("AuditClientUpdateMethodResultEventType"),
				ua.AttributeIDDisplayName: ua.MustVariant("AuditClientUpdateMethodResultEventType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 23606),
		},
		&node{
			id: ua.NewNumericNodeID(0, 2340),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ObjectType_8"),
				ua.AttributeIDBrowseName:  ua.MustVariant("AggregateFunctionType"),
				ua.AttributeIDDisplayName: ua.MustVariant("AggregateFunctionType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 58),
		},
		&node{
			id: ua.NewNumericNodeID(0, 2138),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("VariableType_16"),
				ua.AttributeIDBrowseName:  ua.MustVariant("ServerStatusType"),
				ua.AttributeIDDisplayName: ua.MustVariant("ServerStatusType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 63),
		},
		&node{
			id: ua.NewNumericNodeID(0, 3051),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("VariableType_16"),
				ua.AttributeIDBrowseName:  ua.MustVariant("BuildInfoType"),
				ua.AttributeIDDisplayName: ua.MustVariant("BuildInfoType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 63),
		},
		&node{
			id: ua.NewNumericNodeID(0, 2150),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("VariableType_16"),
				ua.AttributeIDBrowseName:  ua.MustVariant("ServerDiagnosticsSummaryType"),
				ua.AttributeIDDisplayName: ua.MustVariant("ServerDiagnosticsSummaryType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 63),
		},
		&node{
			id: ua.NewNumericNodeID(0, 2164),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("VariableType_16"),
				ua.AttributeIDBrowseName:  ua.MustVariant("SamplingIntervalDiagnosticsArrayType"),
				ua.AttributeIDDisplayName: ua.MustVariant("SamplingIntervalDiagnosticsArrayType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 63),
		},
		&node{
			id: ua.NewNumericNodeID(0, 2165),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("VariableType_16"),
				ua.AttributeIDBrowseName:  ua.MustVariant("SamplingIntervalDiagnosticsType"),
				ua.AttributeIDDisplayName: ua.MustVariant("SamplingIntervalDiagnosticsType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 63),
		},
		&node{
			id: ua.NewNumericNodeID(0, 2171),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("VariableType_16"),
				ua.AttributeIDBrowseName:  ua.MustVariant("SubscriptionDiagnosticsArrayType"),
				ua.AttributeIDDisplayName: ua.MustVariant("SubscriptionDiagnosticsArrayType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 63),
		},
		&node{
			id: ua.NewNumericNodeID(0, 2172),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("VariableType_16"),
				ua.AttributeIDBrowseName:  ua.MustVariant("SubscriptionDiagnosticsType"),
				ua.AttributeIDDisplayName: ua.MustVariant("SubscriptionDiagnosticsType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 63),
		},
		&node{
			id: ua.NewNumericNodeID(0, 2196),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("VariableType_16"),
				ua.AttributeIDBrowseName:  ua.MustVariant("SessionDiagnosticsArrayType"),
				ua.AttributeIDDisplayName: ua.MustVariant("SessionDiagnosticsArrayType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 63),
		},
		&node{
			id: ua.NewNumericNodeID(0, 2197),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("VariableType_16"),
				ua.AttributeIDBrowseName:  ua.MustVariant("SessionDiagnosticsVariableType"),
				ua.AttributeIDDisplayName: ua.MustVariant("SessionDiagnosticsVariableType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 63),
		},
		&node{
			id: ua.NewNumericNodeID(0, 2243),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("VariableType_16"),
				ua.AttributeIDBrowseName:  ua.MustVariant("SessionSecurityDiagnosticsArrayType"),
				ua.AttributeIDDisplayName: ua.MustVariant("SessionSecurityDiagnosticsArrayType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 63),
		},
		&node{
			id: ua.NewNumericNodeID(0, 2244),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("VariableType_16"),
				ua.AttributeIDBrowseName:  ua.MustVariant("SessionSecurityDiagnosticsType"),
				ua.AttributeIDDisplayName: ua.MustVariant("SessionSecurityDiagnosticsType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 63),
		},
		&node{
			id: ua.NewNumericNodeID(0, 11487),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("VariableType_16"),
				ua.AttributeIDBrowseName:  ua.MustVariant("OptionSetType"),
				ua.AttributeIDDisplayName: ua.MustVariant("OptionSetType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 63),
		},
		&node{
			id: ua.NewNumericNodeID(0, 16309),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("VariableType_16"),
				ua.AttributeIDBrowseName:  ua.MustVariant("SelectionListType"),
				ua.AttributeIDDisplayName: ua.MustVariant("SelectionListType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 63),
		},
		&node{
			id: ua.NewNumericNodeID(0, 17986),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("VariableType_16"),
				ua.AttributeIDBrowseName:  ua.MustVariant("AudioVariableType"),
				ua.AttributeIDDisplayName: ua.MustVariant("AudioVariableType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 63),
		},
		&node{
			id: ua.NewNumericNodeID(0, 3048),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("EventTypes"),
				ua.AttributeIDDisplayName: ua.MustVariant("EventTypes"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 31915),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Locations"),
				ua.AttributeIDDisplayName: ua.MustVariant("Locations"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 2253),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Server"),
				ua.AttributeIDDisplayName: ua.MustVariant("Server"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 11312),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Variable_2"),
				ua.AttributeIDBrowseName:  ua.MustVariant("CurrentServerId"),
				ua.AttributeIDDisplayName: ua.MustVariant("CurrentServerId"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 11313),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Variable_2"),
				ua.AttributeIDBrowseName:  ua.MustVariant("RedundantServerArray"),
				ua.AttributeIDDisplayName: ua.MustVariant("RedundantServerArray"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 11314),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Variable_2"),
				ua.AttributeIDBrowseName:  ua.MustVariant("ServerUriArray"),
				ua.AttributeIDDisplayName: ua.MustVariant("ServerUriArray"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 14415),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Variable_2"),
				ua.AttributeIDBrowseName:  ua.MustVariant("ServerNetworkGroups"),
				ua.AttributeIDDisplayName: ua.MustVariant("ServerNetworkGroups"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 11192),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("HistoryServerCapabilities"),
				ua.AttributeIDDisplayName: ua.MustVariant("HistoryServerCapabilities"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 23562),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ReferenceType_32"),
				ua.AttributeIDBrowseName:  ua.MustVariant("IsDeprecated"),
				ua.AttributeIDDisplayName: ua.MustVariant("IsDeprecated"),
				ua.AttributeIDInverseName: ua.MustVariant(&ua.LocalizedText{Locale: "", Text: "Deprecates"}),
			},
			superTypeID: ua.NewNumericNodeID(0, 32),
		},
		&node{
			id: ua.NewNumericNodeID(0, 11737),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("BitFieldMaskDataType"),
				ua.AttributeIDDisplayName: ua.MustVariant("BitFieldMaskDataType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 9),
		},
		&node{
			id: ua.NewNumericNodeID(0, 24263),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("SemanticVersionString"),
				ua.AttributeIDDisplayName: ua.MustVariant("SemanticVersionString"),
			},
			superTypeID: ua.NewNumericNodeID(0, 12),
		},
		&node{
			id: ua.NewNumericNodeID(0, 14533),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("KeyValuePair"),
				ua.AttributeIDDisplayName: ua.MustVariant("KeyValuePair"),
			},
			superTypeID: ua.NewNumericNodeID(0, 22),
		},
		&node{
			id: ua.NewNumericNodeID(0, 16313),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("AdditionalParametersType"),
				ua.AttributeIDDisplayName: ua.MustVariant("AdditionalParametersType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 22),
		},
		&node{
			id: ua.NewNumericNodeID(0, 17548),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("EphemeralKeyType"),
				ua.AttributeIDDisplayName: ua.MustVariant("EphemeralKeyType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 22),
		},
		&node{
			id: ua.NewNumericNodeID(0, 15528),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("EndpointType"),
				ua.AttributeIDDisplayName: ua.MustVariant("EndpointType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 22),
		},
		&node{
			id: ua.NewNumericNodeID(0, 31917),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Handle"),
				ua.AttributeIDDisplayName: ua.MustVariant("Handle"),
			},
			superTypeID: ua.NewNumericNodeID(0, 7),
		},
		&node{
			id: ua.NewNumericNodeID(0, 31918),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("TrimmedString"),
				ua.AttributeIDDisplayName: ua.MustVariant("TrimmedString"),
			},
			superTypeID: ua.NewNumericNodeID(0, 12),
		},
		&node{
			id: ua.NewNumericNodeID(0, 2299),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ObjectType_8"),
				ua.AttributeIDBrowseName:  ua.MustVariant("StateMachineType"),
				ua.AttributeIDDisplayName: ua.MustVariant("StateMachineType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 58),
		},
		&node{
			id: ua.NewNumericNodeID(0, 2755),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("VariableType_16"),
				ua.AttributeIDBrowseName:  ua.MustVariant("StateVariableType"),
				ua.AttributeIDDisplayName: ua.MustVariant("StateVariableType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 63),
		},
		&node{
			id: ua.NewNumericNodeID(0, 2762),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("VariableType_16"),
				ua.AttributeIDBrowseName:  ua.MustVariant("TransitionVariableType"),
				ua.AttributeIDDisplayName: ua.MustVariant("TransitionVariableType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 63),
		},
		&node{
			id: ua.NewNumericNodeID(0, 2760),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("VariableType_16"),
				ua.AttributeIDBrowseName:  ua.MustVariant("FiniteStateVariableType"),
				ua.AttributeIDDisplayName: ua.MustVariant("FiniteStateVariableType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 2755),
		},
		&node{
			id: ua.NewNumericNodeID(0, 2767),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("VariableType_16"),
				ua.AttributeIDBrowseName:  ua.MustVariant("FiniteTransitionVariableType"),
				ua.AttributeIDDisplayName: ua.MustVariant("FiniteTransitionVariableType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 2762),
		},
		&node{
			id: ua.NewNumericNodeID(0, 2307),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ObjectType_8"),
				ua.AttributeIDBrowseName:  ua.MustVariant("StateType"),
				ua.AttributeIDDisplayName: ua.MustVariant("StateType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 58),
		},
		&node{
			id: ua.NewNumericNodeID(0, 2309),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ObjectType_8"),
				ua.AttributeIDBrowseName:  ua.MustVariant("InitialStateType"),
				ua.AttributeIDDisplayName: ua.MustVariant("InitialStateType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 2307),
		},
		&node{
			id: ua.NewNumericNodeID(0, 2310),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ObjectType_8"),
				ua.AttributeIDBrowseName:  ua.MustVariant("TransitionType"),
				ua.AttributeIDDisplayName: ua.MustVariant("TransitionType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 58),
		},
		&node{
			id: ua.NewNumericNodeID(0, 15109),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ObjectType_8"),
				ua.AttributeIDBrowseName:  ua.MustVariant("ChoiceStateType"),
				ua.AttributeIDDisplayName: ua.MustVariant("ChoiceStateType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 2307),
		},
		&node{
			id: ua.NewNumericNodeID(0, 15112),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ReferenceType_32"),
				ua.AttributeIDBrowseName:  ua.MustVariant("HasGuard"),
				ua.AttributeIDDisplayName: ua.MustVariant("HasGuard"),
				ua.AttributeIDInverseName: ua.MustVariant(&ua.LocalizedText{Locale: "", Text: "GuardOf"}),
			},
			superTypeID: ua.NewNumericNodeID(0, 47),
		},
		&node{
			id: ua.NewNumericNodeID(0, 15113),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("VariableType_16"),
				ua.AttributeIDBrowseName:  ua.MustVariant("GuardVariableType"),
				ua.AttributeIDDisplayName: ua.MustVariant("GuardVariableType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 63),
		},
		&node{
			id: ua.NewNumericNodeID(0, 15128),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("VariableType_16"),
				ua.AttributeIDBrowseName:  ua.MustVariant("ExpressionGuardVariableType"),
				ua.AttributeIDDisplayName: ua.MustVariant("ExpressionGuardVariableType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 15113),
		},
		&node{
			id: ua.NewNumericNodeID(0, 15317),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("VariableType_16"),
				ua.AttributeIDBrowseName:  ua.MustVariant("ElseGuardVariableType"),
				ua.AttributeIDDisplayName: ua.MustVariant("ElseGuardVariableType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 15113),
		},
		&node{
			id: ua.NewNumericNodeID(0, 17709),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("VariableType_16"),
				ua.AttributeIDBrowseName:  ua.MustVariant("RationalNumberType"),
				ua.AttributeIDDisplayName: ua.MustVariant("RationalNumberType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 63),
		},
		&node{
			id: ua.NewNumericNodeID(0, 17716),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("VariableType_16"),
				ua.AttributeIDBrowseName:  ua.MustVariant("3DVectorType"),
				ua.AttributeIDDisplayName: ua.MustVariant("3DVectorType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 17714),
		},
		&node{
			id: ua.NewNumericNodeID(0, 18774),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("VariableType_16"),
				ua.AttributeIDBrowseName:  ua.MustVariant("3DCartesianCoordinatesType"),
				ua.AttributeIDDisplayName: ua.MustVariant("3DCartesianCoordinatesType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 18772),
		},
		&node{
			id: ua.NewNumericNodeID(0, 18781),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("VariableType_16"),
				ua.AttributeIDBrowseName:  ua.MustVariant("3DOrientationType"),
				ua.AttributeIDDisplayName: ua.MustVariant("3DOrientationType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 18779),
		},
		&node{
			id: ua.NewNumericNodeID(0, 18791),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("VariableType_16"),
				ua.AttributeIDBrowseName:  ua.MustVariant("3DFrameType"),
				ua.AttributeIDDisplayName: ua.MustVariant("3DFrameType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 18786),
		},
		&node{
			id: ua.NewNumericNodeID(0, 18806),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("RationalNumber"),
				ua.AttributeIDDisplayName: ua.MustVariant("RationalNumber"),
			},
			superTypeID: ua.NewNumericNodeID(0, 22),
		},
		&node{
			id: ua.NewNumericNodeID(0, 18808),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("3DVector"),
				ua.AttributeIDDisplayName: ua.MustVariant("3DVector"),
			},
			superTypeID: ua.NewNumericNodeID(0, 18807),
		},
		&node{
			id: ua.NewNumericNodeID(0, 18810),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("3DCartesianCoordinates"),
				ua.AttributeIDDisplayName: ua.MustVariant("3DCartesianCoordinates"),
			},
			superTypeID: ua.NewNumericNodeID(0, 18809),
		},
		&node{
			id: ua.NewNumericNodeID(0, 18812),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("3DOrientation"),
				ua.AttributeIDDisplayName: ua.MustVariant("3DOrientation"),
			},
			superTypeID: ua.NewNumericNodeID(0, 18811),
		},
		&node{
			id: ua.NewNumericNodeID(0, 18814),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("3DFrame"),
				ua.AttributeIDDisplayName: ua.MustVariant("3DFrame"),
			},
			superTypeID: ua.NewNumericNodeID(0, 18813),
		},
		&node{
			id: ua.NewNumericNodeID(0, 11939),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("OpenFileMode"),
				ua.AttributeIDDisplayName: ua.MustVariant("OpenFileMode"),
			},
			superTypeID: ua.NewNumericNodeID(0, 29),
		},
		&node{
			id: ua.NewNumericNodeID(0, 13353),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ObjectType_8"),
				ua.AttributeIDBrowseName:  ua.MustVariant("FileDirectoryType"),
				ua.AttributeIDDisplayName: ua.MustVariant("FileDirectoryType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 61),
		},
		&node{
			id: ua.NewNumericNodeID(0, 16314),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("FileSystem"),
				ua.AttributeIDDisplayName: ua.MustVariant("FileSystem"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15744),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ObjectType_8"),
				ua.AttributeIDBrowseName:  ua.MustVariant("TemporaryFileTransferType"),
				ua.AttributeIDDisplayName: ua.MustVariant("TemporaryFileTransferType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 58),
		},
		&node{
			id: ua.NewNumericNodeID(0, 15803),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ObjectType_8"),
				ua.AttributeIDBrowseName:  ua.MustVariant("FileTransferStateMachineType"),
				ua.AttributeIDDisplayName: ua.MustVariant("FileTransferStateMachineType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 2771),
		},
		&node{
			id: ua.NewNumericNodeID(0, 15607),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ObjectType_8"),
				ua.AttributeIDBrowseName:  ua.MustVariant("RoleSetType"),
				ua.AttributeIDDisplayName: ua.MustVariant("RoleSetType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 58),
		},
		&node{
			id: ua.NewNumericNodeID(0, 15620),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ObjectType_8"),
				ua.AttributeIDBrowseName:  ua.MustVariant("RoleType"),
				ua.AttributeIDDisplayName: ua.MustVariant("RoleType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 58),
		},
		&node{
			id: ua.NewNumericNodeID(0, 15632),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("IdentityCriteriaType"),
				ua.AttributeIDDisplayName: ua.MustVariant("IdentityCriteriaType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 29),
		},
		&node{
			id: ua.NewNumericNodeID(0, 15634),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("IdentityMappingRuleType"),
				ua.AttributeIDDisplayName: ua.MustVariant("IdentityMappingRuleType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 22),
		},
		&node{
			id: ua.NewNumericNodeID(0, 15644),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Anonymous"),
				ua.AttributeIDDisplayName: ua.MustVariant("Anonymous"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15656),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("AuthenticatedUser"),
				ua.AttributeIDDisplayName: ua.MustVariant("AuthenticatedUser"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15668),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Observer"),
				ua.AttributeIDDisplayName: ua.MustVariant("Observer"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15680),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Operator"),
				ua.AttributeIDDisplayName: ua.MustVariant("Operator"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 16036),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Engineer"),
				ua.AttributeIDDisplayName: ua.MustVariant("Engineer"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15692),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Supervisor"),
				ua.AttributeIDDisplayName: ua.MustVariant("Supervisor"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15716),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("ConfigureAdmin"),
				ua.AttributeIDDisplayName: ua.MustVariant("ConfigureAdmin"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15704),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("SecurityAdmin"),
				ua.AttributeIDDisplayName: ua.MustVariant("SecurityAdmin"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 25565),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("SecurityKeyServerAdmin"),
				ua.AttributeIDDisplayName: ua.MustVariant("SecurityKeyServerAdmin"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 25603),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("SecurityKeyServerAccess"),
				ua.AttributeIDDisplayName: ua.MustVariant("SecurityKeyServerAccess"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 25584),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("SecurityKeyServerPush"),
				ua.AttributeIDDisplayName: ua.MustVariant("SecurityKeyServerPush"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 17591),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ObjectType_8"),
				ua.AttributeIDBrowseName:  ua.MustVariant("DictionaryFolderType"),
				ua.AttributeIDDisplayName: ua.MustVariant("DictionaryFolderType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 61),
		},
		&node{
			id: ua.NewNumericNodeID(0, 17594),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Dictionaries"),
				ua.AttributeIDDisplayName: ua.MustVariant("Dictionaries"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 17597),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ReferenceType_32"),
				ua.AttributeIDBrowseName:  ua.MustVariant("HasDictionaryEntry"),
				ua.AttributeIDDisplayName: ua.MustVariant("HasDictionaryEntry"),
				ua.AttributeIDInverseName: ua.MustVariant(&ua.LocalizedText{Locale: "", Text: "DictionaryEntryOf"}),
			},
			superTypeID: ua.NewNumericNodeID(0, 32),
		},
		&node{
			id: ua.NewNumericNodeID(0, 17598),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ObjectType_8"),
				ua.AttributeIDBrowseName:  ua.MustVariant("IrdiDictionaryEntryType"),
				ua.AttributeIDDisplayName: ua.MustVariant("IrdiDictionaryEntryType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 17589),
		},
		&node{
			id: ua.NewNumericNodeID(0, 17600),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ObjectType_8"),
				ua.AttributeIDBrowseName:  ua.MustVariant("UriDictionaryEntryType"),
				ua.AttributeIDDisplayName: ua.MustVariant("UriDictionaryEntryType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 17589),
		},
		&node{
			id: ua.NewNumericNodeID(0, 17708),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("InterfaceTypes"),
				ua.AttributeIDDisplayName: ua.MustVariant("InterfaceTypes"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 17603),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ReferenceType_32"),
				ua.AttributeIDBrowseName:  ua.MustVariant("HasInterface"),
				ua.AttributeIDDisplayName: ua.MustVariant("HasInterface"),
				ua.AttributeIDInverseName: ua.MustVariant(&ua.LocalizedText{Locale: "", Text: "InterfaceOf"}),
			},
			superTypeID: ua.NewNumericNodeID(0, 32),
		},
		&node{
			id: ua.NewNumericNodeID(0, 17604),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ReferenceType_32"),
				ua.AttributeIDBrowseName:  ua.MustVariant("HasAddIn"),
				ua.AttributeIDDisplayName: ua.MustVariant("HasAddIn"),
				ua.AttributeIDInverseName: ua.MustVariant(&ua.LocalizedText{Locale: "", Text: "AddInOf"}),
			},
			superTypeID: ua.NewNumericNodeID(0, 47),
		},
		&node{
			id: ua.NewNumericNodeID(0, 23498),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("CurrencyUnitType"),
				ua.AttributeIDDisplayName: ua.MustVariant("CurrencyUnitType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 22),
		},
		&node{
			id: ua.NewNumericNodeID(0, 23501),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Variable_2"),
				ua.AttributeIDBrowseName:  ua.MustVariant("CurrencyUnit"),
				ua.AttributeIDDisplayName: ua.MustVariant("CurrencyUnit"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 23518),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ObjectType_8"),
				ua.AttributeIDBrowseName:  ua.MustVariant("OrderedListType"),
				ua.AttributeIDDisplayName: ua.MustVariant("OrderedListType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 58),
		},
		&node{
			id: ua.NewNumericNodeID(0, 2365),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("VariableType_16"),
				ua.AttributeIDBrowseName:  ua.MustVariant("DataItemType"),
				ua.AttributeIDDisplayName: ua.MustVariant("DataItemType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 63),
		},
		&node{
			id: ua.NewNumericNodeID(0, 15318),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("VariableType_16"),
				ua.AttributeIDBrowseName:  ua.MustVariant("BaseAnalogType"),
				ua.AttributeIDDisplayName: ua.MustVariant("BaseAnalogType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 2365),
		},
		&node{
			id: ua.NewNumericNodeID(0, 2368),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("VariableType_16"),
				ua.AttributeIDBrowseName:  ua.MustVariant("AnalogItemType"),
				ua.AttributeIDDisplayName: ua.MustVariant("AnalogItemType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 15318),
		},
		&node{
			id: ua.NewNumericNodeID(0, 17497),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("VariableType_16"),
				ua.AttributeIDBrowseName:  ua.MustVariant("AnalogUnitType"),
				ua.AttributeIDDisplayName: ua.MustVariant("AnalogUnitType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 15318),
		},
		&node{
			id: ua.NewNumericNodeID(0, 17570),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("VariableType_16"),
				ua.AttributeIDBrowseName:  ua.MustVariant("AnalogUnitRangeType"),
				ua.AttributeIDDisplayName: ua.MustVariant("AnalogUnitRangeType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 2368),
		},
		&node{
			id: ua.NewNumericNodeID(0, 2373),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("VariableType_16"),
				ua.AttributeIDBrowseName:  ua.MustVariant("TwoStateDiscreteType"),
				ua.AttributeIDDisplayName: ua.MustVariant("TwoStateDiscreteType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 2372),
		},
		&node{
			id: ua.NewNumericNodeID(0, 2376),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("VariableType_16"),
				ua.AttributeIDBrowseName:  ua.MustVariant("MultiStateDiscreteType"),
				ua.AttributeIDDisplayName: ua.MustVariant("MultiStateDiscreteType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 2372),
		},
		&node{
			id: ua.NewNumericNodeID(0, 11238),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("VariableType_16"),
				ua.AttributeIDBrowseName:  ua.MustVariant("MultiStateValueDiscreteType"),
				ua.AttributeIDDisplayName: ua.MustVariant("MultiStateValueDiscreteType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 2372),
		},
		&node{
			id: ua.NewNumericNodeID(0, 12029),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("VariableType_16"),
				ua.AttributeIDBrowseName:  ua.MustVariant("YArrayItemType"),
				ua.AttributeIDDisplayName: ua.MustVariant("YArrayItemType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 12021),
		},
		&node{
			id: ua.NewNumericNodeID(0, 12038),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("VariableType_16"),
				ua.AttributeIDBrowseName:  ua.MustVariant("XYArrayItemType"),
				ua.AttributeIDDisplayName: ua.MustVariant("XYArrayItemType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 12021),
		},
		&node{
			id: ua.NewNumericNodeID(0, 12047),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("VariableType_16"),
				ua.AttributeIDBrowseName:  ua.MustVariant("ImageItemType"),
				ua.AttributeIDDisplayName: ua.MustVariant("ImageItemType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 12021),
		},
		&node{
			id: ua.NewNumericNodeID(0, 12057),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("VariableType_16"),
				ua.AttributeIDBrowseName:  ua.MustVariant("CubeItemType"),
				ua.AttributeIDDisplayName: ua.MustVariant("CubeItemType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 12021),
		},
		&node{
			id: ua.NewNumericNodeID(0, 12068),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("VariableType_16"),
				ua.AttributeIDBrowseName:  ua.MustVariant("NDimensionArrayItemType"),
				ua.AttributeIDDisplayName: ua.MustVariant("NDimensionArrayItemType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 12021),
		},
		&node{
			id: ua.NewNumericNodeID(0, 8995),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("VariableType_16"),
				ua.AttributeIDBrowseName:  ua.MustVariant("TwoStateVariableType"),
				ua.AttributeIDDisplayName: ua.MustVariant("TwoStateVariableType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 2755),
		},
		&node{
			id: ua.NewNumericNodeID(0, 9002),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("VariableType_16"),
				ua.AttributeIDBrowseName:  ua.MustVariant("ConditionVariableType"),
				ua.AttributeIDDisplayName: ua.MustVariant("ConditionVariableType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 63),
		},
		&node{
			id: ua.NewNumericNodeID(0, 9004),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ReferenceType_32"),
				ua.AttributeIDBrowseName:  ua.MustVariant("HasTrueSubState"),
				ua.AttributeIDDisplayName: ua.MustVariant("HasTrueSubState"),
				ua.AttributeIDInverseName: ua.MustVariant(&ua.LocalizedText{Locale: "", Text: "IsTrueSubStateOf"}),
			},
			superTypeID: ua.NewNumericNodeID(0, 32),
		},
		&node{
			id: ua.NewNumericNodeID(0, 9005),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ReferenceType_32"),
				ua.AttributeIDBrowseName:  ua.MustVariant("HasFalseSubState"),
				ua.AttributeIDDisplayName: ua.MustVariant("HasFalseSubState"),
				ua.AttributeIDInverseName: ua.MustVariant(&ua.LocalizedText{Locale: "", Text: "IsFalseSubStateOf"}),
			},
			superTypeID: ua.NewNumericNodeID(0, 32),
		},
		&node{
			id: ua.NewNumericNodeID(0, 16361),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ReferenceType_32"),
				ua.AttributeIDBrowseName:  ua.MustVariant("HasAlarmSuppressionGroup"),
				ua.AttributeIDDisplayName: ua.MustVariant("HasAlarmSuppressionGroup"),
				ua.AttributeIDInverseName: ua.MustVariant(&ua.LocalizedText{Locale: "", Text: "IsAlarmSuppressionGroupOf"}),
			},
			superTypeID: ua.NewNumericNodeID(0, 47),
		},
		&node{
			id: ua.NewNumericNodeID(0, 16362),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ReferenceType_32"),
				ua.AttributeIDBrowseName:  ua.MustVariant("AlarmGroupMember"),
				ua.AttributeIDDisplayName: ua.MustVariant("AlarmGroupMember"),
				ua.AttributeIDInverseName: ua.MustVariant(&ua.LocalizedText{Locale: "", Text: "MemberOfAlarmGroup"}),
			},
			superTypeID: ua.NewNumericNodeID(0, 35),
		},
		&node{
			id: ua.NewNumericNodeID(0, 32059),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ReferenceType_32"),
				ua.AttributeIDBrowseName:  ua.MustVariant("AlarmSuppressionGroupMember"),
				ua.AttributeIDDisplayName: ua.MustVariant("AlarmSuppressionGroupMember"),
				ua.AttributeIDInverseName: ua.MustVariant(&ua.LocalizedText{Locale: "", Text: "MemberOfAlarmSuppressionGroup"}),
			},
			superTypeID: ua.NewNumericNodeID(0, 16362),
		},
		&node{
			id: ua.NewNumericNodeID(0, 2830),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ObjectType_8"),
				ua.AttributeIDBrowseName:  ua.MustVariant("DialogConditionType"),
				ua.AttributeIDDisplayName: ua.MustVariant("DialogConditionType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 2782),
		},
		&node{
			id: ua.NewNumericNodeID(0, 2881),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ObjectType_8"),
				ua.AttributeIDBrowseName:  ua.MustVariant("AcknowledgeableConditionType"),
				ua.AttributeIDDisplayName: ua.MustVariant("AcknowledgeableConditionType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 2782),
		},
		&node{
			id: ua.NewNumericNodeID(0, 2915),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ObjectType_8"),
				ua.AttributeIDBrowseName:  ua.MustVariant("AlarmConditionType"),
				ua.AttributeIDDisplayName: ua.MustVariant("AlarmConditionType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 2881),
		},
		&node{
			id: ua.NewNumericNodeID(0, 16405),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ObjectType_8"),
				ua.AttributeIDBrowseName:  ua.MustVariant("AlarmGroupType"),
				ua.AttributeIDDisplayName: ua.MustVariant("AlarmGroupType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 61),
		},
		&node{
			id: ua.NewNumericNodeID(0, 32064),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ObjectType_8"),
				ua.AttributeIDBrowseName:  ua.MustVariant("AlarmSuppressionGroupType"),
				ua.AttributeIDDisplayName: ua.MustVariant("AlarmSuppressionGroupType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 16405),
		},
		&node{
			id: ua.NewNumericNodeID(0, 2929),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ObjectType_8"),
				ua.AttributeIDBrowseName:  ua.MustVariant("ShelvedStateMachineType"),
				ua.AttributeIDDisplayName: ua.MustVariant("ShelvedStateMachineType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 2771),
		},
		&node{
			id: ua.NewNumericNodeID(0, 2955),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ObjectType_8"),
				ua.AttributeIDBrowseName:  ua.MustVariant("LimitAlarmType"),
				ua.AttributeIDDisplayName: ua.MustVariant("LimitAlarmType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 2915),
		},
		&node{
			id: ua.NewNumericNodeID(0, 9318),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ObjectType_8"),
				ua.AttributeIDBrowseName:  ua.MustVariant("ExclusiveLimitStateMachineType"),
				ua.AttributeIDDisplayName: ua.MustVariant("ExclusiveLimitStateMachineType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 2771),
		},
		&node{
			id: ua.NewNumericNodeID(0, 9341),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ObjectType_8"),
				ua.AttributeIDBrowseName:  ua.MustVariant("ExclusiveLimitAlarmType"),
				ua.AttributeIDDisplayName: ua.MustVariant("ExclusiveLimitAlarmType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 2955),
		},
		&node{
			id: ua.NewNumericNodeID(0, 9906),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ObjectType_8"),
				ua.AttributeIDBrowseName:  ua.MustVariant("NonExclusiveLimitAlarmType"),
				ua.AttributeIDDisplayName: ua.MustVariant("NonExclusiveLimitAlarmType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 2955),
		},
		&node{
			id: ua.NewNumericNodeID(0, 10060),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ObjectType_8"),
				ua.AttributeIDBrowseName:  ua.MustVariant("NonExclusiveLevelAlarmType"),
				ua.AttributeIDDisplayName: ua.MustVariant("NonExclusiveLevelAlarmType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 9906),
		},
		&node{
			id: ua.NewNumericNodeID(0, 9482),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ObjectType_8"),
				ua.AttributeIDBrowseName:  ua.MustVariant("ExclusiveLevelAlarmType"),
				ua.AttributeIDDisplayName: ua.MustVariant("ExclusiveLevelAlarmType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 9341),
		},
		&node{
			id: ua.NewNumericNodeID(0, 10368),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ObjectType_8"),
				ua.AttributeIDBrowseName:  ua.MustVariant("NonExclusiveDeviationAlarmType"),
				ua.AttributeIDDisplayName: ua.MustVariant("NonExclusiveDeviationAlarmType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 9906),
		},
		&node{
			id: ua.NewNumericNodeID(0, 10214),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ObjectType_8"),
				ua.AttributeIDBrowseName:  ua.MustVariant("NonExclusiveRateOfChangeAlarmType"),
				ua.AttributeIDDisplayName: ua.MustVariant("NonExclusiveRateOfChangeAlarmType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 9906),
		},
		&node{
			id: ua.NewNumericNodeID(0, 9764),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ObjectType_8"),
				ua.AttributeIDBrowseName:  ua.MustVariant("ExclusiveDeviationAlarmType"),
				ua.AttributeIDDisplayName: ua.MustVariant("ExclusiveDeviationAlarmType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 9341),
		},
		&node{
			id: ua.NewNumericNodeID(0, 9623),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ObjectType_8"),
				ua.AttributeIDBrowseName:  ua.MustVariant("ExclusiveRateOfChangeAlarmType"),
				ua.AttributeIDDisplayName: ua.MustVariant("ExclusiveRateOfChangeAlarmType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 9341),
		},
		&node{
			id: ua.NewNumericNodeID(0, 10523),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ObjectType_8"),
				ua.AttributeIDBrowseName:  ua.MustVariant("DiscreteAlarmType"),
				ua.AttributeIDDisplayName: ua.MustVariant("DiscreteAlarmType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 2915),
		},
		&node{
			id: ua.NewNumericNodeID(0, 10637),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ObjectType_8"),
				ua.AttributeIDBrowseName:  ua.MustVariant("OffNormalAlarmType"),
				ua.AttributeIDDisplayName: ua.MustVariant("OffNormalAlarmType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 10523),
		},
		&node{
			id: ua.NewNumericNodeID(0, 11753),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ObjectType_8"),
				ua.AttributeIDBrowseName:  ua.MustVariant("SystemOffNormalAlarmType"),
				ua.AttributeIDDisplayName: ua.MustVariant("SystemOffNormalAlarmType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 10637),
		},
		&node{
			id: ua.NewNumericNodeID(0, 10751),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ObjectType_8"),
				ua.AttributeIDBrowseName:  ua.MustVariant("TripAlarmType"),
				ua.AttributeIDDisplayName: ua.MustVariant("TripAlarmType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 10637),
		},
		&node{
			id: ua.NewNumericNodeID(0, 18347),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ObjectType_8"),
				ua.AttributeIDBrowseName:  ua.MustVariant("InstrumentDiagnosticAlarmType"),
				ua.AttributeIDDisplayName: ua.MustVariant("InstrumentDiagnosticAlarmType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 10637),
		},
		&node{
			id: ua.NewNumericNodeID(0, 18496),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ObjectType_8"),
				ua.AttributeIDBrowseName:  ua.MustVariant("SystemDiagnosticAlarmType"),
				ua.AttributeIDDisplayName: ua.MustVariant("SystemDiagnosticAlarmType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 10637),
		},
		&node{
			id: ua.NewNumericNodeID(0, 13225),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ObjectType_8"),
				ua.AttributeIDBrowseName:  ua.MustVariant("CertificateExpirationAlarmType"),
				ua.AttributeIDDisplayName: ua.MustVariant("CertificateExpirationAlarmType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 11753),
		},
		&node{
			id: ua.NewNumericNodeID(0, 17080),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ObjectType_8"),
				ua.AttributeIDBrowseName:  ua.MustVariant("DiscrepancyAlarmType"),
				ua.AttributeIDDisplayName: ua.MustVariant("DiscrepancyAlarmType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 2915),
		},
		&node{
			id: ua.NewNumericNodeID(0, 2790),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ObjectType_8"),
				ua.AttributeIDBrowseName:  ua.MustVariant("AuditConditionEventType"),
				ua.AttributeIDDisplayName: ua.MustVariant("AuditConditionEventType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 2127),
		},
		&node{
			id: ua.NewNumericNodeID(0, 2803),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ObjectType_8"),
				ua.AttributeIDBrowseName:  ua.MustVariant("AuditConditionEnableEventType"),
				ua.AttributeIDDisplayName: ua.MustVariant("AuditConditionEnableEventType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 2790),
		},
		&node{
			id: ua.NewNumericNodeID(0, 2829),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ObjectType_8"),
				ua.AttributeIDBrowseName:  ua.MustVariant("AuditConditionCommentEventType"),
				ua.AttributeIDDisplayName: ua.MustVariant("AuditConditionCommentEventType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 2790),
		},
		&node{
			id: ua.NewNumericNodeID(0, 8927),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ObjectType_8"),
				ua.AttributeIDBrowseName:  ua.MustVariant("AuditConditionRespondEventType"),
				ua.AttributeIDDisplayName: ua.MustVariant("AuditConditionRespondEventType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 2790),
		},
		&node{
			id: ua.NewNumericNodeID(0, 8944),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ObjectType_8"),
				ua.AttributeIDBrowseName:  ua.MustVariant("AuditConditionAcknowledgeEventType"),
				ua.AttributeIDDisplayName: ua.MustVariant("AuditConditionAcknowledgeEventType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 2790),
		},
		&node{
			id: ua.NewNumericNodeID(0, 8961),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ObjectType_8"),
				ua.AttributeIDBrowseName:  ua.MustVariant("AuditConditionConfirmEventType"),
				ua.AttributeIDDisplayName: ua.MustVariant("AuditConditionConfirmEventType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 2790),
		},
		&node{
			id: ua.NewNumericNodeID(0, 11093),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ObjectType_8"),
				ua.AttributeIDBrowseName:  ua.MustVariant("AuditConditionShelvingEventType"),
				ua.AttributeIDDisplayName: ua.MustVariant("AuditConditionShelvingEventType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 2790),
		},
		&node{
			id: ua.NewNumericNodeID(0, 17225),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ObjectType_8"),
				ua.AttributeIDBrowseName:  ua.MustVariant("AuditConditionSuppressionEventType"),
				ua.AttributeIDDisplayName: ua.MustVariant("AuditConditionSuppressionEventType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 2790),
		},
		&node{
			id: ua.NewNumericNodeID(0, 17242),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ObjectType_8"),
				ua.AttributeIDBrowseName:  ua.MustVariant("AuditConditionSilenceEventType"),
				ua.AttributeIDDisplayName: ua.MustVariant("AuditConditionSilenceEventType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 2790),
		},
		&node{
			id: ua.NewNumericNodeID(0, 15013),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ObjectType_8"),
				ua.AttributeIDBrowseName:  ua.MustVariant("AuditConditionResetEventType"),
				ua.AttributeIDDisplayName: ua.MustVariant("AuditConditionResetEventType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 2790),
		},
		&node{
			id: ua.NewNumericNodeID(0, 17259),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ObjectType_8"),
				ua.AttributeIDBrowseName:  ua.MustVariant("AuditConditionOutOfServiceEventType"),
				ua.AttributeIDDisplayName: ua.MustVariant("AuditConditionOutOfServiceEventType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 2790),
		},
		&node{
			id: ua.NewNumericNodeID(0, 9006),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ReferenceType_32"),
				ua.AttributeIDBrowseName:  ua.MustVariant("HasCondition"),
				ua.AttributeIDDisplayName: ua.MustVariant("HasCondition"),
				ua.AttributeIDInverseName: ua.MustVariant(&ua.LocalizedText{Locale: "", Text: "IsConditionOf"}),
			},
			superTypeID: ua.NewNumericNodeID(0, 32),
		},
		&node{
			id: ua.NewNumericNodeID(0, 17276),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ReferenceType_32"),
				ua.AttributeIDBrowseName:  ua.MustVariant("HasEffectDisable"),
				ua.AttributeIDDisplayName: ua.MustVariant("HasEffectDisable"),
				ua.AttributeIDInverseName: ua.MustVariant(&ua.LocalizedText{Locale: "", Text: "MayBeDisabledBy"}),
			},
			superTypeID: ua.NewNumericNodeID(0, 54),
		},
		&node{
			id: ua.NewNumericNodeID(0, 17983),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ReferenceType_32"),
				ua.AttributeIDBrowseName:  ua.MustVariant("HasEffectEnable"),
				ua.AttributeIDDisplayName: ua.MustVariant("HasEffectEnable"),
				ua.AttributeIDInverseName: ua.MustVariant(&ua.LocalizedText{Locale: "", Text: "MayBeEnabledBy"}),
			},
			superTypeID: ua.NewNumericNodeID(0, 54),
		},
		&node{
			id: ua.NewNumericNodeID(0, 17984),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ReferenceType_32"),
				ua.AttributeIDBrowseName:  ua.MustVariant("HasEffectSuppressed"),
				ua.AttributeIDDisplayName: ua.MustVariant("HasEffectSuppressed"),
				ua.AttributeIDInverseName: ua.MustVariant(&ua.LocalizedText{Locale: "", Text: "MayBeSuppressedBy"}),
			},
			superTypeID: ua.NewNumericNodeID(0, 54),
		},
		&node{
			id: ua.NewNumericNodeID(0, 17985),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ReferenceType_32"),
				ua.AttributeIDBrowseName:  ua.MustVariant("HasEffectUnsuppressed"),
				ua.AttributeIDDisplayName: ua.MustVariant("HasEffectUnsuppressed"),
				ua.AttributeIDInverseName: ua.MustVariant(&ua.LocalizedText{Locale: "", Text: "MayBeUnsuppressedBy"}),
			},
			superTypeID: ua.NewNumericNodeID(0, 54),
		},
		&node{
			id: ua.NewNumericNodeID(0, 17279),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ObjectType_8"),
				ua.AttributeIDBrowseName:  ua.MustVariant("AlarmMetricsType"),
				ua.AttributeIDDisplayName: ua.MustVariant("AlarmMetricsType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 58),
		},
		&node{
			id: ua.NewNumericNodeID(0, 17277),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("VariableType_16"),
				ua.AttributeIDBrowseName:  ua.MustVariant("AlarmRateVariableType"),
				ua.AttributeIDDisplayName: ua.MustVariant("AlarmRateVariableType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 63),
		},
		&node{
			id: ua.NewNumericNodeID(0, 32244),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("VariableType_16"),
				ua.AttributeIDBrowseName:  ua.MustVariant("AlarmStateVariableType"),
				ua.AttributeIDDisplayName: ua.MustVariant("AlarmStateVariableType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 63),
		},
		&node{
			id: ua.NewNumericNodeID(0, 32251),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("AlarmMask"),
				ua.AttributeIDDisplayName: ua.MustVariant("AlarmMask"),
			},
			superTypeID: ua.NewNumericNodeID(0, 5),
		},
		&node{
			id: ua.NewNumericNodeID(0, 2391),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ObjectType_8"),
				ua.AttributeIDBrowseName:  ua.MustVariant("ProgramStateMachineType"),
				ua.AttributeIDDisplayName: ua.MustVariant("ProgramStateMachineType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 2771),
		},
		&node{
			id: ua.NewNumericNodeID(0, 3806),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ObjectType_8"),
				ua.AttributeIDBrowseName:  ua.MustVariant("ProgramTransitionAuditEventType"),
				ua.AttributeIDDisplayName: ua.MustVariant("ProgramTransitionAuditEventType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 2315),
		},
		&node{
			id: ua.NewNumericNodeID(0, 2380),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("VariableType_16"),
				ua.AttributeIDBrowseName:  ua.MustVariant("ProgramDiagnosticType"),
				ua.AttributeIDDisplayName: ua.MustVariant("ProgramDiagnosticType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 63),
		},
		&node{
			id: ua.NewNumericNodeID(0, 15383),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("VariableType_16"),
				ua.AttributeIDBrowseName:  ua.MustVariant("ProgramDiagnostic2Type"),
				ua.AttributeIDDisplayName: ua.MustVariant("ProgramDiagnostic2Type"),
			},
			superTypeID: ua.NewNumericNodeID(0, 63),
		},
		&node{
			id: ua.NewNumericNodeID(0, 11214),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Variable_2"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Annotations"),
				ua.AttributeIDDisplayName: ua.MustVariant("Annotations"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 2318),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ObjectType_8"),
				ua.AttributeIDBrowseName:  ua.MustVariant("HistoricalDataConfigurationType"),
				ua.AttributeIDDisplayName: ua.MustVariant("HistoricalDataConfigurationType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 58),
		},
		&node{
			id: ua.NewNumericNodeID(0, 11202),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("HA Configuration"),
				ua.AttributeIDDisplayName: ua.MustVariant("HA Configuration"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 11215),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Variable_2"),
				ua.AttributeIDBrowseName:  ua.MustVariant("HistoricalEventFilter"),
				ua.AttributeIDDisplayName: ua.MustVariant("HistoricalEventFilter"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 2330),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ObjectType_8"),
				ua.AttributeIDBrowseName:  ua.MustVariant("HistoryServerCapabilitiesType"),
				ua.AttributeIDDisplayName: ua.MustVariant("HistoryServerCapabilitiesType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 58),
		},
		&node{
			id: ua.NewNumericNodeID(0, 12522),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ObjectType_8"),
				ua.AttributeIDBrowseName:  ua.MustVariant("TrustListType"),
				ua.AttributeIDDisplayName: ua.MustVariant("TrustListType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 11575),
		},
		&node{
			id: ua.NewNumericNodeID(0, 23564),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("TrustListValidationOptions"),
				ua.AttributeIDDisplayName: ua.MustVariant("TrustListValidationOptions"),
			},
			superTypeID: ua.NewNumericNodeID(0, 7),
		},
		&node{
			id: ua.NewNumericNodeID(0, 12552),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("TrustListMasks"),
				ua.AttributeIDDisplayName: ua.MustVariant("TrustListMasks"),
			},
			superTypeID: ua.NewNumericNodeID(0, 29),
		},
		&node{
			id: ua.NewNumericNodeID(0, 12554),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("TrustListDataType"),
				ua.AttributeIDDisplayName: ua.MustVariant("TrustListDataType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 22),
		},
		&node{
			id: ua.NewNumericNodeID(0, 19297),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ObjectType_8"),
				ua.AttributeIDBrowseName:  ua.MustVariant("TrustListOutOfDateAlarmType"),
				ua.AttributeIDDisplayName: ua.MustVariant("TrustListOutOfDateAlarmType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 11753),
		},
		&node{
			id: ua.NewNumericNodeID(0, 12555),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ObjectType_8"),
				ua.AttributeIDBrowseName:  ua.MustVariant("CertificateGroupType"),
				ua.AttributeIDDisplayName: ua.MustVariant("CertificateGroupType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 58),
		},
		&node{
			id: ua.NewNumericNodeID(0, 13813),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ObjectType_8"),
				ua.AttributeIDBrowseName:  ua.MustVariant("CertificateGroupFolderType"),
				ua.AttributeIDDisplayName: ua.MustVariant("CertificateGroupFolderType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 61),
		},
		&node{
			id: ua.NewNumericNodeID(0, 12558),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ObjectType_8"),
				ua.AttributeIDBrowseName:  ua.MustVariant("HttpsCertificateType"),
				ua.AttributeIDDisplayName: ua.MustVariant("HttpsCertificateType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 12556),
		},
		&node{
			id: ua.NewNumericNodeID(0, 15181),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ObjectType_8"),
				ua.AttributeIDBrowseName:  ua.MustVariant("UserCredentialCertificateType"),
				ua.AttributeIDDisplayName: ua.MustVariant("UserCredentialCertificateType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 12556),
		},
		&node{
			id: ua.NewNumericNodeID(0, 12559),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ObjectType_8"),
				ua.AttributeIDBrowseName:  ua.MustVariant("RsaMinApplicationCertificateType"),
				ua.AttributeIDDisplayName: ua.MustVariant("RsaMinApplicationCertificateType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 12557),
		},
		&node{
			id: ua.NewNumericNodeID(0, 12560),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ObjectType_8"),
				ua.AttributeIDBrowseName:  ua.MustVariant("RsaSha256ApplicationCertificateType"),
				ua.AttributeIDDisplayName: ua.MustVariant("RsaSha256ApplicationCertificateType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 12557),
		},
		&node{
			id: ua.NewNumericNodeID(0, 23537),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ObjectType_8"),
				ua.AttributeIDBrowseName:  ua.MustVariant("EccApplicationCertificateType"),
				ua.AttributeIDDisplayName: ua.MustVariant("EccApplicationCertificateType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 12557),
		},
		&node{
			id: ua.NewNumericNodeID(0, 23538),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ObjectType_8"),
				ua.AttributeIDBrowseName:  ua.MustVariant("EccNistP256ApplicationCertificateType"),
				ua.AttributeIDDisplayName: ua.MustVariant("EccNistP256ApplicationCertificateType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 23537),
		},
		&node{
			id: ua.NewNumericNodeID(0, 23539),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ObjectType_8"),
				ua.AttributeIDBrowseName:  ua.MustVariant("EccNistP384ApplicationCertificateType"),
				ua.AttributeIDDisplayName: ua.MustVariant("EccNistP384ApplicationCertificateType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 23537),
		},
		&node{
			id: ua.NewNumericNodeID(0, 23540),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ObjectType_8"),
				ua.AttributeIDBrowseName:  ua.MustVariant("EccBrainpoolP256r1ApplicationCertificateType"),
				ua.AttributeIDDisplayName: ua.MustVariant("EccBrainpoolP256r1ApplicationCertificateType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 23537),
		},
		&node{
			id: ua.NewNumericNodeID(0, 23541),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ObjectType_8"),
				ua.AttributeIDBrowseName:  ua.MustVariant("EccBrainpoolP384r1ApplicationCertificateType"),
				ua.AttributeIDDisplayName: ua.MustVariant("EccBrainpoolP384r1ApplicationCertificateType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 23537),
		},
		&node{
			id: ua.NewNumericNodeID(0, 23542),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ObjectType_8"),
				ua.AttributeIDBrowseName:  ua.MustVariant("EccCurve25519ApplicationCertificateType"),
				ua.AttributeIDDisplayName: ua.MustVariant("EccCurve25519ApplicationCertificateType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 23537),
		},
		&node{
			id: ua.NewNumericNodeID(0, 23543),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ObjectType_8"),
				ua.AttributeIDBrowseName:  ua.MustVariant("EccCurve448ApplicationCertificateType"),
				ua.AttributeIDDisplayName: ua.MustVariant("EccCurve448ApplicationCertificateType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 23537),
		},
		&node{
			id: ua.NewNumericNodeID(0, 32285),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("TransactionErrorType"),
				ua.AttributeIDDisplayName: ua.MustVariant("TransactionErrorType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 22),
		},
		&node{
			id: ua.NewNumericNodeID(0, 32286),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ObjectType_8"),
				ua.AttributeIDBrowseName:  ua.MustVariant("TransactionDiagnosticsType"),
				ua.AttributeIDDisplayName: ua.MustVariant("TransactionDiagnosticsType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 58),
		},
		&node{
			id: ua.NewNumericNodeID(0, 12581),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ObjectType_8"),
				ua.AttributeIDBrowseName:  ua.MustVariant("ServerConfigurationType"),
				ua.AttributeIDDisplayName: ua.MustVariant("ServerConfigurationType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 58),
		},
		&node{
			id: ua.NewNumericNodeID(0, 12637),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("ServerConfiguration"),
				ua.AttributeIDDisplayName: ua.MustVariant("ServerConfiguration"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 17496),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ObjectType_8"),
				ua.AttributeIDBrowseName:  ua.MustVariant("KeyCredentialConfigurationFolderType"),
				ua.AttributeIDDisplayName: ua.MustVariant("KeyCredentialConfigurationFolderType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 61),
		},
		&node{
			id: ua.NewNumericNodeID(0, 18155),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("KeyCredentialConfiguration"),
				ua.AttributeIDDisplayName: ua.MustVariant("KeyCredentialConfiguration"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 18001),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ObjectType_8"),
				ua.AttributeIDBrowseName:  ua.MustVariant("KeyCredentialConfigurationType"),
				ua.AttributeIDDisplayName: ua.MustVariant("KeyCredentialConfigurationType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 58),
		},
		&node{
			id: ua.NewNumericNodeID(0, 18029),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ObjectType_8"),
				ua.AttributeIDBrowseName:  ua.MustVariant("KeyCredentialUpdatedAuditEventType"),
				ua.AttributeIDDisplayName: ua.MustVariant("KeyCredentialUpdatedAuditEventType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 18011),
		},
		&node{
			id: ua.NewNumericNodeID(0, 18047),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ObjectType_8"),
				ua.AttributeIDBrowseName:  ua.MustVariant("KeyCredentialDeletedAuditEventType"),
				ua.AttributeIDDisplayName: ua.MustVariant("KeyCredentialDeletedAuditEventType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 18011),
		},
		&node{
			id: ua.NewNumericNodeID(0, 23556),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ObjectType_8"),
				ua.AttributeIDBrowseName:  ua.MustVariant("AuthorizationServicesConfigurationFolderType"),
				ua.AttributeIDDisplayName: ua.MustVariant("AuthorizationServicesConfigurationFolderType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 61),
		},
		&node{
			id: ua.NewNumericNodeID(0, 17732),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("AuthorizationServices"),
				ua.AttributeIDDisplayName: ua.MustVariant("AuthorizationServices"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 17852),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ObjectType_8"),
				ua.AttributeIDBrowseName:  ua.MustVariant("AuthorizationServiceConfigurationType"),
				ua.AttributeIDDisplayName: ua.MustVariant("AuthorizationServiceConfigurationType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 58),
		},
		&node{
			id: ua.NewNumericNodeID(0, 11187),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ObjectType_8"),
				ua.AttributeIDBrowseName:  ua.MustVariant("AggregateConfigurationType"),
				ua.AttributeIDDisplayName: ua.MustVariant("AggregateConfigurationType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 58),
		},
		&node{
			id: ua.NewNumericNodeID(0, 2341),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Interpolative"),
				ua.AttributeIDDisplayName: ua.MustVariant("Interpolative"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 2342),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Average"),
				ua.AttributeIDDisplayName: ua.MustVariant("Average"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 2343),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("TimeAverage"),
				ua.AttributeIDDisplayName: ua.MustVariant("TimeAverage"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 11285),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("TimeAverage2"),
				ua.AttributeIDDisplayName: ua.MustVariant("TimeAverage2"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 2344),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Total"),
				ua.AttributeIDDisplayName: ua.MustVariant("Total"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 11304),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Total2"),
				ua.AttributeIDDisplayName: ua.MustVariant("Total2"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 2346),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Minimum"),
				ua.AttributeIDDisplayName: ua.MustVariant("Minimum"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 2347),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Maximum"),
				ua.AttributeIDDisplayName: ua.MustVariant("Maximum"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 2348),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("MinimumActualTime"),
				ua.AttributeIDDisplayName: ua.MustVariant("MinimumActualTime"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 2349),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("MaximumActualTime"),
				ua.AttributeIDDisplayName: ua.MustVariant("MaximumActualTime"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 2350),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Range"),
				ua.AttributeIDDisplayName: ua.MustVariant("Range"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 11286),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Minimum2"),
				ua.AttributeIDDisplayName: ua.MustVariant("Minimum2"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 11287),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Maximum2"),
				ua.AttributeIDDisplayName: ua.MustVariant("Maximum2"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 11305),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("MinimumActualTime2"),
				ua.AttributeIDDisplayName: ua.MustVariant("MinimumActualTime2"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 11306),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("MaximumActualTime2"),
				ua.AttributeIDDisplayName: ua.MustVariant("MaximumActualTime2"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 11288),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Range2"),
				ua.AttributeIDDisplayName: ua.MustVariant("Range2"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 2351),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("AnnotationCount"),
				ua.AttributeIDDisplayName: ua.MustVariant("AnnotationCount"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 2352),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Count"),
				ua.AttributeIDDisplayName: ua.MustVariant("Count"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 11307),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("DurationInStateZero"),
				ua.AttributeIDDisplayName: ua.MustVariant("DurationInStateZero"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 11308),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("DurationInStateNonZero"),
				ua.AttributeIDDisplayName: ua.MustVariant("DurationInStateNonZero"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 2355),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("NumberOfTransitions"),
				ua.AttributeIDDisplayName: ua.MustVariant("NumberOfTransitions"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 2357),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Start"),
				ua.AttributeIDDisplayName: ua.MustVariant("Start"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 2358),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("End"),
				ua.AttributeIDDisplayName: ua.MustVariant("End"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 2359),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Delta"),
				ua.AttributeIDDisplayName: ua.MustVariant("Delta"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 11505),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("StartBound"),
				ua.AttributeIDDisplayName: ua.MustVariant("StartBound"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 11506),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("EndBound"),
				ua.AttributeIDDisplayName: ua.MustVariant("EndBound"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 11507),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("DeltaBounds"),
				ua.AttributeIDDisplayName: ua.MustVariant("DeltaBounds"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 2360),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("DurationGood"),
				ua.AttributeIDDisplayName: ua.MustVariant("DurationGood"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 2361),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("DurationBad"),
				ua.AttributeIDDisplayName: ua.MustVariant("DurationBad"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 2362),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("PercentGood"),
				ua.AttributeIDDisplayName: ua.MustVariant("PercentGood"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 2363),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("PercentBad"),
				ua.AttributeIDDisplayName: ua.MustVariant("PercentBad"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 2364),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("WorstQuality"),
				ua.AttributeIDDisplayName: ua.MustVariant("WorstQuality"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 11292),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("WorstQuality2"),
				ua.AttributeIDDisplayName: ua.MustVariant("WorstQuality2"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 11426),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("StandardDeviationSample"),
				ua.AttributeIDDisplayName: ua.MustVariant("StandardDeviationSample"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 11427),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("StandardDeviationPopulation"),
				ua.AttributeIDDisplayName: ua.MustVariant("StandardDeviationPopulation"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 11428),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("VarianceSample"),
				ua.AttributeIDDisplayName: ua.MustVariant("VarianceSample"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 11429),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("VariancePopulation"),
				ua.AttributeIDDisplayName: ua.MustVariant("VariancePopulation"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15487),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("StructureDescription"),
				ua.AttributeIDDisplayName: ua.MustVariant("StructureDescription"),
			},
			superTypeID: ua.NewNumericNodeID(0, 14525),
		},
		&node{
			id: ua.NewNumericNodeID(0, 15488),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("EnumDescription"),
				ua.AttributeIDDisplayName: ua.MustVariant("EnumDescription"),
			},
			superTypeID: ua.NewNumericNodeID(0, 14525),
		},
		&node{
			id: ua.NewNumericNodeID(0, 15005),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("SimpleTypeDescription"),
				ua.AttributeIDDisplayName: ua.MustVariant("SimpleTypeDescription"),
			},
			superTypeID: ua.NewNumericNodeID(0, 14525),
		},
		&node{
			id: ua.NewNumericNodeID(0, 15006),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("UABinaryFileDataType"),
				ua.AttributeIDDisplayName: ua.MustVariant("UABinaryFileDataType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 15534),
		},
		&node{
			id: ua.NewNumericNodeID(0, 24105),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("PortableQualifiedName"),
				ua.AttributeIDDisplayName: ua.MustVariant("PortableQualifiedName"),
			},
			superTypeID: ua.NewNumericNodeID(0, 22),
		},
		&node{
			id: ua.NewNumericNodeID(0, 24106),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("PortableNodeId"),
				ua.AttributeIDDisplayName: ua.MustVariant("PortableNodeId"),
			},
			superTypeID: ua.NewNumericNodeID(0, 22),
		},
		&node{
			id: ua.NewNumericNodeID(0, 24107),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("UnsignedRationalNumber"),
				ua.AttributeIDDisplayName: ua.MustVariant("UnsignedRationalNumber"),
			},
			superTypeID: ua.NewNumericNodeID(0, 22),
		},
		&node{
			id: ua.NewNumericNodeID(0, 14647),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("PubSubState"),
				ua.AttributeIDDisplayName: ua.MustVariant("PubSubState"),
			},
			superTypeID: ua.NewNumericNodeID(0, 29),
		},
		&node{
			id: ua.NewNumericNodeID(0, 14523),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("DataSetMetaDataType"),
				ua.AttributeIDDisplayName: ua.MustVariant("DataSetMetaDataType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 15534),
		},
		&node{
			id: ua.NewNumericNodeID(0, 14524),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("FieldMetaData"),
				ua.AttributeIDDisplayName: ua.MustVariant("FieldMetaData"),
			},
			superTypeID: ua.NewNumericNodeID(0, 22),
		},
		&node{
			id: ua.NewNumericNodeID(0, 15904),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("DataSetFieldFlags"),
				ua.AttributeIDDisplayName: ua.MustVariant("DataSetFieldFlags"),
			},
			superTypeID: ua.NewNumericNodeID(0, 5),
		},
		&node{
			id: ua.NewNumericNodeID(0, 14593),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("ConfigurationVersionDataType"),
				ua.AttributeIDDisplayName: ua.MustVariant("ConfigurationVersionDataType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 22),
		},
		&node{
			id: ua.NewNumericNodeID(0, 15578),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("PublishedDataSetDataType"),
				ua.AttributeIDDisplayName: ua.MustVariant("PublishedDataSetDataType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 22),
		},
		&node{
			id: ua.NewNumericNodeID(0, 14273),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("PublishedVariableDataType"),
				ua.AttributeIDDisplayName: ua.MustVariant("PublishedVariableDataType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 22),
		},
		&node{
			id: ua.NewNumericNodeID(0, 15581),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("PublishedDataItemsDataType"),
				ua.AttributeIDDisplayName: ua.MustVariant("PublishedDataItemsDataType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 15580),
		},
		&node{
			id: ua.NewNumericNodeID(0, 15582),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("PublishedEventsDataType"),
				ua.AttributeIDDisplayName: ua.MustVariant("PublishedEventsDataType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 15580),
		},
		&node{
			id: ua.NewNumericNodeID(0, 25269),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("PublishedDataSetCustomSourceDataType"),
				ua.AttributeIDDisplayName: ua.MustVariant("PublishedDataSetCustomSourceDataType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 15580),
		},
		&node{
			id: ua.NewNumericNodeID(0, 15583),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("DataSetFieldContentMask"),
				ua.AttributeIDDisplayName: ua.MustVariant("DataSetFieldContentMask"),
			},
			superTypeID: ua.NewNumericNodeID(0, 7),
		},
		&node{
			id: ua.NewNumericNodeID(0, 15597),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("DataSetWriterDataType"),
				ua.AttributeIDDisplayName: ua.MustVariant("DataSetWriterDataType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 22),
		},
		&node{
			id: ua.NewNumericNodeID(0, 15480),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("WriterGroupDataType"),
				ua.AttributeIDDisplayName: ua.MustVariant("WriterGroupDataType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 15609),
		},
		&node{
			id: ua.NewNumericNodeID(0, 15617),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("PubSubConnectionDataType"),
				ua.AttributeIDDisplayName: ua.MustVariant("PubSubConnectionDataType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 22),
		},
		&node{
			id: ua.NewNumericNodeID(0, 15510),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("NetworkAddressUrlDataType"),
				ua.AttributeIDDisplayName: ua.MustVariant("NetworkAddressUrlDataType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 15502),
		},
		&node{
			id: ua.NewNumericNodeID(0, 15520),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("ReaderGroupDataType"),
				ua.AttributeIDDisplayName: ua.MustVariant("ReaderGroupDataType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 15609),
		},
		&node{
			id: ua.NewNumericNodeID(0, 15623),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("DataSetReaderDataType"),
				ua.AttributeIDDisplayName: ua.MustVariant("DataSetReaderDataType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 22),
		},
		&node{
			id: ua.NewNumericNodeID(0, 15631),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("TargetVariablesDataType"),
				ua.AttributeIDDisplayName: ua.MustVariant("TargetVariablesDataType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 15630),
		},
		&node{
			id: ua.NewNumericNodeID(0, 14744),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("FieldTargetDataType"),
				ua.AttributeIDDisplayName: ua.MustVariant("FieldTargetDataType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 22),
		},
		&node{
			id: ua.NewNumericNodeID(0, 15874),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("OverrideValueHandling"),
				ua.AttributeIDDisplayName: ua.MustVariant("OverrideValueHandling"),
			},
			superTypeID: ua.NewNumericNodeID(0, 29),
		},
		&node{
			id: ua.NewNumericNodeID(0, 15635),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("SubscribedDataSetMirrorDataType"),
				ua.AttributeIDDisplayName: ua.MustVariant("SubscribedDataSetMirrorDataType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 15630),
		},
		&node{
			id: ua.NewNumericNodeID(0, 15530),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("PubSubConfigurationDataType"),
				ua.AttributeIDDisplayName: ua.MustVariant("PubSubConfigurationDataType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 22),
		},
		&node{
			id: ua.NewNumericNodeID(0, 23599),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("StandaloneSubscribedDataSetRefDataType"),
				ua.AttributeIDDisplayName: ua.MustVariant("StandaloneSubscribedDataSetRefDataType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 15630),
		},
		&node{
			id: ua.NewNumericNodeID(0, 23600),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("StandaloneSubscribedDataSetDataType"),
				ua.AttributeIDDisplayName: ua.MustVariant("StandaloneSubscribedDataSetDataType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 15630),
		},
		&node{
			id: ua.NewNumericNodeID(0, 23601),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("SecurityGroupDataType"),
				ua.AttributeIDDisplayName: ua.MustVariant("SecurityGroupDataType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 22),
		},
		&node{
			id: ua.NewNumericNodeID(0, 25270),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("PubSubKeyPushTargetDataType"),
				ua.AttributeIDDisplayName: ua.MustVariant("PubSubKeyPushTargetDataType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 22),
		},
		&node{
			id: ua.NewNumericNodeID(0, 23602),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("PubSubConfiguration2DataType"),
				ua.AttributeIDDisplayName: ua.MustVariant("PubSubConfiguration2DataType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 15530),
		},
		&node{
			id: ua.NewNumericNodeID(0, 20408),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("DataSetOrderingType"),
				ua.AttributeIDDisplayName: ua.MustVariant("DataSetOrderingType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 29),
		},
		&node{
			id: ua.NewNumericNodeID(0, 15642),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("UadpNetworkMessageContentMask"),
				ua.AttributeIDDisplayName: ua.MustVariant("UadpNetworkMessageContentMask"),
			},
			superTypeID: ua.NewNumericNodeID(0, 7),
		},
		&node{
			id: ua.NewNumericNodeID(0, 15645),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("UadpWriterGroupMessageDataType"),
				ua.AttributeIDDisplayName: ua.MustVariant("UadpWriterGroupMessageDataType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 15616),
		},
		&node{
			id: ua.NewNumericNodeID(0, 15646),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("UadpDataSetMessageContentMask"),
				ua.AttributeIDDisplayName: ua.MustVariant("UadpDataSetMessageContentMask"),
			},
			superTypeID: ua.NewNumericNodeID(0, 7),
		},
		&node{
			id: ua.NewNumericNodeID(0, 15652),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("UadpDataSetWriterMessageDataType"),
				ua.AttributeIDDisplayName: ua.MustVariant("UadpDataSetWriterMessageDataType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 15605),
		},
		&node{
			id: ua.NewNumericNodeID(0, 15653),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("UadpDataSetReaderMessageDataType"),
				ua.AttributeIDDisplayName: ua.MustVariant("UadpDataSetReaderMessageDataType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 15629),
		},
		&node{
			id: ua.NewNumericNodeID(0, 15654),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("JsonNetworkMessageContentMask"),
				ua.AttributeIDDisplayName: ua.MustVariant("JsonNetworkMessageContentMask"),
			},
			superTypeID: ua.NewNumericNodeID(0, 7),
		},
		&node{
			id: ua.NewNumericNodeID(0, 15657),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("JsonWriterGroupMessageDataType"),
				ua.AttributeIDDisplayName: ua.MustVariant("JsonWriterGroupMessageDataType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 15616),
		},
		&node{
			id: ua.NewNumericNodeID(0, 15658),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("JsonDataSetMessageContentMask"),
				ua.AttributeIDDisplayName: ua.MustVariant("JsonDataSetMessageContentMask"),
			},
			superTypeID: ua.NewNumericNodeID(0, 7),
		},
		&node{
			id: ua.NewNumericNodeID(0, 15664),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("JsonDataSetWriterMessageDataType"),
				ua.AttributeIDDisplayName: ua.MustVariant("JsonDataSetWriterMessageDataType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 15605),
		},
		&node{
			id: ua.NewNumericNodeID(0, 15665),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("JsonDataSetReaderMessageDataType"),
				ua.AttributeIDDisplayName: ua.MustVariant("JsonDataSetReaderMessageDataType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 15629),
		},
		&node{
			id: ua.NewNumericNodeID(0, 23605),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("TransmitQosPriorityDataType"),
				ua.AttributeIDDisplayName: ua.MustVariant("TransmitQosPriorityDataType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 23604),
		},
		&node{
			id: ua.NewNumericNodeID(0, 23609),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("ReceiveQosPriorityDataType"),
				ua.AttributeIDDisplayName: ua.MustVariant("ReceiveQosPriorityDataType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 23608),
		},
		&node{
			id: ua.NewNumericNodeID(0, 17467),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("DatagramConnectionTransportDataType"),
				ua.AttributeIDDisplayName: ua.MustVariant("DatagramConnectionTransportDataType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 15618),
		},
		&node{
			id: ua.NewNumericNodeID(0, 23612),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("DatagramConnectionTransport2DataType"),
				ua.AttributeIDDisplayName: ua.MustVariant("DatagramConnectionTransport2DataType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 17467),
		},
		&node{
			id: ua.NewNumericNodeID(0, 15532),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("DatagramWriterGroupTransportDataType"),
				ua.AttributeIDDisplayName: ua.MustVariant("DatagramWriterGroupTransportDataType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 15611),
		},
		&node{
			id: ua.NewNumericNodeID(0, 23613),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("DatagramWriterGroupTransport2DataType"),
				ua.AttributeIDDisplayName: ua.MustVariant("DatagramWriterGroupTransport2DataType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 15532),
		},
		&node{
			id: ua.NewNumericNodeID(0, 23614),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("DatagramDataSetReaderTransportDataType"),
				ua.AttributeIDDisplayName: ua.MustVariant("DatagramDataSetReaderTransportDataType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 15628),
		},
		&node{
			id: ua.NewNumericNodeID(0, 15007),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("BrokerConnectionTransportDataType"),
				ua.AttributeIDDisplayName: ua.MustVariant("BrokerConnectionTransportDataType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 15618),
		},
		&node{
			id: ua.NewNumericNodeID(0, 15008),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("BrokerTransportQualityOfService"),
				ua.AttributeIDDisplayName: ua.MustVariant("BrokerTransportQualityOfService"),
			},
			superTypeID: ua.NewNumericNodeID(0, 29),
		},
		&node{
			id: ua.NewNumericNodeID(0, 15667),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("BrokerWriterGroupTransportDataType"),
				ua.AttributeIDDisplayName: ua.MustVariant("BrokerWriterGroupTransportDataType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 15611),
		},
		&node{
			id: ua.NewNumericNodeID(0, 15669),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("BrokerDataSetWriterTransportDataType"),
				ua.AttributeIDDisplayName: ua.MustVariant("BrokerDataSetWriterTransportDataType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 15598),
		},
		&node{
			id: ua.NewNumericNodeID(0, 15670),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("BrokerDataSetReaderTransportDataType"),
				ua.AttributeIDDisplayName: ua.MustVariant("BrokerDataSetReaderTransportDataType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 15628),
		},
		&node{
			id: ua.NewNumericNodeID(0, 15906),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ObjectType_8"),
				ua.AttributeIDBrowseName:  ua.MustVariant("PubSubKeyServiceType"),
				ua.AttributeIDDisplayName: ua.MustVariant("PubSubKeyServiceType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 58),
		},
		&node{
			id: ua.NewNumericNodeID(0, 15452),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ObjectType_8"),
				ua.AttributeIDBrowseName:  ua.MustVariant("SecurityGroupFolderType"),
				ua.AttributeIDDisplayName: ua.MustVariant("SecurityGroupFolderType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 61),
		},
		&node{
			id: ua.NewNumericNodeID(0, 15471),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ObjectType_8"),
				ua.AttributeIDBrowseName:  ua.MustVariant("SecurityGroupType"),
				ua.AttributeIDDisplayName: ua.MustVariant("SecurityGroupType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 58),
		},
		&node{
			id: ua.NewNumericNodeID(0, 25345),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ReferenceType_32"),
				ua.AttributeIDBrowseName:  ua.MustVariant("HasPushedSecurityGroup"),
				ua.AttributeIDDisplayName: ua.MustVariant("HasPushedSecurityGroup"),
				ua.AttributeIDInverseName: ua.MustVariant(&ua.LocalizedText{Locale: "", Text: "HasPushTarget"}),
			},
			superTypeID: ua.NewNumericNodeID(0, 33),
		},
		&node{
			id: ua.NewNumericNodeID(0, 25337),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ObjectType_8"),
				ua.AttributeIDBrowseName:  ua.MustVariant("PubSubKeyPushTargetType"),
				ua.AttributeIDDisplayName: ua.MustVariant("PubSubKeyPushTargetType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 58),
		},
		&node{
			id: ua.NewNumericNodeID(0, 25346),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ObjectType_8"),
				ua.AttributeIDBrowseName:  ua.MustVariant("PubSubKeyPushTargetFolderType"),
				ua.AttributeIDDisplayName: ua.MustVariant("PubSubKeyPushTargetFolderType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 61),
		},
		&node{
			id: ua.NewNumericNodeID(0, 14416),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ObjectType_8"),
				ua.AttributeIDBrowseName:  ua.MustVariant("PublishSubscribeType"),
				ua.AttributeIDDisplayName: ua.MustVariant("PublishSubscribeType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 15906),
		},
		&node{
			id: ua.NewNumericNodeID(0, 14443),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("PublishSubscribe"),
				ua.AttributeIDDisplayName: ua.MustVariant("PublishSubscribe"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 32405),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("DataSetClasses"),
				ua.AttributeIDDisplayName: ua.MustVariant("DataSetClasses"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 14476),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ReferenceType_32"),
				ua.AttributeIDBrowseName:  ua.MustVariant("HasPubSubConnection"),
				ua.AttributeIDDisplayName: ua.MustVariant("HasPubSubConnection"),
				ua.AttributeIDInverseName: ua.MustVariant(&ua.LocalizedText{Locale: "", Text: "PubSubConnectionOf"}),
			},
			superTypeID: ua.NewNumericNodeID(0, 47),
		},
		&node{
			id: ua.NewNumericNodeID(0, 25482),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ObjectType_8"),
				ua.AttributeIDBrowseName:  ua.MustVariant("PubSubConfigurationType"),
				ua.AttributeIDDisplayName: ua.MustVariant("PubSubConfigurationType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 11575),
		},
		&node{
			id: ua.NewNumericNodeID(0, 25517),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("PubSubConfigurationRefMask"),
				ua.AttributeIDDisplayName: ua.MustVariant("PubSubConfigurationRefMask"),
			},
			superTypeID: ua.NewNumericNodeID(0, 7),
		},
		&node{
			id: ua.NewNumericNodeID(0, 25519),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("PubSubConfigurationRefDataType"),
				ua.AttributeIDDisplayName: ua.MustVariant("PubSubConfigurationRefDataType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 22),
		},
		&node{
			id: ua.NewNumericNodeID(0, 25520),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("PubSubConfigurationValueDataType"),
				ua.AttributeIDDisplayName: ua.MustVariant("PubSubConfigurationValueDataType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 22),
		},
		&node{
			id: ua.NewNumericNodeID(0, 14509),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ObjectType_8"),
				ua.AttributeIDBrowseName:  ua.MustVariant("PublishedDataSetType"),
				ua.AttributeIDDisplayName: ua.MustVariant("PublishedDataSetType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 58),
		},
		&node{
			id: ua.NewNumericNodeID(0, 15489),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ObjectType_8"),
				ua.AttributeIDBrowseName:  ua.MustVariant("ExtensionFieldsType"),
				ua.AttributeIDDisplayName: ua.MustVariant("ExtensionFieldsType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 58),
		},
		&node{
			id: ua.NewNumericNodeID(0, 14936),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ReferenceType_32"),
				ua.AttributeIDBrowseName:  ua.MustVariant("DataSetToWriter"),
				ua.AttributeIDDisplayName: ua.MustVariant("DataSetToWriter"),
				ua.AttributeIDInverseName: ua.MustVariant(&ua.LocalizedText{Locale: "", Text: "WriterToDataSet"}),
			},
			superTypeID: ua.NewNumericNodeID(0, 33),
		},
		&node{
			id: ua.NewNumericNodeID(0, 14534),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ObjectType_8"),
				ua.AttributeIDBrowseName:  ua.MustVariant("PublishedDataItemsType"),
				ua.AttributeIDDisplayName: ua.MustVariant("PublishedDataItemsType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 14509),
		},
		&node{
			id: ua.NewNumericNodeID(0, 14572),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ObjectType_8"),
				ua.AttributeIDBrowseName:  ua.MustVariant("PublishedEventsType"),
				ua.AttributeIDDisplayName: ua.MustVariant("PublishedEventsType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 14509),
		},
		&node{
			id: ua.NewNumericNodeID(0, 14477),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ObjectType_8"),
				ua.AttributeIDBrowseName:  ua.MustVariant("DataSetFolderType"),
				ua.AttributeIDDisplayName: ua.MustVariant("DataSetFolderType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 61),
		},
		&node{
			id: ua.NewNumericNodeID(0, 14209),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ObjectType_8"),
				ua.AttributeIDBrowseName:  ua.MustVariant("PubSubConnectionType"),
				ua.AttributeIDDisplayName: ua.MustVariant("PubSubConnectionType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 58),
		},
		&node{
			id: ua.NewNumericNodeID(0, 17725),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ObjectType_8"),
				ua.AttributeIDBrowseName:  ua.MustVariant("WriterGroupType"),
				ua.AttributeIDDisplayName: ua.MustVariant("WriterGroupType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 14232),
		},
		&node{
			id: ua.NewNumericNodeID(0, 15296),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ReferenceType_32"),
				ua.AttributeIDBrowseName:  ua.MustVariant("HasDataSetWriter"),
				ua.AttributeIDDisplayName: ua.MustVariant("HasDataSetWriter"),
				ua.AttributeIDInverseName: ua.MustVariant(&ua.LocalizedText{Locale: "", Text: "IsWriterInGroup"}),
			},
			superTypeID: ua.NewNumericNodeID(0, 47),
		},
		&node{
			id: ua.NewNumericNodeID(0, 18804),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ReferenceType_32"),
				ua.AttributeIDBrowseName:  ua.MustVariant("HasWriterGroup"),
				ua.AttributeIDDisplayName: ua.MustVariant("HasWriterGroup"),
				ua.AttributeIDInverseName: ua.MustVariant(&ua.LocalizedText{Locale: "", Text: "IsWriterGroupOf"}),
			},
			superTypeID: ua.NewNumericNodeID(0, 47),
		},
		&node{
			id: ua.NewNumericNodeID(0, 17999),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ObjectType_8"),
				ua.AttributeIDBrowseName:  ua.MustVariant("ReaderGroupType"),
				ua.AttributeIDDisplayName: ua.MustVariant("ReaderGroupType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 14232),
		},
		&node{
			id: ua.NewNumericNodeID(0, 15297),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ReferenceType_32"),
				ua.AttributeIDBrowseName:  ua.MustVariant("HasDataSetReader"),
				ua.AttributeIDDisplayName: ua.MustVariant("HasDataSetReader"),
				ua.AttributeIDInverseName: ua.MustVariant(&ua.LocalizedText{Locale: "", Text: "IsReaderInGroup"}),
			},
			superTypeID: ua.NewNumericNodeID(0, 47),
		},
		&node{
			id: ua.NewNumericNodeID(0, 18805),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ReferenceType_32"),
				ua.AttributeIDBrowseName:  ua.MustVariant("HasReaderGroup"),
				ua.AttributeIDDisplayName: ua.MustVariant("HasReaderGroup"),
				ua.AttributeIDInverseName: ua.MustVariant(&ua.LocalizedText{Locale: "", Text: "IsReaderGroupOf"}),
			},
			superTypeID: ua.NewNumericNodeID(0, 47),
		},
		&node{
			id: ua.NewNumericNodeID(0, 15298),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ObjectType_8"),
				ua.AttributeIDBrowseName:  ua.MustVariant("DataSetWriterType"),
				ua.AttributeIDDisplayName: ua.MustVariant("DataSetWriterType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 58),
		},
		&node{
			id: ua.NewNumericNodeID(0, 15306),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ObjectType_8"),
				ua.AttributeIDBrowseName:  ua.MustVariant("DataSetReaderType"),
				ua.AttributeIDDisplayName: ua.MustVariant("DataSetReaderType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 58),
		},
		&node{
			id: ua.NewNumericNodeID(0, 15108),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ObjectType_8"),
				ua.AttributeIDBrowseName:  ua.MustVariant("SubscribedDataSetType"),
				ua.AttributeIDDisplayName: ua.MustVariant("SubscribedDataSetType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 58),
		},
		&node{
			id: ua.NewNumericNodeID(0, 15111),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ObjectType_8"),
				ua.AttributeIDBrowseName:  ua.MustVariant("TargetVariablesType"),
				ua.AttributeIDDisplayName: ua.MustVariant("TargetVariablesType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 15108),
		},
		&node{
			id: ua.NewNumericNodeID(0, 15127),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ObjectType_8"),
				ua.AttributeIDBrowseName:  ua.MustVariant("SubscribedDataSetMirrorType"),
				ua.AttributeIDDisplayName: ua.MustVariant("SubscribedDataSetMirrorType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 15108),
		},
		&node{
			id: ua.NewNumericNodeID(0, 23795),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ObjectType_8"),
				ua.AttributeIDBrowseName:  ua.MustVariant("SubscribedDataSetFolderType"),
				ua.AttributeIDDisplayName: ua.MustVariant("SubscribedDataSetFolderType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 61),
		},
		&node{
			id: ua.NewNumericNodeID(0, 23828),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ObjectType_8"),
				ua.AttributeIDBrowseName:  ua.MustVariant("StandaloneSubscribedDataSetType"),
				ua.AttributeIDDisplayName: ua.MustVariant("StandaloneSubscribedDataSetType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 58),
		},
		&node{
			id: ua.NewNumericNodeID(0, 14643),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ObjectType_8"),
				ua.AttributeIDBrowseName:  ua.MustVariant("PubSubStatusType"),
				ua.AttributeIDDisplayName: ua.MustVariant("PubSubStatusType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 58),
		},
		&node{
			id: ua.NewNumericNodeID(0, 19723),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("DiagnosticsLevel"),
				ua.AttributeIDDisplayName: ua.MustVariant("DiagnosticsLevel"),
			},
			superTypeID: ua.NewNumericNodeID(0, 29),
		},
		&node{
			id: ua.NewNumericNodeID(0, 19725),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("VariableType_16"),
				ua.AttributeIDBrowseName:  ua.MustVariant("PubSubDiagnosticsCounterType"),
				ua.AttributeIDDisplayName: ua.MustVariant("PubSubDiagnosticsCounterType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 63),
		},
		&node{
			id: ua.NewNumericNodeID(0, 19730),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("PubSubDiagnosticsCounterClassification"),
				ua.AttributeIDDisplayName: ua.MustVariant("PubSubDiagnosticsCounterClassification"),
			},
			superTypeID: ua.NewNumericNodeID(0, 29),
		},
		&node{
			id: ua.NewNumericNodeID(0, 19732),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ObjectType_8"),
				ua.AttributeIDBrowseName:  ua.MustVariant("PubSubDiagnosticsRootType"),
				ua.AttributeIDDisplayName: ua.MustVariant("PubSubDiagnosticsRootType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 19677),
		},
		&node{
			id: ua.NewNumericNodeID(0, 19786),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ObjectType_8"),
				ua.AttributeIDBrowseName:  ua.MustVariant("PubSubDiagnosticsConnectionType"),
				ua.AttributeIDDisplayName: ua.MustVariant("PubSubDiagnosticsConnectionType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 19677),
		},
		&node{
			id: ua.NewNumericNodeID(0, 19834),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ObjectType_8"),
				ua.AttributeIDBrowseName:  ua.MustVariant("PubSubDiagnosticsWriterGroupType"),
				ua.AttributeIDDisplayName: ua.MustVariant("PubSubDiagnosticsWriterGroupType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 19677),
		},
		&node{
			id: ua.NewNumericNodeID(0, 19903),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ObjectType_8"),
				ua.AttributeIDBrowseName:  ua.MustVariant("PubSubDiagnosticsReaderGroupType"),
				ua.AttributeIDDisplayName: ua.MustVariant("PubSubDiagnosticsReaderGroupType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 19677),
		},
		&node{
			id: ua.NewNumericNodeID(0, 19968),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ObjectType_8"),
				ua.AttributeIDBrowseName:  ua.MustVariant("PubSubDiagnosticsDataSetWriterType"),
				ua.AttributeIDDisplayName: ua.MustVariant("PubSubDiagnosticsDataSetWriterType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 19677),
		},
		&node{
			id: ua.NewNumericNodeID(0, 20027),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ObjectType_8"),
				ua.AttributeIDBrowseName:  ua.MustVariant("PubSubDiagnosticsDataSetReaderType"),
				ua.AttributeIDDisplayName: ua.MustVariant("PubSubDiagnosticsDataSetReaderType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 19677),
		},
		&node{
			id: ua.NewNumericNodeID(0, 23832),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ObjectType_8"),
				ua.AttributeIDBrowseName:  ua.MustVariant("PubSubCapabilitiesType"),
				ua.AttributeIDDisplayName: ua.MustVariant("PubSubCapabilitiesType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 58),
		},
		&node{
			id: ua.NewNumericNodeID(0, 21105),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ObjectType_8"),
				ua.AttributeIDBrowseName:  ua.MustVariant("UadpWriterGroupMessageType"),
				ua.AttributeIDDisplayName: ua.MustVariant("UadpWriterGroupMessageType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 17998),
		},
		&node{
			id: ua.NewNumericNodeID(0, 21111),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ObjectType_8"),
				ua.AttributeIDBrowseName:  ua.MustVariant("UadpDataSetWriterMessageType"),
				ua.AttributeIDDisplayName: ua.MustVariant("UadpDataSetWriterMessageType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 21096),
		},
		&node{
			id: ua.NewNumericNodeID(0, 21116),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ObjectType_8"),
				ua.AttributeIDBrowseName:  ua.MustVariant("UadpDataSetReaderMessageType"),
				ua.AttributeIDDisplayName: ua.MustVariant("UadpDataSetReaderMessageType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 21104),
		},
		&node{
			id: ua.NewNumericNodeID(0, 21126),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ObjectType_8"),
				ua.AttributeIDBrowseName:  ua.MustVariant("JsonWriterGroupMessageType"),
				ua.AttributeIDDisplayName: ua.MustVariant("JsonWriterGroupMessageType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 17998),
		},
		&node{
			id: ua.NewNumericNodeID(0, 21128),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ObjectType_8"),
				ua.AttributeIDBrowseName:  ua.MustVariant("JsonDataSetWriterMessageType"),
				ua.AttributeIDDisplayName: ua.MustVariant("JsonDataSetWriterMessageType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 21096),
		},
		&node{
			id: ua.NewNumericNodeID(0, 21130),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ObjectType_8"),
				ua.AttributeIDBrowseName:  ua.MustVariant("JsonDataSetReaderMessageType"),
				ua.AttributeIDDisplayName: ua.MustVariant("JsonDataSetReaderMessageType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 21104),
		},
		&node{
			id: ua.NewNumericNodeID(0, 15064),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ObjectType_8"),
				ua.AttributeIDBrowseName:  ua.MustVariant("DatagramConnectionTransportType"),
				ua.AttributeIDDisplayName: ua.MustVariant("DatagramConnectionTransportType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 17721),
		},
		&node{
			id: ua.NewNumericNodeID(0, 21133),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ObjectType_8"),
				ua.AttributeIDBrowseName:  ua.MustVariant("DatagramWriterGroupTransportType"),
				ua.AttributeIDDisplayName: ua.MustVariant("DatagramWriterGroupTransportType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 17997),
		},
		&node{
			id: ua.NewNumericNodeID(0, 24016),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ObjectType_8"),
				ua.AttributeIDBrowseName:  ua.MustVariant("DatagramDataSetReaderTransportType"),
				ua.AttributeIDDisplayName: ua.MustVariant("DatagramDataSetReaderTransportType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 15319),
		},
		&node{
			id: ua.NewNumericNodeID(0, 15155),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ObjectType_8"),
				ua.AttributeIDBrowseName:  ua.MustVariant("BrokerConnectionTransportType"),
				ua.AttributeIDDisplayName: ua.MustVariant("BrokerConnectionTransportType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 17721),
		},
		&node{
			id: ua.NewNumericNodeID(0, 21136),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ObjectType_8"),
				ua.AttributeIDBrowseName:  ua.MustVariant("BrokerWriterGroupTransportType"),
				ua.AttributeIDDisplayName: ua.MustVariant("BrokerWriterGroupTransportType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 17997),
		},
		&node{
			id: ua.NewNumericNodeID(0, 21138),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ObjectType_8"),
				ua.AttributeIDBrowseName:  ua.MustVariant("BrokerDataSetWriterTransportType"),
				ua.AttributeIDDisplayName: ua.MustVariant("BrokerDataSetWriterTransportType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 15305),
		},
		&node{
			id: ua.NewNumericNodeID(0, 21142),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ObjectType_8"),
				ua.AttributeIDBrowseName:  ua.MustVariant("BrokerDataSetReaderTransportType"),
				ua.AttributeIDDisplayName: ua.MustVariant("BrokerDataSetReaderTransportType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 15319),
		},
		&node{
			id: ua.NewNumericNodeID(0, 21147),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ObjectType_8"),
				ua.AttributeIDBrowseName:  ua.MustVariant("NetworkAddressUrlType"),
				ua.AttributeIDDisplayName: ua.MustVariant("NetworkAddressUrlType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 21145),
		},
		&node{
			id: ua.NewNumericNodeID(0, 23455),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ObjectType_8"),
				ua.AttributeIDBrowseName:  ua.MustVariant("AliasNameType"),
				ua.AttributeIDDisplayName: ua.MustVariant("AliasNameType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 58),
		},
		&node{
			id: ua.NewNumericNodeID(0, 23456),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ObjectType_8"),
				ua.AttributeIDBrowseName:  ua.MustVariant("AliasNameCategoryType"),
				ua.AttributeIDDisplayName: ua.MustVariant("AliasNameCategoryType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 61),
		},
		&node{
			id: ua.NewNumericNodeID(0, 23468),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("AliasNameDataType"),
				ua.AttributeIDDisplayName: ua.MustVariant("AliasNameDataType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 22),
		},
		&node{
			id: ua.NewNumericNodeID(0, 23469),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ReferenceType_32"),
				ua.AttributeIDBrowseName:  ua.MustVariant("AliasFor"),
				ua.AttributeIDDisplayName: ua.MustVariant("AliasFor"),
				ua.AttributeIDInverseName: ua.MustVariant(&ua.LocalizedText{Locale: "", Text: "HasAlias"}),
			},
			superTypeID: ua.NewNumericNodeID(0, 32),
		},
		&node{
			id: ua.NewNumericNodeID(0, 23470),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Aliases"),
				ua.AttributeIDDisplayName: ua.MustVariant("Aliases"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 23479),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("TagVariables"),
				ua.AttributeIDDisplayName: ua.MustVariant("TagVariables"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 23488),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Topics"),
				ua.AttributeIDDisplayName: ua.MustVariant("Topics"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 24264),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ObjectType_8"),
				ua.AttributeIDBrowseName:  ua.MustVariant("UserManagementType"),
				ua.AttributeIDDisplayName: ua.MustVariant("UserManagementType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 58),
		},
		&node{
			id: ua.NewNumericNodeID(0, 24277),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("PasswordOptionsMask"),
				ua.AttributeIDDisplayName: ua.MustVariant("PasswordOptionsMask"),
			},
			superTypeID: ua.NewNumericNodeID(0, 7),
		},
		&node{
			id: ua.NewNumericNodeID(0, 24279),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("UserConfigurationMask"),
				ua.AttributeIDDisplayName: ua.MustVariant("UserConfigurationMask"),
			},
			superTypeID: ua.NewNumericNodeID(0, 7),
		},
		&node{
			id: ua.NewNumericNodeID(0, 24281),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("UserManagementDataType"),
				ua.AttributeIDDisplayName: ua.MustVariant("UserManagementDataType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 22),
		},
		&node{
			id: ua.NewNumericNodeID(0, 24290),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("UserManagement"),
				ua.AttributeIDDisplayName: ua.MustVariant("UserManagement"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 19077),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("VariableType_16"),
				ua.AttributeIDBrowseName:  ua.MustVariant("MultiStateDictionaryEntryDiscreteBaseType"),
				ua.AttributeIDDisplayName: ua.MustVariant("MultiStateDictionaryEntryDiscreteBaseType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 11238),
		},
		&node{
			id: ua.NewNumericNodeID(0, 19084),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("VariableType_16"),
				ua.AttributeIDBrowseName:  ua.MustVariant("MultiStateDictionaryEntryDiscreteType"),
				ua.AttributeIDDisplayName: ua.MustVariant("MultiStateDictionaryEntryDiscreteType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 19077),
		},
		&node{
			id: ua.NewNumericNodeID(0, 25726),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("EncodedTicket"),
				ua.AttributeIDDisplayName: ua.MustVariant("EncodedTicket"),
			},
			superTypeID: ua.NewNumericNodeID(0, 12),
		},
		&node{
			id: ua.NewNumericNodeID(0, 25731),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ObjectType_8"),
				ua.AttributeIDBrowseName:  ua.MustVariant("ApplicationConfigurationType"),
				ua.AttributeIDDisplayName: ua.MustVariant("ApplicationConfigurationType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 12581),
		},
		&node{
			id: ua.NewNumericNodeID(0, 26871),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ObjectType_8"),
				ua.AttributeIDBrowseName:  ua.MustVariant("ProvisionableDeviceType"),
				ua.AttributeIDDisplayName: ua.MustVariant("ProvisionableDeviceType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 58),
		},
		&node{
			id: ua.NewNumericNodeID(0, 29878),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("ProvisionableDevice"),
				ua.AttributeIDDisplayName: ua.MustVariant("ProvisionableDevice"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 24210),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Duplex"),
				ua.AttributeIDDisplayName: ua.MustVariant("Duplex"),
			},
			superTypeID: ua.NewNumericNodeID(0, 29),
		},
		&node{
			id: ua.NewNumericNodeID(0, 24212),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("InterfaceAdminStatus"),
				ua.AttributeIDDisplayName: ua.MustVariant("InterfaceAdminStatus"),
			},
			superTypeID: ua.NewNumericNodeID(0, 29),
		},
		&node{
			id: ua.NewNumericNodeID(0, 24214),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("InterfaceOperStatus"),
				ua.AttributeIDDisplayName: ua.MustVariant("InterfaceOperStatus"),
			},
			superTypeID: ua.NewNumericNodeID(0, 29),
		},
		&node{
			id: ua.NewNumericNodeID(0, 24216),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("NegotiationStatus"),
				ua.AttributeIDDisplayName: ua.MustVariant("NegotiationStatus"),
			},
			superTypeID: ua.NewNumericNodeID(0, 29),
		},
		&node{
			id: ua.NewNumericNodeID(0, 24218),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("TsnFailureCode"),
				ua.AttributeIDDisplayName: ua.MustVariant("TsnFailureCode"),
			},
			superTypeID: ua.NewNumericNodeID(0, 29),
		},
		&node{
			id: ua.NewNumericNodeID(0, 24220),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("TsnStreamState"),
				ua.AttributeIDDisplayName: ua.MustVariant("TsnStreamState"),
			},
			superTypeID: ua.NewNumericNodeID(0, 29),
		},
		&node{
			id: ua.NewNumericNodeID(0, 24222),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("TsnTalkerStatus"),
				ua.AttributeIDDisplayName: ua.MustVariant("TsnTalkerStatus"),
			},
			superTypeID: ua.NewNumericNodeID(0, 29),
		},
		&node{
			id: ua.NewNumericNodeID(0, 24224),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("TsnListenerStatus"),
				ua.AttributeIDDisplayName: ua.MustVariant("TsnListenerStatus"),
			},
			superTypeID: ua.NewNumericNodeID(0, 29),
		},
		&node{
			id: ua.NewNumericNodeID(0, 25220),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("PriorityMappingEntryType"),
				ua.AttributeIDDisplayName: ua.MustVariant("PriorityMappingEntryType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 22),
		},
		&node{
			id: ua.NewNumericNodeID(0, 24226),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Resources"),
				ua.AttributeIDDisplayName: ua.MustVariant("Resources"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 24227),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Communication"),
				ua.AttributeIDDisplayName: ua.MustVariant("Communication"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 24228),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("MappingTables"),
				ua.AttributeIDDisplayName: ua.MustVariant("MappingTables"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 24229),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("NetworkInterfaces"),
				ua.AttributeIDDisplayName: ua.MustVariant("NetworkInterfaces"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 24230),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Streams"),
				ua.AttributeIDDisplayName: ua.MustVariant("Streams"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 24231),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("TalkerStreams"),
				ua.AttributeIDDisplayName: ua.MustVariant("TalkerStreams"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 24232),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("ListenerStreams"),
				ua.AttributeIDDisplayName: ua.MustVariant("ListenerStreams"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 25221),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ObjectType_8"),
				ua.AttributeIDBrowseName:  ua.MustVariant("IetfBaseNetworkInterfaceType"),
				ua.AttributeIDDisplayName: ua.MustVariant("IetfBaseNetworkInterfaceType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 58),
		},
		&node{
			id: ua.NewNumericNodeID(0, 25227),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ObjectType_8"),
				ua.AttributeIDBrowseName:  ua.MustVariant("PriorityMappingTableType"),
				ua.AttributeIDDisplayName: ua.MustVariant("PriorityMappingTableType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 58),
		},
		&node{
			id: ua.NewNumericNodeID(0, 25237),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ReferenceType_32"),
				ua.AttributeIDBrowseName:  ua.MustVariant("UsesPriorityMappingTable"),
				ua.AttributeIDDisplayName: ua.MustVariant("UsesPriorityMappingTable"),
				ua.AttributeIDInverseName: ua.MustVariant(&ua.LocalizedText{Locale: "", Text: "UsedByNetworkInterface"}),
			},
			superTypeID: ua.NewNumericNodeID(0, 32),
		},
		&node{
			id: ua.NewNumericNodeID(0, 25238),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ReferenceType_32"),
				ua.AttributeIDBrowseName:  ua.MustVariant("HasLowerLayerInterface"),
				ua.AttributeIDDisplayName: ua.MustVariant("HasLowerLayerInterface"),
				ua.AttributeIDInverseName: ua.MustVariant(&ua.LocalizedText{Locale: "", Text: "HasHigherLayerInterface"}),
			},
			superTypeID: ua.NewNumericNodeID(0, 33),
		},
		&node{
			id: ua.NewNumericNodeID(0, 25253),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ReferenceType_32"),
				ua.AttributeIDBrowseName:  ua.MustVariant("IsExecutableOn"),
				ua.AttributeIDDisplayName: ua.MustVariant("IsExecutableOn"),
				ua.AttributeIDInverseName: ua.MustVariant(&ua.LocalizedText{Locale: "", Text: "CanExecute"}),
			},
			superTypeID: ua.NewNumericNodeID(0, 32),
		},
		&node{
			id: ua.NewNumericNodeID(0, 25254),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ReferenceType_32"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Controls"),
				ua.AttributeIDDisplayName: ua.MustVariant("Controls"),
				ua.AttributeIDInverseName: ua.MustVariant(&ua.LocalizedText{Locale: "", Text: "IsControlledBy"}),
			},
			superTypeID: ua.NewNumericNodeID(0, 33),
		},
		&node{
			id: ua.NewNumericNodeID(0, 25255),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ReferenceType_32"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Utilizes"),
				ua.AttributeIDDisplayName: ua.MustVariant("Utilizes"),
				ua.AttributeIDInverseName: ua.MustVariant(&ua.LocalizedText{Locale: "", Text: "IsUtilizedBy"}),
			},
			superTypeID: ua.NewNumericNodeID(0, 32),
		},
		&node{
			id: ua.NewNumericNodeID(0, 25265),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ReferenceType_32"),
				ua.AttributeIDBrowseName:  ua.MustVariant("IsExecutingOn"),
				ua.AttributeIDDisplayName: ua.MustVariant("IsExecutingOn"),
				ua.AttributeIDInverseName: ua.MustVariant(&ua.LocalizedText{Locale: "", Text: "Executes"}),
			},
			superTypeID: ua.NewNumericNodeID(0, 25255),
		},
		&node{
			id: ua.NewNumericNodeID(0, 25256),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ReferenceType_32"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Requires"),
				ua.AttributeIDDisplayName: ua.MustVariant("Requires"),
				ua.AttributeIDInverseName: ua.MustVariant(&ua.LocalizedText{Locale: "", Text: "IsRequiredBy"}),
			},
			superTypeID: ua.NewNumericNodeID(0, 33),
		},
		&node{
			id: ua.NewNumericNodeID(0, 25257),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ReferenceType_32"),
				ua.AttributeIDBrowseName:  ua.MustVariant("IsPhysicallyConnectedTo"),
				ua.AttributeIDDisplayName: ua.MustVariant("IsPhysicallyConnectedTo"),
			},
			superTypeID: ua.NewNumericNodeID(0, 32),
		},
		&node{
			id: ua.NewNumericNodeID(0, 25258),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ReferenceType_32"),
				ua.AttributeIDBrowseName:  ua.MustVariant("RepresentsSameEntityAs"),
				ua.AttributeIDDisplayName: ua.MustVariant("RepresentsSameEntityAs"),
			},
			superTypeID: ua.NewNumericNodeID(0, 32),
		},
		&node{
			id: ua.NewNumericNodeID(0, 25259),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ReferenceType_32"),
				ua.AttributeIDBrowseName:  ua.MustVariant("RepresentsSameHardwareAs"),
				ua.AttributeIDDisplayName: ua.MustVariant("RepresentsSameHardwareAs"),
			},
			superTypeID: ua.NewNumericNodeID(0, 25258),
		},
		&node{
			id: ua.NewNumericNodeID(0, 25260),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ReferenceType_32"),
				ua.AttributeIDBrowseName:  ua.MustVariant("RepresentsSameFunctionalityAs"),
				ua.AttributeIDDisplayName: ua.MustVariant("RepresentsSameFunctionalityAs"),
			},
			superTypeID: ua.NewNumericNodeID(0, 25258),
		},
		&node{
			id: ua.NewNumericNodeID(0, 25261),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ReferenceType_32"),
				ua.AttributeIDBrowseName:  ua.MustVariant("IsHostedBy"),
				ua.AttributeIDDisplayName: ua.MustVariant("IsHostedBy"),
				ua.AttributeIDInverseName: ua.MustVariant(&ua.LocalizedText{Locale: "", Text: "Hosts"}),
			},
			superTypeID: ua.NewNumericNodeID(0, 25255),
		},
		&node{
			id: ua.NewNumericNodeID(0, 25262),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ReferenceType_32"),
				ua.AttributeIDBrowseName:  ua.MustVariant("HasPhysicalComponent"),
				ua.AttributeIDDisplayName: ua.MustVariant("HasPhysicalComponent"),
				ua.AttributeIDInverseName: ua.MustVariant(&ua.LocalizedText{Locale: "", Text: "PhysicalComponentOf"}),
			},
			superTypeID: ua.NewNumericNodeID(0, 47),
		},
		&node{
			id: ua.NewNumericNodeID(0, 25263),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ReferenceType_32"),
				ua.AttributeIDBrowseName:  ua.MustVariant("HasContainedComponent"),
				ua.AttributeIDDisplayName: ua.MustVariant("HasContainedComponent"),
				ua.AttributeIDInverseName: ua.MustVariant(&ua.LocalizedText{Locale: "", Text: "ContainedComponentOf"}),
			},
			superTypeID: ua.NewNumericNodeID(0, 25262),
		},
		&node{
			id: ua.NewNumericNodeID(0, 25264),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ReferenceType_32"),
				ua.AttributeIDBrowseName:  ua.MustVariant("HasAttachedComponent"),
				ua.AttributeIDDisplayName: ua.MustVariant("HasAttachedComponent"),
				ua.AttributeIDInverseName: ua.MustVariant(&ua.LocalizedText{Locale: "", Text: "AttachedComponentOf"}),
			},
			superTypeID: ua.NewNumericNodeID(0, 25262),
		},
		&node{
			id: ua.NewNumericNodeID(0, 32679),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("ReferenceType_32"),
				ua.AttributeIDBrowseName:  ua.MustVariant("HasReferenceDescription"),
				ua.AttributeIDDisplayName: ua.MustVariant("HasReferenceDescription"),
				ua.AttributeIDInverseName: ua.MustVariant(&ua.LocalizedText{Locale: "", Text: "ReferenceDescriptionOf"}),
			},
			superTypeID: ua.NewNumericNodeID(0, 34),
		},
		&node{
			id: ua.NewNumericNodeID(0, 32657),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("VariableType_16"),
				ua.AttributeIDBrowseName:  ua.MustVariant("ReferenceDescriptionVariableType"),
				ua.AttributeIDDisplayName: ua.MustVariant("ReferenceDescriptionVariableType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 63),
		},
		&node{
			id: ua.NewNumericNodeID(0, 32659),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("ReferenceDescriptionDataType"),
				ua.AttributeIDDisplayName: ua.MustVariant("ReferenceDescriptionDataType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 22),
		},
		&node{
			id: ua.NewNumericNodeID(0, 32660),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("ReferenceListEntryDataType"),
				ua.AttributeIDDisplayName: ua.MustVariant("ReferenceListEntryDataType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 22),
		},
		&node{
			id: ua.NewNumericNodeID(0, 256),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("IdType"),
				ua.AttributeIDDisplayName: ua.MustVariant("IdType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 29),
		},
		&node{
			id: ua.NewNumericNodeID(0, 257),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("NodeClass"),
				ua.AttributeIDDisplayName: ua.MustVariant("NodeClass"),
			},
			superTypeID: ua.NewNumericNodeID(0, 29),
		},
		&node{
			id: ua.NewNumericNodeID(0, 94),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("PermissionType"),
				ua.AttributeIDDisplayName: ua.MustVariant("PermissionType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 7),
		},
		&node{
			id: ua.NewNumericNodeID(0, 15031),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("AccessLevelType"),
				ua.AttributeIDDisplayName: ua.MustVariant("AccessLevelType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 3),
		},
		&node{
			id: ua.NewNumericNodeID(0, 15406),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("AccessLevelExType"),
				ua.AttributeIDDisplayName: ua.MustVariant("AccessLevelExType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 7),
		},
		&node{
			id: ua.NewNumericNodeID(0, 15033),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("EventNotifierType"),
				ua.AttributeIDDisplayName: ua.MustVariant("EventNotifierType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 3),
		},
		&node{
			id: ua.NewNumericNodeID(0, 95),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("AccessRestrictionType"),
				ua.AttributeIDDisplayName: ua.MustVariant("AccessRestrictionType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 5),
		},
		&node{
			id: ua.NewNumericNodeID(0, 96),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("RolePermissionType"),
				ua.AttributeIDDisplayName: ua.MustVariant("RolePermissionType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 22),
		},
		&node{
			id: ua.NewNumericNodeID(0, 98),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("StructureType"),
				ua.AttributeIDDisplayName: ua.MustVariant("StructureType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 29),
		},
		&node{
			id: ua.NewNumericNodeID(0, 101),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("StructureField"),
				ua.AttributeIDDisplayName: ua.MustVariant("StructureField"),
			},
			superTypeID: ua.NewNumericNodeID(0, 22),
		},
		&node{
			id: ua.NewNumericNodeID(0, 99),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("StructureDefinition"),
				ua.AttributeIDDisplayName: ua.MustVariant("StructureDefinition"),
			},
			superTypeID: ua.NewNumericNodeID(0, 97),
		},
		&node{
			id: ua.NewNumericNodeID(0, 100),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("EnumDefinition"),
				ua.AttributeIDDisplayName: ua.MustVariant("EnumDefinition"),
			},
			superTypeID: ua.NewNumericNodeID(0, 97),
		},
		&node{
			id: ua.NewNumericNodeID(0, 296),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Argument"),
				ua.AttributeIDDisplayName: ua.MustVariant("Argument"),
			},
			superTypeID: ua.NewNumericNodeID(0, 22),
		},
		&node{
			id: ua.NewNumericNodeID(0, 7594),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("EnumValueType"),
				ua.AttributeIDDisplayName: ua.MustVariant("EnumValueType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 22),
		},
		&node{
			id: ua.NewNumericNodeID(0, 102),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("EnumField"),
				ua.AttributeIDDisplayName: ua.MustVariant("EnumField"),
			},
			superTypeID: ua.NewNumericNodeID(0, 7594),
		},
		&node{
			id: ua.NewNumericNodeID(0, 12877),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("NormalizedString"),
				ua.AttributeIDDisplayName: ua.MustVariant("NormalizedString"),
			},
			superTypeID: ua.NewNumericNodeID(0, 12),
		},
		&node{
			id: ua.NewNumericNodeID(0, 12878),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("DecimalString"),
				ua.AttributeIDDisplayName: ua.MustVariant("DecimalString"),
			},
			superTypeID: ua.NewNumericNodeID(0, 12),
		},
		&node{
			id: ua.NewNumericNodeID(0, 12879),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("DurationString"),
				ua.AttributeIDDisplayName: ua.MustVariant("DurationString"),
			},
			superTypeID: ua.NewNumericNodeID(0, 12),
		},
		&node{
			id: ua.NewNumericNodeID(0, 12880),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("TimeString"),
				ua.AttributeIDDisplayName: ua.MustVariant("TimeString"),
			},
			superTypeID: ua.NewNumericNodeID(0, 12),
		},
		&node{
			id: ua.NewNumericNodeID(0, 12881),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("DateString"),
				ua.AttributeIDDisplayName: ua.MustVariant("DateString"),
			},
			superTypeID: ua.NewNumericNodeID(0, 12),
		},
		&node{
			id: ua.NewNumericNodeID(0, 290),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Duration"),
				ua.AttributeIDDisplayName: ua.MustVariant("Duration"),
			},
			superTypeID: ua.NewNumericNodeID(0, 11),
		},
		&node{
			id: ua.NewNumericNodeID(0, 294),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("UtcTime"),
				ua.AttributeIDDisplayName: ua.MustVariant("UtcTime"),
			},
			superTypeID: ua.NewNumericNodeID(0, 13),
		},
		&node{
			id: ua.NewNumericNodeID(0, 295),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("LocaleId"),
				ua.AttributeIDDisplayName: ua.MustVariant("LocaleId"),
			},
			superTypeID: ua.NewNumericNodeID(0, 12),
		},
		&node{
			id: ua.NewNumericNodeID(0, 8912),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("TimeZoneDataType"),
				ua.AttributeIDDisplayName: ua.MustVariant("TimeZoneDataType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 22),
		},
		&node{
			id: ua.NewNumericNodeID(0, 17588),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Index"),
				ua.AttributeIDDisplayName: ua.MustVariant("Index"),
			},
			superTypeID: ua.NewNumericNodeID(0, 7),
		},
		&node{
			id: ua.NewNumericNodeID(0, 288),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("IntegerId"),
				ua.AttributeIDDisplayName: ua.MustVariant("IntegerId"),
			},
			superTypeID: ua.NewNumericNodeID(0, 7),
		},
		&node{
			id: ua.NewNumericNodeID(0, 307),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("ApplicationType"),
				ua.AttributeIDDisplayName: ua.MustVariant("ApplicationType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 29),
		},
		&node{
			id: ua.NewNumericNodeID(0, 308),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("ApplicationDescription"),
				ua.AttributeIDDisplayName: ua.MustVariant("ApplicationDescription"),
			},
			superTypeID: ua.NewNumericNodeID(0, 22),
		},
		&node{
			id: ua.NewNumericNodeID(0, 20998),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("VersionTime"),
				ua.AttributeIDDisplayName: ua.MustVariant("VersionTime"),
			},
			superTypeID: ua.NewNumericNodeID(0, 7),
		},
		&node{
			id: ua.NewNumericNodeID(0, 12189),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("ServerOnNetwork"),
				ua.AttributeIDDisplayName: ua.MustVariant("ServerOnNetwork"),
			},
			superTypeID: ua.NewNumericNodeID(0, 22),
		},
		&node{
			id: ua.NewNumericNodeID(0, 311),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("ApplicationInstanceCertificate"),
				ua.AttributeIDDisplayName: ua.MustVariant("ApplicationInstanceCertificate"),
			},
			superTypeID: ua.NewNumericNodeID(0, 15),
		},
		&node{
			id: ua.NewNumericNodeID(0, 302),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("MessageSecurityMode"),
				ua.AttributeIDDisplayName: ua.MustVariant("MessageSecurityMode"),
			},
			superTypeID: ua.NewNumericNodeID(0, 29),
		},
		&node{
			id: ua.NewNumericNodeID(0, 303),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("UserTokenType"),
				ua.AttributeIDDisplayName: ua.MustVariant("UserTokenType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 29),
		},
		&node{
			id: ua.NewNumericNodeID(0, 304),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("UserTokenPolicy"),
				ua.AttributeIDDisplayName: ua.MustVariant("UserTokenPolicy"),
			},
			superTypeID: ua.NewNumericNodeID(0, 22),
		},
		&node{
			id: ua.NewNumericNodeID(0, 312),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("EndpointDescription"),
				ua.AttributeIDDisplayName: ua.MustVariant("EndpointDescription"),
			},
			superTypeID: ua.NewNumericNodeID(0, 22),
		},
		&node{
			id: ua.NewNumericNodeID(0, 432),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("RegisteredServer"),
				ua.AttributeIDDisplayName: ua.MustVariant("RegisteredServer"),
			},
			superTypeID: ua.NewNumericNodeID(0, 22),
		},
		&node{
			id: ua.NewNumericNodeID(0, 12890),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("DiscoveryConfiguration"),
				ua.AttributeIDDisplayName: ua.MustVariant("DiscoveryConfiguration"),
			},
			superTypeID: ua.NewNumericNodeID(0, 22),
		},
		&node{
			id: ua.NewNumericNodeID(0, 12891),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("MdnsDiscoveryConfiguration"),
				ua.AttributeIDDisplayName: ua.MustVariant("MdnsDiscoveryConfiguration"),
			},
			superTypeID: ua.NewNumericNodeID(0, 12890),
		},
		&node{
			id: ua.NewNumericNodeID(0, 315),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("SecurityTokenRequestType"),
				ua.AttributeIDDisplayName: ua.MustVariant("SecurityTokenRequestType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 29),
		},
		&node{
			id: ua.NewNumericNodeID(0, 344),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("SignedSoftwareCertificate"),
				ua.AttributeIDDisplayName: ua.MustVariant("SignedSoftwareCertificate"),
			},
			superTypeID: ua.NewNumericNodeID(0, 22),
		},
		&node{
			id: ua.NewNumericNodeID(0, 388),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("SessionAuthenticationToken"),
				ua.AttributeIDDisplayName: ua.MustVariant("SessionAuthenticationToken"),
			},
			superTypeID: ua.NewNumericNodeID(0, 17),
		},
		&node{
			id: ua.NewNumericNodeID(0, 319),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("AnonymousIdentityToken"),
				ua.AttributeIDDisplayName: ua.MustVariant("AnonymousIdentityToken"),
			},
			superTypeID: ua.NewNumericNodeID(0, 316),
		},
		&node{
			id: ua.NewNumericNodeID(0, 322),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("UserNameIdentityToken"),
				ua.AttributeIDDisplayName: ua.MustVariant("UserNameIdentityToken"),
			},
			superTypeID: ua.NewNumericNodeID(0, 316),
		},
		&node{
			id: ua.NewNumericNodeID(0, 325),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("X509IdentityToken"),
				ua.AttributeIDDisplayName: ua.MustVariant("X509IdentityToken"),
			},
			superTypeID: ua.NewNumericNodeID(0, 316),
		},
		&node{
			id: ua.NewNumericNodeID(0, 938),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("IssuedIdentityToken"),
				ua.AttributeIDDisplayName: ua.MustVariant("IssuedIdentityToken"),
			},
			superTypeID: ua.NewNumericNodeID(0, 316),
		},
		&node{
			id: ua.NewNumericNodeID(0, 348),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("NodeAttributesMask"),
				ua.AttributeIDDisplayName: ua.MustVariant("NodeAttributesMask"),
			},
			superTypeID: ua.NewNumericNodeID(0, 29),
		},
		&node{
			id: ua.NewNumericNodeID(0, 376),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("AddNodesItem"),
				ua.AttributeIDDisplayName: ua.MustVariant("AddNodesItem"),
			},
			superTypeID: ua.NewNumericNodeID(0, 22),
		},
		&node{
			id: ua.NewNumericNodeID(0, 379),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("AddReferencesItem"),
				ua.AttributeIDDisplayName: ua.MustVariant("AddReferencesItem"),
			},
			superTypeID: ua.NewNumericNodeID(0, 22),
		},
		&node{
			id: ua.NewNumericNodeID(0, 382),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("DeleteNodesItem"),
				ua.AttributeIDDisplayName: ua.MustVariant("DeleteNodesItem"),
			},
			superTypeID: ua.NewNumericNodeID(0, 22),
		},
		&node{
			id: ua.NewNumericNodeID(0, 385),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("DeleteReferencesItem"),
				ua.AttributeIDDisplayName: ua.MustVariant("DeleteReferencesItem"),
			},
			superTypeID: ua.NewNumericNodeID(0, 22),
		},
		&node{
			id: ua.NewNumericNodeID(0, 347),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("AttributeWriteMask"),
				ua.AttributeIDDisplayName: ua.MustVariant("AttributeWriteMask"),
			},
			superTypeID: ua.NewNumericNodeID(0, 7),
		},
		&node{
			id: ua.NewNumericNodeID(0, 521),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("ContinuationPoint"),
				ua.AttributeIDDisplayName: ua.MustVariant("ContinuationPoint"),
			},
			superTypeID: ua.NewNumericNodeID(0, 15),
		},
		&node{
			id: ua.NewNumericNodeID(0, 537),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("RelativePathElement"),
				ua.AttributeIDDisplayName: ua.MustVariant("RelativePathElement"),
			},
			superTypeID: ua.NewNumericNodeID(0, 22),
		},
		&node{
			id: ua.NewNumericNodeID(0, 540),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("RelativePath"),
				ua.AttributeIDDisplayName: ua.MustVariant("RelativePath"),
			},
			superTypeID: ua.NewNumericNodeID(0, 22),
		},
		&node{
			id: ua.NewNumericNodeID(0, 289),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Counter"),
				ua.AttributeIDDisplayName: ua.MustVariant("Counter"),
			},
			superTypeID: ua.NewNumericNodeID(0, 7),
		},
		&node{
			id: ua.NewNumericNodeID(0, 291),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("NumericRange"),
				ua.AttributeIDDisplayName: ua.MustVariant("NumericRange"),
			},
			superTypeID: ua.NewNumericNodeID(0, 12),
		},
		&node{
			id: ua.NewNumericNodeID(0, 331),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("EndpointConfiguration"),
				ua.AttributeIDDisplayName: ua.MustVariant("EndpointConfiguration"),
			},
			superTypeID: ua.NewNumericNodeID(0, 22),
		},
		&node{
			id: ua.NewNumericNodeID(0, 576),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("FilterOperator"),
				ua.AttributeIDDisplayName: ua.MustVariant("FilterOperator"),
			},
			superTypeID: ua.NewNumericNodeID(0, 29),
		},
		&node{
			id: ua.NewNumericNodeID(0, 583),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("ContentFilterElement"),
				ua.AttributeIDDisplayName: ua.MustVariant("ContentFilterElement"),
			},
			superTypeID: ua.NewNumericNodeID(0, 22),
		},
		&node{
			id: ua.NewNumericNodeID(0, 586),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("ContentFilter"),
				ua.AttributeIDDisplayName: ua.MustVariant("ContentFilter"),
			},
			superTypeID: ua.NewNumericNodeID(0, 22),
		},
		&node{
			id: ua.NewNumericNodeID(0, 589),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("FilterOperand"),
				ua.AttributeIDDisplayName: ua.MustVariant("FilterOperand"),
			},
			superTypeID: ua.NewNumericNodeID(0, 22),
		},
		&node{
			id: ua.NewNumericNodeID(0, 592),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("ElementOperand"),
				ua.AttributeIDDisplayName: ua.MustVariant("ElementOperand"),
			},
			superTypeID: ua.NewNumericNodeID(0, 589),
		},
		&node{
			id: ua.NewNumericNodeID(0, 595),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("LiteralOperand"),
				ua.AttributeIDDisplayName: ua.MustVariant("LiteralOperand"),
			},
			superTypeID: ua.NewNumericNodeID(0, 589),
		},
		&node{
			id: ua.NewNumericNodeID(0, 598),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("AttributeOperand"),
				ua.AttributeIDDisplayName: ua.MustVariant("AttributeOperand"),
			},
			superTypeID: ua.NewNumericNodeID(0, 589),
		},
		&node{
			id: ua.NewNumericNodeID(0, 601),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("SimpleAttributeOperand"),
				ua.AttributeIDDisplayName: ua.MustVariant("SimpleAttributeOperand"),
			},
			superTypeID: ua.NewNumericNodeID(0, 589),
		},
		&node{
			id: ua.NewNumericNodeID(0, 659),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("HistoryEvent"),
				ua.AttributeIDDisplayName: ua.MustVariant("HistoryEvent"),
			},
			superTypeID: ua.NewNumericNodeID(0, 22),
		},
		&node{
			id: ua.NewNumericNodeID(0, 11234),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("HistoryUpdateType"),
				ua.AttributeIDDisplayName: ua.MustVariant("HistoryUpdateType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 29),
		},
		&node{
			id: ua.NewNumericNodeID(0, 11293),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("PerformUpdateType"),
				ua.AttributeIDDisplayName: ua.MustVariant("PerformUpdateType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 29),
		},
		&node{
			id: ua.NewNumericNodeID(0, 719),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("MonitoringFilter"),
				ua.AttributeIDDisplayName: ua.MustVariant("MonitoringFilter"),
			},
			superTypeID: ua.NewNumericNodeID(0, 22),
		},
		&node{
			id: ua.NewNumericNodeID(0, 725),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("EventFilter"),
				ua.AttributeIDDisplayName: ua.MustVariant("EventFilter"),
			},
			superTypeID: ua.NewNumericNodeID(0, 719),
		},
		&node{
			id: ua.NewNumericNodeID(0, 948),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("AggregateConfiguration"),
				ua.AttributeIDDisplayName: ua.MustVariant("AggregateConfiguration"),
			},
			superTypeID: ua.NewNumericNodeID(0, 22),
		},
		&node{
			id: ua.NewNumericNodeID(0, 920),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("HistoryEventFieldList"),
				ua.AttributeIDDisplayName: ua.MustVariant("HistoryEventFieldList"),
			},
			superTypeID: ua.NewNumericNodeID(0, 22),
		},
		&node{
			id: ua.NewNumericNodeID(0, 338),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("BuildInfo"),
				ua.AttributeIDDisplayName: ua.MustVariant("BuildInfo"),
			},
			superTypeID: ua.NewNumericNodeID(0, 22),
		},
		&node{
			id: ua.NewNumericNodeID(0, 851),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("RedundancySupport"),
				ua.AttributeIDDisplayName: ua.MustVariant("RedundancySupport"),
			},
			superTypeID: ua.NewNumericNodeID(0, 29),
		},
		&node{
			id: ua.NewNumericNodeID(0, 852),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("ServerState"),
				ua.AttributeIDDisplayName: ua.MustVariant("ServerState"),
			},
			superTypeID: ua.NewNumericNodeID(0, 29),
		},
		&node{
			id: ua.NewNumericNodeID(0, 853),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("RedundantServerDataType"),
				ua.AttributeIDDisplayName: ua.MustVariant("RedundantServerDataType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 22),
		},
		&node{
			id: ua.NewNumericNodeID(0, 11943),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("EndpointUrlListDataType"),
				ua.AttributeIDDisplayName: ua.MustVariant("EndpointUrlListDataType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 22),
		},
		&node{
			id: ua.NewNumericNodeID(0, 11944),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("NetworkGroupDataType"),
				ua.AttributeIDDisplayName: ua.MustVariant("NetworkGroupDataType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 22),
		},
		&node{
			id: ua.NewNumericNodeID(0, 856),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("SamplingIntervalDiagnosticsDataType"),
				ua.AttributeIDDisplayName: ua.MustVariant("SamplingIntervalDiagnosticsDataType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 22),
		},
		&node{
			id: ua.NewNumericNodeID(0, 859),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("ServerDiagnosticsSummaryDataType"),
				ua.AttributeIDDisplayName: ua.MustVariant("ServerDiagnosticsSummaryDataType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 22),
		},
		&node{
			id: ua.NewNumericNodeID(0, 862),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("ServerStatusDataType"),
				ua.AttributeIDDisplayName: ua.MustVariant("ServerStatusDataType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 22),
		},
		&node{
			id: ua.NewNumericNodeID(0, 865),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("SessionDiagnosticsDataType"),
				ua.AttributeIDDisplayName: ua.MustVariant("SessionDiagnosticsDataType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 22),
		},
		&node{
			id: ua.NewNumericNodeID(0, 868),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("SessionSecurityDiagnosticsDataType"),
				ua.AttributeIDDisplayName: ua.MustVariant("SessionSecurityDiagnosticsDataType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 22),
		},
		&node{
			id: ua.NewNumericNodeID(0, 871),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("ServiceCounterDataType"),
				ua.AttributeIDDisplayName: ua.MustVariant("ServiceCounterDataType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 22),
		},
		&node{
			id: ua.NewNumericNodeID(0, 299),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("StatusResult"),
				ua.AttributeIDDisplayName: ua.MustVariant("StatusResult"),
			},
			superTypeID: ua.NewNumericNodeID(0, 22),
		},
		&node{
			id: ua.NewNumericNodeID(0, 874),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("SubscriptionDiagnosticsDataType"),
				ua.AttributeIDDisplayName: ua.MustVariant("SubscriptionDiagnosticsDataType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 22),
		},
		&node{
			id: ua.NewNumericNodeID(0, 877),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("ModelChangeStructureDataType"),
				ua.AttributeIDDisplayName: ua.MustVariant("ModelChangeStructureDataType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 22),
		},
		&node{
			id: ua.NewNumericNodeID(0, 897),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("SemanticChangeStructureDataType"),
				ua.AttributeIDDisplayName: ua.MustVariant("SemanticChangeStructureDataType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 22),
		},
		&node{
			id: ua.NewNumericNodeID(0, 884),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Range"),
				ua.AttributeIDDisplayName: ua.MustVariant("Range"),
			},
			superTypeID: ua.NewNumericNodeID(0, 22),
		},
		&node{
			id: ua.NewNumericNodeID(0, 887),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("EUInformation"),
				ua.AttributeIDDisplayName: ua.MustVariant("EUInformation"),
			},
			superTypeID: ua.NewNumericNodeID(0, 22),
		},
		&node{
			id: ua.NewNumericNodeID(0, 12077),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("AxisScaleEnumeration"),
				ua.AttributeIDDisplayName: ua.MustVariant("AxisScaleEnumeration"),
			},
			superTypeID: ua.NewNumericNodeID(0, 29),
		},
		&node{
			id: ua.NewNumericNodeID(0, 12171),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("ComplexNumberType"),
				ua.AttributeIDDisplayName: ua.MustVariant("ComplexNumberType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 22),
		},
		&node{
			id: ua.NewNumericNodeID(0, 12172),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("DoubleComplexNumberType"),
				ua.AttributeIDDisplayName: ua.MustVariant("DoubleComplexNumberType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 22),
		},
		&node{
			id: ua.NewNumericNodeID(0, 12079),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("AxisInformation"),
				ua.AttributeIDDisplayName: ua.MustVariant("AxisInformation"),
			},
			superTypeID: ua.NewNumericNodeID(0, 22),
		},
		&node{
			id: ua.NewNumericNodeID(0, 12080),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("XVType"),
				ua.AttributeIDDisplayName: ua.MustVariant("XVType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 22),
		},
		&node{
			id: ua.NewNumericNodeID(0, 894),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("ProgramDiagnosticDataType"),
				ua.AttributeIDDisplayName: ua.MustVariant("ProgramDiagnosticDataType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 22),
		},
		&node{
			id: ua.NewNumericNodeID(0, 24033),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("ProgramDiagnostic2DataType"),
				ua.AttributeIDDisplayName: ua.MustVariant("ProgramDiagnostic2DataType"),
			},
			superTypeID: ua.NewNumericNodeID(0, 22),
		},
		&node{
			id: ua.NewNumericNodeID(0, 891),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Annotation"),
				ua.AttributeIDDisplayName: ua.MustVariant("Annotation"),
			},
			superTypeID: ua.NewNumericNodeID(0, 22),
		},
		&node{
			id: ua.NewNumericNodeID(0, 890),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("DataType_64"),
				ua.AttributeIDBrowseName:  ua.MustVariant("ExceptionDeviationFormat"),
				ua.AttributeIDDisplayName: ua.MustVariant("ExceptionDeviationFormat"),
			},
			superTypeID: ua.NewNumericNodeID(0, 29),
		},
		&node{
			id: ua.NewNumericNodeID(0, 12766),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default Binary"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default Binary"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 14846),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default Binary"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default Binary"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 17537),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default Binary"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default Binary"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 17549),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default Binary"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default Binary"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15671),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default Binary"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default Binary"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 18815),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default Binary"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default Binary"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 18816),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default Binary"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default Binary"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 18817),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default Binary"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default Binary"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 18818),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default Binary"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default Binary"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 18819),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default Binary"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default Binary"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 18820),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default Binary"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default Binary"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 18821),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default Binary"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default Binary"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 18822),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default Binary"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default Binary"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 18823),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default Binary"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default Binary"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15736),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default Binary"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default Binary"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 23507),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default Binary"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default Binary"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 12680),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default Binary"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default Binary"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 32382),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default Binary"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default Binary"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15676),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default Binary"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default Binary"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 125),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default Binary"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default Binary"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 126),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default Binary"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default Binary"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 127),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default Binary"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default Binary"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15421),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default Binary"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default Binary"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15422),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default Binary"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default Binary"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 24108),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default Binary"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default Binary"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 24109),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default Binary"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default Binary"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 24110),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default Binary"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default Binary"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 124),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default Binary"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default Binary"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 14839),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default Binary"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default Binary"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 14847),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default Binary"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default Binary"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15677),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default Binary"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default Binary"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15678),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default Binary"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default Binary"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 14323),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default Binary"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default Binary"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15679),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default Binary"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default Binary"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15681),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default Binary"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default Binary"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 25529),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default Binary"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default Binary"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15682),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default Binary"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default Binary"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15683),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default Binary"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default Binary"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15688),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default Binary"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default Binary"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15689),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default Binary"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default Binary"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 21150),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default Binary"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default Binary"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15691),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default Binary"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default Binary"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15693),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default Binary"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default Binary"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15694),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default Binary"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default Binary"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15695),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default Binary"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default Binary"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 21151),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default Binary"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default Binary"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 21152),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default Binary"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default Binary"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 21153),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default Binary"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default Binary"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15701),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default Binary"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default Binary"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15702),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default Binary"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default Binary"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15703),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default Binary"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default Binary"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15705),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default Binary"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default Binary"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15706),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default Binary"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default Binary"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15707),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default Binary"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default Binary"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15712),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default Binary"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default Binary"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 14848),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default Binary"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default Binary"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15713),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default Binary"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default Binary"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 21154),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default Binary"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default Binary"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 23851),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default Binary"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default Binary"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 23852),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default Binary"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default Binary"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 23853),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default Binary"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default Binary"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 25530),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default Binary"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default Binary"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 23854),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default Binary"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default Binary"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15715),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default Binary"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default Binary"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15717),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default Binary"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default Binary"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15718),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default Binary"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default Binary"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15719),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default Binary"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default Binary"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15724),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default Binary"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default Binary"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15725),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default Binary"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default Binary"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 23855),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default Binary"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default Binary"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 23856),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default Binary"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default Binary"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 23857),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default Binary"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default Binary"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 23860),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default Binary"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default Binary"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 23861),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default Binary"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default Binary"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 17468),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default Binary"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default Binary"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 23864),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default Binary"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default Binary"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 21155),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default Binary"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default Binary"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 23865),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default Binary"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default Binary"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 23866),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default Binary"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default Binary"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15479),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default Binary"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default Binary"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15727),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default Binary"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default Binary"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15729),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default Binary"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default Binary"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15733),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default Binary"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default Binary"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 25531),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default Binary"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default Binary"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 25532),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default Binary"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default Binary"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 23499),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default Binary"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default Binary"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 24292),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default Binary"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default Binary"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 25239),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default Binary"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default Binary"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 32661),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default Binary"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default Binary"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 32662),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default Binary"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default Binary"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 128),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default Binary"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default Binary"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 121),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default Binary"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default Binary"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 14844),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default Binary"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default Binary"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 122),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default Binary"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default Binary"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 123),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default Binary"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default Binary"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 298),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default Binary"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default Binary"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 8251),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default Binary"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default Binary"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 14845),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default Binary"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default Binary"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 12765),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default Binary"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default Binary"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 8917),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default Binary"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default Binary"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 310),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default Binary"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default Binary"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 12207),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default Binary"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default Binary"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 306),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default Binary"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default Binary"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 314),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default Binary"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default Binary"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 434),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default Binary"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default Binary"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 12900),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default Binary"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default Binary"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 12901),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default Binary"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default Binary"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 346),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default Binary"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default Binary"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 318),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default Binary"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default Binary"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 321),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default Binary"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default Binary"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 324),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default Binary"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default Binary"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 327),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default Binary"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default Binary"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 940),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default Binary"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default Binary"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 378),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default Binary"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default Binary"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 381),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default Binary"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default Binary"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 384),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default Binary"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default Binary"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 387),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default Binary"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default Binary"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 539),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default Binary"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default Binary"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 542),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default Binary"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default Binary"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 333),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default Binary"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default Binary"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 585),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default Binary"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default Binary"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 588),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default Binary"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default Binary"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 591),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default Binary"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default Binary"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 594),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default Binary"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default Binary"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 597),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default Binary"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default Binary"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 600),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default Binary"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default Binary"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 603),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default Binary"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default Binary"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 661),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default Binary"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default Binary"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 721),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default Binary"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default Binary"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 727),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default Binary"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default Binary"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 950),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default Binary"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default Binary"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 922),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default Binary"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default Binary"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 340),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default Binary"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default Binary"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 855),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default Binary"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default Binary"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 11957),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default Binary"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default Binary"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 11958),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default Binary"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default Binary"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 858),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default Binary"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default Binary"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 861),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default Binary"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default Binary"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 864),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default Binary"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default Binary"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 867),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default Binary"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default Binary"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 870),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default Binary"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default Binary"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 873),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default Binary"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default Binary"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 301),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default Binary"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default Binary"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 876),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default Binary"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default Binary"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 879),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default Binary"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default Binary"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 899),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default Binary"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default Binary"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 886),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default Binary"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default Binary"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 889),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default Binary"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default Binary"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 12181),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default Binary"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default Binary"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 12182),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default Binary"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default Binary"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 12089),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default Binary"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default Binary"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 12090),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default Binary"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default Binary"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 896),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default Binary"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default Binary"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 24034),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default Binary"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default Binary"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 893),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default Binary"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default Binary"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 7617),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Variable_2"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Opc.Ua"),
				ua.AttributeIDDisplayName: ua.MustVariant("Opc.Ua"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 12758),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default XML"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default XML"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 14802),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default XML"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default XML"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 17541),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default XML"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default XML"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 17553),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default XML"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default XML"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15949),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default XML"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default XML"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 18851),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default XML"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default XML"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 18852),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default XML"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default XML"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 18853),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default XML"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default XML"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 18854),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default XML"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default XML"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 18855),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default XML"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default XML"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 18856),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default XML"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default XML"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 18857),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default XML"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default XML"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 18858),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default XML"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default XML"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 18859),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default XML"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default XML"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15728),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default XML"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default XML"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 23520),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default XML"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default XML"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 12676),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default XML"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default XML"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 32386),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default XML"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default XML"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15950),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default XML"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default XML"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 14796),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default XML"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default XML"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15589),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default XML"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default XML"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15590),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default XML"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default XML"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15529),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default XML"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default XML"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15531),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default XML"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default XML"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 24120),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default XML"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default XML"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 24121),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default XML"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default XML"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 24122),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default XML"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default XML"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 14794),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default XML"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default XML"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 14795),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default XML"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default XML"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 14803),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default XML"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default XML"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15951),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default XML"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default XML"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15952),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default XML"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default XML"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 14319),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default XML"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default XML"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15953),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default XML"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default XML"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15954),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default XML"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default XML"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 25545),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default XML"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default XML"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15955),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default XML"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default XML"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15956),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default XML"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default XML"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15987),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default XML"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default XML"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15988),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default XML"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default XML"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 21174),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default XML"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default XML"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15990),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default XML"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default XML"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15991),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default XML"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default XML"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15992),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default XML"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default XML"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15993),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default XML"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default XML"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 21175),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default XML"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default XML"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 21176),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default XML"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default XML"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 21177),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default XML"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default XML"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15995),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default XML"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default XML"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15996),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default XML"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default XML"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 16007),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default XML"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default XML"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 16008),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default XML"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default XML"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 16009),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default XML"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default XML"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 16010),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default XML"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default XML"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 16011),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default XML"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default XML"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 14804),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default XML"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default XML"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 16012),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default XML"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default XML"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 21178),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default XML"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default XML"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 23919),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default XML"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default XML"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 23920),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default XML"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default XML"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 23921),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default XML"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default XML"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 25546),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default XML"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default XML"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 23922),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default XML"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default XML"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 16014),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default XML"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default XML"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 16015),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default XML"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default XML"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 16016),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default XML"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default XML"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 16017),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default XML"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default XML"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 16018),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default XML"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default XML"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 16019),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default XML"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default XML"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 23923),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default XML"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default XML"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 23924),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default XML"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default XML"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 23925),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default XML"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default XML"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 23928),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default XML"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default XML"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 23929),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default XML"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default XML"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 17472),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default XML"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default XML"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 23932),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default XML"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default XML"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 21179),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default XML"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default XML"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 23933),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default XML"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default XML"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 23934),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default XML"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default XML"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15579),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default XML"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default XML"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 16021),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default XML"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default XML"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 16022),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default XML"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default XML"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 16023),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default XML"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default XML"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 25547),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default XML"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default XML"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 25548),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default XML"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default XML"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 23505),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default XML"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default XML"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 24296),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default XML"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default XML"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 25243),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default XML"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default XML"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 32669),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default XML"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default XML"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 32670),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default XML"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default XML"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 16126),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default XML"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default XML"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 14797),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default XML"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default XML"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 14800),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default XML"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default XML"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 14798),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default XML"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default XML"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 14799),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default XML"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default XML"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 297),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default XML"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default XML"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 7616),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default XML"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default XML"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 14801),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default XML"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default XML"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 12757),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default XML"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default XML"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 8913),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default XML"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default XML"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 309),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default XML"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default XML"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 12195),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default XML"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default XML"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 305),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default XML"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default XML"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 313),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default XML"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default XML"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 433),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default XML"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default XML"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 12892),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default XML"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default XML"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 12893),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default XML"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default XML"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 345),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default XML"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default XML"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 317),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default XML"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default XML"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 320),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default XML"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default XML"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 323),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default XML"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default XML"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 326),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default XML"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default XML"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 939),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default XML"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default XML"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 377),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default XML"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default XML"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 380),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default XML"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default XML"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 383),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default XML"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default XML"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 386),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default XML"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default XML"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 538),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default XML"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default XML"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 541),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default XML"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default XML"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 332),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default XML"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default XML"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 584),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default XML"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default XML"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 587),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default XML"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default XML"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 590),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default XML"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default XML"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 593),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default XML"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default XML"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 596),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default XML"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default XML"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 599),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default XML"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default XML"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 602),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default XML"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default XML"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 660),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default XML"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default XML"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 720),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default XML"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default XML"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 726),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default XML"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default XML"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 949),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default XML"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default XML"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 921),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default XML"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default XML"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 339),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default XML"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default XML"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 854),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default XML"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default XML"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 11949),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default XML"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default XML"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 11950),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default XML"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default XML"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 857),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default XML"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default XML"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 860),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default XML"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default XML"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 863),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default XML"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default XML"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 866),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default XML"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default XML"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 869),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default XML"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default XML"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 872),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default XML"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default XML"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 300),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default XML"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default XML"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 875),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default XML"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default XML"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 878),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default XML"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default XML"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 898),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default XML"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default XML"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 885),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default XML"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default XML"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 888),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default XML"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default XML"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 12173),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default XML"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default XML"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 12174),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default XML"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default XML"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 12081),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default XML"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default XML"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 12082),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default XML"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default XML"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 895),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default XML"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default XML"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 24038),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default XML"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default XML"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 892),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default XML"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default XML"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 8252),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Variable_2"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Opc.Ua"),
				ua.AttributeIDDisplayName: ua.MustVariant("Opc.Ua"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15085),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default JSON"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default JSON"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15041),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default JSON"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default JSON"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 17547),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default JSON"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default JSON"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 17557),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default JSON"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default JSON"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 16150),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default JSON"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default JSON"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 19064),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default JSON"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default JSON"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 19065),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default JSON"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default JSON"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 19066),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default JSON"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default JSON"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 19067),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default JSON"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default JSON"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 19068),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default JSON"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default JSON"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 19069),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default JSON"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default JSON"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 19070),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default JSON"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default JSON"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 19071),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default JSON"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default JSON"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 19072),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default JSON"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default JSON"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15042),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default JSON"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default JSON"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 23528),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default JSON"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default JSON"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15044),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default JSON"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default JSON"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 32390),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default JSON"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default JSON"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 16151),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default JSON"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default JSON"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15057),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default JSON"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default JSON"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15058),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default JSON"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default JSON"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15059),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default JSON"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default JSON"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15700),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default JSON"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default JSON"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15714),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default JSON"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default JSON"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 24132),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default JSON"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default JSON"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 24133),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default JSON"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default JSON"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 24134),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default JSON"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default JSON"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15050),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default JSON"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default JSON"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15051),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default JSON"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default JSON"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15049),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default JSON"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default JSON"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 16152),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default JSON"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default JSON"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 16153),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default JSON"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default JSON"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15060),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default JSON"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default JSON"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 16154),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default JSON"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default JSON"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 16155),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default JSON"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default JSON"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 25561),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default JSON"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default JSON"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 16156),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default JSON"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default JSON"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 16157),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default JSON"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default JSON"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 16158),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default JSON"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default JSON"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 16159),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default JSON"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default JSON"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 21198),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default JSON"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default JSON"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 16161),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default JSON"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default JSON"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 16280),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default JSON"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default JSON"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 16281),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default JSON"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default JSON"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 16282),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default JSON"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default JSON"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 21199),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default JSON"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default JSON"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 21200),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default JSON"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default JSON"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 21201),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default JSON"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default JSON"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 16284),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default JSON"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default JSON"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 16285),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default JSON"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default JSON"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 16286),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default JSON"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default JSON"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 16287),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default JSON"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default JSON"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 16288),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default JSON"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default JSON"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 16308),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default JSON"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default JSON"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 16310),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default JSON"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default JSON"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15061),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default JSON"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default JSON"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 16311),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default JSON"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default JSON"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 21202),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default JSON"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default JSON"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 23987),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default JSON"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default JSON"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 23988),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default JSON"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default JSON"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 23989),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default JSON"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default JSON"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 25562),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default JSON"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default JSON"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 23990),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default JSON"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default JSON"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 16323),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default JSON"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default JSON"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 16391),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default JSON"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default JSON"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 16392),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default JSON"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default JSON"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 16393),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default JSON"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default JSON"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 16394),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default JSON"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default JSON"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 16404),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default JSON"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default JSON"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 23991),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default JSON"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default JSON"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 23992),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default JSON"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default JSON"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 23993),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default JSON"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default JSON"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 23996),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default JSON"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default JSON"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 23997),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default JSON"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default JSON"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 17476),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default JSON"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default JSON"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 24000),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default JSON"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default JSON"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 21203),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default JSON"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default JSON"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 24001),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default JSON"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default JSON"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 24002),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default JSON"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default JSON"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15726),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default JSON"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default JSON"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 16524),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default JSON"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default JSON"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 16525),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default JSON"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default JSON"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 16526),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default JSON"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default JSON"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 25563),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default JSON"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default JSON"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 25564),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default JSON"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default JSON"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 23511),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default JSON"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default JSON"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 24300),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default JSON"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default JSON"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 25247),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default JSON"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default JSON"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 32677),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default JSON"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default JSON"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 32678),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default JSON"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default JSON"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15062),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default JSON"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default JSON"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15063),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default JSON"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default JSON"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15065),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default JSON"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default JSON"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15066),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default JSON"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default JSON"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15067),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default JSON"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default JSON"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15081),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default JSON"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default JSON"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15082),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default JSON"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default JSON"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15083),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default JSON"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default JSON"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15084),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default JSON"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default JSON"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15086),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default JSON"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default JSON"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15087),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default JSON"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default JSON"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15095),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default JSON"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default JSON"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15098),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default JSON"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default JSON"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15099),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default JSON"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default JSON"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15102),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default JSON"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default JSON"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15105),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default JSON"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default JSON"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15106),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default JSON"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default JSON"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15136),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default JSON"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default JSON"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15140),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default JSON"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default JSON"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15141),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default JSON"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default JSON"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15142),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default JSON"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default JSON"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15143),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default JSON"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default JSON"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15144),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default JSON"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default JSON"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15165),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default JSON"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default JSON"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15169),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default JSON"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default JSON"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15172),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default JSON"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default JSON"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15175),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default JSON"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default JSON"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15188),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default JSON"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default JSON"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15189),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default JSON"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default JSON"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15199),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default JSON"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default JSON"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15204),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default JSON"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default JSON"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15205),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default JSON"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default JSON"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15206),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default JSON"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default JSON"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15207),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default JSON"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default JSON"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15208),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default JSON"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default JSON"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15209),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default JSON"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default JSON"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15210),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default JSON"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default JSON"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15273),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default JSON"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default JSON"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15293),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default JSON"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default JSON"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15295),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default JSON"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default JSON"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15304),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default JSON"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default JSON"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15349),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default JSON"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default JSON"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15361),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default JSON"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default JSON"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15362),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default JSON"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default JSON"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15363),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default JSON"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default JSON"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15364),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default JSON"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default JSON"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15365),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default JSON"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default JSON"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15366),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default JSON"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default JSON"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15367),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default JSON"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default JSON"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15368),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default JSON"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default JSON"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15369),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default JSON"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default JSON"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15370),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default JSON"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default JSON"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15371),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default JSON"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default JSON"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15372),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default JSON"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default JSON"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15373),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default JSON"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default JSON"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15374),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default JSON"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default JSON"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15375),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default JSON"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default JSON"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15376),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default JSON"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default JSON"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15377),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default JSON"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default JSON"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15378),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default JSON"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default JSON"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15379),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default JSON"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default JSON"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15380),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default JSON"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default JSON"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15381),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default JSON"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default JSON"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 24042),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default JSON"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default JSON"),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15382),
			attr: map[ua.AttributeID]*ua.Variant{
				ua.AttributeIDNodeClass:   ua.MustVariant("Object_1"),
				ua.AttributeIDBrowseName:  ua.MustVariant("Default JSON"),
				ua.AttributeIDDisplayName: ua.MustVariant("Default JSON"),
			},
		},
	}
}
