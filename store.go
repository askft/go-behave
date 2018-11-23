package behaviortree

// Context is renewed every time the tree is run.
type Context struct {
	tree  *BehaviorTree
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

// Blackboard ...
type Blackboard struct {
	data map[string]interface{}
}

// NewBlackboard instantiates a new blackboard
func NewBlackboard() *Blackboard {
	bb := &Blackboard{}
	bb.data = map[string]interface{}{}
	return bb
}

func (bb *Blackboard) Read(id string) (interface{}, bool) {
	value, ok := bb.data[id]
	return value, ok
}

func (bb *Blackboard) Write(id string, data interface{}) {
	bb.data[id] = data
}
