package decorator

import (
	"github.com/askft/go-behave/core"
)

// Inverter ...
func Inverter(params core.Params, child core.Node) core.Node {
	base := core.NewDecorator("Inverter", params, child)
	return &inverter{Decorator: base}
}

// inverter ...
type inverter struct {
	*core.Decorator
}

// Enter ...
func (d *inverter) Enter(ctx *core.Context) {}

// Tick ...
func (d *inverter) Tick(ctx *core.Context) core.Status {
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
func (d *inverter) Leave(ctx *core.Context) {}
