package core

import (
	"fmt"
	"strings"
)

// Action ...
type Action struct {
	*Node
	Function string
	In, Out  []string
}

// NewAction ...
func NewAction(fn string, in, out []string) *Action {
	return &Action{
		Node:     NewNode(CategoryLeaf),
		Function: fn,
		In:       in,
		Out:      out,
	}
}

// GetChildren returns an empty list of INode, since a leaf has no children.
func (a *Action) GetChildren() []INode {
	return []INode{}
}

// String returns a string representation of the action node.
func (a Action) String() string {
	return fmt.Sprintf("%s (%s : %s)", a.Function,
		strings.Join(a.In, ", "),
		strings.Join(a.Out, ", "),
	)
}
