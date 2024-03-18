package composite

import (
	"context"

	"github.com/jbcpollak/greenstalk/core"
)

// Selector updates each child in order, returning success as soon as
// a child succeeds. If a child returns Running, the selector node
// will resume execution from that child the next tick.
func Selector[Blackboard any](children ...core.Node[Blackboard]) core.Node[Blackboard] {
	base := core.NewComposite("Selector", children)
	return &selector[Blackboard]{Composite: base}
}

type selector[Blackboard any] struct {
	core.Composite[Blackboard]
}

func (s *selector[Blackboard]) Enter(bb Blackboard) {
	s.Composite.CurrentChild = 0
}

func (s *selector[Blackboard]) Tick(bb Blackboard, ctx context.Context, evt core.Event) core.NodeResult {
	for s.CurrentChild < len(s.Children) {
		status := core.Update(s.Children[s.CurrentChild], bb, ctx, evt)
		if status != core.StatusFailure {
			return status
		}
		s.Composite.CurrentChild++
	}
	return core.StatusFailure
}

func (s *selector[Blackboard]) Leave(bb Blackboard) {}
