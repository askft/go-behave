package core

import (
	"fmt"
)

// Decorator base type
type Decorator struct {
	*BaseNode
	Child  Node
	Params Params
}

// NewDecorator ...
func NewDecorator(name string, params map[string]string, child Node) *Decorator {
	return &Decorator{
		BaseNode: newBaseNode(CategoryDecorator, name),
		Child:    child,
		Params:   params,
	}
}

// GetChildren returns a list containing the only child of the decorator node.
func (d *Decorator) GetChildren() []Node {
	return append([]Node{}, d.Child)
}

// String returns a string representation of the decorator node.
func (d *Decorator) String() string {
	return fmt.Sprintf("* %s (%v)", d.Name, d.Params)
}
