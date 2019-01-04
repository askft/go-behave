package decorator

import (
	"github.com/askft/go-behave/core"
)

// UntilFailure ...
func UntilFailure(params core.Params, child core.Node) core.Node {
	base := core.NewDecorator("UntilFailure", params, child)
	return &untilFailure{Decorator: base}
}

// untilFailure ...
type untilFailure struct {
	*core.Decorator
}

// Start ...
func (d *untilFailure) Start(ctx *core.Context) {}

// Tick ...
func (d *untilFailure) Tick(ctx *core.Context) core.Status {
	status := core.Update(d.Child, ctx)
	if status != core.StatusFailure {
		return core.StatusRunning
	}
	return core.StatusSuccess
}

// Stop ...
func (d *untilFailure) Stop(ctx *core.Context) {}
