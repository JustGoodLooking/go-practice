package workerpool

import "sync"

// StartFixedWorkerPool 啟動固定數量的 worker 處理 jobs
func StartFixedWorkerPool(jobs []Job, workerCount int) []Result {
	jobChan := make(chan Job)
	resultChan := make(chan Result)
	var wg sync.WaitGroup

	// 啟動固定數量 worker
	for i := 1; i <= workerCount; i++ {
		wg.Add(1)
		go fixedWorker(i, jobChan, resultChan, &wg)
	}

	// 投遞所有 job
	go func() {
		for _, job := range jobs {
			jobChan <- job
		}
		close(jobChan)
	}()

	// 等待 worker 全部完成後關閉 result channel
	go func() {
		wg.Wait()
		close(resultChan)
	}()

	// 收集所有結果
	var results []Result
	for r := range resultChan {
		results = append(results, r)
	}
	return results
}
