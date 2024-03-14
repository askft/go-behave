package util

import (
	"fmt"
	"strings"

	"github.com/jbcpollak/go-behave/core"

	"github.com/fatih/color"
)

// NodeToString returns a string representation
// of a tree node and all its children.
func NodeToString[Context any](node core.Node[Context]) string {
	var b strings.Builder
	fmt.Println()
	nodeRecurse(node, 0, &b)
	return b.String()
}

func nodeRecurse[Context any](node core.Node[Context], level int, b *strings.Builder) {
	indent := strings.Repeat("    ", level)
	b.WriteString(indent + node.String() + "\n")
	for _, child := range node.GetChildren() {
		nodeRecurse(child, level+1, b)
	}
}

// PrintTreeInColor prints the tree with colors representing node state.
//
// Red = Failure, Yellow = Running, Green = Success, Magenta = Invalid.
func PrintTreeInColor[Context any](node core.Node[Context]) {
	printTreeInColor(node, 0)
}

func printTreeInColor[Context any](node core.Node[Context], level int) {
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
