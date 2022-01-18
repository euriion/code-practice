package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	dir := "/data/personal/aiden/workspace/wp_bitbucket/ds/wp-auto-prod-cate/modeling/fasttext-model/data/prod_titles/20211101_20211130"
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file.Name(), file.IsDir())
	}
}
