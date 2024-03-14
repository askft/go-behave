package decorator

import (
	"github.com/jbcpollak/go-behave/core"
)

// Repeater updates its child n times, at which point the repeater
// returns Success. The repeater runs forever if n == 0.
func Repeater[Blackboard any, Event any](params core.Params, child core.Node[Blackboard, Event]) core.Node[Blackboard, Event] {
	base := core.NewDecorator("Repeater", params, child)
	d := &repeater[Blackboard, Event]{Decorator: base}

	n, err := params.GetInt("n")
	if err != nil {
		panic(err)
	}

	d.n = n
	return d
}

type repeater[Blackboard any, Event any] struct {
	*core.Decorator[Blackboard, Event]
	n int
	i int
}

func (d *repeater[Blackboard, Event]) Enter(bb Blackboard) {
	d.i = 0
}

func (d *repeater[Blackboard, Event]) Tick(bb Blackboard, evt Event) core.NodeResult {
	status := core.Update(d.Child, bb, evt)

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

func (d *repeater[Blackboard, Event]) Leave(bb Blackboard) {}
