package core

type IBase interface {
	GetChildren() []INode
	GetCategory() Category
	GetStatus() Status
	SetStatus(Status)
	String() string
}

type ISpec interface {
	Initialize(args ...interface{})
	Start(*Context)
	Tick(*Context) Status
	Stop(*Context)
}

type INode interface {
	IBase
	ISpec
}

// Node ...
type Node struct {
	Category
	Status
}

// NewNode ...
func NewNode(category Category) *Node {
	return &Node{Category: category}
}

// GetCategory returns the category of the node.
func (n *Node) GetCategory() Category {
	return n.Category
}

// GetStatus ...
func (n *Node) GetStatus() Status {
	return n.Status
}

// SetStatus ...
func (n *Node) SetStatus(status Status) {
	n.Status = status
}
