package core

// Update updates a node by calling its Enter method if it is not running,
// then its Tick method, and finally Leave if it is not still running.
func Update[Blackboard any, Event any](node Node[Blackboard, Event], bb Blackboard, evt Event) NodeResult {

	if node.GetStatus() != StatusRunning {
		node.Enter(bb)
	}

	result := node.Tick(bb, evt)

	status := result.Status()
	node.SetStatus(status)

	if status != StatusRunning {
		node.Leave(bb)
	}

	return result
}
