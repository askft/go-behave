package composite

import (
	"github.com/jbcpollak/go-behave/core"
)

// PersistentSequence updates each child in order. If a child
// returns Failure or Running, this node returns the same value,
// and resumes execution from the same child node the next tick.
func PersistentSequence[Blackboard any, Event any](children ...core.Node[Blackboard, Event]) core.Node[Blackboard, Event] {
	base := core.NewComposite("PersistentSequence", children)
	return &persistentSequence[Blackboard, Event]{Composite: base}
}

type persistentSequence[Blackboard any, Event any] struct {
	*core.Composite[Blackboard, Event]
}

func (s *persistentSequence[Blackboard, Event]) Enter(bb Blackboard) {}

func (s *persistentSequence[Blackboard, Event]) Tick(bb Blackboard, evt Event) core.NodeResult {
	for s.CurrentChild < len(s.Children) {
		status := core.Update(s.Children[s.CurrentChild], bb, evt)
		if status != core.StatusSuccess {
			return status
		}
		s.CurrentChild++
	}
	return core.StatusSuccess
}

func (s *persistentSequence[Blackboard, Event]) Leave(bb Blackboard) {
	s.Composite.CurrentChild = 0
}
