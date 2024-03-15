package core

import (
	"fmt"
)

// Category denotes whether a node is a composite, decorator or leaf.
type Category string

// A list of behavior tree node categories.
const (
	CategoryInvalid   = Category("invalid")
	CategoryComposite = Category("composite")
	CategoryDecorator = Category("decorator")
	CategoryLeaf      = Category("leaf")
)

type NodeResult interface {
	Status() Status
}

// Status denotes the return value of the execution of a node.
type Status int

// A list of possible statuses.
const (
	StatusInvalid Status = iota
	StatusInitialized
	StatusSuccess
	StatusFailure
	StatusRunning
	StatusError
)

func (s Status) Status() Status { return s }

type NodeAsyncRunning func(enqueue func(Event) error) error

func (NodeAsyncRunning) Status() Status { return StatusRunning }

type NodeRuntimeError struct{ error }

func (NodeRuntimeError) Status() Status { return StatusError }

type (
	// Params denotes a list of parameters to a node.
	Params map[string]interface{}

	// Returns is just a type alias for Params.
	Returns = Params
)

func (p Params) GetInt(key string) (int, error) {
	val, ok := p[key]
	if !ok {
		return 0, ErrParamNotFound(key)
	}
	n, ok := val.(int)
	if !ok {
		return 0, ErrInvalidType(key)
	}
	return n, nil
}

func (p Params) GetString(key string) (string, error) {
	val, ok := p[key]
	if !ok {
		return "", ErrParamNotFound(key)
	}
	s, ok := val.(string)
	if !ok {
		return "", ErrInvalidType(key)
	}
	return s, nil
}

func ErrParamNotFound(name string) error {
	return fmt.Errorf("parameter %s not found", name)
}

func ErrInvalidType(name string) error {
	return fmt.Errorf("invalid type for %s", name)
}
