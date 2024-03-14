package behave

import (
	"github.com/jbcpollak/go-behave/core"
	"github.com/jbcpollak/go-behave/internal"
	"github.com/jbcpollak/go-behave/util"
)

// BehaviorTree ...
type BehaviorTree[Context any] struct {
	Root    core.Node[Context]
	Context Context
}

// NewBehaviorTree returns a new BehaviorTree. A data context
// to be propagated down the tree each tick is created.
func NewBehaviorTree[Context any](ctx Context, root core.Node[Context]) (*BehaviorTree[Context], error) {
	var eb internal.ErrorBuilder
	eb.SetMessage("NewBehaviorTree")
	if root == nil {
		eb.Write("Config.Root is nil")
	}
	// if ctx == nil {
	// 	eb.Write("Config.Data is nil")
	// }
	if eb.Error() != nil {
		return nil, eb.Error()
	}
	tree := &BehaviorTree[Context]{
		Root:    root,
		Context: ctx,
	}
	return tree, nil
}

// Update propagates an update call down the behavior tree.
func (bt *BehaviorTree[Context]) Update() core.Status {
	return core.Update(bt.Root, bt.Context)
}

// String creates a string representation of the behavior tree
// by traversing it and writing lexical elements to a string.
func (bt *BehaviorTree[Context]) String() string {
	return util.NodeToString(bt.Root)
}
