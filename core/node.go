package core

import "context"

type Event any

// Preliminary interface to work around intermediate types like
// composite, decorator, etc not inplementing Enter/Tick/Leave
type Walkable[Blackboard any] interface {
	// Automatically implemented by embedding a pointer to a
	// Composite, Decorator or Leaf node in the custom node.
	Status() Status
	SetStatus(Status)
	Category() Category
	String() string

	Walk(WalkFunc[Blackboard], int)
}

type WalkFunc[Blackboard any] func(Walkable[Blackboard], int)

// The Node interface must be satisfied by any custom node.
type Node[Blackboard any] interface {
	Walkable[Blackboard]

	// Must be implemented by the custom node.
	Enter(Blackboard)
	Tick(Blackboard, context.Context, Event) NodeResult
	Leave(Blackboard)
}

// BaseNode contains properties shared by all categories of node.
// Do not use this type directly.
type BaseNode struct {
	category Category
	name     string
	status   Status
}

func newBaseNode(category Category, name string) BaseNode {
	return BaseNode{category: category, name: name}
}

// Status returns the status of this node.
func (n *BaseNode) Status() Status {
	return n.status
}

// SetStatus sets the status of this node.
func (n *BaseNode) SetStatus(status Status) {
	n.status = status
}

// GetCategory returns the category of this node.
func (n *BaseNode) Category() Category {
	return n.category
}
