package generate

import (
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"os"
	"strings"

	graphql "github.com/Wryte/graphql-client-generator/graphql"
)

const (
	structTemplateName = "struct"
	unionTemplateName  = "union"
	enumTemplateName   = "enum"
)

var (
	templateStrings = []string{
		structTemplateName,
		unionTemplateName,
		enumTemplateName,
	}
	tpls = map[string]*template.Template{}
)

// Write writes the generated code to the io writer
func Write(wr io.Writer, s graphql.Schema) error {
	err := loadTemplates()
	if err != nil {
		return fmt.Errorf("failed to load templates: %s", err)
	}

	for _, t := range s.Types {
		if strings.Index(t.Name, "__") != 0 && t.Name != s.MutationType.Name {
			var err error
			switch t.Kind {
			case graphql.TypeKindScalar:
			case graphql.TypeKindEnum:
				etm := newEnumTemplateModel(t)
				err = tpls[enumTemplateName].Execute(wr, etm)
			case graphql.TypeKindInputObject:
				t.Fields = t.InputFields
				fallthrough
			case graphql.TypeKindInterface:
				fallthrough
			case graphql.TypeKindObject:
				stm := newStructTemplateModel(t)
				err = tpls[structTemplateName].Execute(wr, stm)
			case graphql.TypeKindUnion:
				utm := newUnionTemplateModel(t)
				err = tpls[unionTemplateName].Execute(wr, utm)
			default:
				panic(fmt.Sprintf("Unkown type kind of %s\n%+v", t.Kind, t))
			}

			if err != nil {
				return fmt.Errorf("could not generate from type=%s: %s", t.Name, err)
			}
		}
	}

	return nil
}

func loadTemplates() error {
	for _, ts := range templateStrings {
		tpl, err := loadTemplate(ts)

		if err != nil {
			return err
		}

		tpls[ts] = tpl
	}

	return nil
}

func loadTemplate(name string) (*template.Template, error) {
	tmplFile, err := os.Open(fmt.Sprintf("templates/%s.tmpl", name))

	if err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}
	defer tmplFile.Close()

	bytes, err := ioutil.ReadAll(tmplFile)

	if err != nil {
		return nil, fmt.Errorf("error reading whole file: %v", err)
	}

	tmpl, err := template.New(name).Parse(string(bytes))

	if err != nil {
		return nil, fmt.Errorf("error creating model template: %v", err)
	}

	return tmpl, nil
}

func addComments(s, prefix string) string {
	parts := strings.Split(strings.Trim(s, "\n"), "\n")
	return strings.Join(parts, fmt.Sprintf("\n%s// ", prefix))
}
