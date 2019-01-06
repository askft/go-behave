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

// Enter ...
func (d *untilSuccess) Enter(ctx *core.Context) {}

// Tick ...
func (d *untilSuccess) Tick(ctx *core.Context) core.Status {
	status := core.Update(d.Child, ctx)
	if status != core.StatusSuccess {
		return core.StatusRunning
	}
	return core.StatusSuccess
}

// Leave ...
func (d *untilSuccess) Leave(ctx *core.Context) {}
