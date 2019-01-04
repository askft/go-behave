package decorator

import (
	"github.com/askft/go-behave/core"
)

// UntilSuccess ...
func UntilSuccess(params core.Params, child core.Node) core.Node {
	base := core.NewDecorator("UntilSuccess", params, child)
	return &untilSuccess{Decorator: base}
}

// untilSuccess ...
type untilSuccess struct {
	*core.Decorator
}

// Start ...
func (d *untilSuccess) Start(ctx *core.Context) {}

// Tick ...
func (d *untilSuccess) Tick(ctx *core.Context) core.Status {
	status := core.Update(d.Child, ctx)
	if status != core.StatusSuccess {
		return core.StatusRunning
	}
	return core.StatusSuccess
}

// Stop ...
func (d *untilSuccess) Stop(ctx *core.Context) {}
