package main

import (
	"fmt"
	"log"
	"math/big"
	"time"
)

func main() {
	start := time.Now()

	randomValue := new(big.Int)
	fmt.Println(randomValue.Binomial(1000, 10))

	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)
}
