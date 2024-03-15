package core

import (
	"fmt"
)

// Leaf is the base type for any specific leaf node (domain-specific).
// Each leaf node has Params: data keys that the implementation imports
// and Returns: data keys that the implementation exports.
type Leaf[Blackboard any] struct {
	BaseNode
	Params  Params
	Returns Returns
}

// NewLeaf creates a new leaf base node.
func NewLeaf[Blackboard any](name string, params Params, returns Returns) Leaf[Blackboard] {
	return Leaf[Blackboard]{
		BaseNode: newBaseNode(CategoryLeaf, name),
		Params:   params,  // TODO (remove): These are only used for String()
		Returns:  returns, // TODO (remove): These are only used for String()
	}
}

func (c *Leaf[Blackboard]) Walk(walkFn WalkFunc[Blackboard], level int) {
	walkFn(c, level)
}

// String returns a string representation of the leaf node.
func (a *Leaf[Blackboard]) String() string {
	return fmt.Sprintf("! %s (%v : %v)",
		a.name,
		a.Params,
		a.Returns,
	)
}
