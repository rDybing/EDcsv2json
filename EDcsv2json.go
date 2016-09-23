/* small piece of code to convert a rather lenghty (132 lines) csv list
 * of Elite: Dangerous materials to a more managable json format.
 *
 * Original file is *NOT* comma separated, but tab separated...
 *
 * blatantly copied from:
 * https://www.socketloop.com/tutorials/golang-convert-csv-data-to-json-format-and-save-to-file
 * with some minor changes to make it work for my purposes.
 * CC-BY Roy Dybing September 2016
 */

package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
)

type Material struct {
	Name   string
	Type   string
	Rarity string
	How    string
}

func main() {
	// read data from CSV file
	csvFile, err := os.Open("./materials.csv")

	if err != nil {
		fmt.Println(err)
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	reader.Comma = '\t' // Use tab-delimited instead of comma
	reader.FieldsPerRecord = -1
	csvData, err := reader.ReadAll()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var oneRecord Material
	var allRecords []Material

	for _, each := range csvData {
		oneRecord.Name = each[0]
		oneRecord.Type = each[1]
		oneRecord.Rarity = each[2]
		oneRecord.How = each[3]
		allRecords = append(allRecords, oneRecord)
	}

	jsondata, err := json.Marshal(allRecords) // convert to JSON

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// Terminal output (optional, debug only)
	fmt.Println(string(jsondata))
	// JSON file output
	jsonFile, err := os.Create("./materials.json")

	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	jsonFile.Write(jsondata)
	jsonFile.Close()
}
