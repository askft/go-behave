package decorator

import (
	"time"

	"github.com/askft/go-behave/core"
)

// Delayer ...
func Delayer(params core.Params, child core.Node) core.Node {
	base := core.NewDecorator("Delayer", params, child)
	d := &delayer{Decorator: base}

	ms, err := params.GetInt("ms")
	if err != nil {
		panic(err)
	}

	d.delay = time.Duration(ms) * time.Millisecond
	return d
}

// delayer ...
type delayer struct {
	*core.Decorator
	delay time.Duration // delay in milliseconds
	start time.Time
}

// Enter ...
func (d *delayer) Enter(ctx *core.Context) {
	d.start = time.Now()
}

// Tick ...
func (d *delayer) Tick(ctx *core.Context) core.Status {
	if time.Since(d.start) > d.delay {
		return core.Update(d.Child, ctx)
	}
	return core.StatusRunning
}

// Leave ...
func (d *delayer) Leave(ctx *core.Context) {}
