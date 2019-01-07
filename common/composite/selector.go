package composite

import (
	"github.com/askft/go-behave/core"
)

// Selector updates each child in order, returning success as soon as
// a child succeeds. If a child returns Running, the selector node
// will resume execution from that child the next tick.
func Selector(children ...core.Node) core.Node {
	base := core.NewComposite("Selector", children)
	return &selector{Composite: base}
}

type selector struct {
	*core.Composite
}

func (s *selector) Enter(ctx *core.Context) {
	s.Composite.CurrentChild = 0
}

func (s *selector) Tick(ctx *core.Context) core.Status {
	for s.CurrentChild < len(s.Children) {
		status := core.Update(s.Children[s.CurrentChild], ctx)
		if status != core.StatusFailure {
			return status
		}
		s.Composite.CurrentChild++
	}
	return core.StatusFailure
}

func (s *selector) Leave(ctx *core.Context) {}
