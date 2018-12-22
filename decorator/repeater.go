package decorator

// import (
// 	"fmt"

// 	"github.com/alexanderskafte/behaviortree/core"
// )

// // Repeater runs its child until it returns core.StatusRunning.
// // Repeat at most n times, or unlimited if n == 0 (default).
// type Repeater struct {
// 	*core.Decorator
// 	n int
// }

// // Initialize ...
// func (d *Repeater) Initialize(args ...interface{}) {
// 	d.Decorator = args[0].(*core.Decorator)
// }

// // Start ...
// func (d *Repeater) Start(ctx *core.Context) {
// 	d.n = 0
// }

// // Tick ... TODO
// func (d *Repeater) Tick(ctx *core.Context) core.Status {
// 	fmt.Println("Run Repeater")
// 	i := 0
// 	status := core.StatusSuccess
// 	for ; d.n == 0 || i < d.n; i++ {
// 		status = core.Update(d.Child, ctx)
// 		if status == core.StatusRunning {
// 			break
// 		}
// 	}
// 	return status
// }

// // Stop ...
// func (d *Repeater) Stop(ctx *core.Context) {}
