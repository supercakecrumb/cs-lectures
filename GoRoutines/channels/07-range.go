package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("=== Range Over Channels ===\n")

	// GOOD: Range with proper close
	fmt.Println("1. GOOD: Range with closed channel")
	numbers := make(chan int)

	go func() {
		for i := 1; i <= 5; i++ {
			numbers <- i
		}
		close(numbers)
	}()

	for num := range numbers {
		fmt.Printf("Received: %d\n", num)
	}
	fmt.Println("Channel closed, range exited")

	// BAD: Range without closing channel
	fmt.Println("\n2. BAD: Range without close (would deadlock)")
	// Uncomment to see deadlock:
	// unclosed := make(chan int)
	// go func() {
	//     for i := 1; i <= 3; i++ {
	//         unclosed <- i
	//     }
	//     // Missing close(unclosed)!
	// }()
	// for num := range unclosed {
	//     fmt.Println(num)
	// }

	// GOOD: Multiple goroutines with range
	fmt.Println("\n3. GOOD: Multiple producers, one consumer")
	results := make(chan int, 10)

	for i := 1; i <= 3; i++ {
		go func(id int) {
			for j := 1; j <= 3; j++ {
				results <- id*10 + j
				time.Sleep(50 * time.Millisecond)
			}
		}(i)
	}

	go func() {
		time.Sleep(200 * time.Millisecond)
		close(results)
	}()

	for result := range results {
		fmt.Printf("Got: %d\n", result)
	}

	// GOOD: Range with context-like done channel
	fmt.Println("\n4. GOOD: Range with early termination")
	data := make(chan int)
	done := make(chan bool)

	go func() {
		defer close(data)
		for i := 1; i <= 10; i++ {
			select {
			case data <- i:
			case <-done:
				fmt.Println("Producer stopped early")
				return
			}
		}
	}()

	count := 0
	for num := range data {
		fmt.Printf("Processing: %d\n", num)
		count++
		if count >= 5 {
			close(done)
			break
		}
	}
	fmt.Println("Consumer stopped early")
}
