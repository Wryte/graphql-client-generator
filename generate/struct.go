package generate

import (
	"fmt"
	"html/template"
	"strings"

	graphql "github.com/Wryte/graphql-client-generator/graphql"
)

type structTemplateModel struct {
	Name        string
	TitleName   string
	Description template.HTML
	Fields      []structTemplateField
}

type structTemplateField struct {
	Name        string
	TitleName   string
	Type        structTemplateFieldType
	Description template.HTML
}

type structTemplateFieldType struct {
	Name      string
	IsList    bool
	IsPointer bool
}

func newStructTemplateModel(t graphql.Type) structTemplateModel {
	stm := structTemplateModel{
		Name:        t.Name,
		TitleName:   strings.Title(t.Name),
		Description: template.HTML(addComments(t.Description, "")),
	}

	for _, f := range t.Fields {
		var ft structTemplateFieldType
		switch f.Type.Kind {
		case graphql.FieldTypeKindInputObject:
			ft.Name = strings.Title(f.Type.Name)
		case graphql.FieldTypeKindEnum:
			ft.Name = strings.Title(f.Type.Name)
		case graphql.FieldTypeKindObject:
			ft.Name = f.Type.Name
			ft.IsPointer = true
		case graphql.FieldTypeKindScalar:
			n, ok := gqlGoScalarMap[f.Type.Name]
			if !ok {
				panic(fmt.Sprintf("Unknown scalar: %s (type - %s)\n%+v", f.Type.Name, t.Name, f))
			}
			ft.Name = n
		case graphql.FieldTypeKindList:
			navigateOfType(f.Type.OfType, &ft)
			ft.IsList = true
		case graphql.FieldTypeKindNonNull:
			navigateOfType(f.Type.OfType, &ft)
		case graphql.FieldTypeKindUnion:
			ft.Name = f.Type.Name
		case graphql.FieldTypeKindInterface:
			ft.Name = f.Type.Name
		default:
			panic(fmt.Sprintf("Unknown field type kind: %s\n%+v", f.Type.Kind, t))
		}

		if ft.Name != "" {
			stm.Fields = append(
				stm.Fields,
				structTemplateField{
					Name:        f.Name,
					TitleName:   strings.Title(f.Name),
					Type:        ft,
					Description: template.HTML(addComments(f.Description, "\t")),
				},
			)
		} else {
			panic(fmt.Sprintf("Could not retrieve name for field: %s\n%+v", f.Name, t))
		}
	}

	return stm
}

func navigateOfType(o *graphql.OfType, stft *structTemplateFieldType) {
	switch o.Kind {
	case graphql.FieldTypeKindNonNull:
	case graphql.FieldTypeKindInputObject:
		stft.Name = strings.Title(o.Name)
	case graphql.FieldTypeKindInterface:
		stft.Name = strings.Title(o.Name)
	case graphql.FieldTypeKindUnion:
		stft.Name = strings.Title(o.Name)
	case graphql.FieldTypeKindEnum:
		stft.Name = "string"
	case graphql.FieldTypeKindList:
		stft.IsList = true
	case graphql.FieldTypeKindObject:
		if !stft.IsList {
			stft.IsPointer = true
		}
		fallthrough
	case graphql.FieldTypeKindScalar:
		stft.Name = mapToGoScalar(o.Name)
	default:
		panic(fmt.Sprintf("Unknown ofType kind: %s, %s", o.Kind, o.Name))
	}

	if o.OfType == nil {
		return
	}

	navigateOfType(o.OfType, stft)
}

var gqlGoScalarMap = map[string]string{
	"String":          "string",
	"URL":             "string",
	"ID":              "string",
	"DateTime":        "time.Time",
	"Date":            "time.Time",
	"UnsignedInt64":   "int",
	"HTML":            "string",
	"Money":           "float64",
	"Float":           "float64",
	"Int":             "int",
	"FormattedString": "string",
	"JSON":            "string",
	"UtcOffset":       "string",
	"CurrencyCode":    "string",
	"Boolean":         "bool",
	"Decimal":         "string",
	"StorefrontID":    "string",
	"ARN":             "string",
}

func mapToGoScalar(s string) string {
	if m, ok := gqlGoScalarMap[s]; ok {
		return m
	}

	return s
}
