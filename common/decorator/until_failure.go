package decorator

import (
	"github.com/jbcpollak/go-behave/core"
)

// UntilFailure updates its child until it returns Failure.
func UntilFailure[Blackboard any, Event any](params core.Params, child core.Node[Blackboard, Event]) core.Node[Blackboard, Event] {
	base := core.NewDecorator[Blackboard, Event]("UntilFailure", params, child)
	return &untilFailure[Blackboard, Event]{Decorator: base}
}

type untilFailure[Blackboard any, Event any] struct {
	*core.Decorator[Blackboard, Event]
}

func (d *untilFailure[Blackboard, Event]) Enter(bb Blackboard) {}

func (d *untilFailure[Blackboard, Event]) Tick(bb Blackboard, evt Event) core.NodeResult {
	status := core.Update(d.Child, bb, evt)
	if status == core.StatusFailure {
		return core.StatusSuccess
	}
	return core.StatusRunning
}

func (d *untilFailure[Blackboard, Event]) Leave(bb Blackboard) {}
