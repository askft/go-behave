package core

import (
	"fmt"
)

// Composite base type
type Composite struct {
	Node
	Children     []INode
	CurrentChild int
}

// NewComposite ...
func NewComposite() Composite {
	return Composite{
		Node: Node{
			Category:   CategoryComposite,
			StatusChan: make(chan Status, 1),
			// Data:       map[string]interface{}{},
		},
		Children: []INode{},
	}
}

// Init ...
func (c *Composite) Init() {
	fmt.Println("init comp", c.Type, "children", c.Children)
}

// Terminate ..
func (c *Composite) Terminate(ctx *Context) {
	fmt.Println("Terminate comp", c.Type, "child", c.Children)
}

// GetChildren returns a list containing the children of the composite node.
func (c *Composite) GetChildren() []INode {
	return append([]INode{}, c.Children...)
}

// AddChildren ...
func (c *Composite) AddChildren(children ...INode) {
	c.Children = append(c.Children, children...)
}
