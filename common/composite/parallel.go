package composite

import (
	"github.com/jbcpollak/go-behave/core"
)

// Parallel updates all its children in parallel, i.e. every frame.
// It does not retry on nodes that have failed or succeeded.
//
// succ/failReq is the minimum amount of nodes required to
// succeed/fail for the parallel sequence node itself to succeed/fail.
// A value of 0 for either node means that all nodes must succeed/fail.
func Parallel[Context any](succReq, failReq int, children ...core.Node[Context]) core.Node[Context] {
	base := core.NewComposite("Parallel", children)
	if succReq == 0 {
		succReq = len(children)
	}
	if failReq == 0 {
		failReq = len(children)
	}
	return &parallel[Context]{
		base,
		succReq,
		failReq,
		0,
		0,
		make([]bool, len(children)),
	}
}

type parallel[Context any] struct {
	*core.Composite[Context]
	succReq   int
	failReq   int
	succ      int
	fail      int
	completed []bool
}

func (s *parallel[Context]) Enter(ctx Context) {
	s.succ = 0
	s.fail = 0
}

func (s *parallel[Context]) Tick(ctx Context) core.Status {

	// Update every child that has not completed yet every tick.
	for i := 0; i < len(s.Children); i++ {

		// Ignore a child if has already succeeded or failed.
		if s.completed[i] {
			continue
		}

		// Update a child and count whether it succeeded or failed,
		// and mark it as completed in either of those two cases.
		switch core.Update(s.Children[i], ctx) {
		case core.StatusSuccess:
			s.succ++
			s.completed[i] = true
		case core.StatusFailure:
			s.fail++
			s.completed[i] = true
		}
	}

	if s.succ >= s.succReq {
		return core.StatusSuccess
	}
	if s.fail >= s.failReq {
		return core.StatusFailure
	}
	return core.StatusRunning
}

func (s *parallel[Context]) Leave(ctx Context) {}
