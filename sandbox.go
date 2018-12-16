package behaviortree

// /*
// 	Attempt to make use of channels!

// 	https://gamedev.stackexchange.com/q/154119

// 	https://takinginitiative.wordpress.com/2014/02/17/synchronized-behavior-trees/

// */

// // TC just stands for TestComposite or something
// type TC struct{ Composite }
// type TL struct{ Node }

// // Run a selector
// func (tc *TC) Tick(ctx *Context) Status {
// 	for _, child := range tc.Children {
// 		go child.Tick(ctx)
// 		status := <-child.GetChan()

// 		if status == StatusFailure {
// 			return StatusFailure
// 		}
// 	}
// 	return StatusSuccess
// }

// // func (tl *TL) Tick(ctx *Context) Status {

// // }

// // updateFunction is called every frame
// func updateFunction() {
// 	tc := &TC{}
// 	tc.run()
// }

// // to do the below we need to see if a node is fresh
// // or if it has already been started
// func (tc *TC) run() Status {
// 	for _, child := range tc.GetChildren() {
// 		select {
// 		case status := <-child.GetChan(): // success or failure only
// 			return status
// 		default:
// 			return StatusRunning
// 		}
// 	}
// }

// func (tl *TL) run() Status {

// }
