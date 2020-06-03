package behave

import (
	"github.com/askft/go-behave/core"
	"github.com/askft/go-behave/internal"
	"github.com/askft/go-behave/util"
)

// Config describes the configuration of a BehaviorTree object.
type Config struct {
	Owner interface{} // Owner of a tree instance
	Data  interface{} // Global data shared by all entities
	Root  core.Node   // Root node of the tree
}

// BehaviorTree ...
type BehaviorTree struct {
	Root    core.Node
	Context *core.Context
}

// NewBehaviorTree returns a new BehaviorTree. A data context
// to be propagated down the tree each tick is created.
func NewBehaviorTree(cfg Config) (*BehaviorTree, error) {
	var eb internal.ErrorBuilder
	eb.SetMessage("NewBehaviorTree")
	if cfg.Root == nil {
		eb.Write("Config.Root is nil")
	}
	if cfg.Owner == nil {
		eb.Write("Config.Owner is nil")
	}
	if cfg.Data == nil {
		eb.Write("Config.Data is nil")
	}
	if eb.Error() != nil {
		return nil, eb.Error()
	}
	tree := &BehaviorTree{
		Root:    cfg.Root,
		Context: core.NewContext(cfg.Owner, cfg.Data),
	}
	return tree, nil
}

// Update propagates an update call down the behavior tree.
func (bt *BehaviorTree) Update() core.Status {
	return core.Update(bt.Root, bt.Context)
}

// String creates a string representation of the behavior tree
// by traversing it and writing lexical elements to a string.
func (bt *BehaviorTree) String() string {
	return util.NodeToString(bt.Root)
}
