package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ch := make(chan int)
	strCh := make(chan string)

	go func() {
		ch <- 42
		strCh <- "hello"
	}()

	value := <-ch
	message := <-strCh
	fmt.Printf("Received: %d and %s\n", value, message)

	start := time.Now()
	done := make(chan bool)
	go func() {
		fmt.Println("Working...")
		// time.Sleep(1000 * time.Millisecond)
		done <- true
	}()
	<-done
	fmt.Println("Work completed after", time.Since(start))

	numbers := make(chan int)
	go func() {
		for i := 1; i <= 3; i++ {
			numbers <- i
		}
		close(numbers)
	}()

	for n := range numbers {
		fmt.Printf("Got: %d\n", n)
	}
	// return
	ch2 := make(chan bool)
	close(ch2)
	val, ok := <-ch2
	fmt.Printf("Value: %v, Open: %v\n", val, ok)
	// return
	fmt.Println("\n--- Broadcast Signal with Close ---")
	begin := make(chan interface{})
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ { // 0, 1, 2, 3, 4
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			<-begin // Wait for the channel to be closed
			fmt.Printf("%v has begun\n", i)
		}(i)
	}
	fmt.Println("Unblocking goroutines...")
	close(begin) // Closing the channel unblocks all waiting goroutines
	wg.Wait()
}
