package main

import (
	"flag"
	"fmt"
)

func main() {
	file := flag.String("file", "default.txt", "Input file")
	maxcpu := flag.Int("maxcpu", 10, "Max CPU Count")
	isForce := flag.Bool("force", false, "Run force")

	flag.Parse()

	// 포인터 변수이므로 앞에 * 를 붙어 deference 해줘야 한다.
	fmt.Printf("file: %s\n", *file)
	fmt.Printf("maxcpu: %d\n", *maxcpu)
	fmt.Printf("force: %t\n", *isForce)
}

/* 테스트
C> go build exam_arguments.go
C> exam_arguments -file=test.csv -maxtrial=5 -force=true
*/
