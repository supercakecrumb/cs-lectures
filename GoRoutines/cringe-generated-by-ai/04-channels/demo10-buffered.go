// Demo 10: Buffered channel = decoupling sender from receiver
// Purpose: Show that buffered channels allow async send/receive
// Expected output: a, b
package main

import "fmt"

func main() {
	ch := make(chan string, 2) // Buffer size 2

	// Can send without blocking (up to buffer size)
	ch <- "a"
	ch <- "b"
	// ch <- "c"  // Would block - buffer full

	fmt.Println(<-ch) // a
	fmt.Println(<-ch) // b
}

// Key differences from unbuffered:
// - Sender doesn't block until buffer is full
// - Receiver doesn't block until buffer is empty
// - Decouples sender and receiver timing
//
// Buffer size considerations:
// - Size 0 (unbuffered): tight synchronization
// - Size 1: simple async handoff
// - Size N: batch processing, rate limiting
//
// Common pattern: buffered channel as semaphore
// func main() {
//     sem := make(chan struct{}, 3) // Max 3 concurrent
//     for i := 0; i < 10; i++ {
//         sem <- struct{}{}  // Acquire
//         go func(i int) {
//             defer func() { <-sem }()  // Release
//             // Do work with max 3 concurrent
//         }(i)
//     }
// }
