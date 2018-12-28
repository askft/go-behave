package decorator

// import (
// 	"fmt"

// 	"github.com/alexanderskafte/go-behave/core"
// )

// // UntilSuccess ...
// type UntilSuccess struct {
// 	*core.Decorator
// }

// // Initialize ...
// func (d *UntilSuccess) Initialize(args ...interface{}) {
// 	d.Decorator = args[0].(*core.Decorator)
// }

// // Start ...
// func (d *UntilSuccess) Start(ctx *core.Context) {}

// // Tick ...
// func (d *UntilSuccess) Tick(ctx *core.Context) core.Status {
// 	fmt.Println("Run UntilSuccess")
// 	status := core.Update(d.Child, ctx)
// 	if status != core.StatusSuccess {
// 		return core.StatusRunning
// 	}
// 	return core.StatusSuccess
// }

// // Stop ...
// func (d *UntilSuccess) Stop(ctx *core.Context) {}
