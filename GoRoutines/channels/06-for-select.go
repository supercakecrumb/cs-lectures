package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("=== For-Select Pattern ===\n")

	// GOOD: Processing multiple channels until done
	fmt.Println("1. GOOD: For-select with done channel")
	done := make(chan bool)
	jobs := make(chan int, 5)

	go func() {
		for i := 1; i <= 5; i++ {
			jobs <- i
		}
		close(jobs)
	}()

	go func() {
		for {
			select {
			case job, ok := <-jobs:
				if !ok {
					done <- true
					return
				}
				fmt.Printf("Processing job %d\n", job)
				time.Sleep(50 * time.Millisecond)
			}
		}
	}()

	<-done
	fmt.Println("All jobs processed")

	// BAD: Infinite loop without exit condition
	fmt.Println("\n2. BAD: For-select without proper exit")
	counter := 0
	tick := time.NewTicker(50 * time.Millisecond)

	// This would run forever if not for the counter check
	for {
		select {
		case <-tick.C:
			counter++
			fmt.Printf("Tick %d\n", counter)
			if counter >= 3 {
				tick.Stop()
				goto exit // Need explicit exit
			}
		}
	}
exit:

	// GOOD: Multiple channels with timeout
	fmt.Println("\n3. GOOD: Handling multiple channels with timeout")
	c1 := make(chan string)
	c2 := make(chan string)
	done2 := make(chan bool)

	go func() {
		time.Sleep(100 * time.Millisecond)
		c1 <- "from c1"
		time.Sleep(100 * time.Millisecond)
		c2 <- "from c2"
		done2 <- true
	}()

	timeout := time.After(500 * time.Millisecond)
	for {
		select {
		case msg := <-c1:
			fmt.Println("Received:", msg)
		case msg := <-c2:
			fmt.Println("Received:", msg)
		case <-done2:
			fmt.Println("Done signal received")
			return
		case <-timeout:
			fmt.Println("Timeout!")
			return
		}
	}
}
