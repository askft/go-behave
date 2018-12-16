package composite

import (
	. "github.com/alexanderskafte/behaviortree/core"
)

// Sequence ...
type Sequence struct{ Composite }

// Init ...
func (s *Sequence) Init() {
	s.Composite.CurrentChild = 0
}

// Tick ...
func (s *Sequence) Tick(ctx *Context) Status {
	for {
		status := Update(s.Children[s.CurrentChild], ctx)
		if status != StatusSuccess {
			return status
		}
		s.CurrentChild++
		if s.CurrentChild >= len(s.Children) {
			return StatusSuccess
		}
	}
}
