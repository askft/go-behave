package behave

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/jbcpollak/greenstalk/core"
	"github.com/jbcpollak/greenstalk/util"

	// Use dot imports to make a tree definition look nice.
	// Be careful when doing this! These packages export
	// common word identifiers such as "Fail" and "Sequence".
	. "github.com/jbcpollak/greenstalk/common/action"
	. "github.com/jbcpollak/greenstalk/common/composite"
	. "github.com/jbcpollak/greenstalk/common/decorator"
)

// var delayingRoot = Repeater(core.Params{"n": 2},
// 	PersistentSequence(
// 		Delayer(core.Params{"ms": 700},
// 			Succeed(nil, nil),
// 		),
// 		Delayer(core.Params{"ms": 400},
// 			Inverter(nil,
// 				Fail(nil, nil),
// 			),
// 		),
// 	),
// )

type TestBlackboard struct {
	id    int
	count uint
}

type Event struct {
	id string
}

// Counter increments a counter on the blackboard.
func Counter(params core.Params, returns core.Returns) core.Node[TestBlackboard] {
	base := core.NewLeaf[TestBlackboard]("Counter", params, returns)
	return &counter{Leaf: base}
}

// succeed ...
type counter struct {
	core.Leaf[TestBlackboard]
}

// Enter ...
func (a *counter) Enter(bb TestBlackboard) {}

// Tick ...
func (a *counter) Tick(bb TestBlackboard, ctx context.Context, evt core.Event) core.NodeResult {
	fmt.Println("Counter: Incrementing count")
	bb.count++
	return core.StatusSuccess
}

// Leave ...
func (a *counter) Leave(bb TestBlackboard) {}

var synchronousRoot = Sequence[TestBlackboard](
	Repeater(core.Params{"n": 2}, Fail[TestBlackboard](nil, nil)),
	Succeed[TestBlackboard](nil, nil),
)

func TestUpdate(t *testing.T) {
	fmt.Println("Testing synchronous tree...")

	// Synchronous, so does not need to be cancelled.
	ctx := context.Background()

	tree, err := NewBehaviorTree(ctx, synchronousRoot, TestBlackboard{id: 42})
	if err != nil {
		panic(err)
	}

	for {
		evt := Event{}
		status := tree.Update(evt)
		util.PrintTreeInColor(tree.Root)
		fmt.Println()
		if status == core.StatusSuccess {
			break
		}
		time.Sleep(100 * time.Millisecond)
	}

	fmt.Println("Done!")
}

var asynchronousRoot = Sequence[TestBlackboard](
	// Repeater(core.Params{"n": 2}, Fail[TestBlackboard](nil, nil)),
	AsyncDelayer[TestBlackboard](
		core.Params{
			"label": "First",
			"ms":    100,
		},
		Counter(nil, nil),
	),
	AsyncDelayer[TestBlackboard](
		core.Params{
			"label": "Second",
			"ms":    100,
		},
		Counter(nil, nil),
	),
)

func TestEventLoop(t *testing.T) {
	fmt.Println("Testing asynchronous tree...")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	bb := TestBlackboard{id: 42, count: 0}

	tree, err := NewBehaviorTree(ctx, asynchronousRoot, bb)
	if err != nil {
		panic(err)
	}

	evt := Event{"initial event"}
	go tree.EventLoop(evt)

	time.Sleep(50 * time.Millisecond)
	if bb.count != 0 {
		t.Errorf("Expected count to be 0, got %d", bb.count)
	}

	// Sleep a bit more
	time.Sleep(60 * time.Millisecond)
	if bb.count != 1 {
		t.Errorf("Expected count to be 1, got %d", bb.count)
	}

	// Shut it _down_
	cancel()

	// Ensure we shut down before the second tick
	if bb.count != 1 {
		t.Errorf("Expected to shut down before second tick but got %d", bb.count)
	}

	fmt.Println("Done!")
}
