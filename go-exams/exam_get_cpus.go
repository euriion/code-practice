// OS, 아키텍쳐, CPU수, 고루틴 수를 출력한다.
package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("OS\t\t", runtime.GOOS)
	fmt.Println("ARCH\t\t", runtime.GOARCH)
	fmt.Println("CPU\t\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())
}
