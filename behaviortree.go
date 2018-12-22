package behaviortree

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/alexanderskafte/behaviortree/core"
	"github.com/alexanderskafte/behaviortree/lang"
	"github.com/alexanderskafte/behaviortree/registry"
	"github.com/alexanderskafte/behaviortree/store"
)

// Config ...
type Config struct {
	Owner      interface{}
	Store      store.Interface
	Registry   registry.Registry
	Definition string
}

// NewBehaviorTree returns a new behavior tree. It takes three
// parameters: `owner`, the owner of the tree instance; `store`,
// a global store shared by all entities; `treeString`, a
// behavior tree defined as a string.
func NewBehaviorTree(cfg Config) (*BehaviorTree, error) {
	if cfg.Owner == nil {
		return nil, fmt.Errorf("NewBehaviorTree: cfg.Owner is nil")
	}
	if cfg.Store == nil {
		return nil, fmt.Errorf("NewBehaviorTree: cfg.Store is nil")
	}
	if cfg.Registry == nil {
		return nil, fmt.Errorf("NewBehaviorTree: cfg.Registry is nil")
	}
	if cfg.Definition == "" {
		return nil, fmt.Errorf("NewBehaviorTree: cfg.Definition is nil")
	}
	root, err := lang.NewParser(strings.NewReader(cfg.Definition), cfg.Registry).Parse()
	if err != nil {
		return nil, err
	}
	tree := &BehaviorTree{
		Context:  core.NewContext(cfg.Owner, cfg.Store),
		Registry: cfg.Registry,
		Root:     root,
	}
	return tree, nil
}

// BehaviorTree ...
type BehaviorTree struct {
	Context  *core.Context
	Registry registry.Registry
	Root     core.INode
}

// Update propagates an update call down the behavior tree.
func (bt *BehaviorTree) Update() core.Status {
	return core.Update(bt.Root, bt.Context)
}

// String creates a string representation of the behavior tree
// by traversing it and writing lexical elements to a string.
func (bt *BehaviorTree) String() string {
	var b strings.Builder
	fmt.Println()
	nodeRecurse(bt.Root, 0, &b)
	return b.String()
}

func nodeRecurse(node core.INode, level int, b *strings.Builder) {
	indent := strings.Repeat("    ", level)

	if node.GetCategory() == core.CategoryLeaf {
		b.WriteString(indent + node.String())
	} else {
		name := reflect.TypeOf(node).Elem().Name()
		b.WriteString(indent + name)
		b.WriteString(" {")
	}
	b.WriteString("\n")
	children := node.GetChildren()
	for _, child := range children {
		nodeRecurse(child, level+1, b)
	}
	if node.GetCategory() != core.CategoryLeaf {
		b.WriteString(indent + "}\n")
	}
}
