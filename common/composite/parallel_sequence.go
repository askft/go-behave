package composite

import (
	"github.com/askft/go-behave/core"
)

// ParallelSequence creates a new parallelSequence node.
// succ/failReq is the minimum amount of nodes required to
// succeed/fail for the parallel sequence node itself to succeed/fail.
// A value of -1 for either node means that all nodes must succeed/fail.
func ParallelSequence(succReq, failReq int, children ...core.Node) core.Node {
	base := core.NewComposite("ParallelSequence", children)
	if succReq == -1 {
		succReq = len(children)
	}
	if failReq == -1 {
		failReq = len(children)
	}
	return &parallelSequence{base, succReq, failReq}
}

// parallelSequence ...
type parallelSequence struct {
	*core.Composite
	succReq int
	failReq int
}

// Start ...
func (s *parallelSequence) Start(ctx *core.Context) {}

// Tick ...
func (s *parallelSequence) Tick(ctx *core.Context) core.Status {
	successes := 0
	failures := 0
	for i := range s.Children {
		switch core.Update(s.Children[i], ctx) {
		case core.StatusSuccess:
			successes++
		case core.StatusFailure:
			failures++
		}
	}
	if successes >= s.succReq {
		return core.StatusSuccess
	}
	if successes >= s.failReq {
		return core.StatusFailure
	}
	return core.StatusRunning
}

// Stop ...
func (s *parallelSequence) Stop(ctx *core.Context) {}
