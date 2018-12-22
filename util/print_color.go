package util

// PrintTreeInColor prints the tree with colors representing node state.
// Red: Failure  |  Yellow: Running  |  Green: Success  |  Magenta: Invalid.
func PrintTreeInColor(node core.INode, level int) {
	indent := strings.Repeat("    ", level)

	col := colorFor(node.GetStatus())
	color.Set(col)

	if node.GetCategory() == core.CategoryLeaf {
		fmt.Printf(indent + node.String())
		color.Unset()
	} else {
		name := reflect.TypeOf(node).Elem().Name()
		fmt.Printf(indent + name)
		color.Unset()
		fmt.Printf(" {")
	}
	fmt.Printf("\n")
	children := node.GetChildren()
	for _, child := range children {
		PrintTreeInColor(child, level+1)
	}
	if node.GetCategory() != core.CategoryLeaf {
		fmt.Printf(indent + "}\n")
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
