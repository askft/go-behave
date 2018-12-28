package behave

import (
	"github.com/alexanderskafte/go-behave/core"
	"github.com/alexanderskafte/go-behave/gbl"
	"github.com/alexanderskafte/go-behave/store"
	"github.com/alexanderskafte/go-behave/util"
)

// Config describes the configuration of a BehaviorTree object.
type Config struct {
	Owner interface{}     // Owner of a tree instance
	Store store.Interface // Global store shared by all entities
	Root  core.Node       // Root node of the tree
}

// BehaviorTree ...
type BehaviorTree struct {
	Root    core.Node
	context *core.Context
}

// NewBehaviorTree returns a new BehaviorTree. A data context
// to be propagated down the tree each tick is created.
func NewBehaviorTree(cfg Config) (*BehaviorTree, error) {
	var eb util.ErrorBuilder
	eb.SetMessage("NewBehaviorTree")
	if cfg.Owner == nil {
		eb.Write("Config.Owner is nil")
	}
	if cfg.Store == nil {
		eb.Write("Config.Store is nil")
	}
	if cfg.Root == nil {
		eb.Write("Config.Root is nil")
	}
	if eb.Error() != nil {
		return nil, eb.Error()
	}
	tree := &BehaviorTree{
		Root:    cfg.Root,
		context: core.NewContext(cfg.Owner, cfg.Store),
	}
	return tree, nil
}

// Update propagates an update call down the behavior tree.
func (bt *BehaviorTree) Update() core.Status {
	return core.Update(bt.Root, bt.context)
}

// String creates a string representation of the behavior tree
// by traversing it and writing lexical elements to a string.
func (bt *BehaviorTree) String() string {
	return util.NodeToString(bt.Root)
}

// NewNode takes a set of nodes defined in a registry, and a definition
// string for a behavior tree, compiles the string and returns the root
// node if the compilation was successful, else an error.
func NewNode(reg *gbl.Registry, def string) (core.Node, error) {
	node, err := gbl.NewParser(reg).Compile(def)
	if err != nil {
		return nil, err
	}
	return node, nil
}
