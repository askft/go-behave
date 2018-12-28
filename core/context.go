package core

import (
	"github.com/alexanderskafte/go-behave/store"
)

// Context is data implicitly shared by all nodes in a behavior tree
// since a Context instance is propagated through the tree each tick.
type Context struct {
	Owner interface{}
	Store store.Interface
}

// NewContext creates context containing references to an owner and a store.
func NewContext(owner interface{}, store store.Interface) *Context {
	return &Context{
		owner,
		store,
	}
}
