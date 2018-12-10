// slice literals is like array without length.
package main
import (
	"fmt"
)
func main(){
	// Declaring a slice
	s_slice := []int{1,2,3,4,5,6}
	// Declarting an array
	s_array := [3]int{1,2,3}
	fmt.Println(s_slice)
	fmt.Println(s_array)
	fmt.Printf("Length=%v and capacity=%v",len(s_slice),cap(s_slice))
	fmt.Println()
	s_slice = s_slice[:4]
	fmt.Printf("Length=%v and capacity=%v",len(s_slice),cap(s_slice))
	fmt.Println()

}