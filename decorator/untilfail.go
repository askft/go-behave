package decorator

import (
	"fmt"

	. "github.com/alexanderskafte/behaviortree/core"
)

// UntilFail ...
type UntilFail struct {
	Decorator
}

// Tick ... TODO
func (d *UntilFail) Tick(ctx *Context) Status {
	fmt.Println("Run UntilFail")
	if Update(d.Child, ctx) != StatusFailure {
		return StatusRunning
	}
	return StatusSuccess
}
