package composite

import (
	"context"
	"math/rand"

	"github.com/jbcpollak/greenstalk/core"
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

func (s *randomSequence[Blackboard]) Tick(bb Blackboard, ctx context.Context, evt core.Event) core.NodeResult {
	for s.CurrentChild < len(s.Children) {
		result := core.Update(s.Children[s.CurrentChild], bb, ctx, evt)
		if result.Status() != core.StatusSuccess {
			return result
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
