package main

import (
	"fmt"
	fasttext "github.com/bountylabs/go-fasttext"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

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

func predictProductCategory(model *fasttext.Model, inputFilePath string) (fasttext.Prediction, error) {
	//@description: 상품 카테고리를 예측한다.
	//* @param: model: 예측 모델
	//* @param: inputFilePath: 예측할 파일 경로
	predictResult, err := model.Predict("바나나")
	if err != nil {
		log.Fatal(err)
	}

	return predictResult[0], err
}

func main() {

	// set max cpu
	runtime.GOMAXPROCS(runtime.NumCPU() - 2)

	inputDataDir := "/data/personal/aiden/workspace/wp_bitbucket/ds/wp-auto-prod-cate/modeling/fasttext-model/data/prod_titles/20211101_20211130"
	log.Printf("Get file list from directory: %s\n", inputDataDir)
	fileList, err := getFileList(inputDataDir)
	if err != nil {
		log.Fatal(err)
	}
	if len(fileList) == 0 {
		log.Fatal("No file found in directory: ", inputDataDir)
		fmt.Println("Please check your input directory", inputDataDir)
		os.Exit(0)
	}
	log.Printf("Total %d input files\n", len(fileList))
	start := time.Now()
	modelFilePath := "/data/personal/aiden/workspace/wp_bitbucket/ds/wp-auto-prod-cate/modeling/fasttext-model/data/molecule/model/20211101_20211130/20211101_20211130.v01.bin"
	model := fasttext.Open(modelFilePath)
	elapsed := time.Since(start)
	log.Printf("Loading model took %s\n", elapsed)
	log.Println("Predicting category on input files")
	for _, file := range fileList {
		fmt.Println(filepath.Base(file))
		outputFileName := file + ".predicted"
		fmt.Printf("%s --> %s \n", file, outputFileName)
		pred, err := predictProductCategory(model, file)
		if err != nil {
			log.Fatal(err)
		}
		label := pred.Label
		prob := pred.Probability
		fmt.Printf("%s, %f \n", label, prob)
	}
}
