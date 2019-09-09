package composite

import (
	"github.com/askft/go-behave/core"
)

// ActiveSequence ticks each child in order. Returns success if
// all children succeed in one tick, else returns the status of
// the non-succeeding node. Restarts iteration the next tick.
func ActiveSequence(children ...core.Node) core.Node {
	base := core.NewComposite("ActiveSequence", children)
	return &activeSequence{Composite: base}
}

type activeSequence struct {
	*core.Composite
}

func (s *activeSequence) Enter(ctx *core.Context) {}

func (s *activeSequence) Tick(ctx *core.Context) core.Status {
	for i := 0; i < len(s.Children); i++ {
		status := core.Update(s.Children[i], ctx)
		if status != core.StatusSuccess {
			return status
		}
	}
	return core.StatusSuccess
}

func (s *activeSequence) Leave(ctx *core.Context) {}
