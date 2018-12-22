package action

import (
	"github.com/alexanderskafte/behaviortree/core"
)

// Fail ...
type Fail struct {
	*core.Action
}

// Initialize ...
func (a *Fail) Initialize(args ...interface{}) {
	a.Action = args[0].(*core.Action)
}

// Start ...
func (a *Fail) Start(ctx *core.Context) {}

// Tick ...
func (a *Fail) Tick(ctx *core.Context) core.Status {
	return core.StatusFailure
}

// Stop ...
func (a *Fail) Stop(ctx *core.Context) {}
