package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU() - 2) // CPU 코어 수를 구하고 -2를 해서 사용할 코어 수를 지정한다.
	fmt.Println(runtime.GOMAXPROCS(0))       // 설정 값 출력

	s := "Hello, world!"

	for i := 0; i < 100; i++ {
		go func(n int) { // 익명 함수를 고루틴으로 실행
			fmt.Println(s, n)
		}(i)
	}

	fmt.Scanln()
}
