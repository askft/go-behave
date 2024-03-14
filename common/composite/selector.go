package composite

import (
	"github.com/jbcpollak/go-behave/core"
)

// Selector updates each child in order, returning success as soon as
// a child succeeds. If a child returns Running, the selector node
// will resume execution from that child the next tick.
func Selector[Context any](children ...core.Node[Context]) core.Node[Context] {
	base := core.NewComposite("Selector", children)
	return &selector[Context]{Composite: base}
}

type selector[Context any] struct {
	*core.Composite[Context]
}

func (s *selector[Context]) Enter(ctx Context) {
	s.Composite.CurrentChild = 0
}

func (s *selector[Context]) Tick(ctx Context) core.Status {
	for s.CurrentChild < len(s.Children) {
		status := core.Update(s.Children[s.CurrentChild], ctx)
		if status != core.StatusFailure {
			return status
		}
		s.Composite.CurrentChild++
	}
	return core.StatusFailure
}

func (s *selector[Context]) Leave(ctx Context) {}
