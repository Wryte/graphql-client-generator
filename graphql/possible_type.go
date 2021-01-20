package graphql

// PossibleType a possible type for a GrahpQL union
type PossibleType struct {
	Kind   FieldTypeKind `json:"kind"`
	Name   string        `json:"name"`
	OfType *OfType       `json:"ofType"`
}
