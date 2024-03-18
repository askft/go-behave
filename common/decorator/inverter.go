package decorator

import (
	"github.com/jbcpollak/greenstalk/core"
)

// Inverter ...
func Inverter[Blackboard any](params core.Params, child core.Node[Blackboard]) core.Node[Blackboard] {
	base := core.NewDecorator("Inverter", params, child)
	return &inverter[Blackboard]{Decorator: base}
}

// inverter ...
type inverter[Blackboard any] struct {
	core.Decorator[Blackboard]
}

// Enter ...
func (d *inverter[Blackboard]) Enter(bb Blackboard) {}

// Tick ...
func (d *inverter[Blackboard]) Tick(bb Blackboard, evt core.Event) core.NodeResult {
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
func (d *inverter[Blackboard]) Leave(bb Blackboard) {}
