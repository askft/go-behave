// Package gbl (short for "Go Behave Language") provides tools for
// compiling behavior trees defined in GBL syntax. Usage looks like
// the following:
//
//  var reg *gbl.Registry
//  var def string
//
//  // Initialize function registry and definition string.
//
//  root, err := gbl.NewParser(reg).Compile(def)
//
// The returned node may later be used as a root node in a behavior tree.
package gbl
