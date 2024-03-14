package composite

import (
	"math/rand"
	"time"

	"github.com/jbcpollak/go-behave/core"
)

// RandomSelector creates a new random selector node.
func RandomSelector[Context any](children ...core.Node[Context]) core.Node[Context] {
	base := core.NewComposite("RandomSelector", children)
	return &randomSelector[Context]{Composite: base}
}

// randomSelector ...
type randomSelector[Context any] struct {
	*core.Composite[Context]
}

// Enter ...
func (s *randomSelector[Context]) Enter(ctx Context) {}

// Tick ...
func (s *randomSelector[Context]) Tick(ctx Context) core.Status {
	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(len(s.Children))
	child := s.Children[index]
	return core.Update(child, ctx)
}

// Leave ...
func (s *randomSelector[Context]) Leave(ctx Context) {}
