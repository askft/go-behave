package composite

import (
	"math/rand"

	"github.com/jbcpollak/go-behave/core"
)

// RandomSequence works just like Sequence, except it shuffles
// the order of its children every time it is re-updated.
func RandomSequence[Blackboard any](children ...core.Node[Blackboard]) core.Node[Blackboard] {
	base := core.NewComposite("RandomSequence", children)
	return &randomSequence[Blackboard]{Composite: base}
}

type randomSequence[Blackboard any] struct {
	core.Composite[Blackboard]
}

func (s *randomSequence[Blackboard]) Enter(bb Blackboard) {
	shuffle(s.Children)
}

func (s *randomSequence[Blackboard]) Tick(bb Blackboard, evt core.Event) core.NodeResult {
	for s.CurrentChild < len(s.Children) {
		status := core.Update(s.Children[s.CurrentChild], bb, evt)
		if status != core.StatusSuccess {
			return status
		}
		s.Composite.CurrentChild++
	}
	return core.StatusSuccess
}

func (s *randomSequence[Blackboard]) Leave(bb Blackboard) {
	s.Composite.CurrentChild = 0
}

func shuffle[Blackboard any](nodes []core.Node[Blackboard]) {
	rand.Shuffle(len(nodes), func(i, j int) {
		nodes[i], nodes[j] = nodes[j], nodes[i]
	})
}
