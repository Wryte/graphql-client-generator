package graphql

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
