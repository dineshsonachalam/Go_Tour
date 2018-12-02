package main

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Printf("Type:%T Value:%v\n", ToBe, ToBe)
}

// Here %T -> Type of a value.
// %V -> Value in default format.
