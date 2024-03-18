package decorator

import (
	"context"
	"fmt"
	"time"

	"github.com/jbcpollak/greenstalk/core"
)

// AsyncDelayer ...
func AsyncDelayer[Blackboard any](params core.Params, child core.Node[Blackboard]) core.Node[Blackboard] {
	ms, err := params.GetInt("ms")
	if err != nil {
		panic(err)
	}

	label, err := params.GetString("label")
	if err != nil {
		panic(err)
	}

	base := core.NewDecorator("AsyncDelayer "+label, params, child)

	d := &asyncdelayer[Blackboard]{
		Decorator: base,
	}

	d.delay = time.Duration(ms) * time.Millisecond
	return d
}

// delayer ...
type asyncdelayer[Blackboard any] struct {
	core.Decorator[Blackboard]
	delay time.Duration // delay in milliseconds
	start time.Time
}

// Enter ...
func (d *asyncdelayer[Blackboard]) Enter(bb Blackboard) {
	d.start = time.Now()
	d.SetStatus(core.StatusInitialized)

	fmt.Printf("%s Entered\n", d.BaseNode.Name())
}

type DelayerFinishedEvent struct {
	start time.Time
}

func (d *asyncdelayer[Blackboard]) doDelay(ctx context.Context, enqueue core.EnqueueFn) error {
	t := time.NewTimer(d.delay)
	select {
	case <-ctx.Done():
		t.Stop()
		return fmt.Errorf("Interrupted")
	case <-t.C:
	}

	return enqueue(DelayerFinishedEvent{d.start})
}

// Tick ...
func (d *asyncdelayer[Blackboard]) Tick(bb Blackboard, ctx context.Context, evt core.Event) core.NodeResult {
	fmt.Printf("%s: Tick\n", d.BaseNode.Name())

	if _, ok := evt.(DelayerFinishedEvent); ok {
		fmt.Printf("%s: Calling child\n", d.BaseNode.Name())
		return core.Update(d.Child, bb, ctx, evt)
	} else if d.Status() == core.StatusInitialized {
		fmt.Printf("%s: Returning AsyncRunning\n", d.BaseNode.Name())

		return core.NodeAsyncRunning(
			func(enqueue core.EnqueueFn) error {
				return d.doDelay(ctx, enqueue)
			},
		)
	} else {
		return core.StatusFailure
	}
}

// Leave ...
func (d *asyncdelayer[Blackboard]) Leave(bb Blackboard) {}
