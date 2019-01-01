package core

// Context is data implicitly shared by all nodes in a behavior tree
// since a Context instance is propagated through the tree each tick.
type Context struct {
	Owner interface{}
	Data  interface{}
}

// NewContext creates context containing references to an owner and a store.
func NewContext(owner, data interface{}) *Context {
	return &Context{
		owner,
		data,
	}
}
