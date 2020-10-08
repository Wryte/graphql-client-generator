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
	GoJSONTag    template.HTML
}

func newStructTemplateModel(t graphql.Type) structTemplateModel {
	stm := structTemplateModel{
		Name:        t.Name,
		GoName:      makeExportedName(t.Name),
		Description: template.HTML(addComments(t.Description, "")),
	}

	for _, f := range t.Fields {
		var (
			prefix     string
			graphQLTag string
			jsonTag    = f.Name
		)

		if !f.IsNonNull() || f.IsObject() {
			prefix = "*"
		}

		if f.IsList() {
			prefix = "[]"
		}

		if t.IsInputObject() {
			if f.IsNonNull() {
				graphQLTag = ` gql:"required"`
			} else {
				jsonTag = fmt.Sprintf("%s,omitempty", jsonTag)
			}
		}

		fieldType := mapToGoScalar(f.TypeName())
		if fieldType == f.TypeName() {
			fieldType = makeExportedName(fieldType)
		}

		stm.Fields = append(
			stm.Fields,
			structTemplateField{
				GoJSONTag:    template.HTML(jsonTag),
				GoName:       makeExportedName(f.Name),
				Description:  template.HTML(addComments(f.Description, "\t")),
				GoTypeDef:    fmt.Sprintf("%s%s", prefix, fieldType),
				GoGraphQLTag: template.HTML(graphQLTag),
			},
		)
	}

	return stm
}
