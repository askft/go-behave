package core

// Decorator base type
type Decorator struct {
	*Node
	Child INode
}

// NewDecorator ...
func NewDecorator() *Decorator {
	return &Decorator{
		Node: NewNode(CategoryDecorator),
	}
}

// GetChildren returns a list containing the only child of the decorator node.
func (d *Decorator) GetChildren() []INode {
	return append([]INode{}, d.Child)
}

// SetChild ...
func (d *Decorator) SetChild(child INode) {
	d.Child = child
}

// String returns a string representation of the decorator node.
func (d *Decorator) String() string {
	return "Decorator {" + d.Child.String() + "}"
}
