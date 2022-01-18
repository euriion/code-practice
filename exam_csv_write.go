package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	csvFile, err := os.Create("output.csv")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer csvFile.Close()

	records := [][]string{{"item1", "value1"}, {"item2", "value2"}, {"item3", "value3"}}

	writer := csv.NewWriter(csvFile)
	for _, record := range records {
		err := writer.Write(record)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
	}
	writer.Flush()
}
