package amockgen

// MockImplDesc is the description of a mock implementation.
type MockImplDesc struct {
	InterfaceType string
	Package       string
	Name          string
	Methods       []MethoDesc
}

// MethoDesc is the description of a method.
type MethoDesc struct {
	Name       string
	Params     []VarDesc
	ReturnVars []VarDesc
}

// VarDesc is the description of a variable.
type VarDesc struct {
	Name      string
	Type      string
	Interface bool
}
