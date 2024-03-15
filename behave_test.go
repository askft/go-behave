package behave

import (
	"fmt"
	"testing"
	"time"

	"github.com/jbcpollak/go-behave/core"
	"github.com/jbcpollak/go-behave/util"

	// Use dot imports to make a tree definition look nice.
	// Be careful when doing this! These packages export
	// common word identifiers such as "Fail" and "Sequence".
	. "github.com/jbcpollak/go-behave/common/action"
	. "github.com/jbcpollak/go-behave/common/composite"
	. "github.com/jbcpollak/go-behave/common/decorator"
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
	id int
}

type Event struct {
	id string
}

var synchronousRoot = Sequence[TestBlackboard](
	Repeater(core.Params{"n": 2}, Fail[TestBlackboard](nil, nil)),
	Succeed[TestBlackboard](nil, nil),
)

func TestUpdate(t *testing.T) {
	fmt.Println("Testing tree...")

	tree, err := NewBehaviorTree(TestBlackboard{id: 42}, synchronousRoot)
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
		core.Params{"ms": 1000},
		Succeed[TestBlackboard](nil, nil),
	),
)

func TestEventLoop(t *testing.T) {
	fmt.Println("Testing tree...")

	tree, err := NewBehaviorTree(TestBlackboard{id: 42}, asynchronousRoot)
	if err != nil {
		panic(err)
	}

	evt := Event{"initial event"}
	go tree.EventLoop(evt)

	time.Sleep(2 * time.Minute)

	fmt.Println("Done!")
}
