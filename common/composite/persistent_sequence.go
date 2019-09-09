package composite

import (
	"github.com/askft/go-behave/core"
)

// PersistentSequence updates each child in order. If a child
// returns Failure or Running, this node returns the same value,
// and resumes execution from the same child node the next tick.
func PersistentSequence(children ...core.Node) core.Node {
	base := core.NewComposite("PersistentSequence", children)
	return &persistentSequence{Composite: base}
}

type persistentSequence struct {
	*core.Composite
}

func (s *persistentSequence) Enter(ctx *core.Context) {}

func (s *persistentSequence) Tick(ctx *core.Context) core.Status {
	for s.CurrentChild < len(s.Children) {
		status := core.Update(s.Children[s.CurrentChild], ctx)
		if status != core.StatusSuccess {
			return status
		}
		s.CurrentChild++
	}
	return core.StatusSuccess
}

func (s *persistentSequence) Leave(ctx *core.Context) {
	s.Composite.CurrentChild = 0
}
