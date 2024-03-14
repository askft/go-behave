package composite

import (
	"math/rand"

	"github.com/jbcpollak/go-behave/core"
)

// RandomSequence works just like Sequence, except it shuffles
// the order of its children every time it is re-updated.
func RandomSequence[Context any](children ...core.Node[Context]) core.Node[Context] {
	base := core.NewComposite("RandomSequence", children)
	return &randomSequence[Context]{Composite: base}
}

type randomSequence[Context any] struct {
	*core.Composite[Context]
}

func (s *randomSequence[Context]) Enter(ctx Context) {
	shuffle(s.Children)
}

func (s *randomSequence[Context]) Tick(ctx Context) core.Status {
	for s.CurrentChild < len(s.Children) {
		status := core.Update(s.Children[s.CurrentChild], ctx)
		if status != core.StatusSuccess {
			return status
		}
		s.Composite.CurrentChild++
	}
	return core.StatusSuccess
}

func (s *randomSequence[Context]) Leave(ctx Context) {
	s.Composite.CurrentChild = 0
}

func shuffle[Context any](nodes []core.Node[Context]) {
	rand.Shuffle(len(nodes), func(i, j int) {
		nodes[i], nodes[j] = nodes[j], nodes[i]
	})
}
