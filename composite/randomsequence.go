package composite

import (
	"math/rand"

	. "github.com/alexanderskafte/behaviortree/core"
)

// Init ...
func (s *RandomSequence) Init() {
	shuffle(s.Children)
}

// RandomSequence ...
type RandomSequence struct {
	Composite
}

// Tick ...
func (s *RandomSequence) Tick(ctx *Context) Status {
	for {
		status := Update(s.Children[s.CurrentChild], ctx)
		if status != StatusSuccess {
			return status
		}
		s.Composite.CurrentChild++
		if s.CurrentChild >= len(s.Children) {
			return StatusSuccess
		}
	}
}

func shuffle(nodes []INode) {
	rand.Shuffle(len(nodes), func(i, j int) {
		nodes[i], nodes[j] = nodes[j], nodes[i]
	})
}
