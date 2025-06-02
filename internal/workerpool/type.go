package workerpool

// Job 表示要處理的工作（這邊只用 ID 示範）
type Job struct {
	ID int
}

// Result 表示 worker 處理完的結果
type Result struct {
	JobID  int
	Output string
	Error  error
}
