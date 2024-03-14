package decorator

import (
	"github.com/jbcpollak/go-behave/core"
)

// Repeater updates its child n times, at which point the repeater
// returns Success. The repeater runs forever if n == 0.
func Repeater[Context any](params core.Params, child core.Node[Context]) core.Node[Context] {
	base := core.NewDecorator("Repeater", params, child)
	d := &repeater[Context]{Decorator: base}

	n, err := params.GetInt("n")
	if err != nil {
		panic(err)
	}

	d.n = n
	return d
}

type repeater[Context any] struct {
	*core.Decorator[Context]
	n int
	i int
}

func (d *repeater[Context]) Enter(ctx Context) {
	d.i = 0
}

func (d *repeater[Context]) Tick(ctx Context) core.Status {
	status := core.Update(d.Child, ctx)

	if status == core.StatusRunning {
		return core.StatusRunning
	}

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

func (d *repeater[Context]) Leave(ctx Context) {}
