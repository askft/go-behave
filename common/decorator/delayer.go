package decorator

import (
	"time"

	"github.com/jbcpollak/go-behave/core"
)

// Delayer ...
func Delayer[Context any](params core.Params, child core.Node[Context]) core.Node[Context] {
	base := core.NewDecorator("Delayer", params, child)
	d := &delayer[Context]{Decorator: base}

	ms, err := params.GetInt("ms")
	if err != nil {
		panic(err)
	}

	d.delay = time.Duration(ms) * time.Millisecond
	return d
}

// delayer ...
type delayer[Context any] struct {
	*core.Decorator[Context]
	delay time.Duration // delay in milliseconds
	start time.Time
}

// Enter ...
func (d *delayer[Context]) Enter(ctx Context) {
	d.start = time.Now()
}

// Tick ...
func (d *delayer[Context]) Tick(ctx Context) core.Status {
	if time.Since(d.start) > d.delay {
		return core.Update(d.Child, ctx)
	}
	return core.StatusRunning
}

// Leave ...
func (d *delayer[Context]) Leave(ctx Context) {}
