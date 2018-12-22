package core

// Update updates a node.
func Update(node INode, ctx *Context) Status {
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
