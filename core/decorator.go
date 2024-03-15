package core

import (
	"fmt"
)

// Decorator is the base type for any specific decorator node. Such a node
// may be domain-specific, but usually one of the common nodes will be used,
// such as Inverter or Repeater. Each decorator node has Params: a key-value
// map used for setting variables for a specific decorator node, for instance
// Params{"n": 5} for a Repeater node or Params{"ms": 500} for a
// Delayer node.
type Decorator[Blackboard any] struct {
	BaseNode
	Child  Node[Blackboard]
	Params Params
}

// NewDecorator creates a new decorator base node.
func NewDecorator[Blackboard any](name string, params Params, child Node[Blackboard]) Decorator[Blackboard] {
	return Decorator[Blackboard]{
		BaseNode: newBaseNode(CategoryDecorator, name),
		Child:    child,
		Params:   params, // TODO (remove): This is only used for String()
	}
}

func (c *Decorator[Blackboard]) Walk(walkFn WalkFunc[Blackboard], level int) {
	walkFn(c, level)
	c.Child.Walk(walkFn, level+1)
}

// String returns a string representation of the decorator node.
func (d *Decorator[Blackboard]) String() string {
	return fmt.Sprintf("* %s (%v)", d.name, d.Params)
}
