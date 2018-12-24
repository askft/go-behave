package composite

import (
	"github.com/alexanderskafte/behaviortree/core"
)

// selector ...
type selector struct {
	*core.Composite
}

// Selector creates a new selector node.
func Selector(children ...core.INode) core.INode {
	base := core.NewComposite("Selector")
	base.Children = children
	return &selector{Composite: base}
}

// Start ...
func (s *selector) Start(ctx *core.Context) {
	s.Composite.CurrentChild = 0
}

// Tick ...
func (s *selector) Tick(ctx *core.Context) core.Status {
	for {
		status := core.Update(s.Children[s.CurrentChild], ctx)
		if status != core.StatusFailure {
			return status
		}
		s.Composite.CurrentChild++
		if s.CurrentChild >= len(s.Children) {
			return core.StatusFailure
		}
	}
}

// Stop ...
func (s *selector) Stop(ctx *core.Context) {}
