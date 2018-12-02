package main

import (
	"fmt"
)

func main() {
	// 1. Println can handle one or more arguments. It inserts a new new line after we use it.
	fmt.Println("hello world", 100, 200, 300)
	// 2. Print writes data to console without trailing new line.
	fmt.Print("Rock the world", 400, 500)
	fmt.Println()
	// 3. Printf method accepts a format string, we use codes "%s" and "%d" in string to indicate insertion point for values
	name := "Mike"
	weight := 65
	fmt.Printf("Hello I'm, %s and I my weight is %d", name, weight)
	fmt.Println()
	// %v ->can handle ints,bools, string and other values.
	fmt.Printf("Hey %v, your weight is %v", name, weight)
	fmt.Println()
	// 4. Sprintf it works the same way as fmt.Println, but returns a string
	result := fmt.Sprintf("I saw %v, and he weighed %v kgs", name, weight)
	fmt.Println(result)
}
