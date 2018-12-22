package action

import (
	"github.com/alexanderskafte/behaviortree/core"
)

// Succeed ...
type Succeed struct {
	*core.Action
}

// Initialize ...
func (a *Succeed) Initialize(args ...interface{}) {
	a.Action = args[0].(*core.Action)
}

// Start ...
func (a *Succeed) Start(ctx *core.Context) {}

// Tick ...
func (a *Succeed) Tick(ctx *core.Context) core.Status {
	return core.StatusSuccess
}

// Stop ...
func (a *Succeed) Stop(ctx *core.Context) {}
