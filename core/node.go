package core

// INode ...
type INode interface {
	Init()
	Tick(*Context) Status
	Terminate(*Context)
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

// GetCategory returns the category of the node.
func (n *Node) GetCategory() Category {
	return n.Category
}

// GetType returns the specific type of the node.
func (n *Node) GetType() Type {
	return n.Type
}

func (n *Node) GetChan() chan Status {
	return n.StatusChan
}
