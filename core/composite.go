package core

// Composite base type
type Composite struct {
	*BaseNode
	Children     []Node
	CurrentChild int
}

// NewComposite ...
func NewComposite(name string, children []Node) *Composite {
	return &Composite{
		BaseNode: newBaseNode(CategoryComposite, name),
		Children: children,
	}
}

// GetChildren returns a list containing the children of the composite node.
func (c *Composite) GetChildren() []Node {
	return append([]Node{}, c.Children...)
}

// String returns a string representation of the composite node.
func (c *Composite) String() string {
	return "+ " + c.Name
}
