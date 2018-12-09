package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	response, err := http.Get("https://gist.githubusercontent.com/dineshsonachalam/4533c711e8c41e83ca379e231a8eeeb7/raw/ba00cbd3fe964e08aa5e7ee8aa0d59398dedbeb4/data.json")
	if err != nil {
		fmt.Println("There is an error!")
	}
	defer response.Body.Close()

	var view []interface{}
	json.NewDecoder(response.Body).Decode(&view)
	log.Printf(" [x] Pulled JSON: %s", view)
	fmt.Println("---------------------------------------------")
	for row_counter, data := range view {
		md, _ := data.(map[string]interface{})
		fmt.Println(md["shortname"])
		_ = row_counter
	}

	/*
		for _, record := range view {
			log.Printf(" [===>] Record: %s", record)

			// for key, val := range record {
			// 	log.Printf(" [========>] %s = %s", key, val)
			// }
		}
	*/
}
