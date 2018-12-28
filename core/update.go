package core

// Update updates a node by calling its Start method if it is not running,
// then its Tick method, and finally Stop if it is not still running.
func Update(node Node, ctx *Context) Status {
	status := node.GetStatus()
	if status != StatusRunning {
		node.Start(ctx)
	}
	status = node.Tick(ctx)
	node.SetStatus(status)
	if status != StatusRunning {
		node.Stop(ctx)
	}
	return status
}
