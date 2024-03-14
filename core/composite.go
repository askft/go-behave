package core

// Composite is the base type for any specific composite node. Such a node
// may be domain-specific, but usually one of the common nodes will be used,
// such as Sequence or Selector.
type Composite[Blackboard any, Event any] struct {
	*BaseNode
	Children     []Node[Blackboard, Event]
	CurrentChild int // TODO - move into instance nodes
}

// NewComposite creates a new composite base node.
func NewComposite[Blackboard any, Event any](name string, children []Node[Blackboard, Event]) *Composite[Blackboard, Event] {
	return &Composite[Blackboard, Event]{
		BaseNode: newBaseNode(CategoryComposite, name),
		Children: children,
	}
}

// GetChildren returns a list containing the children of the composite node.
func (c *Composite[Blackboard, Event]) GetChildren() []Node[Blackboard, Event] {
	return append([]Node[Blackboard, Event]{}, c.Children...)
}

// String returns a string representation of the composite node.
func (c *Composite[Blackboard, Event]) String() string {
	return "+ " + c.name
}
