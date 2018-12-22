package decorator

import (
	"fmt"

	"github.com/alexanderskafte/behaviortree/core"
)

// Inverter ...
type Inverter struct {
	*core.Decorator
}

// Initialize ...
func (d *Inverter) Initialize(args ...interface{}) {
	d.Decorator = args[0].(*core.Decorator)
}

// Start ...
func (d *Inverter) Start(ctx *core.Context) {}

// Tick ...
func (d *Inverter) Tick(ctx *core.Context) core.Status {
	fmt.Println("Run Inverter")
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
func (d *Inverter) Stop(ctx *core.Context) {}
