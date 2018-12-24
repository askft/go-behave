package core

// Category denotes whether a node is a composite, decorator or leaf.
type Category string

// A list of behavior tree node categories.
const (
	CategoryInvalid   = Category("invalid")
	CategoryComposite = Category("composite")
	CategoryDecorator = Category("decorator")
	CategoryLeaf      = Category("leaf")
)

// Status denotes the return value of the execution of a node.
type Status int

// A list of possible statuses.
const (
	StatusInvalid Status = iota
	StatusSuccess
	StatusFailure
	StatusRunning
)

type (
	// CompositeFn ...
	CompositeFn func(...INode) INode

	// DecoratorFn ...
	DecoratorFn func(Params, INode) INode

	// ActionFn ...
	ActionFn func([]string, []string) INode
)

type (
	// Params ...
	Params map[string]string

	// Returns ...
	Returns map[string]string
)
