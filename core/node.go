package core

// INode ...
type INode interface {
	GetChildren() []INode
	GetCategory() Category
	GetStatus() Status
	SetStatus(Status)
	String() string

	Start(*Context)
	Tick(*Context) Status
	Stop(*Context)
}

// Node ...
type Node struct {
	Category
	Status
	Name string
}

// NewNode ...
func NewNode(category Category, name string) *Node {
	return &Node{Category: category, Name: name}
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
