package main

import (
	"fmt"
	"time"

	"github.com/askft/go-behave"
	"github.com/askft/go-behave/core"
	"github.com/askft/go-behave/store"
	"github.com/askft/go-behave/util"

	// Use dot imports to make a tree definition look nice.
	// Be careful when doing this! These packages export
	// common word identifiers such as "Fail" and "Sequence".
	. "github.com/askft/go-behave/common/action"
	. "github.com/askft/go-behave/common/composite"
	. "github.com/askft/go-behave/common/decorator"
)

// someRoot defines a node structure.
var someRoot = Repeater(core.Params{"n": 2},
	Sequence(
		Delayer(core.Params{"ms": 700},
			Succeed(nil, nil),
		),
		Delayer(core.Params{"ms": 400},
			Inverter(nil,
				Fail(nil, nil),
			),
		),
	),
)

// ID is a simple type only used as tree owner for testing.
// In a real scenario, the owner would be an actual entity
// with some interesting state and functionality.
type ID int

// String returns a string representation of ID.
func (id ID) String() string { return fmt.Sprint(int(id)) }

func main() {
	testTree(someRoot)
}

func testTree(root core.Node) {
	fmt.Println("Testing tree...")

	tree, err := behave.NewBehaviorTree(
		behave.Config{
			Owner: ID(1337),
			Data:  store.NewBlackboard(),
			Root:  root,
		},
	)
	if err != nil {
		panic(err)
	}

	ticker := time.NewTicker(100 * time.Millisecond)
	for {
		status := tree.Update()
		select {
		case <-ticker.C:
			util.PrintTreeInColor(tree.Root)
			fmt.Println()
		default:
		}
		if status == core.StatusSuccess {
			break
		}
	}
	util.PrintTreeInColor(tree.Root)

	fmt.Println("Done!")
}

// An example of a tree for an entity that will attack the nearest target.
// Of course, you'd need to implement all of the leaf nodes here yourself.

// var attackBT =
// 		+ Sequence {
// 			? EnemyInAggroRange ( : target )
// 			! SetTarget ( target : )
// 			+ Selector {
// 				+ Sequence {
// 					? TargetInAttackRange (:)
// 					! Attack (:)
// 				}
// 				! MoveTowardTarget (:)
// 			}
// 		}
// 		`
