package composite

import (
	"fmt"
	"math/rand"

	"github.com/alexanderskafte/behaviortree/core"
)

// Initialize ...
func (s *RandomSequence) Initialize(args ...interface{}) {
	s.Composite = args[0].(*core.Composite)
}

// Start ...
func (s *RandomSequence) Start(ctx *core.Context) {
	shuffle(s.Children)
}

// RandomSequence ...
type RandomSequence struct {
	*core.Composite
}

// Tick ...
func (s *RandomSequence) Tick(ctx *core.Context) core.Status {
	fmt.Println("Run RandomSequence")
	for {
		status := core.Update(s.Children[s.CurrentChild], ctx)
		if status != core.StatusSuccess {
			return status
		}
		s.Composite.CurrentChild++
		if s.CurrentChild >= len(s.Children) {
			return core.StatusSuccess
		}
	}
}

// Stop ...
func (s *RandomSequence) Stop(ctx *core.Context) {
	s.Composite.CurrentChild = 0
}

func shuffle(nodes []core.INode) {
	rand.Shuffle(len(nodes), func(i, j int) {
		nodes[i], nodes[j] = nodes[j], nodes[i]
	})
}
