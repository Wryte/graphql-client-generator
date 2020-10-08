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

// Field a field on a GraphQL type
type Field struct {
	Name              string    `json:"name"`
	Description       string    `json:"description"`
	Type              FieldType `json:"type"`
	IsDeprecated      bool      `json:"isDeprecated"`
	DeprecationReason string    `json:"deprecationReason"`
	Args              []Arg     `json:"args"`
}

// TypeName name of the field's type
func (f *Field) TypeName() string {
	if f.Type.Name != "" {
		return f.Type.Name
	}

	if f.Type.OfType == nil {
		return ""
	}

	return f.Type.OfType.TypeName()
}

// IsNonNull check if field is non null
func (f *Field) IsNonNull() bool {
	return f.Type.Kind == FieldTypeKindNonNull
}

// IsInterface check if field is interface
func (f *Field) IsInterface() bool {
	if f.Type.Kind == FieldTypeKindInterface {
		return true
	}

	if f.Type.OfType == nil {
		return false
	}

	return f.Type.OfType.IsInterface()
}

// IsEnum check if field is enum
func (f *Field) IsEnum() bool {
	if f.Type.Kind == FieldTypeKindEnum {
		return true
	}

	if f.Type.OfType == nil {
		return false
	}

	return f.Type.OfType.IsEnum()
}

// IsObject check if field is object
func (f *Field) IsObject() bool {
	if f.Type.Kind == FieldTypeKindObject {
		return true
	}

	if f.Type.OfType == nil {
		return false
	}

	return f.Type.OfType.IsObject()
}

// IsScalar check if field is scalar
func (f *Field) IsScalar() bool {
	if f.Type.Kind == FieldTypeKindScalar {
		return true
	}

	if f.Type.OfType == nil {
		return false
	}

	return f.Type.OfType.IsScalar()
}

// IsList check if field is list
func (f *Field) IsList() bool {
	if f.Type.Kind == FieldTypeKindList {
		return true
	}

	if f.Type.OfType == nil {
		return false
	}

	return f.Type.OfType.IsList()
}

// IsUnion check if field is union
func (f *Field) IsUnion() bool {
	if f.Type.Kind == FieldTypeKindUnion {
		return true
	}

	if f.Type.OfType == nil {
		return false
	}

	return f.Type.OfType.IsUnion()
}

// IsInputObject check if field is input object
func (f *Field) IsInputObject() bool {
	if f.Type.Kind == FieldTypeKindInputObject {
		return true
	}

	if f.Type.OfType == nil {
		return false
	}

	return f.Type.OfType.IsInputObject()
}

// PossibleType a possible type for a GrahpQL union
type PossibleType struct {
	Kind   FieldTypeKind `json:"kind"`
	Name   string        `json:"name"`
	OfType *OfType       `json:"ofType"`
}

// EnumValue a value of an enum kind
type EnumValue struct {
	Name              string `json:"name"`
	Description       string `json:"description"`
	IsDeprecated      bool   `json:"isDeprecated"`
	DeprecationReason string `json:"deprecationReason"`
}

// Arg a mutation argument
type Arg struct {
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	Type         FieldType `json:"type"`
	DefaultValue string
}

// TypeName name of the field's type
func (a *Arg) TypeName() string {
	if a.Type.Name != "" {
		return a.Type.Name
	}

	if a.Type.OfType == nil {
		return ""
	}

	return a.Type.OfType.TypeName()
}

// IsNonNull check if field is non null
func (a *Arg) IsNonNull() bool {
	return a.Type.Kind == FieldTypeKindNonNull
}

// IsInterface check if field is interface
func (a *Arg) IsInterface() bool {
	if a.Type.Kind == FieldTypeKindInterface {
		return true
	}

	if a.Type.OfType == nil {
		return false
	}

	return a.Type.OfType.IsInterface()
}

// IsEnum check if field is enum
func (a *Arg) IsEnum() bool {
	if a.Type.Kind == FieldTypeKindEnum {
		return true
	}

	if a.Type.OfType == nil {
		return false
	}

	return a.Type.OfType.IsEnum()
}

// IsObject check if field is object
func (a *Arg) IsObject() bool {
	if a.Type.Kind == FieldTypeKindObject {
		return true
	}

	if a.Type.OfType == nil {
		return false
	}

	return a.Type.OfType.IsObject()
}

// IsScalar check if field is scalar
func (a *Arg) IsScalar() bool {
	if a.Type.Kind == FieldTypeKindScalar {
		return true
	}

	if a.Type.OfType == nil {
		return false
	}

	return a.Type.OfType.IsScalar()
}

// IsList check if field is list
func (a *Arg) IsList() bool {
	if a.Type.Kind == FieldTypeKindList {
		return true
	}

	if a.Type.OfType == nil {
		return false
	}

	return a.Type.OfType.IsList()
}

// IsUnion check if field is union
func (a *Arg) IsUnion() bool {
	if a.Type.Kind == FieldTypeKindUnion {
		return true
	}

	if a.Type.OfType == nil {
		return false
	}

	return a.Type.OfType.IsUnion()
}

// IsInputObject check if field is input object
func (a *Arg) IsInputObject() bool {
	if a.Type.Kind == FieldTypeKindInputObject {
		return true
	}

	if a.Type.OfType == nil {
		return false
	}

	return a.Type.OfType.IsInputObject()
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
	// FieldTypeKindInputObject a field type kind Input Object
	FieldTypeKindInputObject FieldTypeKind = "INPUT_OBJECT"
)

// OfType a nested GraphQl type
type OfType struct {
	Kind   FieldTypeKind `json:"kind"`
	Name   string        `json:"name"`
	OfType *OfType       `json:"ofType"`
}

// TypeName name of the type
func (o *OfType) TypeName() string {
	if o.Name != "" {
		return o.Name
	}

	if o.OfType == nil {
		return ""
	}

	return o.OfType.TypeName()
}

// IsNonNull navigate of types for non null
func (o *OfType) IsNonNull() bool {
	if o.Kind == FieldTypeKindNonNull {
		return true
	}

	if o.OfType == nil {
		return false
	}

	return o.OfType.IsNonNull()
}

// IsInterface navigate of types for interface
func (o *OfType) IsInterface() bool {
	if o.Kind == FieldTypeKindInterface {
		return true
	}

	if o.OfType == nil {
		return false
	}

	return o.OfType.IsInterface()
}

// IsEnum navigate of types for Enum
func (o *OfType) IsEnum() bool {
	if o.Kind == FieldTypeKindEnum {
		return true
	}

	if o.OfType == nil {
		return false
	}

	return o.OfType.IsEnum()
}

// IsObject navigate of types for object
func (o *OfType) IsObject() bool {
	if o.Kind == FieldTypeKindObject {
		return true
	}

	if o.OfType == nil {
		return false
	}

	return o.OfType.IsObject()
}

// IsScalar navigate of types for scalar
func (o *OfType) IsScalar() bool {
	if o.Kind == FieldTypeKindScalar {
		return true
	}

	if o.OfType == nil {
		return false
	}

	return o.OfType.IsScalar()
}

// IsList navigate of types for list
func (o *OfType) IsList() bool {
	if o.Kind == FieldTypeKindList {
		return true
	}

	if o.OfType == nil {
		return false
	}

	return o.OfType.IsList()
}

// IsUnion navigate of types for union
func (o *OfType) IsUnion() bool {
	if o.Kind == FieldTypeKindUnion {
		return true
	}

	if o.OfType == nil {
		return false
	}

	return o.OfType.IsUnion()
}

// IsInputObject navigate of types for input object
func (o *OfType) IsInputObject() bool {
	if o.Kind == FieldTypeKindInputObject {
		return true
	}

	if o.OfType == nil {
		return false
	}

	return o.OfType.IsInputObject()
}
