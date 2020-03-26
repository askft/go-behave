package decorator

import "github.com/askft/go-behave/core"

// While node repeats the conditions and runs the action if the condition succeeds.
// The action is started after the first success of the condition.
// While succeeds if the action succeeds and fails when any child fails.
// While returns Running if any child is running.
func While(params core.Params, cond, action core.Node) core.Node {

	base := core.NewDecorator("While", params, cond)
	d := &while{
		Decorator: base,
		action:    action, // action to run after condition succeeds
	}
	return d
}

type while struct {
	*core.Decorator
	action core.Node
}

func (d *while) Enter(ctx *core.Context) {

}

func (d *while) Tick(ctx *core.Context) core.Status {

	// check the condition
	status := core.Update(d.Child, ctx)

	switch status {
	case core.StatusRunning:
		return core.StatusRunning
	case core.StatusFailure:
		return core.StatusFailure
	case core.StatusInvalid:
		return core.StatusInvalid
	}

	// here condition is successful
	actionStatus := core.Update(d.action, ctx)

	return actionStatus

}

func (d *while) Leave(ctx *core.Context) {}
