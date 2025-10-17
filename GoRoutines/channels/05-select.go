package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("=== Select Statement ===\n")

	// GOOD: Basic select usage
	fmt.Println("1. GOOD: Basic select with timeout")
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(100 * time.Millisecond)
		c1 <- "result from c1"
	}()

	go func() {
		time.Sleep(200 * time.Millisecond)
		c2 <- "result from c2"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("Received:", msg1)
		case msg2 := <-c2:
			fmt.Println("Received:", msg2)
		case <-time.After(300 * time.Millisecond):
			fmt.Println("Timeout!")
		}
	}

	// GOOD: Non-blocking receive with default
	fmt.Println("\n2. GOOD: Non-blocking operations")
	messages := make(chan string, 1)

	select {
	case msg := <-messages:
		fmt.Println("Received message:", msg)
	default:
		fmt.Println("No message received (non-blocking)")
	}

	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("Sent message:", msg)
	default:
		fmt.Println("No message sent")
	}

	// BAD: Select without default can block forever
	fmt.Println("\n3. BAD: Potential deadlock without timeout/default")
	// Uncomment to see deadlock:
	// emptyChannel := make(chan int)
	// select {
	// case <-emptyChannel:
	//     fmt.Println("Never happens")
	// }

	// GOOD: Select with done channel pattern
	fmt.Println("\n4. GOOD: Cancellation with done channel")
	done := make(chan bool)
	values := make(chan int)

	go func() {
		for {
			select {
			case <-done:
				fmt.Println("Worker stopped")
				return
			case v := <-values:
				fmt.Printf("Processing: %d\n", v)
			}
		}
	}()

	for i := 0; i < 3; i++ {
		values <- i
		time.Sleep(50 * time.Millisecond)
	}

	done <- true
	time.Sleep(100 * time.Millisecond)

	// BAD: Forgetting to handle all cases
	fmt.Println("\n5. BAD: Not handling all important channels")
	data := make(chan int, 1)
	errors := make(chan error, 1)

	go func() {
		// Simulating work that might error
		errors <- fmt.Errorf("something went wrong")
	}()

	// This only checks data, ignoring errors!
	select {
	case d := <-data:
		fmt.Println("Got data:", d)
	case <-time.After(100 * time.Millisecond):
		fmt.Println("Timeout - but error channel was ignored!")
	}

	// Drain error channel to prevent goroutine leak
	<-errors
}
