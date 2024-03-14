package composite

import (
	"math/rand"

	"github.com/jbcpollak/go-behave/core"
)

// RandomSequence works just like Sequence, except it shuffles
// the order of its children every time it is re-updated.
func RandomSequence[Blackboard any, Event any](children ...core.Node[Blackboard, Event]) core.Node[Blackboard, Event] {
	base := core.NewComposite("RandomSequence", children)
	return &randomSequence[Blackboard, Event]{Composite: base}
}

type randomSequence[Blackboard any, Event any] struct {
	*core.Composite[Blackboard, Event]
}

func (s *randomSequence[Blackboard, Event]) Enter(bb Blackboard) {
	shuffle(s.Children)
}

func (s *randomSequence[Blackboard, Event]) Tick(bb Blackboard, evt Event) core.NodeResult {
	for s.CurrentChild < len(s.Children) {
		status := core.Update(s.Children[s.CurrentChild], bb, evt)
		if status != core.StatusSuccess {
			return status
		}
		s.Composite.CurrentChild++
	}
	return core.StatusSuccess
}

func (s *randomSequence[Blackboard, Event]) Leave(bb Blackboard) {
	s.Composite.CurrentChild = 0
}

func shuffle[Blackboard any, Event any](nodes []core.Node[Blackboard, Event]) {
	rand.Shuffle(len(nodes), func(i, j int) {
		nodes[i], nodes[j] = nodes[j], nodes[i]
	})
}
