package generate

import (
	"fmt"
	"html/template"
	"strings"

	graphql "github.com/Wryte/graphql-client-generator/graphql"
)

type mutationTemplateModel struct {
	TitleName      string
	Description    template.HTML
	Args           []mutationTemplateArg
	ReturnType     string
	ArgsDefinition string
}

type mutationTemplateArg struct {
	Name                string
	Description         template.HTML
	GoParamDef          string
	MutationParamDef    string
	MutationArgumentDef string
}

func newMutationTemplateModel(s graphql.Schema, f graphql.Field) mutationTemplateModel {
	mt := mutationTemplateModel{
		TitleName:   strings.Title(f.Name),
		Description: template.HTML(addComments(f.Description, "")),
		ReturnType:  f.Type.Name,
	}

	for _, a := range f.Args {
		var (
			prefix = "*"
			suffix string
		)

		if a.IsNonNull() {
			prefix = ""
			suffix = "!"
		}

		if a.IsList() {
			prefix = "[]"
		}

		mt.Args = append(mt.Args, mutationTemplateArg{
			Name:                a.Name,
			Description:         template.HTML(addComments(a.Description, "\t")),
			GoParamDef:          fmt.Sprintf("%s %s%s", a.Name, prefix, mapToGoScalar(a.TypeName())),
			MutationParamDef:    fmt.Sprintf("$%s: %s%s,", a.Name, a.TypeName(), suffix),
			MutationArgumentDef: fmt.Sprintf("%s: $%s,", a.Name, a.Name),
		})
	}

	// TODO generate response structure from returntype
	// rt := s.Type(f.Type.Name)

	return mt
}
