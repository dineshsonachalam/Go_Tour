package main
import (
	"fmt"
)
type Vertex struct{
	Lat, Long float64
}
func main() {
	var m = map[string]Vertex{
		"Bell": Vertex {1.23,2.45,},
		"Caterpillar":Vertex {3.14,4.12},
	}
	fmt.Println(m)
}
// map[Bell:{1.23 2.45} Caterpillar:{3.14 4.12}]