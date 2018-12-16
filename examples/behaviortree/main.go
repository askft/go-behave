package main

import (
	"fmt"
	"strings"

	bt "github.com/alexanderskafte/behaviortree"
)

// Example of how I want to be able to define the tree.
// '?' defines a condition, '!' an action.
var exampleTree = `
	Selector {
		Sequence {
			? TargetNearby
			! TargetSelect (   : t1)
			! TargetAttack (t1 :   )
		}
		RandomSelector {
			! Sleep
			! Smoke
		}
	}
	`

func main() {
	testScanner()
}

func testScanner() {
	fmt.Println("Testing scanner...")
	r := strings.NewReader(exampleTree)
	s := bt.NewScanner(r)

	for {
		tok, lit := s.Scan()
		if bt.TokenIsEOF(tok) {
			break
		}
		if bt.TokenIsWhitespace(tok) {
			continue
		}
		if bt.TokenIsInvalid(tok) {
			fmt.Printf("[ Invalid token %q ]\n", lit)
			continue
		}
		fmt.Println(tok, lit)
	}
	fmt.Println("Done!")
}

func testTree() {
	fmt.Println("Testing tree...")

	tree, err := bt.NewBehaviorTree(exampleTree)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(tree)
	}

	blackboard := bt.NewBlackboard()
	tree.Tick(blackboard)

	fmt.Println("Done!")
}
