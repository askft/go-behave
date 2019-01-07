package decorator

import (
	"github.com/askft/go-behave/core"
)

// UntilFailure updates its child until it returns Failure.
func UntilFailure(params core.Params, child core.Node) core.Node {
	base := core.NewDecorator("UntilFailure", params, child)
	return &untilFailure{Decorator: base}
}

type untilFailure struct {
	*core.Decorator
}

func (d *untilFailure) Enter(ctx *core.Context) {}

func (d *untilFailure) Tick(ctx *core.Context) core.Status {
	status := core.Update(d.Child, ctx)
	if status == core.StatusFailure {
		return core.StatusSuccess
	}
	return core.StatusRunning
}

func (d *untilFailure) Leave(ctx *core.Context) {}
