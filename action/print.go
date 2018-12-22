package action

import (
	"fmt"

	"github.com/alexanderskafte/behaviortree/core"
)

// Print ...
type Print struct {
	*core.Action
}

// Initialize ...
func (a *Print) Initialize(args ...interface{}) {
	a.Action = args[0].(*core.Action)
}

// Start ...
func (a *Print) Start(ctx *core.Context) {}

// Tick ...
func (a *Print) Tick(ctx *core.Context) core.Status {
	fmt.Printf("Print: %s, %s\n", a.In, a.Out)
	return core.StatusSuccess
}

// Stop ...
func (a *Print) Stop(ctx *core.Context) {}
