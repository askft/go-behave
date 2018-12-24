package core

import (
	"fmt"
)

// Decorator base type
type Decorator struct {
	*Node
	Child  INode
	Params Params
}

// NewDecorator ...
func NewDecorator(name string, params map[string]string) *Decorator {
	return &Decorator{
		Node:   NewNode(CategoryDecorator, name),
		Params: params,
	}
}

// GetChildren returns a list containing the only child of the decorator node.
func (d *Decorator) GetChildren() []INode {
	return append([]INode{}, d.Child)
}

// String returns a string representation of the decorator node.
func (d *Decorator) String() string {
	return fmt.Sprintf("* %s (%v)", d.Name, d.Params)
}
