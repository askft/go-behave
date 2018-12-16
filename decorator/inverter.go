package decorator

import (
	"fmt"

	. "github.com/alexanderskafte/behaviortree/core"
)

// Inverter ...
type Inverter struct {
	Decorator
}

// Tick ...
func (i *Inverter) Tick(ctx *Context) Status {
	fmt.Println("Run Inverter")
	switch Update(i.Child, ctx) {
	case StatusSuccess:
		return StatusFailure
	case StatusFailure:
		return StatusSuccess
	default:
		return StatusRunning
	}
}
