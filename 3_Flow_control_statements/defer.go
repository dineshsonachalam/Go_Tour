package main

import (
	"fmt"
)

func main() {
	// defer statement -> defers the execution of function until surrounding function returns.
	defer fmt.Println("world")
	fmt.Println("Hello")
}
