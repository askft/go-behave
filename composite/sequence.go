package composite

import (
	"fmt"
	"github.com/alexanderskafte/behaviortree/core"
)

// Sequence ...
type Sequence struct {
	*core.Composite
}

// Initialize ...
func (s *Sequence) Initialize(args ...interface{}) {
	s.Composite = args[0].(*core.Composite)
}

// Start ...
func (s *Sequence) Start(ctx *core.Context) {
	s.Composite.CurrentChild = 0
}

// Tick ...
func (s *Sequence) Tick(ctx *core.Context) core.Status {
	fmt.Println("Run Sequence")
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
func (s *Sequence) Stop(ctx *core.Context) {}
