package generate

import (
	"html/template"
	"strings"

	graphql "github.com/Wryte/graphql-client-generator/graphql"
)

type unionTemplateModel struct {
	Name          string
	TitleName     string
	FirstChar     string
	Description   template.HTML
	PossibleTypes []unionTemplatePossibleType
}

type unionTemplatePossibleType struct {
	Kind      graphql.FieldTypeKind
	Name      string
	TitleName string
	ofType    *graphql.OfType
}

func newUnionTemplateModel(t graphql.Type) unionTemplateModel {
	ut := unionTemplateModel{
		Name:        t.Name,
		TitleName:   strings.Title(t.Name),
		FirstChar:   strings.ToLower(strings.Split(t.Name, "")[0]),
		Description: template.HTML(addComments(t.Description, "")),
	}

	for _, pt := range t.PossibleTypes {
		ut.PossibleTypes = append(ut.PossibleTypes, unionTemplatePossibleType{
			Kind:      pt.Kind,
			Name:      pt.Name,
			TitleName: strings.Title(pt.Name),
		})
	}

	return ut
}
