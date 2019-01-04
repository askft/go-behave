package core

import (
	"fmt"
)

// Leaf is the base type for any specific leaf node (domain-specific).
// Each leaf node has Params: data keys that the implementation imports
// and Returns: data keys that the implementation exports.
type Leaf struct {
	*BaseNode
	Params  Params
	Returns Returns
}

// NewLeaf creates a new leaf base node.
func NewLeaf(name string, params Params, returns Returns) *Leaf {
	return &Leaf{
		BaseNode: newBaseNode(CategoryLeaf, name),
		Params:   params,  // TODO (remove): These are only used for String()
		Returns:  returns, // TODO (remove): These are only used for String()
	}
}

// GetChildren returns an empty list of Node, since a leaf has no children.
// This method is required for Leaf in order to implement IBase.
func (a *Leaf) GetChildren() []Node {
	return []Node{}
}

// String returns a string representation of the leaf node.
func (a *Leaf) String() string {
	return fmt.Sprintf("! %s (%v : %v)",
		a.name,
		a.Params,
		a.Returns,
	)
}
