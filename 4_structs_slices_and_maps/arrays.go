package main
import (
	"fmt"
)
func main(){
	var a[2] string
	a[0] = "hello"
	a[1] = "world"
	primes := [3]string{"Tom","John","Mike"}
	fmt.Println(a)
	fmt.Println(primes)
	var s []string = primes[1:3]
	fmt.Println(s)
}