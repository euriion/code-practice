package main

import (
	"fmt"
	"runtime"
	"time"
)

func foo() {
	time.Sleep(3)
	fmt.Println("foo")
}

func bar() {
	time.Sleep(5)
	fmt.Println("bar")
}

func main() {
	fmt.Println("OS\t\t", runtime.GOOS)                 // windows
	fmt.Println("ARCH\t\t", runtime.GOARCH)             // amd64
	fmt.Println("CPU\t\t", runtime.NumCPU())            // 12
	fmt.Println("Goroutines\t", runtime.NumGoroutine()) // 1

	go foo()
	fmt.Println("CPU\t\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine()) // 2
	bar()                                               // bar가 끝나면 foo가 어찌 되었던 일단 끝난다. 기다리게 하려면 sync 시켜야 한다.
	fmt.Println("CPU\t\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine()) // 2
}
