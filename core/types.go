package core

// Category denotes whether a node is a composite, decorator or leaf.
type Category string

const (
	CategoryComposite = Category("composite")
	CategoryDecorator = Category("decorator")
	CategoryLeaf      = Category("leaf")
)

// Type denotes the specific type of the node of any Category.
type Type string

// A list of behavior tree node types.
const (
	TypeInvalid = Type("invalid")

	// Composite types
	TypeSequence = Type("Sequence")
	TypeSelector = Type("Selector")

	// Decorator types
	TypeInverter = Type("Inverter")

	// Action types
	// Condition types
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
