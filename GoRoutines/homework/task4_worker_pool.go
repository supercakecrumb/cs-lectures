package homework

import "context"

// Task 4: Worker Pool Pattern
//
// OBJECTIVE: Process jobs with fixed number of workers
//
// PACKAGES TO USE:
// - sync.WaitGroup: https://pkg.go.dev/sync#WaitGroup
// - context: https://pkg.go.dev/context
//   ctx.Done() returns <-chan struct{} that closes on cancellation
//
// HINT: Create jobs channel, start workers, send jobs, collect results

// WorkerPool processes jobs using fixed number of workers
func WorkerPool(jobs []int, numWorkers int) []int {
	// TODO: Implement worker pool
	// 1. Create jobs channel and results channel
	// 2. Start fixed number of worker goroutines
	// 3. Each worker reads from jobs channel, processes job, sends to results
	// 4. Send all jobs to jobs channel and close it
	// 5. Collect all results from results channel
	// 6. Return results slice
	return nil
}

// WorkerPoolWithContext adds cancellation support
func WorkerPoolWithContext(ctx context.Context, jobs []int, numWorkers int) ([]int, error) {
	// TODO: Implement worker pool with context cancellation
	// 1. Create jobs and results channels
	// 2. Start workers that check context cancellation in select
	// 3. Send jobs or stop if context is cancelled
	// 4. Collect results or return error if cancelled
	// 5. Return partial results and context error if cancelled
	return nil, nil
}
