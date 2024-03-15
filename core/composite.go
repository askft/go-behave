package core

// Composite is the base type for any specific composite node. Such a node
// may be domain-specific, but usually one of the common nodes will be used,
// such as Sequence or Selector.
type Composite[Blackboard any] struct {
	*BaseNode
	Children     []Node[Blackboard]
	CurrentChild int // TODO - move into instance nodes
}

// NewComposite creates a new composite base node.
func NewComposite[Blackboard any](name string, children []Node[Blackboard]) *Composite[Blackboard] {
	return &Composite[Blackboard]{
		BaseNode: newBaseNode(CategoryComposite, name),
		Children: children,
	}
}

// GetChildren returns a list containing the children of the composite node.
func (c *Composite[Blackboard]) GetChildren() []Node[Blackboard] {
	return append([]Node[Blackboard]{}, c.Children...)
}

// String returns a string representation of the composite node.
func (c *Composite[Blackboard]) String() string {
	return "+ " + c.name
}
