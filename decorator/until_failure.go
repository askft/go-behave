package decorator

// import (
// 	"fmt"

// 	"github.com/alexanderskafte/go-behave/core"
// )

// // UntilFailure ...
// type UntilFailure struct {
// 	*core.Decorator
// }

// // Initialize ...
// func (d *UntilFailure) Initialize(args ...interface{}) {
// 	d.Decorator = args[0].(*core.Decorator)
// }

// // Start ...
// func (d *UntilFailure) Start(ctx *core.Context) {}

// // Tick ...
// func (d *UntilFailure) Tick(ctx *core.Context) core.Status {
// 	fmt.Println("Run UntilFailure")
// 	status := core.Update(d.Child, ctx)
// 	if status != core.StatusFailure {
// 		return core.StatusRunning
// 	}
// 	return core.StatusSuccess
// }

// // Stop ...
// func (d *UntilFailure) Stop(ctx *core.Context) {}
