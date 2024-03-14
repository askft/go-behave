package composite

import (
	"github.com/jbcpollak/go-behave/core"
)

// Selector updates each child in order, returning success as soon as
// a child succeeds. If a child returns Running, the selector node
// will resume execution from that child the next tick.
func Selector[Blackboard any, Event any](children ...core.Node[Blackboard, Event]) core.Node[Blackboard, Event] {
	base := core.NewComposite("Selector", children)
	return &selector[Blackboard, Event]{Composite: base}
}

type selector[Blackboard any, Event any] struct {
	*core.Composite[Blackboard, Event]
}

func (s *selector[Blackboard, Event]) Enter(bb Blackboard) {
	s.Composite.CurrentChild = 0
}

func (s *selector[Blackboard, Event]) Tick(bb Blackboard, evt Event) core.NodeResult {
	for s.CurrentChild < len(s.Children) {
		status := core.Update(s.Children[s.CurrentChild], bb, evt)
		if status != core.StatusFailure {
			return status
		}
		s.Composite.CurrentChild++
	}
	return core.StatusFailure
}

func (s *selector[Blackboard, Event]) Leave(bb Blackboard) {}
