// Type conversion: expression Type(value) -> converts value(v) to Type(T)
package main

import (
	"fmt"
)

func main() {
	i := 42 // int
	fmt.Printf("Type: %T Value: %v\n", i, i)
	f := float64(i) // converting int i to float
	fmt.Printf("Type: %T Value: %v\n", f, f)
}
