package graphql

// Schema the representation of the GraphQL schema
type Schema struct {
	QueryType struct {
		Name string `json:"name"`
	} `json:"queryType"`
	MutationType struct {
		Name string `json:"name"`
	} `json:"mutationType"`
	Types []Type `json:"types"`
}

// Type retrieve a type by name
func (s *Schema) Type(name string) *Type {
	for _, t := range s.Types {
		if t.Name == name {
			return &t
		}
	}

	return nil
}

// Type a type in the GraphQL schema
type Type struct {
	Kind          TypeKind       `json:"kind"`
	Name          string         `json:"name"`
	Description   string         `json:"description"`
	Fields        []Field        `json:"fields"`
	PossibleTypes []PossibleType `json:"possibleTypes"`
}

// NonNullFields retrieve non null fields
func (t *Type) NonNullFields() []Field {
	var fs []Field

	for _, f := range t.Fields {
		if f.Type.Kind == FieldTypeKindNonNull {
			fs = append(fs, f)
		}
	}

	return fs
}

// TypeKind a type's kind
type TypeKind string

const (
	// TypeKindObject an Object type kind
	TypeKindObject TypeKind = "OBJECT"
	// TypeKindInputObject an InputObject type kind
	TypeKindInputObject TypeKind = "INPUT_OBJECT"
	// TypeKindEnum an Enum type kind
	TypeKindEnum TypeKind = "ENUM"
	// TypeKindUnion a Union type kind
	TypeKindUnion TypeKind = "UNION"
	// TypeKindInterface an interface type kind
	TypeKindInterface TypeKind = "INTERFACE"
)

// Field a field on a GraphQL type
type Field struct {
	Name              string    `json:"name"`
	Description       string    `json:"description"`
	Type              FieldType `json:"type"`
	IsDeprecated      bool      `json:"isDeprecated"`
	DeprecationReason string    `json:"deprecationReason"`
}

// PossibleType a possible type for a GrahpQL uniob
type PossibleType struct {
	Kind   FieldTypeKind `json:"kind"`
	Name   string        `json:"name"`
	OfType *OfType       `json:"ofType"`
}

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
)

// OfType a nested GraphQl type
type OfType struct {
	Kind   FieldTypeKind `json:"kind"`
	Name   string        `json:"name"`
	OfType *OfType       `json:"ofType"`
}
