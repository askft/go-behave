package decorator

import (
	"github.com/askft/go-behave/core"
)

// UntilSuccess updates its child until it returns Success.
func UntilSuccess(params core.Params, child core.Node) core.Node {
	base := core.NewDecorator("UntilSuccess", params, child)
	return &untilSuccess{Decorator: base}
}

type untilSuccess struct {
	*core.Decorator
}

func (d *untilSuccess) Enter(ctx *core.Context) {}

func (d *untilSuccess) Tick(ctx *core.Context) core.Status {
	status := core.Update(d.Child, ctx)
	if status == core.StatusSuccess {
		return core.StatusSuccess
	}
	return core.StatusRunning
}

func (d *untilSuccess) Leave(ctx *core.Context) {}
