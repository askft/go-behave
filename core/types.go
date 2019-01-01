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
	// CompositeFn specifies the kind of a composite constructor function.
	CompositeFn = func(...Node) Node

	// DecoratorFn specifies the kind of a decorator constructor function.
	DecoratorFn = func(Params, Node) Node

	// ActionFn specifies the kind of a action constructor function.
	ActionFn = func(Params, Returns) Node
)

type (
	// Params ...
	Params map[string]string

	// Returns ...
	Returns map[string]string
)
