// struct is a collection of fields

package main

import "fmt"

type vertex struct {
	x int
	y int
}

func main() {
	v := vertex{x: 1, y: 2}
	fmt.Println(v)
	fmt.Println(v.x)
	fmt.Println(v.y)
}
