package decorator

import (
	"github.com/askft/go-behave/core"
)

// Repeater updates its child n times, at which point the repeater
// returns Success. The repeater runs forever if n == 0.
func Repeater(params core.Params, child core.Node) core.Node {
	base := core.NewDecorator("Repeater", params, child)
	d := &repeater{Decorator: base}

	n, err := params.GetInt("n")
	if err != nil {
		panic(err)
	}

	d.n = n
	return d
}

type repeater struct {
	*core.Decorator
	n int
	i int
}

func (d *repeater) Enter(ctx *core.Context) {
	d.i = 0
}

func (d *repeater) Tick(ctx *core.Context) core.Status {
	_ = core.Update(d.Child, ctx)

	// Run forever if n == 0.
	if d.n == 0 {
		return core.StatusRunning
	}

	d.i++
	if d.i < d.n {
		return core.StatusRunning
	}

	// At this point, the repeater has updated its child n times.
	return core.StatusSuccess
}

func (d *repeater) Leave(ctx *core.Context) {}
