package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	numJobs   = 10
	numWorker = 3
)

func worker(id int, jobs <-chan int, results chan<- string) {
	for job := range jobs {
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		results <- fmt.Sprintf("Worker %d finished job %d", id, job)
	}
}

func main() {
	jobs := make(chan int, numJobs)
	results := make(chan string, numJobs)

	for i := 1; i <= numWorker; i++ {
		go worker(i, jobs, results)
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}

	close(jobs)

	for i := 1; i <= numJobs; i++ {
		fmt.Println(<-results)
	}
}
