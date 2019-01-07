package composite

import (
	"github.com/askft/go-behave/core"
)

// Sequence updates each child in order, returning success only if
// all children succeed. If a child returns Running, the sequence node
// will resume execution from that child the next tick.
func Sequence(children ...core.Node) core.Node {
	base := core.NewComposite("Sequence", children)
	return &sequence{Composite: base}
}

type sequence struct {
	*core.Composite
}

func (s *sequence) Enter(ctx *core.Context) {
	s.Composite.CurrentChild = 0
}

func (s *sequence) Tick(ctx *core.Context) core.Status {
	for s.CurrentChild < len(s.Children) {
		status := core.Update(s.Children[s.CurrentChild], ctx)
		if status != core.StatusSuccess {
			return status
		}
		s.CurrentChild++
	}
	return core.StatusSuccess
}

func (s *sequence) Leave(ctx *core.Context) {}
