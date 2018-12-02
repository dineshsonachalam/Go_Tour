package main

import (
	"fmt"
)

const Pi = 3.14

func main() {
	fmt.Printf("Pi is %v", Pi)
	fmt.Println()
	const s = "Hello world"
	const Truth = true
	fmt.Printf("%v %v", s, Truth)
}
