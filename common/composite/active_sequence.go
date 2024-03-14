package composite

import (
	"github.com/jbcpollak/go-behave/core"
)

// ActiveSequence ticks each child in order. Returns success if
// all children succeed in one tick, else returns the status of
// the non-succeeding node. Restarts iteration the next tick.
func ActiveSequence[Context any](children ...core.Node[Context]) core.Node[Context] {
	base := core.NewComposite("ActiveSequence", children)
	return &activeSequence[Context]{Composite: base}
}

type activeSequence[Context any] struct {
	*core.Composite[Context]
}

func (s *activeSequence[Context]) Enter(ctx Context) {}

func (s *activeSequence[Context]) Tick(ctx Context) core.Status {
	for i := 0; i < len(s.Children); i++ {
		status := core.Update(s.Children[i], ctx)
		if status != core.StatusSuccess {
			return status
		}
	}
	return core.StatusSuccess
}

func (s *activeSequence[Context]) Leave(ctx Context) {}
