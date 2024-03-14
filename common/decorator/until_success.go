package decorator

import (
	"github.com/jbcpollak/go-behave/core"
)

// UntilSuccess updates its child until it returns Success.
func UntilSuccess[Blackboard any, Event any](params core.Params, child core.Node[Blackboard, Event]) core.Node[Blackboard, Event] {
	base := core.NewDecorator("UntilSuccess", params, child)
	return &untilSuccess[Blackboard, Event]{Decorator: base}
}

type untilSuccess[Blackboard any, Event any] struct {
	*core.Decorator[Blackboard, Event]
}

func (d *untilSuccess[Blackboard, Event]) Enter(bb Blackboard) {}

func (d *untilSuccess[Blackboard, Event]) Tick(bb Blackboard, evt Event) core.NodeResult {
	status := core.Update(d.Child, bb, evt)
	if status == core.StatusSuccess {
		return core.StatusSuccess
	}
	return core.StatusRunning
}

func (d *untilSuccess[Blackboard, Event]) Leave(bb Blackboard) {}
