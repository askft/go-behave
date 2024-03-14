package decorator

import (
	"github.com/jbcpollak/go-behave/core"
)

// UntilSuccess updates its child until it returns Success.
func UntilSuccess[Context any](params core.Params, child core.Node[Context]) core.Node[Context] {
	base := core.NewDecorator("UntilSuccess", params, child)
	return &untilSuccess[Context]{Decorator: base}
}

type untilSuccess[Context any] struct {
	*core.Decorator[Context]
}

func (d *untilSuccess[Context]) Enter(ctx Context) {}

func (d *untilSuccess[Context]) Tick(ctx Context) core.Status {
	status := core.Update(d.Child, ctx)
	if status == core.StatusSuccess {
		return core.StatusSuccess
	}
	return core.StatusRunning
}

func (d *untilSuccess[Context]) Leave(ctx Context) {}
