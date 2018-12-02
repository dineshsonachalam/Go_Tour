// Variables declared without an explicit inital value are given their zero value.
// 0 -> for numeric values.
// false -> for boolean type.
// "" -> (the empty strings) for strings
package main

import "fmt"

func main() {
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q", i, f, b, s)
}

// %q	a single-quoted character literal safely escaped with Go syntax.
