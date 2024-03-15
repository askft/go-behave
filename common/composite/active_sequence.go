package composite

import (
	"github.com/jbcpollak/go-behave/core"
)

// ActiveSequence ticks each child in order. Returns success if
// all children succeed in one tick, else returns the status of
// the non-succeeding node. Restarts iteration the next tick.
func ActiveSequence[Blackboard any](children ...core.Node[Blackboard]) core.Node[Blackboard] {
	base := core.NewComposite("ActiveSequence", children)
	return &activeSequence[Blackboard]{Composite: base}
}

type activeSequence[Blackboard any] struct {
	core.Composite[Blackboard]
}

func (s *activeSequence[Blackboard]) Enter(bb Blackboard) {}

func (s *activeSequence[Blackboard]) Tick(bb Blackboard, evt core.Event) core.NodeResult {
	for i := 0; i < len(s.Children); i++ {
		result := core.Update(s.Children[i], bb, evt)
		if result.Status() != core.StatusSuccess {
			return result
		}
	}
	return core.StatusSuccess
}

func (s *activeSequence[Blackboard]) Leave(bb Blackboard) {}
