package simplevm

type Context struct {
	returnPc int
	locals   []int
}

// NewContext is a function which create a new context.
func NewContext(retPc, nlocals int) *Context {
	return &Context{returnPc: retPc, locals: make([]int, nlocals)}
}

// GetReturnPC is a method to get a return program counter.
func (c *Context) GetReturnPC() int {
	return c.returnPc
}
