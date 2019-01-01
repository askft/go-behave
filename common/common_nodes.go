// Package common provides a registry containing number of useful
// predefined nodes. Subpackages provides the nodes themselves.
package common

import (
	"github.com/askft/go-behave/common/action"
	"github.com/askft/go-behave/common/composite"
	"github.com/askft/go-behave/common/decorator"
	"github.com/askft/go-behave/core"
	"github.com/askft/go-behave/gbl"
)

// NodeRegistry returns a Registry with
// a set of predefined nodes registered.
func NodeRegistry() *gbl.Registry {
	r := gbl.NewRegistry()
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
