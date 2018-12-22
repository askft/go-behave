package core

import (
	"strings"
)

// Composite base type
type Composite struct {
	*Node
	Children     []INode
	CurrentChild int
}

// NewComposite ...
func NewComposite() *Composite {
	return &Composite{
		Node:     NewNode(CategoryComposite),
		Children: []INode{},
	}
}

// GetChildren returns a list containing the children of the composite node.
func (c *Composite) GetChildren() []INode {
	return append([]INode{}, c.Children...)
}

// AddChildren ...
func (c *Composite) AddChildren(children ...INode) {
	c.Children = append(c.Children, children...)
}

// String returns a string representation of the composite node.
func (c *Composite) String() string {
	ss := make([]string, len(c.Children))
	for i, child := range c.Children {
		ss[i] = child.String()
	}
	return strings.Join(ss, " BLOOP ")
}
