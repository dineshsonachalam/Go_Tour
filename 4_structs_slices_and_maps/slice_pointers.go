package main
import (
		"fmt"
)
func main(){
	names := [4]string{"1","2","3","4"}
	a:=names[0:2]
	b:=names[1:3]
	b[0] ="XXX"
	fmt.Println(a,b)
}