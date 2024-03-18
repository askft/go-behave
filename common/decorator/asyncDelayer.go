package decorator

import (
	"context"
	"fmt"
	"time"

	"github.com/jbcpollak/greenstalk/core"
)

// AsyncDelayer ...
func AsyncDelayer[Blackboard any](params core.Params, child core.Node[Blackboard]) core.Node[Blackboard] {
	base := core.NewDecorator("Delayer", params, child)
	d := &asyncdelayer[Blackboard]{Decorator: base}

	ms, err := params.GetInt("ms")
	if err != nil {
		panic(err)
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

	fmt.Printf("AsyncDelayer Entered")
}

type DelayerFinishedEvent struct {
	start time.Time
}

func (d *asyncdelayer[Blackboard]) doDelay(enqueue func(core.Event) error) error {
	time.Sleep(d.delay)

	return enqueue(DelayerFinishedEvent{d.start})
}

// Tick ...
func (d *asyncdelayer[Blackboard]) Tick(bb Blackboard, ctx context.Context, evt core.Event) core.NodeResult {

	if _, ok := evt.(DelayerFinishedEvent); ok {
		fmt.Printf("asyncdelayer: Calling child")
		return core.Update(d.Child, bb, ctx, evt)
	} else if d.Status() == core.StatusInitialized {
		fmt.Printf("asyncdelayer: Returning AsyncRunning")

		return core.NodeAsyncRunning(d.doDelay)
	} else {
		return core.StatusFailure
	}
}

// Leave ...
func (d *asyncdelayer[Blackboard]) Leave(bb Blackboard) {}
