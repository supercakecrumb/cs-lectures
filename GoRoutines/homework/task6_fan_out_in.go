package homework

// Task 6: Fan-Out/Fan-In Pattern
//
// OBJECTIVE: Distribute work to workers (fan-out), collect results (fan-in)
//
// PACKAGES TO USE:
// - sync.WaitGroup: https://pkg.go.dev/sync#WaitGroup
// - Channels for distribution and collection
//
// HINT: Create input channel, start workers reading from it, merge outputs

// FanOutFanIn distributes work across workers and collects results
func FanOutFanIn(numbers []int, numWorkers int) []int {
	// TODO: Implement fan-out/fan-in pattern
	// 1. Create input channel for distributing work
	// 2. Start multiple worker goroutines reading from input channel
	// 3. Each worker processes numbers and sends to its output channel
	// 4. Merge all worker output channels into single channel
	// 5. Collect all results from merged channel
	// 6. Return results slice
	return nil
}
