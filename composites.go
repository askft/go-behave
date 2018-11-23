package behaviortree

import (
	"math/rand"
	"time"
)

// Sequence ...
type Sequence struct{ Composite }

// RandomSequence ...
type RandomSequence struct{ Composite }

// Selector ...
type Selector struct{ Composite }

// RandomSelector ...
type RandomSelector struct{ Composite }

// Run ...
func (s *Sequence) Run(ctx *Context) Status {
	for _, child := range s.Children {
		status := child.Run(ctx)
		if status != StatusSuccess {
			return status
		}
	}
	return StatusSuccess
}

// Run ...
func (s *RandomSequence) Run(ctx *Context) Status {
	shuffle(s.Children)
	for _, child := range s.Children {
		status := child.Run(ctx)
		if status != StatusSuccess {
			return status
		}
	}
	return StatusSuccess
}

// Run ...
func (s *Selector) Run(ctx *Context) Status {
	for _, child := range s.Children {
		status := child.Run(ctx)
		if status != StatusFailure {
			return status
		}
	}
	return StatusFailure
}

// Run ...
func (s *RandomSelector) Run(ctx *Context) Status {
	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(len(s.Children))
	child := s.Children[index]
	return child.Run(ctx)
}

// ParallelSequence ...
// type ParallelSequence struct {
// 	Composite
// 	wg sync.WaitGroup
// }

// Run ...
// func (s *ParallelSequence) Run(ctx *Context) {
// 	for _, child := range s.Children {
// 		s.wg.Add(1)
// 		go func() {
// 			s.wg.Wait()
// 			child.Run(ctx)
// 		}()
// 	}
// }
