package composite

import (
	"github.com/jbcpollak/go-behave/core"
)

// Sequence updates each child in order, returning success only if
// all children succeed. If a child returns Running, the sequence node
// will resume execution from that child the next tick.
func Sequence[Blackboard any, Event any](children ...core.Node[Blackboard, Event]) core.Node[Blackboard, Event] {
	base := core.NewComposite("Sequence", children)
	return &sequence[Blackboard, Event]{Composite: base}
}

type sequence[Blackboard any, Event any] struct {
	*core.Composite[Blackboard, Event]
}

func (s *sequence[Blackboard, Event]) Enter(bb Blackboard) {
	s.Composite.CurrentChild = 0
}

func (s *sequence[Blackboard, Event]) Tick(bb Blackboard, evt Event) core.NodeResult {
	for s.CurrentChild < len(s.Children) {
		status := core.Update(s.Children[s.CurrentChild], bb, evt)
		if status != core.StatusSuccess {
			return status
		}
		s.CurrentChild++
	}
	return core.StatusSuccess
}

func (s *sequence[Blackboard, Event]) Leave(bb Blackboard) {}
