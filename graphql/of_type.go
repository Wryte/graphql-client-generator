package graphql

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
