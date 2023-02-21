package main

import (
	"github.com/ymz-ncnk/dvargen"
)

func main() {
	err := dvargen.Generate(dvargen.DVarDesc{
		Dir:     "templates/go/",
		Varname: "templates",
		Package: "text_template",
	})
	if err != nil {
		panic(err)
	}
}
