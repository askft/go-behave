// Package common provides a registry containing number of useful
// predefined nodes. Subpackages provides the nodes themselves.
package common

import (
	"github.com/alexanderskafte/go-behave/common/action"
	"github.com/alexanderskafte/go-behave/common/composite"
	"github.com/alexanderskafte/go-behave/common/decorator"
	"github.com/alexanderskafte/go-behave/core"
	"github.com/alexanderskafte/go-behave/gbl"
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
