package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/xuri/excelize/v2"
)

func main() {
	file, err := excelize.OpenFile("menu.xlsx")
	if err != nil {
		log.Fatal("Failed to open file:", err)
	}
	defer file.Close()

	columns, err := file.GetCols("Sheet1")
	if err != nil {
		log.Fatal("Failed to get rows:", err)
	}

	var jsonData []byte
	jsonData, err = json.MarshalIndent(columns, "", "    ")
	if err != nil {
		log.Fatal("Failed to marshal to JSON:", err)
	}

	fmt.Println(string(jsonData))

	json_file, err := os.Create("json_data.json")
	if err != nil {
		log.Fatal("Failed to create file:", err)
	}

	json_file.Write(jsonData)
	json_file.Close()

}
