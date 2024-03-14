package action

import (
	"github.com/jbcpollak/go-behave/core"
)

// Succeed returns a new succeed node, which always succeeds in one tick.
func Succeed[Context any](params core.Params, returns core.Returns) core.Node[Context] {
	base := core.NewLeaf[Context]("Succeed", params, returns)
	return &succeed[Context]{Leaf: base}
}

// succeed ...
type succeed[Context any] struct {
	*core.Leaf[Context]
}

// Enter ...
func (a *succeed[Context]) Enter(ctx Context) {}

// Tick ...
func (a *succeed[Context]) Tick(ctx Context) core.Status {
	return core.StatusSuccess
}

// Leave ...
func (a *succeed[Context]) Leave(ctx Context) {}
