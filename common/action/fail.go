package action

import (
	"github.com/jbcpollak/go-behave/core"
)

// Fail returns a new fail node, which always fails in one tick.
func Fail[Context any](params core.Params, returns core.Returns) core.Node[Context] {
	base := core.NewLeaf[Context]("Fail", params, returns)
	return &fail[Context]{Leaf: base}
}

// fail ...
type fail[Context any] struct {
	*core.Leaf[Context]
}

// Enter ...
func (a *fail[Context]) Enter(ctx Context) {}

// Tick ...
func (a *fail[Context]) Tick(ctx Context) core.Status {
	return core.StatusFailure
}

// Leave ...
func (a *fail[Context]) Leave(ctx Context) {}
