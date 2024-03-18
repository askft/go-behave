package composite

import (
	"context"
	"math/rand"
	"time"

	"github.com/jbcpollak/greenstalk/core"
)

// RandomSelector creates a new random selector node.
func RandomSelector[Blackboard any](children ...core.Node[Blackboard]) core.Node[Blackboard] {
	base := core.NewComposite("RandomSelector", children)
	return &randomSelector[Blackboard]{Composite: base}
}

// randomSelector ...
type randomSelector[Blackboard any] struct {
	core.Composite[Blackboard]
}

// Enter ...
func (s *randomSelector[Blackboard]) Enter(bb Blackboard) {}

// Tick ...
func (s *randomSelector[Blackboard]) Tick(bb Blackboard, ctx context.Context, evt core.Event) core.NodeResult {
	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(len(s.Children))
	child := s.Children[index]
	return core.Update(child, bb, ctx, evt)
}

// Leave ...
func (s *randomSelector[Blackboard]) Leave(bb Blackboard) {}
