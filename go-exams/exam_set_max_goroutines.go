// 최대 동시 사용 코어 수를 설정하는 코드
package main

import "runtime"

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU() - 2)
}
