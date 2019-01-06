package composite

import (
	"math/rand"

	"github.com/askft/go-behave/core"
)

// RandomSequence creates a new random sequence node.
func RandomSequence(children ...core.Node) core.Node {
	base := core.NewComposite("RandomSequence", children)
	return &randomSequence{Composite: base}
}

// Enter ...
func (s *randomSequence) Enter(ctx *core.Context) {
	shuffle(s.Children)
}

// randomSequence ...
type randomSequence struct {
	*core.Composite
}

// Tick ...
func (s *randomSequence) Tick(ctx *core.Context) core.Status {
	for {
		status := core.Update(s.Children[s.CurrentChild], ctx)
		if status != core.StatusSuccess {
			return status
		}
		s.Composite.CurrentChild++
		if s.CurrentChild >= len(s.Children) {
			return core.StatusSuccess
		}
	}
}

// Leave ...
func (s *randomSequence) Leave(ctx *core.Context) {
	s.Composite.CurrentChild = 0
}

func shuffle(nodes []core.Node) {
	rand.Shuffle(len(nodes), func(i, j int) {
		nodes[i], nodes[j] = nodes[j], nodes[i]
	})
}
