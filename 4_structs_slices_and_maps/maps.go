package main
import (
	"fmt"
)
type Vertex struct{
	Lat,Long float64
}
func main(){
	m:= make(map[string]Vertex)
	m["bell"] = Vertex{ 1.27, 3.14}
	fmt.Println(m)
}
// map[bell:{1.27 3.14}]