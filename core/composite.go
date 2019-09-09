package core

// Composite is the base type for any specific composite node. Such a node
// may be domain-specific, but usually one of the common nodes will be used,
// such as Sequence or Selector.
type Composite struct {
	*BaseNode
	Children     []Node
	CurrentChild int // TODO - move into instance nodes
}

// NewComposite creates a new composite base node.
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
	return "+ " + c.name
}
