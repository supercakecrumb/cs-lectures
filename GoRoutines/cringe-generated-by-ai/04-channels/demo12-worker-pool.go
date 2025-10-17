// Demo 12: Worker Pool (fan-out/fan-in pattern)
// Purpose: Show practical pattern for parallel task processing
// Expected output: Results from 5 jobs processed by 3 workers
package main

import (
	"fmt"
	"sync"
)

func main() {
	jobs := make(chan int, 5)
	results := make(chan int, 5)

	// Start worker pool
	var wg sync.WaitGroup
	workers := 3
	for w := 1; w <= workers; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &wg)
	}

	// Send jobs
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs) // Signal no more jobs

	// Close results after all workers finish
	go func() {
		wg.Wait()
		close(results)
	}()

	// Collect results
	for r := range results {
		fmt.Println("result:", r)
	}
}

func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range jobs { // Read until jobs is closed
		fmt.Printf("worker %d processing job %d\n", id, j)
		results <- j * 2 // Send result
	}
	fmt.Printf("worker %d done\n", id)
}

// Key pattern elements:
// 1. jobs channel: distributes work to workers (fan-out)
// 2. results channel: collects outputs (fan-in)
// 3. WaitGroup: tracks worker completion
// 4. Separate goroutine closes results after all workers done
//
// Critical points:
// - Producer closes jobs channel
// - Workers read from jobs until closed
// - Separate goroutine closes results after WaitGroup
// - Main reads results until closed
//
// Common mistakes:
// ❌ Closing results in main (before workers finish)
// ❌ Not closing jobs (workers hang forever)
// ❌ Closing results in worker (multiple closes = panic)
//
// Real-world uses:
// - Parallel HTTP requests
// - Batch database operations
// - Image processing pipeline
// - Log aggregation
