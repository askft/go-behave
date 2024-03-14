package decorator

import (
	"github.com/jbcpollak/go-behave/core"
)

// Inverter ...
func Inverter[Context any](params core.Params, child core.Node[Context]) core.Node[Context] {
	base := core.NewDecorator("Inverter", params, child)
	return &inverter[Context]{Decorator: base}
}

// inverter ...
type inverter[Context any] struct {
	*core.Decorator[Context]
}

// Enter ...
func (d *inverter[Context]) Enter(ctx Context) {}

// Tick ...
func (d *inverter[Context]) Tick(ctx Context) core.Status {
	switch core.Update(d.Child, ctx) {
	case core.StatusSuccess:
		return core.StatusFailure
	case core.StatusFailure:
		return core.StatusSuccess
	default:
		return core.StatusRunning
	}
}

// Leave ...
func (d *inverter[Context]) Leave(ctx Context) {}
