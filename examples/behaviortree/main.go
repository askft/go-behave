package main

import (
	"fmt"
	"strings"

	bt "behaviortree"
)

var exampleTree = `
    Sequence {
		Action a1.js
		Action a2.js
        Selector {
            Action a3.js
			Action a4.js
        }
    }`

// Main ...
func main() {
	fmt.Println("Started program!")

	tree, err := bt.NewParser(strings.NewReader(exampleTree)).Parse()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(tree)
	}

	blackboard := bt.NewBlackboard()
	tree.Run(blackboard)

	fmt.Println("Finished program!")
}
