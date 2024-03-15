package core

type Event interface{}

// The Node interface must be satisfied by any custom node.
type Node[Blackboard any] interface {

	// Automatically implemented by embedding a pointer to a
	// Composite, Decorator or Leaf node in the custom node.
	GetStatus() Status
	SetStatus(Status)
	GetCategory() Category
	GetChildren() []Node[Blackboard]
	String() string

	// Must be implemented by the custom node.
	Enter(Blackboard)
	Tick(Blackboard, Event) NodeResult
	Leave(Blackboard)
}

// BaseNode contains properties shared by all categories of node.
// Do not use this type directly.
type BaseNode struct {
	category Category
	name     string
	status   Status
}

func newBaseNode(category Category, name string) *BaseNode {
	return &BaseNode{category: category, name: name}
}

// GetStatus returns the status of this node.
func (n *BaseNode) GetStatus() Status {
	return n.status
}

// SetStatus sets the status of this node.
func (n *BaseNode) SetStatus(status Status) {
	n.status = status
}

// GetCategory returns the category of this node.
func (n *BaseNode) GetCategory() Category {
	return n.category
}
