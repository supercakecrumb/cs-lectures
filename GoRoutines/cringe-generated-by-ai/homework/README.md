# Homework Assignments

Complete these three assignments to solidify your understanding of Go concurrency.

## Assignment 1: Rate Limiter with Semaphore
**Difficulty:** Medium  
**Time:** 30-45 minutes

Build a concurrent rate limiter using a buffered channel as a semaphore.

**Requirements:**
- Implement a function that processes 20 tasks
- Limit to maximum 5 concurrent tasks at any time
- Each task should simulate work with `time.Sleep()`
- Print when each task starts and finishes
- Measure total execution time

**Expected Behavior:**
- First 5 tasks start immediately
- As tasks complete, new ones start
- Never more than 5 running simultaneously
- Total time ≈ (20 tasks / 5 concurrent) × task duration

**Starter Code:**
```go
package main

import (
    "fmt"
    "time"
)

func main() {
    start := time.Now()
    
    // TODO: Create semaphore channel
    
    for i := 1; i <= 20; i++ {
        // TODO: Acquire semaphore
        go func(id int) {
            // TODO: Release semaphore with defer
            processTask(id)
        }(i)
    }
    
    // TODO: Wait for all tasks to complete
    
    fmt.Printf("Total time: %v\n", time.Since(start))
}

func processTask(id int) {
    fmt.Printf("Task %d started\n", id)
    time.Sleep(500 * time.Millisecond)
    fmt.Printf("Task %d finished\n", id)
}
```

**Bonus Challenges:**
1. Add metrics: track how long each task waited before starting
2. Implement dynamic semaphore size (configurable at runtime)
3. Add graceful shutdown with context cancellation

---

## Assignment 2: Pipeline with Context
**Difficulty:** Hard  
**Time:** 60-90 minutes

Build a 3-stage data processing pipeline with context-based cancellation.

**Requirements:**
- **Stage 1 (Generator):** Generate numbers 1-100
- **Stage 2 (Filter):** Keep only even numbers
- **Stage 3 (Square):** Square each number
- Support cancellation via `context.Context`
- Clean up all goroutines when cancelled
- Print results as they arrive

**Architecture:**
```
Generator -> [chan] -> Filter -> [chan] -> Square -> [chan] -> Consumer
     ↓           ↓         ↓         ↓        ↓         ↓
   context    context   context   context  context   context
```

**Starter Code:**
```go
package main

import (
    "context"
    "fmt"
    "time"
)

func main() {
    ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
    defer cancel()
    
    // TODO: Build pipeline
    // gen := generate(ctx, 1, 100)
    // filtered := filter(ctx, gen)
    // squared := square(ctx, filtered)
    
    // TODO: Consume results
    // for result := range squared {
    //     fmt.Println(result)
    // }
}

func generate(ctx context.Context, start, end int) <-chan int {
    out := make(chan int)
    go func() {
        defer close(out)
        for i := start; i <= end; i++ {
            select {
            case out <- i:
                time.Sleep(50 * time.Millisecond) // Simulate work
            case <-ctx.Done():
                return
            }
        }
    }()
    return out
}

// TODO: Implement filter(ctx, in) <-chan int
// TODO: Implement square(ctx, in) <-chan int
```

**Requirements:**
- All stages must respect context cancellation
- No goroutine leaks (verify with `-race` flag)
- Proper channel closing
- Handle context timeout gracefully

**Bonus Challenges:**
1. Add a fourth stage (e.g., sum accumulator)
2. Implement fan-out: multiple filter workers
3. Add metrics: track throughput of each stage
4. Implement backpressure handling

---

## Assignment 3: Trace Analysis
**Difficulty:** Medium  
**Time:** 30-45 minutes

Profile your own concurrent code using Go's execution tracer.

**Requirements:**
1. Take any concurrent program you've written (or use worker pool demo)
2. Add trace instrumentation
3. Generate and analyze trace file
4. Answer analysis questions

**Steps:**

### 1. Add Tracing to Your Code
```go
package main

import (
    "os"
    "runtime/trace"
)

func main() {
    // Create trace file
    f, err := os.Create("trace.out")
    if err != nil {
        panic(err)
    }
    defer f.Close()
    
    // Start tracing
    if err := trace.Start(f); err != nil {
        panic(err)
    }
    defer trace.Stop()
    
    // Your concurrent code here
    // ...
}
```

### 2. Run and Generate Trace
```bash
go run your_program.go
go tool trace trace.out
```

### 3. Analysis Questions

Answer these questions by examining the trace:

**Goroutine Analysis:**
1. How many goroutines were created?
2. What's the maximum number of goroutines running simultaneously?
3. Are there any goroutines that never ran? (goroutine leaks)

**Scheduler Analysis:**
4. How many OS threads (M) were used?
5. How many processors (P) were active?
6. Do you see work-stealing in action?

**Performance Analysis:**
7. What percentage of time is spent in:
   - Running goroutines
   - Waiting for channels
   - Garbage collection
8. Are there any blocking operations that could be optimized?

**Concurrency Patterns:**
9. Can you identify the worker pool pattern in the trace?
10. How long does each worker spend idle vs. working?

### 4. Optimization Exercise

Based on your trace analysis:
1. Identify one bottleneck
2. Propose an optimization
3. Implement it
4. Generate new trace and compare

**Deliverable:**
Write a short report (300-500 words) with:
- Screenshots of interesting trace views
- Answers to analysis questions
- Your optimization and its impact

---

## Submission Guidelines

For each assignment, submit:
1. **Source code** with comments explaining your approach
2. **Test output** showing your program works correctly
3. **README** explaining:
   - How to run your code
   - Design decisions you made
   - Challenges you encountered
   - What you learned

## Evaluation Criteria

- **Correctness:** Does it work as specified?
- **Concurrency Safety:** No race conditions (test with `-race`)
- **Resource Management:** Proper cleanup, no leaks
- **Code Quality:** Clear, idiomatic Go
- **Understanding:** Comments show you understand why, not just what

## Resources

- [Go Concurrency Patterns](https://go.dev/blog/pipelines)
- [Context Package](https://pkg.go.dev/context)
- [Execution Tracer](https://go.dev/blog/execution-tracer)
- [Effective Go - Concurrency](https://go.dev/doc/effective_go#concurrency)

## Getting Help

If stuck:
1. Review the demo code in this repository
2. Check the Go documentation
3. Use `go run -race` to detect issues
4. Ask specific questions with code examples

Good luck! 🚀