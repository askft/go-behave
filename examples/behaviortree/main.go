package main

import (
	"fmt"

	bt "github.com/alexanderskafte/behaviortree"
)

// Example of how I want to be able to define the tree.
// '?' defines a condition, '!' an action.
var exampleTree = `
	Selector {
		Sequence {
			? TargetNearby
			! TargetSelect (-> t1)
			! TargetAttack (t1 ->)
		}
		RandomSelector {
			! Sleep
			! Smoke
		}
	}
`

// Main ...
func main() {
	fmt.Println("Started program!")

	tree, err := bt.NewBehaviorTree(exampleTree)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(tree)
	}

	blackboard := bt.NewBlackboard()
	tree.Run(blackboard)

	fmt.Println("Finished program!")
}
