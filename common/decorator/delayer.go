package decorator

import (
	"context"
	"time"

	"github.com/jbcpollak/greenstalk/core"
)

// Delayer ...
func Delayer[Blackboard any](params core.Params, child core.Node[Blackboard]) core.Node[Blackboard] {
	base := core.NewDecorator("Delayer", params, child)
	d := &delayer[Blackboard]{Decorator: base}

	delay, err := params.GetInt("delay")
	if err != nil {
		panic(err)
	}

	d.delay = time.Duration(delay)
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
func (d *delayer[Blackboard]) Tick(bb Blackboard, ctx context.Context, evt core.Event) core.NodeResult {
	if time.Since(d.start) > d.delay {
		return core.Update(d.Child, bb, ctx, evt)
	}
	return core.StatusRunning
}

// Leave ...
func (d *delayer[Blackboard]) Leave(bb Blackboard) {}
