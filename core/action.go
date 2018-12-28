package core

import (
	"fmt"
)

// Action is the base type for any specific action node (domain-specific).
// Each action node has Params: data keys that the implementation imports
// and Returns: data keys that the implementation exports.
type Action struct {
	*BaseNode
	Params  []string
	Returns []string
}

// NewAction creates a new action base node.
func NewAction(name string, params, returns []string) *Action {
	return &Action{
		BaseNode: newBaseNode(CategoryLeaf, name),
		Params:   params,
		Returns:  returns,
	}
}

// GetChildren returns an empty list of Node, since a leaf has no children.
// This method is required for Action in order to implement IBase.
func (a *Action) GetChildren() []Node {
	return []Node{}
}

// String returns a string representation of the action node.
func (a *Action) String() string {
	return fmt.Sprintf("! %s (%v : %v)",
		a.name,
		a.Params,
		a.Returns,
	)
}
