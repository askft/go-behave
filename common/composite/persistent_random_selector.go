package composite

import (
	"math/rand"

	"github.com/askft/go-behave/core"
)

// PersistentRandomSelector updates each child in order. If a child
// returns Failure or Running, this node returns the same value,
// and resumes execution from the same child node the next tick.
// The child to Update is chosen randomly
func PersistentRandomSelector(children ...core.Node) core.Node {
	base := core.NewComposite("PersistentRandomSelector", children)
	return &persistentRandomSelector{Composite: base}
}

type persistentRandomSelector struct {
	*core.Composite
}

func (s *persistentRandomSelector) Enter(ctx *core.Context) {
	// rand.Seed(time.Now().UnixNano())
	index := rand.Intn(len(s.Children))
	s.Composite.CurrentChild = index
}

func (s *persistentRandomSelector) Tick(ctx *core.Context) core.Status {
	return core.Update(s.Children[s.Composite.CurrentChild], ctx)
}

func (s *persistentRandomSelector) Leave(ctx *core.Context) {
	// s.Composite.CurrentChild = 0
}
