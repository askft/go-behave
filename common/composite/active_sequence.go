package composite

import (
	"github.com/jbcpollak/go-behave/core"
)

// ActiveSequence ticks each child in order. Returns success if
// all children succeed in one tick, else returns the status of
// the non-succeeding node. Restarts iteration the next tick.
func ActiveSequence[Blackboard any, Event any](children ...core.Node[Blackboard, Event]) core.Node[Blackboard, Event] {
	base := core.NewComposite("ActiveSequence", children)
	return &activeSequence[Blackboard, Event]{Composite: base}
}

type activeSequence[Blackboard any, Event any] struct {
	*core.Composite[Blackboard, Event]
}

func (s *activeSequence[Blackboard, Event]) Enter(bb Blackboard) {}

func (s *activeSequence[Blackboard, Event]) Tick(bb Blackboard, evt Event) core.NodeResult {
	for i := 0; i < len(s.Children); i++ {
		status := core.Update(s.Children[i], bb, evt)
		if status != core.StatusSuccess {
			return status
		}
	}
	return core.StatusSuccess
}

func (s *activeSequence[Blackboard, Event]) Leave(bb Blackboard) {}
