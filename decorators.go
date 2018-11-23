package behaviortree

import (
	"fmt"
)

// Inverter ...
type Inverter struct{ Decorator }

// Run ...
func (i *Inverter) Run(ctx *Context) Status {
	fmt.Println("Run Inverter")
	switch i.Child.Run(ctx) {
	case StatusSuccess:
		return StatusFailure
	case StatusFailure:
		return StatusSuccess
	default:
		return StatusRunning
	}
}

// Repeater runs its child until it returns a StatusRunning.
// Repeat at most n times, or unlimited if n == 0 (default).
type Repeater struct {
	Decorator
	n int
}

// Run ... TODO
func (d *Repeater) Run(ctx *Context) Status {
	fmt.Println("Run Repeater")
	i := 0
	status := StatusSuccess
	for ; d.n == 0 || i < d.n; i++ {
		status = d.Child.Run(ctx)
		if status == StatusRunning {
			break
		}
	}
	return status
}

// UntilFail ...
type UntilFail struct{ Decorator }

// Run ... TODO
func (d *UntilFail) Run(ctx *Context) Status {
	fmt.Println("Run UntilFail")
	for {
		if d.Child.Run(ctx) != StatusFailure {
			break
		} else {
			return StatusRunning
		}
	}
	return StatusSuccess
}
