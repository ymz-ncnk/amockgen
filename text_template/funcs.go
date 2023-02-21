package text_template

import (
	"bytes"
	template_mod "text/template"
)

// MakeIncludeFunc makes the template include func.
func MakeIncludeFunc(tmpl *template_mod.Template) func(string, interface{}) (
	string, error) {
	return func(name string, pipeline interface{}) (string, error) {
		var buf bytes.Buffer
		buf.WriteString("\n")
		if err := tmpl.ExecuteTemplate(&buf, name, pipeline); err != nil {
			return "", err
		}
		return buf.String(), nil
	}
}
