package decorator

import (
	"time"

	"github.com/jbcpollak/go-behave/core"
)

// AsyncDelayer ...
func AsyncDelayer[Blackboard any, Event any](params core.Params, child core.Node[Blackboard, Event]) core.Node[Blackboard, Event] {
	base := core.NewDecorator("Delayer", params, child)
	d := &asyncdelayer[Blackboard, Event]{Decorator: base}

	ms, err := params.GetInt("ms")
	if err != nil {
		panic(err)
	}

	d.delay = time.Duration(ms) * time.Millisecond
	return d
}

type DelayerFinishedEvent struct {
	start time.Time
}

// delayer ...
type asyncdelayer[Blackboard any, Event any] struct {
	*core.Decorator[Blackboard, Event]
	delay time.Duration // delay in milliseconds
	start time.Time
}

// Enter ...
func (d *asyncdelayer[Blackboard, Event]) Enter(bb Blackboard) {
	d.start = time.Now()
	d.SetStatus(core.StatusInitialized)
}

func (d *asyncdelayer[Blackboard, Event]) doDelay(enqueue func(Event) error) error {
	time.Sleep(d.delay)
	enqueue(DelayerFinishedEvent{d.start})
	return nil
}

// Tick ...
func (d *asyncdelayer[Blackboard, Event]) Tick(bb Blackboard, evt Event) core.NodeResult {

	if _, ok := evt.(DelayerFinishedEvent); ok {
		return core.StatusSuccess
	} else if d.GetStatus() == core.StatusInitialized {
		return core.NodeAsyncRunning(d.doDelay)
	} else {
		return core.StatusFailure
	}
}

// Leave ...
func (d *asyncdelayer[Blackboard, Event]) Leave(bb Blackboard) {}
