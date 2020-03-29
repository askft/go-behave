package composite

import (
	"log"

	"github.com/askft/go-behave/core"
)

// Parallel updates all its children in parallel, i.e. every frame.
// It does not retry on nodes that have failed or succeeded.
//
// succ/failReq is the minimum amount of nodes required to
// succeed/fail for the parallel sequence node itself to succeed/fail.
// A value of 0 for either node means that all nodes must succeed/fail.
func Parallel(succReq, failReq int, children ...core.Node) core.Node {
	base := core.NewComposite("Parallel", children)
	if succReq == 0 {
		succReq = len(children)
	}
	if failReq == 0 {
		failReq = len(children)
	}
	return &parallel{
		base,
		succReq,
		failReq,
		0,
		0,
		make([]bool, len(children)),
	}
}

type parallel struct {
	*core.Composite
	succReq   int
	failReq   int
	succ      int
	fail      int
	completed []bool
}

func (s *parallel) Enter(ctx *core.Context) {
	s.succ = 0
	s.fail = 0
}

func (s *parallel) Tick(ctx *core.Context) core.Status {

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
			log.Println("success")
			s.succ++
			s.completed[i] = true
		case core.StatusFailure:
			log.Println("failure")
			s.fail++
			s.completed[i] = true
		}
	}

	log.Printf("s.succ: %v; s.succReq: %v", s.succ, s.succReq)
	if s.succ >= s.succReq {
		return core.StatusSuccess
	}
	log.Printf("s.fail: %v; s.failReq: %v", s.fail, s.failReq)
	if s.fail >= s.failReq {
		return core.StatusFailure
	}
	return core.StatusRunning
}

func (s *parallel) Leave(ctx *core.Context) {}
