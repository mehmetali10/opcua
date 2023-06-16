package refs

import (
	"github.com/gopcua/opcua/id"
	"github.com/gopcua/opcua/ua"
)

// HasSubtype returns a HasSubtype reference.
func HasSubtype(typeID *ua.ExpandedNodeID) *ua.ReferenceDescription {
	return &ua.ReferenceDescription{
		ReferenceTypeID: ua.NewNumericNodeID(0, id.HasSubtype),
		TypeDefinition:  typeID,
		IsForward:       true,
	}
}
