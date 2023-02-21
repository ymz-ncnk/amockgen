package amockgen

// MakeCallParams makes a list of parameters for calling a function.
func MakeCallParams(params []VarDesc) (result string) {
	for i := 0; i < len(params); i++ {
		result += params[i].Name
		if params[i].Interface {
			result += "Val"
		}
		result += ","
	}
	return
}

// MakeCallParams makes a list of function parameters.
func MakeParams(params []VarDesc) (result string) {
	for i := 0; i < len(params); i++ {
		result += params[i].Name + " " + params[i].Type + ","
	}
	return
}

// MakeCallParams makes a list of function return values.
func MakeReturnVars(returnVars []VarDesc) (result string) {
	for i := 0; i < len(returnVars); i++ {
		result += returnVars[i].Name + " " + returnVars[i].Type + ","
	}
	return
}

// MakeMethodTmplData makes a data for method template.
func MakeMethodTmplData(iDesc MockImplDesc, mDesc MethoDesc) struct {
	MethoDesc    MethoDesc
	MockImplName string
} {
	return struct {
		MethoDesc    MethoDesc
		MockImplName string
	}{
		MethoDesc:    mDesc,
		MockImplName: iDesc.Name,
	}
}
