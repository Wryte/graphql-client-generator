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
