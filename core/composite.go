package core

// Composite is the base type for any specific composite node. Such a node
// may be domain-specific, but usually one of the common nodes will be used,
// such as Sequence or Selector.
type Composite[Context any] struct {
	*BaseNode
	Children     []Node[Context]
	CurrentChild int // TODO - move into instance nodes
}

// NewComposite creates a new composite base node.
func NewComposite[Context any](name string, children []Node[Context]) *Composite[Context] {
	return &Composite[Context]{
		BaseNode: newBaseNode(CategoryComposite, name),
		Children: children,
	}
}

// GetChildren returns a list containing the children of the composite node.
func (c *Composite[Context]) GetChildren() []Node[Context] {
	return append([]Node[Context]{}, c.Children...)
}

// String returns a string representation of the composite node.
func (c *Composite[Context]) String() string {
	return "+ " + c.name
}
