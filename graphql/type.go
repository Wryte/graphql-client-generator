package graphql

// Type a type in the GraphQL schema
type Type struct {
	Kind          TypeKind       `json:"kind"`
	Name          string         `json:"name"`
	Description   string         `json:"description"`
	Fields        []Field        `json:"fields"`
	PossibleTypes []PossibleType `json:"possibleTypes"`
	EnumValues    []EnumValue    `json:"enumValues"`
	InputFields   []Field        `json:"inputFields"`
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
	// TypeKindScalar an scalar type kind
	TypeKindScalar TypeKind = "SCALAR"
)

// IsInputObject check if the type is an input object
func (t *Type) IsInputObject() bool {
	return t.Kind == TypeKindInputObject
}
