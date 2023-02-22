//go:generate go run gen/dvar.go
package amockgen

// AMockGen is a code generator for AMock.
type AMockGen interface {
	// Generate generates code from the mock implementation description.
	Generate(iDesc MockImplDesc) (data []byte, err error)
}
