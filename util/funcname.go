package util

import (
	"runtime"
	"strings"
)

// FuncName returns the name of the function that calls this function.
// This is useful when you want to define a constructor for a custom
// node without having to explicitly set its name.
// Example:
//  func SomeNode( args ) core.Node {
//  	name := util.FuncName() // will be "SomeNode"
//  	base := core.NewT(name) // T is Composite, Decorator or Action
//  	return &someNode{T: base}
//  }
func FuncName() string {
	pc, _, _, ok := runtime.Caller(1)
	if !ok {
		return "unknown"
	}
	me := runtime.FuncForPC(pc)
	if me == nil {
		return "unnamed"
	}
	tokens := strings.Split(me.Name(), ".")
	return tokens[len(tokens)-1]
}
