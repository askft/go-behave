package core

import (
	"github.com/alexanderskafte/go-behave/store"
)

// Context ...
type Context struct {
	Owner interface{}
	Store store.Interface
}

// NewContext ...
func NewContext(owner interface{}, store store.Interface) *Context {
	return &Context{
		owner,
		store,
	}
}
