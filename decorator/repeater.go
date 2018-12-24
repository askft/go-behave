package decorator

import (
	"strconv"

	"github.com/alexanderskafte/behaviortree/core"
)

// repeater runs its child either until it returns core.StatusRunning
// or until it has run n times. Runs forever if n == 0.
type repeater struct {
	*core.Decorator
	n int
	i int
}

// Repeater ...
func Repeater(params core.Params, child core.INode) core.INode {
	base := core.NewDecorator("Repeater", params)
	base.Child = child
	d := &repeater{Decorator: base}

	n, err := strconv.Atoi(d.Params["n"])
	if err != nil {
		panic(err) // TODO
	}
	d.n = n
	return d
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
