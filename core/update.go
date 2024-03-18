package core

import "context"

// Update updates a node by calling its Enter method if it is not running,
// then its Tick method, and finally Leave if it is not still running.
func Update[Blackboard any](node Node[Blackboard], bb Blackboard, ctx context.Context, evt Event) NodeResult {

	if node.Status() != StatusRunning {
		node.Enter(bb)
	}

	result := node.Tick(bb, ctx, evt)

	status := result.Status()
	node.SetStatus(status)

	if status != StatusRunning {
		node.Leave(bb)
	}

	return result
}
