package composite

import (
	"math/rand"

	"github.com/askft/go-behave/core"
)

// RandomSequence works just like Sequence, except it shuffles
// the order of its children every time it is re-updated.
func RandomSequence(children ...core.Node) core.Node {
	base := core.NewComposite("RandomSequence", children)
	return &randomSequence{Composite: base}
}

type randomSequence struct {
	*core.Composite
}

func (s *randomSequence) Enter(ctx *core.Context) {
	shuffle(s.Children)
}

func (s *randomSequence) Tick(ctx *core.Context) core.Status {
	for s.CurrentChild < len(s.Children) {
		status := core.Update(s.Children[s.CurrentChild], ctx)
		if status != core.StatusSuccess {
			return status
		}
		s.Composite.CurrentChild++
	}
	return core.StatusSuccess
}

func (s *randomSequence) Leave(ctx *core.Context) {
	s.Composite.CurrentChild = 0
}

func shuffle(nodes []core.Node) {
	rand.Shuffle(len(nodes), func(i, j int) {
		nodes[i], nodes[j] = nodes[j], nodes[i]
	})
}
