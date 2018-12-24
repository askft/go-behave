package main

// https://www.gamasutra.com/blogs/ChrisSimpson/20140717/221339/Behavior_trees_for_AI_How_they_work.php
// TODO: Decorator section at https://github.com/libgdx/gdx-ai/wiki/Behavior-Trees
// Good tutorial: https://github.com/adnzzzzZ/blog/issues/3

import (
	"fmt"
	"strings"
	"time"

	"github.com/alexanderskafte/behaviortree"
	"github.com/alexanderskafte/behaviortree/core"
	"github.com/alexanderskafte/behaviortree/lang"
	"github.com/alexanderskafte/behaviortree/registry"
	"github.com/alexanderskafte/behaviortree/store"
	"github.com/alexanderskafte/behaviortree/util"
)

// ID is a simple type only used as tree owner for testing.
// In a real scenario, the owner would be an actual entity
// with some interesting stae and functionality.
type ID int

// String returns a string representation of ID.
func (id ID) String() string { return fmt.Sprint(int(id)) }

func main() {
	// testScanner()
	// testParser()
	testTree(someRoot)
}

func testTree(root core.INode) {
	fmt.Println("Testing tree...")

	tree, err := behaviortree.NewBehaviorTree(
		behaviortree.Config{
			Owner:      ID(1337),
			Store:      store.NewBlackboard(),
			FnRegistry: registry.NewDefault(),
			Definition: someTreeStr,
		},
	)
	if err != nil {
		panic(err)
	}

	if root != nil {
		fmt.Println("Using root created in Go code.")
		tree.Root = root
	}

	ticker := time.NewTicker(100 * time.Millisecond)
	for {
		status := tree.Update()
		select {
		case <-ticker.C:
			util.PrintTreeInColor(tree.Root, 0)
			fmt.Println()
		default:
		}
		if status == core.StatusSuccess {
			break
		}
	}
	util.PrintTreeInColor(tree.Root, 0)

	fmt.Println("Done!")
}

func testScanner() {
	fmt.Println("Testing scanner...")
	r := strings.NewReader(someTreeStr)
	s := lang.NewScanner(r)

	for {
		tok, lit := s.Scan()
		if tok.IsEOF() {
			break
		}
		if tok.IsWhitespace() {
			continue
		}
		if tok.IsInvalid() {
			fmt.Printf("[ Invalid token %q ]\n", lit)
			continue
		}
		fmt.Printf("%-15s%s\n", tok, lit)
	}
	fmt.Println("Done scanning!")
}

func testParser() {
	fmt.Println("Testing parser...")
	node, err := lang.NewParser(registry.NewDefault()).Compile(someTreeStr)
	if err != nil {
		panic(err)
	}
	fmt.Println(behaviortree.NodeToString(node))
	fmt.Println("Done parsing!")
}
