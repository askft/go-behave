package registry

import (
	"fmt"
	"reflect"

	"github.com/alexanderskafte/behaviortree/action"
	"github.com/alexanderskafte/behaviortree/composite"
	"github.com/alexanderskafte/behaviortree/condition"
	"github.com/alexanderskafte/behaviortree/core"
	"github.com/alexanderskafte/behaviortree/decorator"
)

// Registry ...
type Registry map[string]reflect.Type

// NewDefault returns a structure of all
// the nodes that are defined by default.
func NewDefault() Registry {
	nodes := Registry{}

	nodes.Register(&composite.Sequence{})
	nodes.Register(&composite.Selector{})
	nodes.Register(&composite.RandomSequence{})
	nodes.Register(&composite.RandomSelector{})

	nodes.Register(&decorator.Inverter{})
	nodes.Register(&decorator.UntilFailure{})
	nodes.Register(&decorator.UntilSuccess{})

	nodes.Register(&condition.TargetNearby{})

	nodes.Register(&action.Succeed{})
	nodes.Register(&action.Fail{})
	nodes.Register(&action.Print{})

	return nodes
}

// Register ...
func (r Registry) Register(node core.ISpec) {
	elem := reflect.TypeOf(node).Elem()
	r[elem.Name()] = elem
}

// New creates a new node struct pointer identified by `name`.
func (r Registry) New(name string) (interface{}, error) {
	if node, ok := r[name]; ok {
		return reflect.New(node).Interface(), nil
	}
	return nil, fmt.Errorf("Could not create node for name %s", name)
}

// Contains returns true if the node map contains a node identified by `name`.
func (r Registry) Contains(name string) bool {
	_, ok := r[name]
	return ok
}

// Sugar, baby!

// TODO - Change to return ISpec instead of INode.

// TODO - All of these are actually the same...

// NewComposite ...
func (r Registry) NewComposite(name string, base core.IBase) (core.INode, error) {
	any, err := r.New(name)
	if err != nil {
		return nil, err
	}
	spec := any.(core.INode)
	spec.Initialize(base)
	return spec, err
}

// NewDecorator ...
func (r Registry) NewDecorator(name string, base core.IBase) (core.INode, error) {
	any, err := r.New(name)
	if err != nil {
		return nil, err
	}
	spec := any.(core.INode)
	spec.Initialize(base)
	return spec, err
}

// NewAction ...
func (r Registry) NewAction(name string, base core.IBase) (core.INode, error) {
	any, err := r.New(name)
	if err != nil {
		return nil, err
	}
	spec := any.(core.INode)
	spec.Initialize(base)
	return spec, err
}
