package core

import (
	"fmt"
)

// Leaf ...
type Leaf struct {
	Node
	Action string
}

// NewLeaf ...
func NewLeaf(t Type, action string) *Leaf {
	return &Leaf{
		Node: Node{
			Category:   CategoryLeaf,
			Type:       t,
			StatusChan: make(chan Status, 1),
			// Data:       map[string]interface{}{},
		},
		Action: action,
	}
}

// Init ...
func (l *Leaf) Init() {
	fmt.Println("Init leaf", l.Type, "action", l.Action)
}

// Terminate ...
func (l *Leaf) Terminate(ctx *Context) {
	fmt.Println("Terminate leaf", l.Type, "action", l.Action)
}

// GetChildren returns an empty list of INode, since a leaf has no children.
func (l *Leaf) GetChildren() []INode {
	return []INode{}
}

// Tick ...
func (l *Leaf) Tick(ctx *Context) Status {
	fmt.Println("Run leaf", l.Type, "action", l.Action)
	return StatusSuccess
}
