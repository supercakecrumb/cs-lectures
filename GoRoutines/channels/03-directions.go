package main

import "fmt"

func sender(ch chan<- string) {
	ch <- "Hello"
	// <-ch // Compile error
}

func receiver(ch <-chan string) {
	msg := <-ch
	fmt.Println("Received:", msg)
	// ch <- "World" // Compile error
}

func generator(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func main() {
	fmt.Println("=== Channel Directions ===\n")

	ch := make(chan string)
	go sender(ch)
	receiver(ch)

	fmt.Println("\n--- Pipeline Example ---")
	numbers := generator(2, 3, 4)
	squares := square(numbers)

	for result := range squares {
		fmt.Printf("Square: %d\n", result)
	}

	fmt.Println("\n--- Simple Direction Example ---")
	work := make(chan int)
	result := make(chan int)

	go produce(work)
	go process(work, result)

	for i := 0; i < 5; i++ {
		fmt.Printf("Result: %d\n", <-result)
	}
}

func produce(out chan<- int) {
	for i := 1; i <= 5; i++ {
		out <- i
	}
	close(out)
}

func process(in <-chan int, out chan<- int) {
	for n := range in {
		out <- n * 2
	}
	close(out)
}
