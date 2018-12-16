package core

import (
	"fmt"

	"github.com/alexanderskafte/behaviortree/store"
)

// Context is renewed every time the tree is run.
type Context struct {
	owner fmt.Stringer
	store store.Interface
}

// NewContext ...
func NewContext(owner fmt.Stringer, store store.Interface) *Context {
	if owner == nil {
		panic("owner is nil")
	}
	if store == nil {
		panic("board is nil")
	}
	return &Context{owner, store}
}
