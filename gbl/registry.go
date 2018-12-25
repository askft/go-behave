package gbl

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"

	"github.com/pkg/errors"

	"github.com/alexanderskafte/go-behave/core"
)

// Registry allows for registration of node constructor functions by name.
type Registry struct {
	categoryFor map[string]core.Category
	cs          map[string]core.CompositeFn
	ds          map[string]core.DecoratorFn
	as          map[string]core.ActionFn
}

// NewRegistry returns a new empty registry.
func NewRegistry() *Registry {
	return &Registry{
		categoryFor: map[string]core.Category{},
		cs:          map[string]core.CompositeFn{},
		ds:          map[string]core.DecoratorFn{},
		as:          map[string]core.ActionFn{},
	}
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
		r.cs[name] = fn.(core.CompositeFn)
	case core.CategoryDecorator:
		r.ds[name] = fn.(core.DecoratorFn)
	case core.CategoryLeaf:
		r.as[name] = fn.(core.ActionFn)
	default:
		panic(fmt.Errorf("invalid category (category = %s", category))
	}
	r.categoryFor[name] = category
}

// Merge creates a new registry that is the union of r and s.
// Returns the registry and an error if the same name referring
// to a function exists in both r and s. The registry is valid,
// but any name that occurs in both r and s refers to the function
// in s, as the function referred to in r is overwritten.
func (r *Registry) Merge(s *Registry) (*Registry, error) {
	t := NewRegistry()
	var err error
	e1 := t.merge(r)
	e2 := t.merge(s)
	err = errors.Wrap(e1, e1.Error())
	err = errors.Wrap(e2, e2.Error())
	return t, err
}

func (r *Registry) merge(s *Registry) error {
	var err error
	for name, category := range s.categoryFor {
		if r.contains(name) {
			msg := fmt.Sprintf(
				"registry already contains name %q for category %q",
				name, category,
			)
			if err == nil {
				err = errors.New(msg)
			} else {
				err = errors.Wrap(err, msg)
			}
		}
	}
	for name, fn := range s.cs {
		r.cs[name] = fn
	}
	for name, fn := range s.ds {
		r.ds[name] = fn
	}
	for name, fn := range s.as {
		r.as[name] = fn
	}
	return err
}

func (r *Registry) contains(name string) bool {
	_, ok := r.categoryFor[name]
	return ok
}

// Get returns the function and category for a name if it exists,
// otherwise an error.
func (r *Registry) Get(name string) (interface{}, core.Category, error) {
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
		return nil,
			core.CategoryInvalid,
			fmt.Errorf("function %s not found in registry", name)
	}
	return fn, r.categoryFor[name], nil
}

// TODO: Remove? This is unused and probably not necessary since
// a `core.Node` has a `GetCategory()` method.
// // CategoryFor returns the category to which the type `name` belongs.
// func (r *Registry) CategoryFor(name string) (core.Category, error) {
// 	category, ok := r.categoryFor[name]
// 	if !ok {
// 		return core.CategoryInvalid, fmt.Errorf("%q not found in registry", name)
// 	}
// 	return category, nil
// }

// TODO: Remove. Old code.
// keys1 := []string{}
// keys2 := []string{}
// for name := range r1.categoryFor {
// 	keys1 = append(keys1, name)
// }
// for name := range r2.categoryFor {
// 	keys2 = append(keys2, name)
// }
// for _, name := range keys1 {
// 	r.categoryFor[name] = r1.categoryFor[name]
// 	r.cs[name] = r1.cs[name]
// 	r.ds[name] = r1.ds[name]
// 	r.as[name] = r1.as[name]
// }
// for _, name := range keys2 {
// 	r.categoryFor[name] = r2.categoryFor[name]
// 	r.cs[name] = r2.cs[name]
// 	r.ds[name] = r2.ds[name]
// 	r.as[name] = r2.as[name]
// }

// r := NewRegistry()
// for name := range r1.categoryFor {
// 	r.categoryFor[name] = r1.categoryFor[name]
// 	r.cs[name] = r1.cs[name]
// 	r.ds[name] = r1.ds[name]
// 	r.as[name] = r1.as[name]
// }
// var err error
// for name := range r2.categoryFor {
// 	if r.contains(name) {
// 		msg := fmt.Sprintf("registry already contains name %q", name)
// 		if err == nil {
// 			err = errors.New(msg)
// 		} else {
// 			err = errors.Wrap(err, msg)
// 		}
// 	}
// 	r.categoryFor[name] = r2.categoryFor[name]
// 	r.cs[name] = r2.cs[name]
// 	r.ds[name] = r2.ds[name]
// 	r.as[name] = r2.as[name]
// }
// return r, err
