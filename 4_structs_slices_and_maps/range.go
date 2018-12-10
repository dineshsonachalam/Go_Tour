package main
import (
	"fmt"
)
func main(){
	var s = [] string{"one","two"}
	for row_counter,data := range s{
		fmt.Println(row_counter,data)
	}
}
/*
0 one
1 two
*/