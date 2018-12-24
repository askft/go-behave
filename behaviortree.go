package behaviortree

import (
	"fmt"
	"strings"

	"github.com/alexanderskafte/behaviortree/core"
	"github.com/alexanderskafte/behaviortree/lang"
	"github.com/alexanderskafte/behaviortree/registry"
	"github.com/alexanderskafte/behaviortree/store"
	"github.com/alexanderskafte/behaviortree/util"
)

// Config ...
type Config struct {
	Owner      interface{}
	Store      store.Interface
	FnRegistry *registry.Registry
	Definition string
}

// NewBehaviorTree returns a new behavior tree. It takes three
// parameters: `owner`, the owner of the tree instance; `store`,
// a global store shared by all entities; `treeString`, a
// behavior tree defined as a string.
func NewBehaviorTree(cfg Config) (*BehaviorTree, error) {
	var eb util.ErrorBuilder
	eb.SetMessage("NewBehaviorTree")
	if cfg.Owner == nil {
		eb.Write("Config.Owner is nil")
	}
	if cfg.Store == nil {
		eb.Write("Config.Store is nil")
	}
	// if cfg.Registry == nil {
	// 	eb.Write("Config.Registry is nil")
	// }
	if cfg.FnRegistry == nil {
		eb.Write("Config.FnRegistry is nil")
	}
	if cfg.Definition == "" {
		eb.Write("Config.Definition is nil")
	}
	if eb.Error() != nil {
		return nil, eb.Error()
	}
	root, err := lang.NewParser(cfg.FnRegistry).Compile(cfg.Definition)
	if err != nil {
		return nil, err
	}
	tree := &BehaviorTree{
		Context:    core.NewContext(cfg.Owner, cfg.Store),
		FnRegistry: cfg.FnRegistry,
		Root:       root,
	}
	return tree, nil
}

// BehaviorTree ...
type BehaviorTree struct {
	Context    *core.Context
	FnRegistry *registry.Registry
	Root       core.INode
}

// Update propagates an update call down the behavior tree.
func (bt *BehaviorTree) Update() core.Status {
	return core.Update(bt.Root, bt.Context)
}

// String creates a string representation of the behavior tree
// by traversing it and writing lexical elements to a string.
func (bt *BehaviorTree) String() string {
	// var b strings.Builder
	// fmt.Println()
	// nodeRecurse(bt.Root, 0, &b)
	// return b.String()
	return NodeToString(bt.Root)
}

// NodeToString returns a string representation of a node,
// including those of all its children.
func NodeToString(node core.INode) string {
	var b strings.Builder
	fmt.Println()
	nodeRecurse(node, 0, &b)
	return b.String()
}

func nodeRecurse(node core.INode, level int, b *strings.Builder) {
	indent := strings.Repeat("    ", level)
	b.WriteString(indent + node.String() + "\n")
	for _, child := range node.GetChildren() {
		nodeRecurse(child, level+1, b)
	}
}
