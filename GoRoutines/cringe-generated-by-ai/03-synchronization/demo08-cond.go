// Demo 8: Condition variable for event waiting
// Purpose: Show sync.Cond for complex state synchronization
// Expected output: Waits ~1 second, then prints "ready!"
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var mu sync.Mutex
	cond := sync.NewCond(&mu)
	ready := false

	// Producer goroutine
	go func() {
		time.Sleep(time.Second)
		mu.Lock()
		ready = true
		mu.Unlock()
		cond.Signal() // Wake up one waiting goroutine
	}()

	// Consumer waits for ready
	mu.Lock()
	for !ready { // Always check condition in a loop
		cond.Wait() // Atomically unlocks mu and waits
	}
	mu.Unlock()

	fmt.Println("ready!")
}

// Key points:
// - Cond.Wait() atomically unlocks mutex and blocks
// - When signaled, Wait() reacquires lock before returning
// - Always check condition in a loop (spurious wakeups)
//
// Signal vs Broadcast:
// - Signal(): wakes one waiting goroutine
// - Broadcast(): wakes all waiting goroutines
//
// When to use Cond:
// - Complex state conditions
// - Multiple goroutines waiting for same event
// - Need to wake specific number of waiters
//
// Usually channels are simpler for most cases!
