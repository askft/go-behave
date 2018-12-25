// Package store specifies the behavior required from a data storage
// system for use in a behavior tree. The package also provides a
// type that implements the basic interface, the Blackboard type.
package store

// Interface for a store. A type that implements the interface
// can be used for reading and writing data.
type Interface interface {
	Read(string) (interface{}, error)
	Write(string, interface{})
}
