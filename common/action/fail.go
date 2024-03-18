package action

import (
	"context"

	"github.com/jbcpollak/greenstalk/core"
)

// Fail returns a new fail node, which always fails in one tick.
func Fail[Blackboard any](params core.Params, returns core.Returns) core.Node[Blackboard] {
	base := core.NewLeaf[Blackboard]("Fail", params, returns)
	return &fail[Blackboard]{Leaf: base}
}

// fail ...
type fail[Blackboard any] struct {
	core.Leaf[Blackboard]
}

// Enter ...
func (a *fail[Blackboard]) Enter(bb Blackboard) {}

// Tick ...
func (a *fail[Blackboard]) Tick(bb Blackboard, ctx context.Context, evt core.Event) core.NodeResult {
	return core.StatusFailure
}

// Leave ...
func (a *fail[Blackboard]) Leave(bb Blackboard) {}
