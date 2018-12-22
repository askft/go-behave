package core

// Category denotes whether a node is a composite, decorator or leaf.
type Category string

// A list of behavior tree node categories.
const (
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

// // ActionFn ...
// type ActionFn func(*Context) command.Command

// // Type denotes the specific type of the node of any Category.
// type Type string

// // A list of behavior tree node types.
// const (
// 	TypeInvalid = Type("invalid")
// 	TypeCustom  = Type("Custom")

// 	// Composite types
// 	TypeSequence = Type("Sequence")
// 	TypeSelector = Type("Selector")

// 	// Decorator types
// 	TypeInverter = Type("Inverter")

// 	// Leaf types
// 	TypeCondition = Type("Condition")
// 	TypeAction    = Type("Action")
// )
