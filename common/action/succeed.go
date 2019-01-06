package action

import (
	"github.com/askft/go-behave/core"
)

// Succeed returns a new succeed node, which always succeeds in one tick.
func Succeed(params core.Params, returns core.Returns) core.Node {
	base := core.NewLeaf("Succeed", params, returns)
	return &succeed{Leaf: base}
}

// succeed ...
type succeed struct {
	*core.Leaf
}

// Enter ...
func (a *succeed) Enter(ctx *core.Context) {}

// Tick ...
func (a *succeed) Tick(ctx *core.Context) core.Status {
	return core.StatusSuccess
}

// Leave ...
func (a *succeed) Leave(ctx *core.Context) {}
