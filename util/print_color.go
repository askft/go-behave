package util

import (
	"fmt"
	"strings"

	"github.com/alexanderskafte/behaviortree/core"

	"github.com/fatih/color"
)

// PrintTreeInColor prints the tree with colors representing node state.
//
// Red = Failure, Yellow = Running, Green = Success, Magenta = Invalid.
func PrintTreeInColor(node core.INode, level int) {
	indent := strings.Repeat("    ", level)
	color.Set(colorFor(node.GetStatus()))
	fmt.Println(indent + node.String())
	color.Unset()
	for _, child := range node.GetChildren() {
		PrintTreeInColor(child, level+1)
	}
}

func colorFor(status core.Status) color.Attribute {
	switch status {
	case core.StatusFailure:
		return color.FgRed
	case core.StatusRunning:
		return color.FgYellow
	case core.StatusSuccess:
		return color.FgGreen
	case core.StatusInvalid:
		return color.FgMagenta
	default:
		panic("invalid color")
	}
}
