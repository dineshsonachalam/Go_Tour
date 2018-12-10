package main
import(
	"fmt"	
)
func main() {
	var s[] int
	printSlice(s)
	s = append(s,1)
	printSlice(s)
	s = append(s,2,3,4,5)
	printSlice(s)
}
func printSlice(s []int){
	fmt.Println(s,len(s),cap(s))
}
/*
	[] 0 0
	[1] 1 1
	[1 2 3 4 5] 5 6
*/