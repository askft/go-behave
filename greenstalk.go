package behave

import (
	"context"
	"fmt"

	"github.com/jbcpollak/greenstalk/core"
	"github.com/jbcpollak/greenstalk/internal"
	"github.com/jbcpollak/greenstalk/util"
	"github.com/rs/zerolog/log"
)

// BehaviorTree ...
type BehaviorTree[Blackboard any] struct {
	ctx        context.Context
	Root       core.Node[Blackboard]
	Blackboard Blackboard
	events     chan core.Event
}

// NewBehaviorTree returns a new BehaviorTree. A data Blackboard
// to be propagated down the tree each tick is created.
func NewBehaviorTree[Blackboard any](ctx context.Context, root core.Node[Blackboard], bb Blackboard) (*BehaviorTree[Blackboard], error) {
	var eb internal.ErrorBuilder
	eb.SetMessage("NewBehaviorTree")
	if root == nil {
		eb.Write("Config.Root is nil")
	}

	if eb.Error() != nil {
		return nil, eb.Error()
	}
	tree := &BehaviorTree[Blackboard]{
		ctx:        ctx,
		Root:       root,
		Blackboard: bb,
		events:     make(chan core.Event, 100 /* arbitrary */),
	}
	return tree, nil
}

// Update propagates an update call down the behavior tree.
func (bt *BehaviorTree[Blackboard]) Update(evt core.Event) core.Status {
	result := core.Update(bt.Root, bt.Blackboard, bt.ctx, evt)

	status := result.Status()
	switch status {
	case core.StatusSuccess:
		// whatever
	case core.StatusFailure:
		// whatever
	case core.StatusRunning:
		if asyncRunning, ok := result.(core.NodeAsyncRunning); ok {
			go asyncRunning(bt.ctx, func(evt core.Event) error {
				bt.events <- evt
				return nil
			})
		}
	case core.StatusError:
		if status, ok := result.(core.NodeRuntimeError); ok {
			panic(status.Err)
		}
	default:
		panic(fmt.Errorf("invalid status %v", status))
	}

	return status
}

func (bt *BehaviorTree[Blackboard]) EventLoop(evt core.Event) {
	defer close(bt.events)

	// Put the first event on the queue.
	bt.events <- evt

	for evt := range bt.events {
		log.Info().Msgf("Event: %v", evt)
		bt.Update(evt)
		util.PrintTreeInColor(bt.Root)
	}
}

// String creates a string representation of the behavior tree
// by traversing it and writing lexical elements to a string.
func (bt *BehaviorTree[Blackboard]) String() string {
	return util.NodeToString(bt.Root)
}
