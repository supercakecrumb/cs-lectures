package homework

import (
	"fmt"
	"sync"
	"time"
)

// ERROR 1: Easy - Missing WaitGroup.Wait()
// Bug: Goroutines may not complete before function returns
func PrintNumbers() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			fmt.Println(n)
		}(i)
	}

	// BUG: Missing wg.Wait()
	fmt.Println("Done")
}

// ERROR 2: Easy-Medium - Loop variable capture bug
// Bug: All goroutines will likely print the same (last) value
func PrintSquares(numbers []int) {
	var wg sync.WaitGroup

	for _, num := range numbers {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// BUG: Capturing loop variable by reference
			fmt.Printf("%d squared is %d\n", num, num*num)
		}()
	}

	wg.Wait()
}

// ERROR 3: Medium - Channel not closed, causing deadlock
// Bug: Range over channel will block forever
func SumChannel(numbers []int) int {
	ch := make(chan int)

	go func() {
		for _, num := range numbers {
			ch <- num
		}
		// BUG: Missing close(ch)
	}()

	sum := 0
	for num := range ch {
		sum += num
	}

	return sum
}

// ERROR 4: Medium-Hard - Race condition on shared variable
// Bug: Multiple goroutines accessing counter without synchronization
func ConcurrentCounter(n int) int {
	counter := 0
	var wg sync.WaitGroup

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// BUG: Race condition - no mutex protection
			counter++
		}()
	}

	wg.Wait()
	return counter
}

// ERROR 5: Hard - Goroutine leak with blocked channel
// Bug: Goroutine will block forever if channel is not read
func ProcessData(data []int) []int {
	results := make(chan int)

	go func() {
		for _, d := range data {
			// BUG: Unbuffered channel with no reader will block
			results <- d * 2
		}
		close(results)
	}()

	// BUG: Only reading first result, rest of goroutine is blocked
	output := []int{}
	if len(data) > 0 {
		output = append(output, <-results)
	}

	return output
}

// ERROR 6: Hard - Select without default in tight loop
// Bug: Will block if no data available
func MonitorChannel(input <-chan int, duration time.Duration) []int {
	results := []int{}
	timeout := time.After(duration)

	for {
		select {
		case val := <-input:
			results = append(results, val)
		case <-timeout:
			return results
			// BUG: Missing default case causes blocking
		}
	}
}

// ERROR 7: Hard - Deadlock with mutual channel dependency
// Bug: Two goroutines waiting on each other
func DeadlockExample() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		// BUG: Waiting to send on ch1 before receiving from ch2
		ch1 <- 1
		<-ch2
	}()

	go func() {
		// BUG: Waiting to send on ch2 before receiving from ch1
		ch2 <- 2
		<-ch1
	}()

	time.Sleep(100 * time.Millisecond)
}
