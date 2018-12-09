package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	response, err := http.Get("https://gist.githubusercontent.com/dineshsonachalam/4533c711e8c41e83ca379e231a8eeeb7/raw/081aa24a0d5d2fd63c8401e1f34294d5b1330abc/data.json")
	if err != nil {
		fmt.Println("There is an error!")
	}
	defer response.Body.Close()

	var view []interface{}
	json.NewDecoder(response.Body).Decode(&view)
	log.Printf(" [x] Pulled JSON: %s", view)
	fmt.Println("---------------------------------------------")
	for row_counter, data := range view {
		// md, _ := data.(map[string]interface{})
		// fmt.Println(md["shortname"])
		// _ = row_counter
		_ = row_counter
		fmt.Println(data)
		// for key, val := range data {
		// 	fmt.Println(key, ":", val)
		// }

	}

}
