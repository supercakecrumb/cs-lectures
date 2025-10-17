// Demo 11: Closing channels and range
// Purpose: Show proper channel closing and iteration patterns
// Expected output: 1, 2, 3
package main

import "fmt"

func main() {
	ch := make(chan int, 3)

	// Send values
	for i := 1; i <= 3; i++ {
		ch <- i
	}
	close(ch) // Signal no more values

	// Range reads until channel is closed
	for v := range ch {
		fmt.Println(v)
	}

	// Reading from closed channel returns zero value
	v, ok := <-ch
	fmt.Printf("After close: v=%d, ok=%t\n", v, ok)
}

// Key rules for closing channels:
// 1. Only the SENDER should close
// 2. Closing is optional (for signaling "done")
// 3. Send to closed channel = PANIC
// 4. Receive from closed channel = zero value + ok=false
// 5. Close closed channel = PANIC
//
// Common pattern: close to signal completion
// func producer(ch chan int) {
//     defer close(ch)  // Ensure close on return
//     for i := 0; i < 10; i++ {
//         ch <- i
//     }
// }
//
// func consumer(ch chan int) {
//     for v := range ch {  // Exits when ch is closed
//         fmt.Println(v)
//     }
// }
//
// When NOT to close:
// - Multiple senders (who closes?)
// - Channel is just for signaling (use context instead)
// - Unclear ownership
