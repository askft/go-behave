package composite

// RandomSelector ...
// type RandomSelector struct{ Composite }

// Tick ...
// func (s *RandomSelector) Tick(ctx *Context) Status {
// 	rand.Seed(time.Now().UnixNano())
// 	index := rand.Intn(len(s.Children))
// 	child := s.Children[index]
// 	return child.Tick(ctx)
// }

// ParallelSequence ...
// type ParallelSequence struct {
// 	Composite
// 	wg sync.WaitGroup
// }

// Tick ...
// func (s *ParallelSequence) Tick(ctx *Context) {
// 	for _, child := range s.Children {
// 		s.wg.Add(1)
// 		go func() {
// 			s.wg.Wait()
// 			child.Tick(ctx)
// 		}()
// 	}
// }

// Old Sequence
// for _, child := range s.Children {
// 	status := child.Tick(ctx)
// 	if status != StatusSuccess {
// 		return status
// 	}
// }
// return StatusSuccess

// func (s *Selector) Tick(ctx *Context) Status {
// 	for _, child := range s.Children {
// 		status := child.Tick(ctx)
// 		if status != StatusFailure {
// 			return status
// 		}
// 	}
// 	return StatusFailure
// }
