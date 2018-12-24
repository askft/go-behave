package action

import (
	"github.com/alexanderskafte/behaviortree/core"
)

// succeed ...
type succeed struct {
	*core.Action
}

// Succeed returns a new succeed node.
func Succeed(params, returns []string) core.INode {
	base := core.NewAction("Succeed", params, returns)
	return &succeed{Action: base}
}

// Start ...
func (a *succeed) Start(ctx *core.Context) {}

// Tick ...
func (a *succeed) Tick(ctx *core.Context) core.Status {
	return core.StatusSuccess
}

// Stop ...
func (a *succeed) Stop(ctx *core.Context) {}
