package main

import (
	"fmt"
	"go-practice/internal/workerpool"
)

func main() {
	// 準備 10 筆 job
	var jobs []workerpool.Job
	for i := 1; i <= 10; i++ {
		jobs = append(jobs, workerpool.Job{ID: i})
	}

	// 啟動 worker pool 處理這些 job
	results := workerpool.StartFixedWorkerPool(jobs, 3)

	// 印出處理結果
	for _, res := range results {
		fmt.Println(res.Output)
	}
}

//package main
//
//import (
//	"fmt"
//	"math/rand"
//	"time"
//)
//
//const (
//	numJobs   = 10
//	numWorker = 3
//)
//
//func worker(id int, jobs <-chan int, results chan<- string) {
//	for job := range jobs {
//		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
//		results <- fmt.Sprintf("Worker %d finished job %d", id, job)
//	}
//}
//
//func main() {
//	jobs := make(chan int, numJobs)
//	results := make(chan string, numJobs)
//
//	for i := 1; i <= numWorker; i++ {
//		go worker(i, jobs, results)
//	}
//
//	for j := 1; j <= numJobs; j++ {
//		jobs <- j
//	}
//
//	close(jobs)
//
//	for i := 1; i <= numJobs; i++ {
//		fmt.Println(<-results)
//	}
//}
