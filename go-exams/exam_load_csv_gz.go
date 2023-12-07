package main

import (
	"compress/gzip"
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("product_titles.csv.gz")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	gzipReader, err := gzip.NewReader(file)

	if err != nil {
		log.Fatal(err)
	}

	defer gzipReader.Close()

	csvReader := csv.NewReader(gzipReader)
	csvReader.Comma = '\t'
	csvReader.FieldsPerRecord = -1
	record, err := csvReader.Read()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d 개\n", len(record))

	for _, v := range record {
		fmt.Println(v)
	}
}
