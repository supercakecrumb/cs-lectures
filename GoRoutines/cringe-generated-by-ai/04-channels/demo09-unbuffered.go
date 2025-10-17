// Demo 9: Unbuffered channel = synchronization point
// Purpose: Show that unbuffered channels block until both sides are ready
// Expected output: 42
package main

import "fmt"

func main() {
	ch := make(chan int) // Unbuffered channel

	go func() {
		fmt.Println("Goroutine: about to send")
		ch <- 42 // Blocks until someone receives
		fmt.Println("Goroutine: sent successfully")
	}()

	fmt.Println("Main: about to receive")
	v := <-ch // Blocks until someone sends
	fmt.Println("Main: received", v)
}

// Key points:
// - Unbuffered channel has capacity 0
// - Send blocks until receive is ready
// - Receive blocks until send is ready
// - Provides synchronization: both sides must be ready
//
// This is a "rendezvous" - sender and receiver meet at the channel
//
// Deadlock example (commented to avoid hanging):
// func main() {
//     ch := make(chan int)
//     ch <- 42  // Deadlock! No one to receive
//     fmt.Println(<-ch)
// }
