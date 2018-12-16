package composite

import (
	. "github.com/alexanderskafte/behaviortree/core"
)

// Selector ...
type Selector struct {
	Composite
}

// Init ...
func (s *Selector) Init() {
	s.Composite.CurrentChild = 0
}

// Tick ...
func (s *Selector) Tick(ctx *Context) Status {
	for {
		status := Update(s.Children[s.CurrentChild], ctx)
		if status != StatusFailure {
			return status
		}
		s.Composite.CurrentChild++
		if s.CurrentChild >= len(s.Children) {
			return StatusFailure
		}
	}
}
