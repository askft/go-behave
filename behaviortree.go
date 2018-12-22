package behaviortree

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/alexanderskafte/behaviortree/core"
	"github.com/alexanderskafte/behaviortree/lang"
	"github.com/alexanderskafte/behaviortree/registry"
	"github.com/alexanderskafte/behaviortree/store"

	"github.com/fatih/color"
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

// // EmitCommands ... TODO
// func (bt *BehaviorTree) EmitCommands() []command.Command {
// 	data, ok := bt.Context.Store.Read(bt.Context.Owner.(fmt.Stringer).String() + ".actions")
// 	if !ok {
// 		return []command.Command{}
// 	}
// 	return data.([]command.Command)
// }

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

	col := colorFor(node.GetStatus())
	color.Set(col)

	if node.GetCategory() == core.CategoryLeaf {
		fmt.Printf(indent + node.String())
		color.Unset()
	} else {
		name := reflect.TypeOf(node).Elem().Name()
		fmt.Printf(indent + name)
		color.Unset()
		fmt.Printf(" {")
	}
	fmt.Printf("\n")
	children := node.GetChildren()
	for _, child := range children {
		nodeRecurse(child, level+1, b)
	}
	if node.GetCategory() != core.CategoryLeaf {
		fmt.Printf(indent + "}\n")
	}
}

func colorFor(status core.Status) color.Attribute {
	switch status {
	case core.StatusFailure:
		return color.FgRed
	case core.StatusRunning:
		return color.FgYellow
	case core.StatusSuccess:
		return color.FgGreen
	case core.StatusInvalid:
		return color.FgMagenta
	default:
		panic("invalid color")
	}
}
