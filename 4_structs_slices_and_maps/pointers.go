// Pointer stores the address of another variable.
// & operator -> Generates pointer to its operand
// * operator -> Denotes pointers to underlying value.
package main

import (
	"fmt"
)

func main() {
	i := 42
	p := &i
	fmt.Printf("Address of i:%v and value of i:%v", p, *p)
	fmt.Println()
	*p = 100
	fmt.Printf("Address of i:%v and address of p:%v", &i, p)
	fmt.Println()
	fmt.Printf("New value of i: %v", i)
}
