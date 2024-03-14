package composite

import (
	"github.com/jbcpollak/go-behave/core"
)

// Sequence updates each child in order, returning success only if
// all children succeed. If a child returns Running, the sequence node
// will resume execution from that child the next tick.
func Sequence[Context any](children ...core.Node[Context]) core.Node[Context] {
	base := core.NewComposite("Sequence", children)
	return &sequence[Context]{Composite: base}
}

type sequence[Context any] struct {
	*core.Composite[Context]
}

func (s *sequence[Context]) Enter(ctx Context) {
	s.Composite.CurrentChild = 0
}

func (s *sequence[Context]) Tick(ctx Context) core.Status {
	for s.CurrentChild < len(s.Children) {
		status := core.Update(s.Children[s.CurrentChild], ctx)
		if status != core.StatusSuccess {
			return status
		}
		s.CurrentChild++
	}
	return core.StatusSuccess
}

func (s *sequence[Context]) Leave(ctx Context) {}
