package decorator

import (
	"github.com/askft/go-behave/core"
)

// Repeater ...
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

// repeater runs its child either until it returns core.StatusRunning
// or until it has run n times. Runs forever if n == 0.
type repeater struct {
	*core.Decorator
	n int
	i int
}

// Start ...
func (d *repeater) Start(ctx *core.Context) {
	d.i = 0
}

// Tick ...
func (d *repeater) Tick(ctx *core.Context) core.Status {
	status := core.StatusSuccess
	for ; d.n == 0 || d.i < d.n; d.i++ {
		status = core.Update(d.Child, ctx)
		if status == core.StatusRunning {
			break
		}
	}
	return status
}

// Stop ...
func (d *repeater) Stop(ctx *core.Context) {}
