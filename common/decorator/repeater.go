package decorator

import (
	"context"

	"github.com/jbcpollak/greenstalk/core"
	"github.com/rs/zerolog/log"
)

// Repeater updates its child n times, at which point the repeater
// returns Success. The repeater runs forever if n == 0.
func Repeater[Blackboard any](params core.Params, child core.Node[Blackboard]) core.Node[Blackboard] {
	base := core.NewDecorator("Repeater", params, child)
	d := &repeater[Blackboard]{Decorator: base}

	n, err := params.GetInt("n")
	if err != nil {
		panic(err)
	}

	d.n = n
	return d
}

type repeater[Blackboard any] struct {
	core.Decorator[Blackboard]
	n int
	i int
}

func (d *repeater[Blackboard]) Enter(bb Blackboard) {
	d.i = 0
}

func (d *repeater[Blackboard]) Tick(bb Blackboard, ctx context.Context, evt core.Event) core.NodeResult {
	log.Info().Msg("Repeater: Calling child")
	status := core.Update(d.Child, bb, ctx, evt)

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

func (d *repeater[Blackboard]) Leave(bb Blackboard) {}
