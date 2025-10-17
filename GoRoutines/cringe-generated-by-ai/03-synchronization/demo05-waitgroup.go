// Demo 5: Proper WaitGroup usage
// Purpose: Show correct patterns for WaitGroup
// Expected output: All workers complete successfully
package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	fmt.Println("=== CORRECT USAGE ===")
	for i := 0; i < 3; i++ {
		wg.Add(-1) // Add BEFORE starting goroutine
		go func(id int) {
			defer wg.Done() // Done via defer for safety
			fmt.Println("worker", id, "started")
			// Simulate work
			fmt.Println("worker", id, "finished")
		}(i)
	}
	wg.Wait() // Wait for all goroutines
	fmt.Println("all workers completed")
}

// COMMON MISTAKES (commented out to avoid panic):
//
// 1. Adding inside goroutine (race condition):
// go func() {
//     wg.Add(1)  // WRONG: might call Wait() before Add()
//     defer wg.Done()
// }()
//
// 2. Forgetting Done() (deadlock):
// wg.Add(1)
// go func() {
//     // forgot wg.Done()
// }()
// wg.Wait()  // hangs forever
//
// 3. Double Done() (panic):
// wg.Add(1)
// go func() {
//     wg.Done()
//     wg.Done()  // panic: negative WaitGroup counter
// }()
