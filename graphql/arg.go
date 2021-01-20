package graphql

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
