package action

import (
	"github.com/askft/go-behave/core"
)

// Fail returns a new fail node, which always fails in one tick.
func Fail(params core.Params, returns core.Returns) core.Node {
	base := core.NewAction("Fail", params, returns)
	return &fail{Action: base}
}

// fail ...
type fail struct {
	*core.Action
}

// Start ...
func (a *fail) Start(ctx *core.Context) {}

// Tick ...
func (a *fail) Tick(ctx *core.Context) core.Status {
	return core.StatusFailure
}

// Stop ...
func (a *fail) Stop(ctx *core.Context) {}
