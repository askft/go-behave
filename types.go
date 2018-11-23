package behaviortree

// Category denotes whether a node is a composite, decorator or leaf.
type Category string

const (
	cComposite = Category("composite")
	cDecorator = Category("decorator")
	cLeaf      = Category("leaf")
)

// Type denotes the specific type of the node of any Category.
type Type string

const (
	tInvalid = Type("invalid")

	// Composite types
	tSequence = Type("Sequence")
	tSelector = Type("Selector")

	// Decorator types
	tInverter = Type("Inverter")

	// Action types
	// Condition types
)

// Status denotes the return value of the execution of a node.
type Status int

const (
	StatusInvalid Status = iota
	StatusSuccess
	StatusFailure
	StatusRunning
)
