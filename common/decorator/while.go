package decorator

import "github.com/jbcpollak/go-behave/core"

// While node repeats the conditions and runs the action if the condition succeeds.
// The action is started after the first success of the condition.
// While succeeds if the action succeeds and fails when any child fails.
// While returns Running if any child is running.
//
// An alternative implementation of the While behavior is:
//
//	UntilFailure {
//	    Sequence {
//	        Condition (custom node)
//	        Action (custom node)
//	    }
//	}
//
// This also allows you to have multiple conditions and multiple actions (just put them after each other in the sequence).
//
// This implementation is taken from https://github.com/DanTulovsky/go-behave/blob/master/common/decorator/while.go
// See also https://github.com/askft/go-behave/pull/2
func While[Blackboard any, Event any](params core.Params, cond, action core.Node[Blackboard, Event]) core.Node[Blackboard, Event] {

	base := core.NewDecorator("While", params, cond)
	d := &while[Blackboard, Event]{
		Decorator: base,
		action:    action, // action to run after condition succeeds
	}
	return d
}

type while[Blackboard any, Event any] struct {
	*core.Decorator[Blackboard, Event]
	action core.Node[Blackboard, Event]
}

func (d *while[Blackboard, Event]) Enter(bb Blackboard) {

}

func (d *while[Blackboard, Event]) Tick(bb Blackboard, evt Event) core.NodeResult {

	// check the condition
	status := core.Update(d.Child, bb, evt)

	switch status {
	case core.StatusRunning:
		return core.StatusRunning
	case core.StatusFailure:
		return core.StatusFailure
	case core.StatusInvalid:
		return core.StatusInvalid
	}

	// here condition is successful
	actionStatus := core.Update(d.action, bb, evt)

	return actionStatus

}

func (d *while[Blackboard, Event]) Leave(bb Blackboard) {}
