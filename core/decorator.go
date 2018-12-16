package core

import (
	"fmt"
)

// Decorator base type
type Decorator struct {
	Node
	Child INode
}

// NewDecorator ...
func NewDecorator() Decorator {
	return Decorator{
		Node: Node{
			Category:   CategoryDecorator,
			StatusChan: make(chan Status, 1),
			// Data:       map[string]interface{}{},
		},
	}
}

// Init ...
func (d *Decorator) Init() {
	fmt.Println("Init deco", d.Type, "child", d.Child)
}

// Terminate ..
func (d *Decorator) Terminate(ctx *Context) {
	fmt.Println("Terminate deco", d.Type, "child", d.Child)
}

// GetChildren returns a list containing the only child of the decorator node.
func (d *Decorator) GetChildren() []INode {
	return append([]INode{}, d.Child)
}

// SetChild ...
func (d *Decorator) SetChild(child INode) {
	d.Child = child
}
