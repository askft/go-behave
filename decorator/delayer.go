package decorator

import (
	"fmt"
	"strconv"
	"time"

	"github.com/alexanderskafte/go-behave/core"
)

// delayer ...
type delayer struct {
	*core.Decorator
	delay time.Duration // delay in milliseconds
	start time.Time
}

// Delayer ...
func Delayer(params core.Params, child core.Node) core.Node {
	base := core.NewDecorator("Delayer", params, child)
	d := &delayer{Decorator: base}

	str, ok := params["ms"]
	if !ok {
		panic(fmt.Errorf("ms not found in params to %s", d.String()))
	}
	delay, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	d.delay = time.Duration(delay) * time.Millisecond
	return d
}

// Start ...
func (d *delayer) Start(ctx *core.Context) {
	d.start = time.Now()
}

// Tick ...
func (d *delayer) Tick(ctx *core.Context) core.Status {
	if time.Since(d.start) > d.delay {
		return core.Update(d.Child, ctx)
	}
	return core.StatusRunning
}

// Stop ...
func (d *delayer) Stop(ctx *core.Context) {}
