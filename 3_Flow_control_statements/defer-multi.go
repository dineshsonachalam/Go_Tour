package main

import (
	"fmt"
)

func main() {
	// Deferred function calls are pushed onto a stack. When function returns, deferred calls executed in last-in-first-out order.
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
}
