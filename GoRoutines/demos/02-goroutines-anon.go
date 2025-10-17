package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("=== Goroutine Closure Example ===\n")

	// varExample()
	// varExampleWithPass()

	// sliceExample()
	sliceExampleWithPass()
}

func varExample() {
	var wg sync.WaitGroup
	salutation := "hello"
	wg.Add(1)
	go func() {
		defer wg.Done()
		salutation = "welcome"
	}()
	wg.Wait()
	fmt.Println(salutation)
}

func varExampleWithPass() {
	var wg sync.WaitGroup
	salutation := "hello"
	wg.Add(1)
	go func(badzim string) {
		defer wg.Done()
		badzim = "welcome"
		fmt.Println("goroutine:", badzim)
	}(salutation)
	wg.Wait()
	fmt.Println("varExampleWithPass:", salutation)
}

func sliceExample() {
	var wg sync.WaitGroup

	for i, salutation := range []string{"hello", "greetings", "good day"} {
		fmt.Println("main:", i, salutation)
		wg.Add(1)
		go func(idx int) {
			// time.Sleep(1000 * time.Millisecond)
			defer wg.Done()
			fmt.Println("goroutine:", idx, salutation)
		}(i)
	}

	wg.Wait()
}

func sliceExampleWithPass() {
	var wg sync.WaitGroup

	for _, salutation := range []string{"hello", "greetings", "good day"} {
		wg.Add(1)
		go func(greeting string) {
			defer wg.Done()
			fmt.Println(greeting)
		}(salutation)
	}

	wg.Wait()
	fmt.Println("Notice: Each goroutine printed its own greeting!")
}
