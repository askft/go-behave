package composite

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/alexanderskafte/behaviortree/core"
)

// RandomSelector ...
type RandomSelector struct {
	*core.Composite
}

// Initialize ...
func (s *RandomSelector) Initialize(args ...interface{}) {
	s.Composite = args[0].(*core.Composite)
}

// Start ...
func (s *RandomSelector) Start(ctx *core.Context) {}

// Tick ...
func (s *RandomSelector) Tick(ctx *core.Context) core.Status {
	fmt.Println("Run RandomSelector")
	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(len(s.Children))
	child := s.Children[index]
	return core.Update(child, ctx)
}

// Stop ...
func (s *RandomSelector) Stop(ctx *core.Context) {}
