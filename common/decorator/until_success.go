package decorator

import (
	"github.com/jbcpollak/go-behave/core"
)

// UntilSuccess updates its child until it returns Success.
func UntilSuccess[Blackboard any](params core.Params, child core.Node[Blackboard]) core.Node[Blackboard] {
	base := core.NewDecorator("UntilSuccess", params, child)
	return &untilSuccess[Blackboard]{Decorator: base}
}

type untilSuccess[Blackboard any] struct {
	*core.Decorator[Blackboard]
}

func (d *untilSuccess[Blackboard]) Enter(bb Blackboard) {}

func (d *untilSuccess[Blackboard]) Tick(bb Blackboard, evt core.Event) core.NodeResult {
	status := core.Update(d.Child, bb, evt)
	if status == core.StatusSuccess {
		return core.StatusSuccess
	}
	return core.StatusRunning
}

func (d *untilSuccess[Blackboard]) Leave(bb Blackboard) {}
