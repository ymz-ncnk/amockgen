package text_template

import (
	"bytes"
	template_mod "text/template"

	"github.com/ymz-ncnk/amockgen"
)

const baseTmplFile = "mock_implementation.go.tmpl"

// New creates a new AMockGen.
func New() (AMockGen, error) {
	baseTmpl := template_mod.New("base")
	registerFuncs(baseTmpl)
	err := registerTemplates(baseTmpl)
	if err != nil {
		return AMockGen{}, err
	}
	return AMockGen{baseTmpl}, err
}

// AMockGen is a code generator for AMock.
type AMockGen struct {
	baseTmpl *template_mod.Template
}

// Generate generates code from the mock implementation description.
func (aMockGen AMockGen) Generate(iDesc amockgen.MockImplDesc) (
	data []byte, err error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	err = aMockGen.baseTmpl.ExecuteTemplate(buf, baseTmplFile, iDesc)
	if err != nil {
		return
	}
	data = buf.Bytes()
	return
}

func registerFuncs(tmpl *template_mod.Template) {
	tmpl.Funcs(map[string]interface{}{
		"MakeCallParams":     amockgen.MakeCallParams,
		"MakeParams":         amockgen.MakeParams,
		"MakeReturnVars":     amockgen.MakeReturnVars,
		"MakeMethodTmplData": amockgen.MakeMethodTmplData,
		"include":            MakeIncludeFunc(tmpl),
	})
}

func registerTemplates(tmpl *template_mod.Template) (err error) {
	var childTmpl *template_mod.Template
	for name, template := range templates {
		childTmpl = tmpl.New(name)
		_, err = childTmpl.Parse(template)
		if err != nil {
			return
		}
	}
	return nil
}
