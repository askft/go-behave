package composite

import (
	"github.com/alexanderskafte/go-behave/core"
)

// sequence ...
type sequence struct {
	*core.Composite
}

// Sequence creates a new sequence node.
func Sequence(children ...core.Node) core.Node {
	base := core.NewComposite("Sequence", children)
	return &sequence{Composite: base}
}

// Start ...
func (s *sequence) Start(ctx *core.Context) {
	s.Composite.CurrentChild = 0
}

// Tick ...
func (s *sequence) Tick(ctx *core.Context) core.Status {
	for {
		status := core.Update(s.Children[s.CurrentChild], ctx)
		if status != core.StatusSuccess {
			return status
		}
		s.CurrentChild++
		if s.CurrentChild >= len(s.Children) {
			return core.StatusSuccess
		}
	}
}

// Stop ...
func (s *sequence) Stop(ctx *core.Context) {}
