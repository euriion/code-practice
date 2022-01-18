// csv.gz 파일에서 데이터를 한 줄씩 읽어오는 방법
package main

import (
	"compress/gzip"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
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

	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		// 첫번째 필드만 꺼내오기
		fmt.Println("record count:", len(record))
		if len(record) < 1 {
			log.Println("record is empty")
			continue
		}
		field1 := strings.Trim(record[0], " ")
		if len(field1) < 1 {
			log.Println("field1 is empty")
			continue
		}

		fmt.Println("field1:", field1)
	}
}
