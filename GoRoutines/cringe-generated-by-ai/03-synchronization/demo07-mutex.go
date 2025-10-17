// Demo 7: Mutex fixes the race
// Purpose: Show how Mutex protects shared memory
// Expected output: x = 1000 (always)
// Run with: go run -race demo07-mutex.go (no race detected)
package main

import (
	"fmt"
	"sync"
	"time"
)

var x int
var mu sync.RWMutex

func main() {
	start := time.Now()
	var (
		x  int
		mu sync.RWMutex
	)

	for i := 0; i < 1000; i++ {
		go func() {
			mu.Lock() // Acquire lock reader == 0
			x++       // Critical section
			// read
			// add
			// assign
			mu.Unlock() // Release lock
		}()
	}

	time.Sleep(100 * time.Microsecond)

	mu.Lock()
	fmt.Println("x =", x) // Always 1000
	mu.Unlock()

	fmt.Println("total:", time.Since(start))
}

func printx() {
	mu.RLock()
	fmt.Println(x)
}

// Key points:
// - Mutex ensures only one goroutine accesses x at a time
// - Lock/Unlock creates a critical section
// - Always unlock (prefer defer mu.Unlock() for safety)
//
// When to use sync.RWMutex instead:
// - Many readers, few writers
// - RLock() allows concurrent reads
// - Lock() still exclusive for writes
//
// Example:
//   var mu sync.RWMutex
//
//   // Multiple readers can hold RLock simultaneously
//   mu.RLock()
//   value := sharedData
//   mu.RUnlock()
//
//   // Writer needs exclusive Lock
//   mu.Lock()
//   sharedData = newValue
//   mu.Unlock()
