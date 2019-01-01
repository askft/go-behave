package util

import (
	"fmt"
	"strings"

	"github.com/askft/go-behave/core"

	"github.com/fatih/color"
)

// NodeToString returns a string representation
// of a tree node and all its children.
func NodeToString(node core.Node) string {
	var b strings.Builder
	fmt.Println()
	nodeRecurse(node, 0, &b)
	return b.String()
}

func nodeRecurse(node core.Node, level int, b *strings.Builder) {
	indent := strings.Repeat("    ", level)
	b.WriteString(indent + node.String() + "\n")
	for _, child := range node.GetChildren() {
		nodeRecurse(child, level+1, b)
	}
}

// PrintTreeInColor prints the tree with colors representing node state.
//
// Red = Failure, Yellow = Running, Green = Success, Magenta = Invalid.
func PrintTreeInColor(node core.Node) {
	printTreeInColor(node, 0)
}

func printTreeInColor(node core.Node, level int) {
	indent := strings.Repeat("    ", level)
	color.Set(colorFor[node.GetStatus()])
	fmt.Println(indent + node.String())
	color.Unset()
	for _, child := range node.GetChildren() {
		printTreeInColor(child, level+1)
	}
}

var colorFor = map[core.Status]color.Attribute{
	core.StatusFailure: color.FgRed,
	core.StatusRunning: color.FgYellow,
	core.StatusSuccess: color.FgGreen,
	core.StatusInvalid: color.FgMagenta,
}
