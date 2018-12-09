package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	//"encoding/json"

	"os"
)

func ReadCsvFile(filePath string) []map[string]interface{} {
	// Load a csv file.
	f, _ := os.Open(filePath)
	// Create a new reader.
	r := csv.NewReader(bufio.NewReader(f))
	result, _ := r.ReadAll()
	parsedData := make([]map[string]interface{}, 0, 0)

	header_name := result[0]

	for row_counter, row := range result {

		if row_counter != 0 {
			var singleMap = make(map[string]interface{})
			for col_counter, col := range row {
				singleMap[header_name[col_counter]] = col
			}
			parsedData = append(parsedData, singleMap)
		}

	}

	return parsedData

}
func main() {
	parsedData := ReadCsvFile("./data/csv_data.csv")
	fmt.Println("parsedData:", parsedData)
	response, err := http.Get("https://gist.githubusercontent.com/dineshsonachalam/4533c711e8c41e83ca379e231a8eeeb7/raw/081aa24a0d5d2fd63c8401e1f34294d5b1330abc/data.json")
	if err != nil {
		fmt.Println("There is an error!")
	}
	defer response.Body.Close()

	var view []interface{}
	json.NewDecoder(response.Body).Decode(&view)
	fmt.Println("-----------------------------------------------------------------------------------")
	log.Printf(" [x] Pulled JSON: %s", view)

	update := false

	for _, pData := range parsedData {
		update = false
		_ = pData
		for _, cData := range view {
			fmt.Println(cData)
			cData, ok := cData.(map[string]interface{})
			// fmt.Println("cData:", cData)
			// fmt.Println("ok:", ok)
			_ = ok
			for key, val := range cData {
				fmt.Println(key, " ", val)
			}
			fmt.Println("*********************")
		}
		fmt.Println("-------------------------------------------------------")
	}
	_ = update
}

/*
dinesh@dinesh-Inspiron-5559:~/Desktop/Golang_Tour/4_structs_slices_and_maps$ go run test2.go
parsedData: [map[price:100 id:1 name:Ikea] map[id:2 name:Oriflame price:120] map[id:3 name:Walmart price:130]]
-----------------------------------------------------------------------------------
2018/12/09 20:12:56  [x] Pulled JSON: [map[name:Barot Bellingham shortname:Barot_Bellingham reknown:Royal Academy of Painting and Sculpture bio:Barot has just finished id:5] map[id:6 name:Jonathan G. Ferrar II shortname:Jonathan_Ferrar reknown:Artist to Watch in 2012 bio:The Artist to Watch in 2012mixed media, was ]]
map[id:5 name:Barot Bellingham shortname:Barot_Bellingham reknown:Royal Academy of Painting and Sculpture bio:Barot has just finished]
id   5
name   Barot Bellingham
shortname   Barot_Bellingham
reknown   Royal Academy of Painting and Sculpture
bio   Barot has just finished
*********************
map[id:6 name:Jonathan G. Ferrar II shortname:Jonathan_Ferrar reknown:Artist to Watch in 2012 bio:The Artist to Watch in 2012mixed media, was ]
id   6
name   Jonathan G. Ferrar II
shortname   Jonathan_Ferrar
reknown   Artist to Watch in 2012
bio   The Artist to Watch in 2012mixed media, was
*********************
-------------------------------------------------------
map[id:5 name:Barot Bellingham shortname:Barot_Bellingham reknown:Royal Academy of Painting and Sculpture bio:Barot has just finished]
name   Barot Bellingham
shortname   Barot_Bellingham
reknown   Royal Academy of Painting and Sculpture
bio   Barot has just finished
id   5
*********************
map[id:6 name:Jonathan G. Ferrar II shortname:Jonathan_Ferrar reknown:Artist to Watch in 2012 bio:The Artist to Watch in 2012mixed media, was ]
bio   The Artist to Watch in 2012mixed media, was
id   6
name   Jonathan G. Ferrar II
shortname   Jonathan_Ferrar
reknown   Artist to Watch in 2012
*********************
-------------------------------------------------------
map[id:5 name:Barot Bellingham shortname:Barot_Bellingham reknown:Royal Academy of Painting and Sculpture bio:Barot has just finished]
bio   Barot has just finished
id   5
name   Barot Bellingham
shortname   Barot_Bellingham
reknown   Royal Academy of Painting and Sculpture
*********************
map[id:6 name:Jonathan G. Ferrar II shortname:Jonathan_Ferrar reknown:Artist to Watch in 2012 bio:The Artist to Watch in 2012mixed media, was ]
reknown   Artist to Watch in 2012
bio   The Artist to Watch in 2012mixed media, was
id   6
name   Jonathan G. Ferrar II
shortname   Jonathan_Ferrar
*********************
-------------------------------------------------------

*/
