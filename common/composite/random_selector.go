package composite

import (
	"math/rand"
	"time"

	"github.com/askft/go-behave/core"
)

// RandomSelector creates a new random selector node.
func RandomSelector(children ...core.Node) core.Node {
	base := core.NewComposite("RandomSelector", children)
	return &randomSelector{Composite: base}
}

// randomSelector ...
type randomSelector struct {
	*core.Composite
}

// Start ...
func (s *randomSelector) Start(ctx *core.Context) {}

// Tick ...
func (s *randomSelector) Tick(ctx *core.Context) core.Status {
	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(len(s.Children))
	child := s.Children[index]
	return core.Update(child, ctx)
}

// Stop ...
func (s *randomSelector) Stop(ctx *core.Context) {}
