package decorator

import (
	"fmt"

	. "github.com/alexanderskafte/behaviortree/core"
)

// Repeater runs its child until it returns a StatusRunning.
// Repeat at most n times, or unlimited if n == 0 (default).
type Repeater struct {
	Decorator
	n int
}

// Tick ... TODO
func (d *Repeater) Tick(ctx *Context) Status {
	fmt.Println("Run Repeater")
	i := 0
	status := StatusSuccess
	for ; d.n == 0 || i < d.n; i++ {
		status = d.Child.Tick(ctx)
		if status == StatusRunning {
			break
		}
	}
	return status
}
