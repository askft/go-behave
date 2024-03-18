package action

import (
	"context"

	"github.com/jbcpollak/greenstalk/core"
)

// Succeed returns a new succeed node, which always succeeds in one tick.
func Succeed[Blackboard any](params core.Params, returns core.Returns) core.Node[Blackboard] {
	base := core.NewLeaf[Blackboard]("Succeed", params, returns)
	return &succeed[Blackboard]{Leaf: base}
}

// succeed ...
type succeed[Blackboard any] struct {
	core.Leaf[Blackboard]
}

// Enter ...
func (a *succeed[Blackboard]) Enter(bb Blackboard) {}

// Tick ...
func (a *succeed[Blackboard]) Tick(bb Blackboard, ctx context.Context, evt core.Event) core.NodeResult {
	return core.StatusSuccess
}

// Leave ...
func (a *succeed[Blackboard]) Leave(bb Blackboard) {}
