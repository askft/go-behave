package decorator

import (
	"github.com/jbcpollak/go-behave/core"
)

// Inverter ...
func Inverter[Blackboard any, Event any](params core.Params, child core.Node[Blackboard, Event]) core.Node[Blackboard, Event] {
	base := core.NewDecorator("Inverter", params, child)
	return &inverter[Blackboard, Event]{Decorator: base}
}

// inverter ...
type inverter[Blackboard any, Event any] struct {
	*core.Decorator[Blackboard, Event]
}

// Enter ...
func (d *inverter[Blackboard, Event]) Enter(bb Blackboard) {}

// Tick ...
func (d *inverter[Blackboard, Event]) Tick(bb Blackboard, evt Event) core.NodeResult {
	switch core.Update(d.Child, bb, evt) {
	case core.StatusSuccess:
		return core.StatusFailure
	case core.StatusFailure:
		return core.StatusSuccess
	default:
		return core.StatusRunning
	}
}

// Leave ...
func (d *inverter[Blackboard, Event]) Leave(bb Blackboard) {}
