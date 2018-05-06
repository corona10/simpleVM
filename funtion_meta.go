package simplevm

type FunctionMetaData struct {
	name    string
	nargs   int
	nlocals int
	address int
}

// GetName is a method for get a name of function from meta data.
func (m *FunctionMetaData) GetName() string {
	return m.name
}

// GetNArgs is a method for get a number of arguments from function.
func (m *FunctionMetaData) GetNArgs() int {
	return m.nargs
}

// GetNLocals is a method for get a number of local variables from function.
func (m *FunctionMetaData) GetNLocals() int {
	return m.nlocals
}

// GetAddress is a method for get a address of funtion.
func (m *FunctionMetaData) GetAddress() int {
	return m.address
}
