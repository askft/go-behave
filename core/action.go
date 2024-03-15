package core

import (
	"fmt"
)

// Leaf is the base type for any specific leaf node (domain-specific).
// Each leaf node has Params: data keys that the implementation imports
// and Returns: data keys that the implementation exports.
type Leaf[Blackboard any] struct {
	*BaseNode
	Params  Params
	Returns Returns
}

// NewLeaf creates a new leaf base node.
func NewLeaf[Blackboard any](name string, params Params, returns Returns) *Leaf[Blackboard] {
	return &Leaf[Blackboard]{
		BaseNode: newBaseNode(CategoryLeaf, name),
		Params:   params,  // TODO (remove): These are only used for String()
		Returns:  returns, // TODO (remove): These are only used for String()
	}
}

// GetChildren returns an empty list of Node, since a leaf has no children.
// This method is required for Leaf in order to implement Node.
func (a *Leaf[Blackboard]) GetChildren() []Node[Blackboard] {
	return []Node[Blackboard]{}
}

// String returns a string representation of the leaf node.
func (a *Leaf[Blackboard]) String() string {
	return fmt.Sprintf("! %s (%v : %v)",
		a.name,
		a.Params,
		a.Returns,
	)
}
