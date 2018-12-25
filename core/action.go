package core

import (
	"fmt"
)

// Action ...
type Action struct {
	*BaseNode
	Params  []string
	Returns []string
}

// NewAction ...
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
		a.Name,
		a.Params,
		a.Returns,
	)
}
