package homework

// Task 5: Rate Limiter Pattern
//
// OBJECTIVE: Process items with rate limiting
//
// PACKAGES TO USE:
// - time.NewTicker: https://pkg.go.dev/time#NewTicker
//   ticker := time.NewTicker(100 * time.Millisecond)
//   defer ticker.Stop()
//   <-ticker.C  // wait for tick
//
// HINT: Create ticker, wait for tick before processing each item

// RateLimitedProcessor processes items with rate limiting
func RateLimitedProcessor(items []string, maxPerSecond int) <-chan string {
	// TODO: Implement rate limiting
	// 1. Calculate interval between items based on maxPerSecond
	// 2. Create ticker with that interval
	// 3. Create output channel
	// 4. Launch goroutine that waits for ticker before processing each item
	// 5. Send processed items to output channel
	// 6. Close output channel and stop ticker when done
	// 7. Return output channel

	// Return closed channel to prevent hanging in tests
	ch := make(chan string)
	close(ch)
	return ch
}
