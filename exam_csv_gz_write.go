// An example of writing to a csv.gzip compressed file.
package main

import (
	"compress/gzip"
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main() {
	// raw 파일 생성으로 열기
	csvGzipFile, err := os.Create("output.csv.gz")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer csvGzipFile.Close()

	// GZIP writer 열기
	gzipWriter := gzip.NewWriter(csvGzipFile)
	if err != nil {
		log.Fatal(err)
	}
	defer gzipWriter.Close()

	// CSV writer 열기
	csvFileWriter := csv.NewWriter(gzipWriter)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer csvFileWriter.Flush()

	records := [][]string{{"item1", "value1"}, {"item2", "value2"}, {"item3", "value3"}}
	// 한줄씩 데이터를 CSV에 저장
	for _, record := range records {
		err := csvFileWriter.Write(record)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
	}
}
