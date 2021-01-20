package graphql

// FieldType a GraphQL field's type
type FieldType struct {
	Kind   FieldTypeKind `json:"kind"`
	Name   string        `json:"name"`
	OfType *OfType       `json:"ofType"`
}

// FieldTypeKind a GraphQl field type kind
type FieldTypeKind string

const (
	// FieldTypeKindEnum a field type kind Enum
	FieldTypeKindEnum FieldTypeKind = "ENUM"
	// FieldTypeKindInterface a field type kind Interface
	FieldTypeKindInterface FieldTypeKind = "INTERFACE"
	// FieldTypeKindList a field type kind List
	FieldTypeKindList FieldTypeKind = "LIST"
	// FieldTypeKindNonNull a field type kind NonNull
	FieldTypeKindNonNull FieldTypeKind = "NON_NULL"
	// FieldTypeKindObject a field type kind Object
	FieldTypeKindObject FieldTypeKind = "OBJECT"
	// FieldTypeKindScalar a field type kind Scalar
	FieldTypeKindScalar FieldTypeKind = "SCALAR"
	// FieldTypeKindUnion a field type kind Union
	FieldTypeKindUnion FieldTypeKind = "UNION"
	// FieldTypeKindInputObject a field type kind Input Object
	FieldTypeKindInputObject FieldTypeKind = "INPUT_OBJECT"
)
