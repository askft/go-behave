package core

// Composite base type
type Composite struct {
	*Node
	Children     []INode
	CurrentChild int
}

// NewComposite ...
func NewComposite(name string) *Composite {
	return &Composite{
		Node:     NewNode(CategoryComposite, name),
		Children: []INode{},
	}
}

// GetChildren returns a list containing the children of the composite node.
func (c *Composite) GetChildren() []INode {
	return append([]INode{}, c.Children...)
}

// String returns a string representation of the composite node.
func (c *Composite) String() string {
	return "+ " + c.Name
}
