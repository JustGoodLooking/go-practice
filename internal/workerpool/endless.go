package workerpool

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"
)

func StartEndlessWorkerPool() {
	const numWorkers = 3

	jobs := make(chan Job, 100)
	results := make(chan Result, 100)
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup

	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go endlessWorker(ctx, i, jobs, results, &wg)
	}

	go func() {
		jobID := 1
		for {
			select {
			case <-ctx.Done():
				close(jobs)
				return
			default:
				fmt.Printf("send job %d \n", jobID)
				jobs <- Job{ID: jobID}
				jobID++
				time.Sleep(300 * time.Millisecond)
			}
		}
	}()

	go func() {
		for r := range results {
			fmt.Println(r.Output)
		}
	}()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	<-sig
	fmt.Println("Interrupt received, shutting down...")

	cancel()
	wg.Wait()
	close(results)
	fmt.Println("All done ✅")
}
