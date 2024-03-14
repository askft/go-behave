package action

import (
	"github.com/jbcpollak/go-behave/core"
)

// Succeed returns a new succeed node, which always succeeds in one tick.
func Succeed[Blackboard any, Event any](params core.Params, returns core.Returns) core.Node[Blackboard, Event] {
	base := core.NewLeaf[Blackboard, Event]("Succeed", params, returns)
	return &succeed[Blackboard, Event]{Leaf: base}
}

// succeed ...
type succeed[Blackboard any, Event any] struct {
	*core.Leaf[Blackboard, Event]
}

// Enter ...
func (a *succeed[Blackboard, Event]) Enter(bb Blackboard) {}

// Tick ...
func (a *succeed[Blackboard, Event]) Tick(bb Blackboard, evt Event) core.NodeResult {
	return core.StatusSuccess
}

// Leave ...
func (a *succeed[Blackboard, Event]) Leave(bb Blackboard) {}
