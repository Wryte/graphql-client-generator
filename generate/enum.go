package generate

import (
	"fmt"
	"html/template"
	"strings"

	graphql "github.com/Wryte/graphql-client-generator/graphql"
)

type enumTemplateModel struct {
	GoName      string
	Description template.HTML
	Values      []enumTemplateValue
}

type enumTemplateValue struct {
	GoName      string
	Description template.HTML
	Value       string
}

func newEnumTemplateModel(t graphql.Type) enumTemplateModel {
	et := enumTemplateModel{
		GoName:      makeExportedName(t.Name),
		Description: template.HTML(addComments(t.Description, "")),
	}

	for _, ev := range t.EnumValues {
		et.Values = append(et.Values, enumTemplateValue{
			GoName:      makeExportedName(shoutingSnakeToPascal(ev.Name)),
			Description: template.HTML(addComments(ev.Description, "\t")),
			Value:       ev.Name,
		})
	}

	return et
}

func shoutingSnakeToPascal(s string) string {
	parts := strings.Split(s, "_")
	var r string

	for _, p := range parts {
		r = fmt.Sprintf("%s%s", r, strings.Title(strings.ToLower(p)))
	}

	return r
}
