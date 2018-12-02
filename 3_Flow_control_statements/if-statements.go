package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) float64 {
	if x == 7 {
		fmt.Printf("The square of %v is %v", x, math.Sqrt(x))
	}
	return math.Sqrt(x)
}
func main() {
	fmt.Printf("The square root is %v", sqrt(49))
}
