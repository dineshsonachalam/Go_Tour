// Every program starts is made up of packages
// Program using packages with import paths "fmt" and "math/rand"
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("My favourite number is:", rand.Intn(10))
}
