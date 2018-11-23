package behaviortree

// Context is renewed every time the tree is run.
type Context struct {
	tree  *BehaviorTree // TODO: Remove? Is this really needed?
	board *Blackboard
}

// NewContext ...
func NewContext(tree *BehaviorTree, board *Blackboard) *Context {
	if tree == nil {
		panic("tree is nil")
	}
	if board == nil {
		panic("board is nil")
	}
	return &Context{tree, board}
}
