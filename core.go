package behaviortree

import (
	"fmt"
	"math/rand"
	"strings"
)

// https://www.gamasutra.com/blogs/ChrisSimpson/20140717/221339/Behavior_trees_for_AI_How_they_work.php
// TODO: Decorator section at https://github.com/libgdx/gdx-ai/wiki/Behavior-Trees

// BehaviorTree ...
type BehaviorTree struct {
	Child INode
}

// ParseTreeString takes a string defining a behavior tree
// and returns a new Behavior Tree object.
func NewBehaviorTree(treeString string) (*BehaviorTree, error) {
	return NewParser(strings.NewReader(treeString)).Parse()
}

func (bt *BehaviorTree) String() string {
	var b strings.Builder
	nodeRecurse(bt.Child, 0, &b)
	return b.String()
}

func nodeRecurse(node INode, level int, b *strings.Builder) {
	indent := strings.Repeat("  ", level)
	b.WriteString(indent + string(node.GetType()))
	if node.GetCategory() != cLeaf {
		b.WriteString(" {")
	}
	b.WriteString("\n")
	children := node.GetChildren()
	for _, child := range children {
		nodeRecurse(child, level+1, b)
	}
	if node.GetCategory() != cLeaf {
		b.WriteString(indent + "}\n")
	}
}

// Run ...
func (bt *BehaviorTree) Run(bb *Blackboard) Status {
	ctx := NewContext(bt, bb)
	status := bt.Child.Run(ctx)
	return status
}

// Node ---------------------------------------------------

// INode ...
type INode interface {
	Init()
	Run(*Context) Status
	GetChildren() []INode
	GetCategory() Category
	GetType() Type
	GetChan() chan Status
}

// Node ...
type Node struct {
	Category
	Type
	StatusChan chan Status
	// Data       map[string]interface{}
}

// GetCategory ...
func (n *Node) GetCategory() Category { return n.Category }

// GetType ...
func (n *Node) GetType() Type { return n.Type }

func (n *Node) GetChan() chan Status {
	return n.StatusChan
}

// Composite nodes ----------------------------------------

// Composite base type
type Composite struct {
	Node
	Children []INode
}

// NewComposite ...
func NewComposite() Composite {
	return Composite{
		Node: Node{
			Category:   cComposite,
			StatusChan: make(chan Status, 1),
			// Data:       map[string]interface{}{},
		},
		Children: []INode{},
	}
}

// Init ...
func (c *Composite) Init() {
	fmt.Println("init comp", c.Type, "children", c.Children)
}

// GetChildren returns a list containing the children of the composite node.
func (c *Composite) GetChildren() []INode {
	return append([]INode{}, c.Children...)
}

// AddChildren ...
func (c *Composite) AddChildren(children ...INode) {
	c.Children = append(c.Children, children...)
}

// Decorator nodes ----------------------------------------

// Decorator base type
type Decorator struct {
	Node
	Child INode
}

// NewDecorator ...
func NewDecorator() Decorator {
	return Decorator{
		Node: Node{
			Category:   cDecorator,
			StatusChan: make(chan Status, 1),
			// Data:       map[string]interface{}{},
		},
	}
}

// Init ...
func (d *Decorator) Init() {
	fmt.Println("Init deco", d.Type, "child", d.Child)
}

// GetChildren returns a list containing the only child of the decorator node.
func (d *Decorator) GetChildren() []INode {
	return append([]INode{}, d.Child)
}

// SetChild ...
func (d *Decorator) SetChild(child INode) {
	d.Child = child
}

// Leaf nodes ---------------------------------------------

// Leaf ...
type Leaf struct {
	Node
	Action string
}

// NewLeaf ...
func NewLeaf(t Type, action string) *Leaf {
	return &Leaf{
		Node: Node{
			Category:   cLeaf,
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

// GetChildren returns an empty list of INode, since a leaf has no children.
func (l *Leaf) GetChildren() []INode {
	return []INode{}
}

// Run ...
func (l *Leaf) Run(ctx *Context) Status {
	fmt.Println("Run leaf", l.Type, "action", l.Action)
	return StatusSuccess
}

// Utility functions --------------------------------------

func shuffle(nodes []INode) {
	rand.Shuffle(len(nodes), func(i, j int) {
		nodes[i], nodes[j] = nodes[j], nodes[i]
	})
}

// Other --------------------------------------------------

// type taskWait struct {
// 	duration time.Duration
// 	visited  bool
// 	timer    *time.Timer
// }

// func (t *taskWait) init() {
// 	t.timer = time.NewTimer(t.duration)
// }

// func (t *taskWait) Run(ctx *Context) Status {
// 	if !t.visited {
// 		t.visited = true
// 		t.timer = time.NewTimer(t.duration)
// 	}
// 	select {
// 	case <-t.timer.C:
// 		t.timer.Stop()
// 		t.visited = false
// 		fmt.Println("task done")
// 		return StatusSuccess
// 	default:
// 		return StatusRunning
// 	}
// }

// Functions ----------------------------------------------

// type actionFn func(in, out interface{}) Status

// type leaf struct {
// 	fn      actionFn
// 	in, out interface{}
// }

// func (l *leaf) Run(ctx *Context) Status {
// 	return l.fn(l.in, l.out)
// }

// type Action struct {
// 	leaf
// 	data interface{}
// }

// type condition struct {
// 	leaf
// 	data interface{}
// }
