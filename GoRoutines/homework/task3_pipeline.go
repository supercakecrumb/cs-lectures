package homework

// Task 3: Pipeline Pattern
//
// OBJECTIVE: Create 3-stage pipeline: generate → square → filter even
//
// PACKAGES TO USE:
// - Channels: https://go.dev/tour/concurrency/2
//   out := make(chan int)
//   go func() { for i := 0; i < n; i++ { out <- i }; close(out) }()
//   for val := range ch { /* process */ }
//
// HINT: Each stage returns <-chan int, chain them together

// ProcessPipeline creates a 3-stage pipeline
func ProcessPipeline(n int) <-chan int {
	// TODO: Chain the three stages together
	// 1. Call generate function to create first stage
	// 2. Pass its output channel to square function
	// 3. Pass square output channel to filterEven function
	// 4. Return the final output channel
	genChan := generate(n)
	squareChan := square(genChan)
	return filterEven(squareChan)
}

func generate(n int) <-chan int {
	// TODO: Generate numbers from 1 to n
	// 1. Create output channel
	// 2. Launch goroutine that loops from 1 to n
	// 3. Send each number to output channel
	// 4. Close channel when loop completes
	// 5. Return output channel immediately
	out := make(chan int)

	// Handle edge case of n <= 0
	if n <= 0 {
		close(out)
		return out
	}

	go func() {
		for i := 1; i <= n; i++ {
			out <- i
		}
		close(out)
	}()

	return out
}

func square(in <-chan int) <-chan int {
	// TODO: Square each number from input channel
	// 1. Create output channel
	// 2. Launch goroutine that ranges over input channel
	// 3. For each number, calculate its square
	// 4. Send squared number to output channel
	// 5. Close output channel when input is exhausted
	// 6. Return output channel immediately
	out := make(chan int)

	go func() {
		for num := range in {
			out <- num * num
		}
		close(out)
	}()

	return out
}

func filterEven(in <-chan int) <-chan int {
	// TODO: Filter and pass only even numbers
	// 1. Create output channel
	// 2. Launch goroutine that ranges over input channel
	// 3. For each number, check if it's even
	// 4. Send only even numbers to output channel
	// 5. Close output channel when input is exhausted
	// 6. Return output channel immediately
	out := make(chan int)

	go func() {
		for num := range in {
			if num%2 == 0 {
				out <- num
			}
		}
		close(out)
	}()

	return out
}
