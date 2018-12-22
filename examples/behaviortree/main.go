package main

// https://www.gamasutra.com/blogs/ChrisSimpson/20140717/221339/Behavior_trees_for_AI_How_they_work.php
// TODO: Decorator section at https://github.com/libgdx/gdx-ai/wiki/Behavior-Trees
// Good tutorial: https://github.com/adnzzzzZ/blog/issues/3

import (
	"fmt"
	"strings"

	"github.com/alexanderskafte/behaviortree"
	"github.com/alexanderskafte/behaviortree/lang"
	"github.com/alexanderskafte/behaviortree/registry"
	"github.com/alexanderskafte/behaviortree/store"
)

var exampleTree = `
	+ RandomSequence {
		+ Sequence {
			! Succeed (asd : sdf )
		}
		* UntilSuccess {
			! Fail (qwe : wer) 
		}
	}
	`

type ID int

func (id ID) String() string { return fmt.Sprint(id) }

func main() {
	// testScanner()
	testParser()
	testTree()
}

func testTree() {
	fmt.Println("Testing tree...")

	tree, err := behaviortree.NewBehaviorTree(
		behaviortree.Config{
			Owner:      ID(0),
			Store:      store.NewBlackboard(),
			Registry:   registry.NewDefault(),
			Definition: exampleTree,
		})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(tree)

	for i := 0; i < 5; i++ {
		tree.Update()
		fmt.Println(tree)
	}
	fmt.Println(tree)

	fmt.Println("Done!")
}

func testScanner() {
	fmt.Println("Testing scanner...")
	r := strings.NewReader(exampleTree)
	s := lang.NewScanner(r)

	for {
		tok, lit := s.Scan()
		if lang.TokenIsEOF(tok) {
			break
		}
		if lang.TokenIsWhitespace(tok) {
			continue
		}
		if lang.TokenIsInvalid(tok) {
			fmt.Printf("[ Invalid token %q ]\n", lit)
			continue
		}
		fmt.Println(tok, "\t", lit)
	}
	fmt.Println("Done scanning!")
}

func testParser() {
	fmt.Println("Testing parser...")
	r := strings.NewReader(exampleTree)
	p := lang.NewParser(r, registry.NewDefault())
	node, err := p.Parse()
	if err != nil {
		panic(err)
	}
	fmt.Println(node)
	fmt.Println("Done parsing!")
}
