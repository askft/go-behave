package decorator

import (
	"time"

	"github.com/jbcpollak/go-behave/core"
)

// Delayer ...
func Delayer[Blackboard any, Event any](params core.Params, child core.Node[Blackboard, Event]) core.Node[Blackboard, Event] {
	base := core.NewDecorator("Delayer", params, child)
	d := &delayer[Blackboard, Event]{Decorator: base}

	ms, err := params.GetInt("ms")
	if err != nil {
		panic(err)
	}

	d.delay = time.Duration(ms) * time.Millisecond
	return d
}

// delayer ...
type delayer[Blackboard any, Event any] struct {
	*core.Decorator[Blackboard, Event]
	delay time.Duration // delay in milliseconds
	start time.Time
}

// Enter ...
func (d *delayer[Blackboard, Event]) Enter(bb Blackboard) {
	d.start = time.Now()
}

// Tick ...
func (d *delayer[Blackboard, Event]) Tick(bb Blackboard, evt Event) core.NodeResult {
	if time.Since(d.start) > d.delay {
		return core.Update(d.Child, bb, evt)
	}
	return core.StatusRunning
}

// Leave ...
func (d *delayer[Blackboard, Event]) Leave(bb Blackboard) {}
