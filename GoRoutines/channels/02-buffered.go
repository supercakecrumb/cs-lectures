package main

import (
	"bytes"
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Println("=== Buffered Channels ===\n")

	ch := make(chan int, 2)

	ch <- 1
	ch <- 2
	fmt.Println("Sent 2 values without blocking")

	// ch <- 3 // DEADLOCK! Buffer is full

	fmt.Printf("Received: %d\n", <-ch)
	fmt.Printf("Received: %d\n", <-ch)

	fmt.Println("\n--- Deadlock Examples ---")

	// unbuffered := make(chan int)
	// unbuffered <- 1 // DEADLOCK!

	// empty := make(chan int)
	// <-empty // DEADLOCK!

	// full := make(chan int, 1)
	// full <- 1
	// full <- 2 // DEADLOCK!

	fmt.Println("\n--- Worker Pool ---")

	jobs := make(chan int, 5)
	results := make(chan int, 5)

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	for r := 1; r <= 5; r++ {
		fmt.Printf("Result: %d\n", <-results)
	}

	fmt.Println("\n--- Rate Limiting ---")
	requests := make(chan int, 3)

	for i := 1; i <= 3; i++ {
		requests <- i
	}

	ticker := time.NewTicker(200 * time.Millisecond)
	defer ticker.Stop()

	for i := 1; i <= 3; i++ {
		<-ticker.C
		req := <-requests
		fmt.Printf("Processing request %d at %s\n", req, time.Now().Format("15:04:05"))
	}

	fmt.Println("\n--- Producer/Consumer ---")
	var stdoutBuff bytes.Buffer
	defer stdoutBuff.WriteTo(os.Stdout)
	intStream := make(chan int, 4)
	go func() {
		defer close(intStream)
		defer fmt.Fprintln(&stdoutBuff, "Producer Done.")
		for i := 0; i < 5; i++ {
			fmt.Fprintf(&stdoutBuff, "Sending: %d\n", i)
			intStream <- i
		}
	}()
	for integer := range intStream {
		fmt.Fprintf(&stdoutBuff, "Received %v.\n", integer)
	}
}

func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, job)
		time.Sleep(100 * time.Millisecond)
		results <- job * 2
	}
}
