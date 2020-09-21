package generate

import (
	"fmt"
	"html/template"

	graphql "github.com/Wryte/graphql-client-generator/graphql"
)

type structTemplateModel struct {
	Name        string
	GoName      string
	Description template.HTML
	Fields      []structTemplateField
}

type structTemplateField struct {
	Name         string
	GoName       string
	Description  template.HTML
	GoTypeDef    string
	GoGraphQLTag template.HTML
}

func newStructTemplateModel(t graphql.Type) structTemplateModel {
	stm := structTemplateModel{
		Name:        t.Name,
		GoName:      makeExportedName(t.Name),
		Description: template.HTML(addComments(t.Description, "")),
	}

	for _, f := range t.Fields {
		var prefix string
		var graphQLTag string
		if !f.IsNonNull() || f.IsObject() {
			prefix = "*"
		}
		if f.IsList() {
			prefix = "[]"
		}
		if f.IsNonNull() && f.IsInputObject() {
			graphQLTag = ` gql:"required"`
		}

		fieldType := mapToGoScalar(f.TypeName())
		if fieldType == f.TypeName() {
			fieldType = makeExportedName(fieldType)
		}

		stm.Fields = append(
			stm.Fields,
			structTemplateField{
				Name:         f.Name,
				GoName:       makeExportedName(f.Name),
				Description:  template.HTML(addComments(f.Description, "\t")),
				GoTypeDef:    fmt.Sprintf("%s%s", prefix, fieldType),
				GoGraphQLTag: template.HTML(graphQLTag),
			},
		)
	}

	return stm
}
