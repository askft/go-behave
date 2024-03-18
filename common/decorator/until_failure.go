package decorator

import (
	"github.com/jbcpollak/greenstalk/core"
)

// UntilFailure updates its child until it returns Failure.
func UntilFailure[Blackboard any](params core.Params, child core.Node[Blackboard]) core.Node[Blackboard] {
	base := core.NewDecorator[Blackboard]("UntilFailure", params, child)
	return &untilFailure[Blackboard]{Decorator: base}
}

type untilFailure[Blackboard any] struct {
	core.Decorator[Blackboard]
}

func (d *untilFailure[Blackboard]) Enter(bb Blackboard) {}

func (d *untilFailure[Blackboard]) Tick(bb Blackboard, evt core.Event) core.NodeResult {
	status := core.Update(d.Child, bb, evt)
	if status == core.StatusFailure {
		return core.StatusSuccess
	}
	return core.StatusRunning
}

func (d *untilFailure[Blackboard]) Leave(bb Blackboard) {}
