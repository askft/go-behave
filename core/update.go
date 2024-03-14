package core

import (
	"fmt"
)

// Update updates a node by calling its Enter method if it is not running,
// then its Tick method, and finally Leave if it is not still running.
func Update[Blackboard any, Event any](node Node[Blackboard, Event], bb Blackboard, evt Event) Status {

	if node.GetStatus() != StatusRunning {
		node.Enter(bb)
	}

	result := node.Tick(bb, evt)

	status := result.status()
	switch status {
	case StatusSuccess:
		// whatever
	case StatusFailure:
		// whatever
	case StatusRunning:
		if asyncRunning, ok := result.(NodeAsyncRunning[Event]); ok {
			go asyncRunning(func(evt Event) error {
				return fmt.Errorf("TODO: enqueue event %v", evt)
			})
		}
	case StatusError:
		if err, ok := result.(NodeRuntimeError); ok {
			panic(err)
		}
	default:
		panic(fmt.Errorf("invalid status %v", status))
	}
	node.SetStatus(status)

	if status != StatusRunning {
		node.Leave(bb)
	}

	return status
}
