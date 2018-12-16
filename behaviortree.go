package behaviortree

import (
	"fmt"
	"strings"

	"github.com/alexanderskafte/behaviortree/core"
	"github.com/alexanderskafte/behaviortree/lang"
	"github.com/alexanderskafte/behaviortree/store"
)

// NewBehaviorTree returns a new behavior tree. It takes three
// parameters: `id`, identifying the tree instance; `store`,
// a global store shared by all entities; `treeString`, a
// behavior tree defined as a string.
func NewBehaviorTree(id fmt.Stringer, store store.Interface, treeString string,
) (
	*BehaviorTree, error,
) {
	node, err := lang.NewParser(strings.NewReader(treeString)).Parse()
	if err != nil {
		return nil, err
	}
	tree := &BehaviorTree{
		Context: core.NewContext(id, store),
		Child:   node,
	}
	return tree, nil
}

// BehaviorTree ...
type BehaviorTree struct {
	Context *core.Context
	Child   core.INode
}

// Update propagates an update call down the behavior tree.
func (bt *BehaviorTree) Update() core.Status {
	return core.Update(bt.Child, bt.Context)
}

// String creates a string representation of the behavior tree
// by traversing it and writing lexical elements to a string.
func (bt *BehaviorTree) String() string {
	var b strings.Builder
	nodeRecurse(bt.Child, 0, &b)
	return b.String()
}

func nodeRecurse(node core.INode, level int, b *strings.Builder) {
	indent := strings.Repeat("  ", level)
	b.WriteString(indent + string(node.GetType()))
	if node.GetCategory() != core.CategoryLeaf {
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
