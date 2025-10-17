package homework

import (
	"errors"
	"time"
)

// Task 7: Timeout Pattern
//
// OBJECTIVE: Process data with timeout constraint
//
// PACKAGES TO USE:
// - time.After: https://pkg.go.dev/time#After
//   timeout := time.After(5 * time.Second)
//   select { case <-timeout: /* handle timeout */ }
//
// HINT: Use select with time.After and done channel

var ErrTimeout = errors.New("processing timeout exceeded")

// ProcessWithTimeout processes data with a timeout
func ProcessWithTimeout(data []int, timeout time.Duration) (int, error) {
	// TODO: Implement processing with timeout
	// 1. Create done channel for completion signal
	// 2. Launch goroutine to process all data items
	// 3. Use select statement with timeout channel
	// 4. If done channel receives first, return count
	// 5. If timeout channel receives first, return error
	// 6. Return processed count or timeout error
	return 0, nil
}
