package composite

import (
	"math/rand"
	"time"

	"github.com/jbcpollak/go-behave/core"
)

// RandomSelector creates a new random selector node.
func RandomSelector[Blackboard any, Event any](children ...core.Node[Blackboard, Event]) core.Node[Blackboard, Event] {
	base := core.NewComposite("RandomSelector", children)
	return &randomSelector[Blackboard, Event]{Composite: base}
}

// randomSelector ...
type randomSelector[Blackboard any, Event any] struct {
	*core.Composite[Blackboard, Event]
}

// Enter ...
func (s *randomSelector[Blackboard, Event]) Enter(bb Blackboard) {}

// Tick ...
func (s *randomSelector[Blackboard, Event]) Tick(bb Blackboard, evt Event) core.NodeResult {
	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(len(s.Children))
	child := s.Children[index]
	return core.Update(child, bb, evt)
}

// Leave ...
func (s *randomSelector[Blackboard, Event]) Leave(bb Blackboard) {}
