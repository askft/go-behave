package core

// https://www.gamasutra.com/blogs/ChrisSimpson/20140717/221339/Behavior_trees_for_AI_How_they_work.php
// TODO: Decorator section at https://github.com/libgdx/gdx-ai/wiki/Behavior-Trees
// Good tutorial: https://github.com/adnzzzzZ/blog/issues/3

// Update updates a node.
func Update(node INode, ctx *Context) Status {
	status := <-node.GetChan() // NOTE; TODO; might block!!!
	if status != StatusRunning {
		node.Init() // TODO have context
	}
	status = node.Tick(ctx)
	if status != StatusRunning {
		node.Terminate(ctx)
	}
	return status
}

// TODO
// There shouldn't be Initialize and Terminate methods for
// the categorical nodes, only for specific ones.

// type taskWait struct {
// 	duration time.Duration
// 	visited  bool
// 	timer    *time.Timer
// }

// func (t *taskWait) init() {
// 	t.timer = time.NewTimer(t.duration)
// }

// func (t *taskWait) Tick(ctx *Context) Status {
// 	if !t.visited {
// 		t.visited = true
// 		t.timer = time.NewTimer(t.duration)
// 	}
// 	select {
// 	case <-t.timer.C:
// 		t.timer.Stop()
// 		t.visited = false
// 		fmt.Println("task done")
// 		return StatusSuccess
// 	default:
// 		return StatusRunning
// 	}
// }

// Functions ----------------------------------------------

// type actionFn func(in, out interface{}) Status

// type leaf struct {
// 	fn      actionFn
// 	in, out interface{}
// }

// func (l *leaf) Tick(ctx *Context) Status {
// 	return l.fn(l.in, l.out)
// }

// type Action struct {
// 	leaf
// 	data interface{}
// }

// type condition struct {
// 	leaf
// 	data interface{}
// }
