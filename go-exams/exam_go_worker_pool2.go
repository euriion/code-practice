package main

import "fmt"
import "time"

//다음은 여러개의 인스턴스를 동시에 실행시킬 워커입니다. 이 워커들은 jobs 채널을 통해 작업을 받으며 작업의 결괏값을 results로 보냅니다. 비용이 큰 작업을 시뮬레이션 하기 위해 각 잡마다 1초의 딜레이를 줄겁니다.
func worker(id int, jobs <-chan int, results chan<- int) {

	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func main() {
	//워커풀을 사용하기 위해선 워커에 작업을 보내고 그 결괏값들을 수집해야합니다. 이를 위해 2개의 채널을 만듭니다.

	jobs := make(chan int, 100)
	results := make(chan int, 100)
	//다음은 3개의 워커를 실행시키는데, 처음에는 잡이 없기 때문에 각 워커는 블로킹 됩니다.

	for work := 1; work <= 3; work++ {
		go worker(work, jobs, results)
	}
	//5개의 잡을 보내고 작업을 다 보냈음을 알리기 위해 채널을 close합니다.

	for jobId := 1; jobId <= 5; jobId++ {
		jobs <- jobId
	}
	close(jobs)
	//마지막으로 모든 작업의 결괏값들을 가져옵니다.

	for a := 1; a <= 5; a++ {
		println("a:", a)
		println("result: ", <-results)
	}
}
