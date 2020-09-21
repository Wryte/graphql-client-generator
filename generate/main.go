package generate

import (
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"unicode"

	graphql "github.com/Wryte/graphql-client-generator/graphql"
)

const (
	structTemplateName   = "struct"
	unionTemplateName    = "union"
	enumTemplateName     = "enum"
	mutationTemplateName = "mutation"
)

var (
	templateStrings = []string{
		structTemplateName,
		unionTemplateName,
		enumTemplateName,
		mutationTemplateName,
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
		if t.Name == s.MutationType.Name {
			// for _, f := range t.Fields {
			// 	mtm := newMutationTemplateModel(s, f)
			// 	err = tpls[mutationTemplateName].Execute(wr, mtm)

			// 	if err != nil {
			// 		return fmt.Errorf("could not generate from mutation=%s: %s", f.Name, err)
			// 	}
			// }
		} else if strings.Index(t.Name, "__") != 0 {
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

var gqlGoScalarMap = map[string]string{
	"String":          "string",
	"URL":             "string",
	"ID":              "string",
	"DateTime":        "time.Time",
	"Date":            "string",
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

func makeExportedName(name string) string {
	return lintName(strings.Title(name))
}

// lintName returns a different name if it should be different.
func lintName(name string) (should string) {
	// Fast path for simple cases: "_" and all lowercase.
	if name == "_" {
		return name
	}
	allLower := true
	for _, r := range name {
		if !unicode.IsLower(r) {
			allLower = false
			break
		}
	}
	if allLower {
		return name
	}

	// Split camelCase at any lower->upper transition, and split on underscores.
	// Check each word for common initialisms.
	runes := []rune(name)
	w, i := 0, 0 // index of start of word, scan
	for i+1 <= len(runes) {
		eow := false // whether we hit the end of a word
		if i+1 == len(runes) {
			eow = true
		} else if runes[i+1] == '_' {
			// underscore; shift the remainder forward over any run of underscores
			eow = true
			n := 1
			for i+n+1 < len(runes) && runes[i+n+1] == '_' {
				n++
			}

			// Leave at most one underscore if the underscore is between two digits
			if i+n+1 < len(runes) && unicode.IsDigit(runes[i]) && unicode.IsDigit(runes[i+n+1]) {
				n--
			}

			copy(runes[i+1:], runes[i+n+1:])
			runes = runes[:len(runes)-n]
		} else if unicode.IsLower(runes[i]) && !unicode.IsLower(runes[i+1]) {
			// lower->non-lower
			eow = true
		}
		i++
		if !eow {
			continue
		}

		// [w,i) is a word.
		word := string(runes[w:i])
		if u := strings.ToUpper(word); commonInitialisms[u] {
			// Keep consistent case, which is lowercase only at the start.
			if w == 0 && unicode.IsLower(runes[w]) {
				u = strings.ToLower(u)
			}
			// All the common initialisms are ASCII,
			// so we can replace the bytes exactly.
			copy(runes[w:], []rune(u))
		} else if w > 0 && strings.ToLower(word) == word {
			// already all lowercase, and not the first word, so uppercase the first character.
			runes[w] = unicode.ToUpper(runes[w])
		}
		w = i
	}
	return string(runes)
}

// commonInitialisms is a set of common initialisms.
// Only add entries that are highly unlikely to be non-initialisms.
// For instance, "ID" is fine (Freudian code is rare), but "AND" is not.
var commonInitialisms = map[string]bool{
	"ACL":   true,
	"API":   true,
	"ASCII": true,
	"CPU":   true,
	"CSS":   true,
	"DNS":   true,
	"EOF":   true,
	"GUID":  true,
	"HTML":  true,
	"HTTP":  true,
	"HTTPS": true,
	"ID":    true,
	"IP":    true,
	"JSON":  true,
	"LHS":   true,
	"QPS":   true,
	"RAM":   true,
	"RHS":   true,
	"RPC":   true,
	"SLA":   true,
	"SMTP":  true,
	"SQL":   true,
	"SSH":   true,
	"TCP":   true,
	"TLS":   true,
	"TTL":   true,
	"UDP":   true,
	"UI":    true,
	"UID":   true,
	"UUID":  true,
	"URI":   true,
	"URL":   true,
	"UTF8":  true,
	"VM":    true,
	"XML":   true,
	"XMPP":  true,
	"XSRF":  true,
	"XSS":   true,
}
