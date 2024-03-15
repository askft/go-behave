package composite

import (
	"github.com/jbcpollak/go-behave/core"
)

// PersistentSequence updates each child in order. If a child
// returns Failure or Running, this node returns the same value,
// and resumes execution from the same child node the next tick.
func PersistentSequence[Blackboard any](children ...core.Node[Blackboard]) core.Node[Blackboard] {
	base := core.NewComposite("PersistentSequence", children)
	return &persistentSequence[Blackboard]{Composite: base}
}

type persistentSequence[Blackboard any] struct {
	core.Composite[Blackboard]
}

func (s *persistentSequence[Blackboard]) Enter(bb Blackboard) {}

func (s *persistentSequence[Blackboard]) Tick(bb Blackboard, evt core.Event) core.NodeResult {
	for s.CurrentChild < len(s.Children) {
		result := core.Update(s.Children[s.CurrentChild], bb, evt)
		if result.Status() != core.StatusSuccess {
			return result
		}
		s.CurrentChild++
	}
	return core.StatusSuccess
}

func (s *persistentSequence[Blackboard]) Leave(bb Blackboard) {
	s.Composite.CurrentChild = 0
}
