package composite

import (
	"github.com/jbcpollak/go-behave/core"
)

// PersistentSequence updates each child in order. If a child
// returns Failure or Running, this node returns the same value,
// and resumes execution from the same child node the next tick.
func PersistentSequence[Context any](children ...core.Node[Context]) core.Node[Context] {
	base := core.NewComposite("PersistentSequence", children)
	return &persistentSequence[Context]{Composite: base}
}

type persistentSequence[Context any] struct {
	*core.Composite[Context]
}

func (s *persistentSequence[Context]) Enter(ctx Context) {}

func (s *persistentSequence[Context]) Tick(ctx Context) core.Status {
	for s.CurrentChild < len(s.Children) {
		status := core.Update(s.Children[s.CurrentChild], ctx)
		if status != core.StatusSuccess {
			return status
		}
		s.CurrentChild++
	}
	return core.StatusSuccess
}

func (s *persistentSequence[Context]) Leave(ctx Context) {
	s.Composite.CurrentChild = 0
}
