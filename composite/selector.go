package composite

import (
	"fmt"
	"github.com/alexanderskafte/behaviortree/core"
)

// Selector ...
type Selector struct {
	*core.Composite
}

// Initialize ...
func (s *Selector) Initialize(args ...interface{}) {
	s.Composite = args[0].(*core.Composite)
}

// Start ...
func (s *Selector) Start(ctx *core.Context) {
	s.Composite.CurrentChild = 0
}

// Tick ...
func (s *Selector) Tick(ctx *core.Context) core.Status {
	fmt.Println("Run Sequence")
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
func (s *Selector) Stop(ctx *core.Context) {}
