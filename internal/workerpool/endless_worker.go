package workerpool

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func endlessWorker(ctx context.Context, id int, jobs <-chan Job, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done() // worker 執行結束時，自動回報給 WaitGroup，確保主程式不會卡住

	for {
		select {
		//case <-ctx.Done():
		//	fmt.Printf("Worker %d shutting down\n", id)
		//	return
		case job, ok := <-jobs:
			if !ok {
				fmt.Printf("Worker %d: jobs channel closed\n", id)
				return
			}
			time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
			output := fmt.Sprintf("Worker %d finished job %d", id, job.ID)
			results <- Result{JobID: job.ID, Output: output}
		}
	}
}
