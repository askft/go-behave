package registry

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"

	"github.com/alexanderskafte/behaviortree/action"
	"github.com/alexanderskafte/behaviortree/composite"
	"github.com/alexanderskafte/behaviortree/core"
	"github.com/alexanderskafte/behaviortree/decorator"
)

// Registry ...
type Registry struct {
	// functionFor map[string]interface{}
	categoryFor map[string]core.Category

	cs map[string]core.CompositeFn
	ds map[string]core.DecoratorFn
	as map[string]core.ActionFn
}

// NewEmpty returns a new empty registry.
func NewEmpty() *Registry {
	return &Registry{
		// functionFor: map[string]interface{}{},
		categoryFor: map[string]core.Category{},
		cs:          map[string]core.CompositeFn{},
		ds:          map[string]core.DecoratorFn{},
		as:          map[string]core.ActionFn{},
	}
}

// NewDefault returns a structure of all
// the nodes that are defined by default.
func NewDefault() *Registry {
	r := NewEmpty()
	r.Register(core.CategoryComposite,
		composite.Sequence,
		composite.Selector,
		composite.RandomSequence,
		composite.RandomSelector,
	)
	r.Register(core.CategoryDecorator,
		decorator.Delayer,
		decorator.Inverter,
		decorator.Repeater,
	)
	r.Register(core.CategoryLeaf,
		action.Succeed,
		action.Fail,
	)
	return r
}

// Register registers each function in `fns` by its name. In other words,
// the registry maps function names to function handles.
func (r *Registry) Register(category core.Category, fns ...interface{}) {
	for _, fn := range fns {
		r.registerOne(category, fn)
	}
}

func (r *Registry) registerOne(category core.Category, fn interface{}) {

	// Get name of function
	fullName := runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name()
	tokens := strings.Split(fullName, ".")
	name := tokens[len(tokens)-1]

	switch category {
	case core.CategoryComposite:
		r.cs[name] = core.CompositeFn(fn.(func(...core.INode) core.INode))
	case core.CategoryDecorator:
		r.ds[name] = core.DecoratorFn(fn.(func(core.Params, core.INode) core.INode))
	case core.CategoryLeaf:
		r.as[name] = core.ActionFn(fn.(func([]string, []string) core.INode))
	default:
		panic(fmt.Errorf("invalid category (category = %s", category))
	}
	r.categoryFor[name] = category
}

// GetFunction ...
func (r *Registry) GetFunction(name string) (interface{}, error) {
	var fn interface{}
	var ok bool
	switch r.categoryFor[name] {
	case core.CategoryComposite:
		fn, ok = r.cs[name]
	case core.CategoryDecorator:
		fn, ok = r.ds[name]
	case core.CategoryLeaf:
		fn, ok = r.as[name]
	}
	if !ok {
		return nil, fmt.Errorf("function %s not found in registry", name)
	}
	return fn, nil
}

// CategoryFor returns the category to which the type `name` belongs.
func (r *Registry) CategoryFor(name string) (core.Category, error) {
	category, ok := r.categoryFor[name]
	if !ok {
		return core.CategoryInvalid, fmt.Errorf("%q not found in registry", name)
	}
	return category, nil
}
