package decorator

import (
	"github.com/alexanderskafte/behaviortree/core"
)

// inverter ...
type inverter struct {
	*core.Decorator
}

// Inverter ...
func Inverter(params core.Params, child core.INode) core.INode {
	base := core.NewDecorator("Inverter", params)
	base.Child = child
	return &delayer{Decorator: base}
}

// Start ...
func (d *inverter) Start(ctx *core.Context) {}

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

// Stop ...
func (d *inverter) Stop(ctx *core.Context) {}
