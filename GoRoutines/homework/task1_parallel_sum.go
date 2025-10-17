package homework

// Task 1: Concurrent Computing - Parallel Sum
//
// OBJECTIVE: Calculate sum of numbers using parallel goroutines
//
// PACKAGES TO USE:
// - sync.WaitGroup: https://pkg.go.dev/sync#WaitGroup
//   wg.Add(1)    // increment counter
//   wg.Done()    // decrement counter
//   wg.Wait()    // block until counter is 0
//
// - Channels: https://go.dev/tour/concurrency/2
//   ch := make(chan int, 10)  // buffered channel
//   ch <- 42                   // send
//   result := <-ch             // receive
//   close(ch)                  // close channel
//
// HINT: Split slice into chunks, sum each chunk in goroutine, collect results

// ParallelSum calculates the sum of numbers in parallel chunks
func ParallelSum(numbers []int, workers int) int {
	// TODO: Implement parallel sum
	// 1. Handle edge cases for empty slice or invalid worker count
	// 2. Calculate chunk size by dividing total numbers by worker count
	// 3. Create buffered channel for collecting partial sums from workers
	// 4. Create WaitGroup to track worker completion
	// 5. For each worker:
	//    a. Calculate start and end indices for this worker's chunk
	//    b. Increment WaitGroup counter
	//    c. Launch goroutine that sums its chunk and sends result to channel
	// 6. Launch separate goroutine to wait for all workers and close results channel
	// 7. Collect all partial sums from results channel
	// 8. Return total sum
	return 0
}

// SquareSum calculates sum of squares in parallel
func SquareSum(numbers []int, workers int) int {
	// TODO: Implement parallel sum of squares
	// 1. Same algorithm as ParallelSum
	// 2. But square each number before adding to the sum
	return 0
}
