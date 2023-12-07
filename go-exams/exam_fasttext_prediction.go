// FastText prediction
package main

import (
	"fmt"
	fasttext "github.com/bountylabs/go-fasttext"
)

func main() {
	model := fasttext.Open("/data/personal/aiden/workspace/wp_bitbucket/ds/wp-auto-prod-cate/modeling/fasttext-model/data/molecule/model/20211101_20211130/20211101_20211130.v01.bin")
	fmt.Println("Prediction...")
	predictResult, err := model.Predict("분류할 입력 텍스트입니다.")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Label: %s, Prob: %fs\n", predictResult[0].Label, predictResult[0].Probability)
}
