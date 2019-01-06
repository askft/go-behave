package core

// Update updates a node by calling its Enter method if it is not running,
// then its Tick method, and finally Leave if it is not still running.
func Update(node Node, ctx *Context) Status {
	status := node.GetStatus()
	if status != StatusRunning {
		node.Enter(ctx)
	}
	status = node.Tick(ctx)
	node.SetStatus(status)
	if status != StatusRunning {
		node.Leave(ctx)
	}
	return status
}
