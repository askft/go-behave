package core

import (
	"fmt"
)

// Decorator is the base type for any specific decorator node. Such a node
// may be domain-specific, but usually one of the common nodes will be used,
// such as Inverter or Repeater. Each decorator node has Params: a key-value
// map used for setting variables for a specific decorator node, for instance
// Params{"n": 5} for a Repeater node or Params{"ms": 500"} for a
// Delayer node.
type Decorator struct {
	*BaseNode
	Child  Node
	Params Params
}

// NewDecorator creates a new decorator base node.
func NewDecorator(name string, params Params, child Node) *Decorator {
	return &Decorator{
		BaseNode: newBaseNode(CategoryDecorator, name),
		Child:    child,
		Params:   params, // TODO (remove): This is only used for String()
	}
}

// GetChildren returns a list containing the only child of the decorator node.
func (d *Decorator) GetChildren() []Node {
	return append([]Node{}, d.Child)
}

// String returns a string representation of the decorator node.
func (d *Decorator) String() string {
	return fmt.Sprintf("* %s (%v)", d.name, d.Params)
}
