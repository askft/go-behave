package action

import (
	"github.com/jbcpollak/go-behave/core"
)

// Fail returns a new fail node, which always fails in one tick.
func Fail[Blackboard any, Event any](params core.Params, returns core.Returns) core.Node[Blackboard, Event] {
	base := core.NewLeaf[Blackboard, Event]("Fail", params, returns)
	return &fail[Blackboard, Event]{Leaf: base}
}

// fail ...
type fail[Blackboard any, Event any] struct {
	*core.Leaf[Blackboard, Event]
}

// Enter ...
func (a *fail[Blackboard, Event]) Enter(bb Blackboard) {}

// Tick ...
func (a *fail[Blackboard, Event]) Tick(bb Blackboard, evt Event) core.NodeResult {
	return core.StatusFailure
}

// Leave ...
func (a *fail[Blackboard, Event]) Leave(bb Blackboard) {}
