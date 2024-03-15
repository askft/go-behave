package decorator

import (
	"time"

	"github.com/jbcpollak/go-behave/core"
)

// Delayer ...
func Delayer[Blackboard any](params core.Params, child core.Node[Blackboard]) core.Node[Blackboard] {
	base := core.NewDecorator("Delayer", params, child)
	d := &delayer[Blackboard]{Decorator: base}

	ms, err := params.GetInt("ms")
	if err != nil {
		panic(err)
	}

	d.delay = time.Duration(ms) * time.Millisecond
	return d
}

// delayer ...
type delayer[Blackboard any] struct {
	core.Decorator[Blackboard]
	delay time.Duration // delay in milliseconds
	start time.Time
}

// Enter ...
func (d *delayer[Blackboard]) Enter(bb Blackboard) {
	d.start = time.Now()
}

// Tick ...
func (d *delayer[Blackboard]) Tick(bb Blackboard, evt core.Event) core.NodeResult {
	if time.Since(d.start) > d.delay {
		return core.Update(d.Child, bb, evt)
	}
	return core.StatusRunning
}

// Leave ...
func (d *delayer[Blackboard]) Leave(bb Blackboard) {}
