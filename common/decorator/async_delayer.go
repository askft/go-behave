package decorator

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jbcpollak/greenstalk/core"
	"github.com/rs/zerolog/log"
)

// AsyncDelayer ...
func AsyncDelayer[Blackboard any](params core.Params, child core.Node[Blackboard]) core.Node[Blackboard] {
	v, err := params.Get("delay")
	if err != nil {
		panic(err)
	}
	delay, ok := v.(time.Duration)
	if !ok {
		panic(fmt.Errorf("delay must be a time.Duration"))
	}

	label, err := params.GetString("label")
	if err != nil {
		panic(err)
	}

	base := core.NewDecorator("AsyncDelayer "+label, params, child)

	d := &asyncdelayer[Blackboard]{
		Decorator: base,
	}

	d.delay = time.Duration(delay)
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

	log.Info().Msgf("%s Entered", d.BaseNode.Name())
}

type DelayerFinishedEvent struct {
	targetNodeId uuid.UUID
	start        time.Time
}

func (e DelayerFinishedEvent) TargetNodeId() uuid.UUID {
	return e.targetNodeId
}

func (d *asyncdelayer[Blackboard]) doDelay(ctx context.Context, enqueue core.EnqueueFn) error {
	t := time.NewTimer(d.delay)
	select {
	case <-ctx.Done():
		t.Stop()
		return fmt.Errorf("Interrupted")
	case <-t.C:
		log.Info().Msgf("Delayed: %v", time.Since(d.start))
		return enqueue(DelayerFinishedEvent{d.Id(), d.start})
	}

}

// Tick ...
func (d *asyncdelayer[Blackboard]) Tick(bb Blackboard, ctx context.Context, evt core.Event) core.NodeResult {
	log.Info().Msgf("%s: Tick", d.Name())

	if dfe, ok := evt.(DelayerFinishedEvent); ok {
		if dfe.TargetNodeId() == d.Id() {
			log.Info().Msgf("%s: DelayerFinishedEvent", d.Name())
			return core.Update(d.Child, bb, ctx, evt)
		}
	}

	if d.Status() == core.StatusInitialized {
		log.Info().Msgf("%s: Returning AsyncRunning", d.Name())

		return core.NodeAsyncRunning(d.doDelay)
	} else {
		return core.StatusFailure
	}
}

// Leave ...
func (d *asyncdelayer[Blackboard]) Leave(bb Blackboard) {}
