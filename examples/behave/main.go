package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/askft/go-behave"
	"github.com/askft/go-behave/common"
	"github.com/askft/go-behave/core"
	"github.com/askft/go-behave/gbl"
	"github.com/askft/go-behave/store"
	"github.com/askft/go-behave/util"

	// Use dot imports to make a tree definition look nice.
	// Be careful when doing this! These packages export
	// common word identifiers such as "Fail" and "Sequence".
	. "github.com/askft/go-behave/common/action"
	. "github.com/askft/go-behave/common/composite"
	. "github.com/askft/go-behave/common/decorator"
)

// The two trees below are equivalent.

// rootGBL defines a node structure in GBL code.
var rootGBL = `
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

// rootGo defines a node structure directly in Go code.
var someRoot = Repeater(core.Params{"n": 2},
	Sequence(
		Delayer(core.Params{"ms": 700},
			Succeed(nil, nil),
		),
		Delayer(core.Params{"ms": 400},
			Succeed(nil, nil),
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
	testScanner()
	testParser()
	testTree(someRoot)
	testCircularDependency()
}

func testCircularDependency() {

}

func testScanner() {
	fmt.Println("Testing scanner...")
	r := strings.NewReader(rootGBL)
	s := gbl.NewScanner(r)

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
	reg := common.NodeRegistry()
	node, err := gbl.NewParser(reg).Compile(rootGBL)
	if err != nil {
		panic(err)
	}
	fmt.Println(util.NodeToString(node))
	fmt.Println("Done parsing!")
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
