package main
import (
	"fmt"
)
func main(){
	// Slice created with built in make function -> used to create dynamically sized array.
	// Make function allocates zeroed array amd returns slices that refers to the array.
	a := make([] int,5) 
	fmt.Println(a,len(a),cap(a)) // [0 0 0 0 0] 5 5
	b := make([] int,0,5)
	fmt.Println(b,len(b),cap(b)) // [] 0 5
	c := b[:2]
	fmt.Println(c,len(c),cap(c)) // [0 0] 2 5
	d := a[:2]
	fmt.Println(d,len(d),cap(d)) // [0 0] 2 5
}