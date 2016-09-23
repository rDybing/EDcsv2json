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

	// sanity check
	// NOTE : You can stream the JSON data to http service as well instead of saving to file
	fmt.Println(string(jsondata))

	// now write to JSON file

	jsonFile, err := os.Create("./materials.json")

	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	jsonFile.Write(jsondata)
	jsonFile.Close()
}
