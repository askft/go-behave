package core

// Update updates a node by calling its Enter method if it is not running,
// then its Tick method, and finally Leave if it is not still running.
func Update[Context any](node Node[Context], ctx Context) Status {

	if node.GetStatus() != StatusRunning {
		node.Enter(ctx)
	}

	status := node.Tick(ctx)
	node.SetStatus(status)

	if status != StatusRunning {
		node.Leave(ctx)
	}

	return status
}
