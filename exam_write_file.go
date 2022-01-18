package main

import (
	"bufio"
	"bytes"
	"compress/gzip"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	//flag.Parse() // get the arguments from command line
	//
	//filename := flag.Arg(0)
	//
	//if filename == "" {
	//	fmt.Println("Usage : go-gzip sourcefile")
	//	os.Exit(1)
	//}

	filename := "./exam_write_file.go"
	rawFile, err := os.Open(filename)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer rawFile.Close()

	// calculate the buffer size for rawFile
	info, _ := rawFile.Stat()

	var size int64 = info.Size()
	rawBytes := make([]byte, size)

	// read rawFile content into buffer
	buffer := bufio.NewReader(rawFile)
	_, err = buffer.Read(rawBytes)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var buf bytes.Buffer
	writer := gzip.NewWriter(&buf)
	writer.Write(rawBytes)
	writer.Close()

	err = ioutil.WriteFile(filename+".gz", buf.Bytes(), info.Mode())
	// use 0666 to replace info.Mode() if you prefer

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("%s compressed to %s\n", filename, filename+".gz")

}
