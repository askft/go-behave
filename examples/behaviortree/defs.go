package main

// I use dot imports to make the tree definition look nicer. Sorry!
import (
	. "github.com/alexanderskafte/behaviortree/action"
	. "github.com/alexanderskafte/behaviortree/composite"
	"github.com/alexanderskafte/behaviortree/core"
	. "github.com/alexanderskafte/behaviortree/decorator"
)

// The trees below are equivalent.

var someTreeStr = `
* Repeater (n = #2) {
	+ Sequence {
		* Delayer (ms = #700) {
			! Succeed (:)
		}
		* Delayer (ms = #400) {
			! Succeed (:)
		}
	}
}
`

var someRoot = Repeater(core.Params{"n": "2"},
	Sequence(
		Delayer(core.Params{"ms": "700"},
			Succeed(nil, nil),
		),
		Delayer(core.Params{"ms": "400"},
			Succeed(nil, nil),
		),
	),
)

// An example of a tree for an entity that will attack the nearest target.
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
