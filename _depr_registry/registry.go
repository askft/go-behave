package deprregistry

// import (
// 	"fmt"
// 	"reflect"

// 	"github.com/alexanderskafte/behaviortree/action"
// 	"github.com/alexanderskafte/behaviortree/composite"
// 	"github.com/alexanderskafte/behaviortree/core"
// 	"github.com/alexanderskafte/behaviortree/decorator"
// )

// // Registry ...
// type Registry struct {
// 	structFor   map[string]reflect.Type
// 	categoryFor map[string]core.Category
// }

// // NewEmpty returns a new empty oldregistry.
// func NewEmpty() *Registry {
// 	return &Registry{
// 		structFor:   map[string]reflect.Type{},
// 		categoryFor: map[string]core.Category{},
// 	}
// }

// // NewDefault returns a structure of all
// // the nodes that are defined by default.
// func NewDefault() *Registry {
// 	r := NewEmpty()
// 	r.Register(core.CategoryComposite,
// 		&composite.RandomSelector{},
// 		&composite.RandomSequence{},
// 		&composite.Selector{},
// 		&composite.Sequence{},
// 	)
// 	r.Register(core.CategoryDecorator,
// 		&decorator.Delayer{},
// 		&decorator.Inverter{},
// 		&decorator.Repeater{},
// 		&decorator.UntilFailure{},
// 		&decorator.UntilSuccess{},
// 	)
// 	r.Register(core.CategoryLeaf,
// 		&action.Succeed{},
// 		&action.Fail{},
// 		&action.Print{},
// 	)
// 	return r
// }

// // Register ...
// func (r *Registry) Register(category core.Category, nodes ...core.ISpec) {
// 	for _, node := range nodes {
// 		r.registerOne(category, node)
// 	}
// }

// // Register ...
// func (r *Registry) registerOne(category core.Category, node core.ISpec) {
// 	elem := reflect.TypeOf(node).Elem()
// 	r.structFor[elem.Name()] = elem
// 	r.categoryFor[elem.Name()] = category
// }

// // New creates a new node struct pointer identified by `name`.
// func (r *Registry) New(name string) (interface{}, error) {
// 	if node, ok := r.structFor[name]; ok {
// 		return reflect.New(node).Interface(), nil
// 	}
// 	return nil, fmt.Errorf("Could not create node for name %s", name)
// }

// // Contains returns true if the node map contains a node identified by `name`.
// func (r *Registry) Contains(name string) bool {
// 	_, ok := r.structFor[name]
// 	return ok
// }

// // CategoryFor returns the category to which the type `name` belongs.
// func (r *Registry) CategoryFor(name string) (core.Category, error) {
// 	category, ok := r.categoryFor[name]
// 	if !ok {
// 		return core.CategoryInvalid, fmt.Errorf("%q not found in oldregistry", name)
// 	}
// 	return category, nil
// }

// // Sugar, baby!

// // TODO - Change to return ISpec instead of INode.

// // TODO - All of these are actually the same...

// // TODO - Initialize should return error

// // NewComposite ...
// func (r Registry) NewComposite(name string, base core.IBase) (core.INode, error) {
// 	any, err := r.New(name)
// 	if err != nil {
// 		return nil, err
// 	}
// 	spec := any.(core.INode)
// 	spec.Initialize(base)
// 	return spec, err
// }

// // NewDecorator ...
// func (r Registry) NewDecorator(name string, base core.IBase) (core.INode, error) {
// 	any, err := r.New(name)
// 	if err != nil {
// 		return nil, err
// 	}
// 	spec := any.(core.INode)
// 	spec.Initialize(base)
// 	return spec, err
// }

// // NewAction ...
// func (r Registry) NewAction(name string, base core.IBase) (core.INode, error) {
// 	any, err := r.New(name)
// 	if err != nil {
// 		return nil, err
// 	}
// 	spec := any.(core.INode)
// 	spec.Initialize(base)
// 	return spec, err
// }
