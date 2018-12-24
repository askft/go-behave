package action

import (
	"github.com/alexanderskafte/behaviortree/core"
)

// fail ...
type fail struct {
	*core.Action
}

// Fail returns a new fail node.
func Fail(params, returns []string) core.INode {
	base := core.NewAction("Fail", params, returns)
	return &fail{Action: base}
}

// Start ...
func (a *fail) Start(ctx *core.Context) {}

// Tick ...
func (a *fail) Tick(ctx *core.Context) core.Status {
	return core.StatusFailure
}

// Stop ...
func (a *fail) Stop(ctx *core.Context) {}
