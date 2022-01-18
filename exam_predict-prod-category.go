package main

import (
	"compress/gzip"
	"encoding/csv"
	"fmt"
	fasttext "github.com/bountylabs/go-fasttext"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

// "github.com/bountylabs/go-fasttext"
// https://pkg.go.dev/github.com/bountylabs/go-fasttext#Model.Predict

func getFileList(dir string) ([]string, error) {
	//@description: 디렉토리 안에 있는 파일 목록을 반환한다.
	//* @return: String Array

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}
	var fileList []string
	for _, file := range files {
		parts := [...]string{dir, file.Name()}
		filePath := filepath.FromSlash(strings.Join(parts[:], "/"))
		fileList = append(fileList, filePath)
	}

	return fileList, err
}

func makePredictedOutputFile(model *fasttext.Model, inputFilePath string, outputFilePath string) (bool, error) {
	//@description: 상품 카테고리를 예측한다.
	//* @param: model: 예측 모델
	//* @param: inputFilePath: 예측할 파일 경로
	//* @param: outputFilePath: 예측된 파일 경로
	//* @return: bool: 성공 여부

	inputFileHandle, err := os.Open(inputFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer inputFileHandle.Close()
	gzipReader, err := gzip.NewReader(inputFileHandle)
	if err != nil {
		log.Fatal(err)
	}
	defer gzipReader.Close()

	csvReader := csv.NewReader(gzipReader)
	csvReader.Comma = '\t'
	csvReader.FieldsPerRecord = -1

	outputFileHandle, err := os.Open(outputFilePath)
	if err != nil {
		log.Fatal(err)
	}

	gzipWriter := gzip.NewWriter(outputFileHandle)
	if err != nil {
		log.Fatal(err)
	}

	csvWriter := csv.NewWriter(gzipWriter)
	if err != nil {
		log.Fatal(err)
	}

	// read all csv lines
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

		//predictResults, err := model.Predict("바나나")
		predictResults, err := model.Predict(field1)
		if err != nil {
			log.Fatal(err)
		}
		var predictLabel string
		var predictProb float32

		if len(predictResults) < 1 {
			log.Println("predictResults is empty")
			predictLabel = ""
			predictProb = 0.0
		} else {
			predictLabel = predictResults[0].Label
			predictProb = predictResults[0].Probability
		}
		outputRecord := []string{field1, predictLabel, fmt.Sprintf("%f", predictProb)}
		csvWriter.Write(outputRecord)
	}

	return true, nil
}

func main() {

	//fmt.Println("Product auto categorization")
	//fmt.Println("Loading FastText model...")
	//model := fasttext.Open("/data/personal/aiden/workspace/wp_bitbucket/ds/wp-auto-prod-cate/modeling/fasttext-model/data/molecule/model/20211101_20211130/20211101_20211130.v01.bin")
	//fmt.Println("Prediction...")
	//predictResult, err := model.Predict("바나나")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//label := strings.Replace(predictResult[0].Label, "__label__", "", -1)
	//prob := predictResult[0].Probability
	//fmt.Printf("Label: %s, Prob: %fs\n", label, prob)

	// set max cpu automatically
	maxCores := runtime.NumCPU() - 2
	runtime.GOMAXPROCS(maxCores)

	inputDataDir := "/data/personal/aiden/workspace/wp_bitbucket/ds/wp-auto-prod-cate/modeling/fasttext-model/data/prod_titles/20211101_20211130"

	// Check if the input directory exists
	if _, err := os.Stat(inputDataDir); os.IsNotExist(err) {
		fmt.Println("Input directory does not exist.", inputDataDir)
		os.Exit(1)
	}

	// Get inputFileName list in the input directory
	log.Printf("Get inputFileName list from directory: %s\n", inputDataDir)
	inputFileList, err := getFileList(inputDataDir)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Total %d input files\n", len(inputFileList))

	// Load FastText model
	log.Println("Loading FastText model...")
	start := time.Now()
	modelFilePath := "/data/personal/aiden/workspace/wp_bitbucket/ds/wp-auto-prod-cate/modeling/fasttext-model/data/molecule/model/20211101_20211130/20211101_20211130.v01.bin"
	model := fasttext.Open(modelFilePath)
	elapsed := time.Since(start)
	log.Println("Loading model took: ", elapsed)

	// Process each inputFileName
	log.Println("Predicting category on input files")
	for _, inputFileName := range inputFileList {
		fmt.Println(filepath.Base(inputFileName))
		outputFileName := inputFileName + ".predicted"
		fmt.Printf("%s --> %s \n", inputFileName, outputFileName)
		procResult, err := makePredictedOutputFile(model, inputFileName, outputFileName)
		if err != nil {
			log.Fatal(err)
		}
		if procResult {
			fmt.Println("Success")
		} else {
			fmt.Println("Fail")
		}
		//pred, err := makePredictedOutputFile(model, inputFileName)
		//if err != nil {
		//	log.Fatal(err)
		//}
		//label := pred.Label
		//prob := pred.Probability
		//fmt.Printf("%s, %f \n", label, prob)
	}
}
