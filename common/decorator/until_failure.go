package decorator

import (
	"github.com/jbcpollak/go-behave/core"
)

// UntilFailure updates its child until it returns Failure.
func UntilFailure[Context any](params core.Params, child core.Node[Context]) core.Node[Context] {
	base := core.NewDecorator[Context]("UntilFailure", params, child)
	return &untilFailure[Context]{Decorator: base}
}

type untilFailure[Context any] struct {
	*core.Decorator[Context]
}

func (d *untilFailure[Context]) Enter(ctx Context) {}

func (d *untilFailure[Context]) Tick(ctx Context) core.Status {
	status := core.Update(d.Child, ctx)
	if status == core.StatusFailure {
		return core.StatusSuccess
	}
	return core.StatusRunning
}

func (d *untilFailure[Context]) Leave(ctx Context) {}
