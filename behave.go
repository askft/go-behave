package behave

import (
	"github.com/jbcpollak/go-behave/core"
	"github.com/jbcpollak/go-behave/internal"
	"github.com/jbcpollak/go-behave/util"
)

// BehaviorTree ...
type BehaviorTree[Blackboard any, Event any] struct {
	Root       core.Node[Blackboard, Event]
	Blackboard Blackboard
}

// NewBehaviorTree returns a new BehaviorTree. A data Blackboard
// to be propagated down the tree each tick is created.
func NewBehaviorTree[Blackboard any, Event any](bb Blackboard, root core.Node[Blackboard, Event]) (*BehaviorTree[Blackboard, Event], error) {
	var eb internal.ErrorBuilder
	eb.SetMessage("NewBehaviorTree")
	if root == nil {
		eb.Write("Config.Root is nil")
	}
	// if bb == nil {
	// 	eb.Write("Config.Data is nil")
	// }
	if eb.Error() != nil {
		return nil, eb.Error()
	}
	tree := &BehaviorTree[Blackboard, Event]{
		Root:       root,
		Blackboard: bb,
	}
	return tree, nil
}

// Update propagates an update call down the behavior tree.
func (bt *BehaviorTree[Blackboard, Event]) Update(evt Event) core.Status {
	return core.Update(bt.Root, bt.Blackboard, evt)
}

// String creates a string representation of the behavior tree
// by traversing it and writing lexical elements to a string.
func (bt *BehaviorTree[Blackboard, Event]) String() string {
	return util.NodeToString(bt.Root)
}
