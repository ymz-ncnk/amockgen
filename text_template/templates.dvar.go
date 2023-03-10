// Code generated by dvargen. DO NOT EDIT.

package text_template

var templates map[string]string

func init() {
	templates = make(map[string]string)
	templates["method.go.tmpl"] = `{{- /* MethoDesc, MockImplName */ -}}
func (mock {{.MockImplName}}) {{.MethoDesc.Name}}({{ MakeParams .MethoDesc.Params }}) ({{ MakeReturnVars .MethoDesc.ReturnVars }}) {
	{{- range $index, $vDesc := .MethoDesc.Params }}
		{{- if $vDesc.Interface }}
			var {{$vDesc.Name}}Val reflect.Value
			if {{$vDesc.Name}} == nil {
				{{$vDesc.Name}}Val = reflect.Zero(reflect.TypeOf((*{{$vDesc.Type}})(nil)).Elem())
			} else {
				{{$vDesc.Name}}Val = reflect.ValueOf({{$vDesc.Name}})
			}
		{{- end }}
	{{- end }}
	{{- if eq (len .MethoDesc.ReturnVars) 0 }}
		_, err := mock.Call("{{.MethoDesc.Name}}", {{MakeCallParams .MethoDesc.Params}})
		if err != nil {
			panic(err)
		}
	{{- else }}
  	result, err := mock.Call("{{.MethoDesc.Name}}", {{MakeCallParams .MethoDesc.Params}})
		if err != nil {
			panic(err)
		}
		{{- range $index, $vDesc := .MethoDesc.ReturnVars }}
			{{- if $vDesc.Interface }}
				{{$vDesc.Name}}, _ = result[{{$index}}].({{$vDesc.Type}})	
			{{- else }}
				{{$vDesc.Name}} = result[{{$index}}].({{$vDesc.Type}})
			{{- end }}
		{{- end }}
		return		
	{{- end }}
}`
	templates["mock_implementation.go.tmpl"] = `{{- /* MockImplDesc */ -}}
// Code generated by amockgen. DO NOT EDIT.

package {{.Package}}

import amock_core "github.com/ymz-ncnk/amock/core"

// New creates a new {{.Name}}.
func New{{.Name}}() {{.Name}} {
	return {{.Name}} {
		Mock: amock_core.New("{{.Name}}"),
	}
}

// {{.Name}} is a mock implementation of the {{.InterfaceType}}.
type {{.Name}} struct {
	*amock_core.Mock
}

{{- $iDesc := . }}
{{- range $index, $mDesc := .Methods }}
	{{ include "register_method.go.tmpl" (MakeMethodTmplData $iDesc $mDesc) }}
	{{ include "register_n_method.go.tmpl" (MakeMethodTmplData $iDesc $mDesc) }}
	{{ include "unregister_method.go.tmpl" (MakeMethodTmplData $iDesc $mDesc) }}
{{- end }}

{{- range $index, $mDesc := .Methods }}
	{{ include "method.go.tmpl" (MakeMethodTmplData $iDesc $mDesc) }}
{{- end }}`
	templates["register_method.go.tmpl"] = `{{- /* MethoDesc, MockImplName */ -}}
// Register{{.MethoDesc.Name}} registers a function as a single {{.MethoDesc.Name}}() method call.
func (mock {{.MockImplName}}) Register{{.MethoDesc.Name}}(
  fn func({{ MakeParams .MethoDesc.Params }}) ({{ MakeReturnVars .MethoDesc.ReturnVars }})) {{.MockImplName}} {
  mock.Register("{{.MethoDesc.Name}}", fn)
  return mock
}`
	templates["register_n_method.go.tmpl"] = `{{- /* MethoDesc, MockImplName */ -}}
// RegisterN{{.MethoDesc.Name}} registers a function as n {{.MethoDesc.Name}}() method calls.
func (mock {{.MockImplName}}) RegisterN{{.MethoDesc.Name}}(n int,
  fn func({{ MakeParams .MethoDesc.Params }}) ({{ MakeReturnVars .MethoDesc.ReturnVars }})) {{.MockImplName}} {
  mock.RegisterN("{{.MethoDesc.Name}}", n, fn)
  return mock
}`
	templates["unregister_method.go.tmpl"] = `{{- /* MethoDesc, MockImplName */ -}}
// Unregister{{.MethoDesc.Name}} unregisters {{.MethoDesc.Name}}() method calls.
func (mock {{.MockImplName}}) Unregister{{.MethoDesc.Name}}() {{.MockImplName}} {
  mock.Unregister("{{.MethoDesc.Name}}")
  return mock
}`
}
