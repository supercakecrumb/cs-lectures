# Go Concurrency Lesson

A comprehensive, hands-on lesson on Go concurrency covering goroutines, channels, and synchronization primitives. This repository contains 12 progressive demos, detailed explanations, presentation slides, and practical homework assignments.

## 📚 What You'll Learn

- Why concurrency matters and when to use it
- Creating and managing goroutines safely
- Synchronization primitives (WaitGroup, Mutex, Cond)
- Channel-based communication patterns
- Worker pools and pipelines
- Race condition detection and prevention
- Real-world concurrency patterns

## 🎯 Prerequisites

- Basic Go knowledge (syntax, functions, structs)
- Go 1.21+ installed
- Familiarity with command line

## 📖 Repository Structure

```
.
├── README.md                    # This file
├── LESSON_PLAN.md              # Detailed lesson architecture
├── teacher-guide.md            # Teaching instructions and timing
├── slides.md                   # Presentation slides with diagrams
├── go.mod                      # Go module file
│
├── 01-motivation/              # Part 1: Why concurrency?
│   ├── README.md
│   ├── demo01-sequential.go    # Sequential delays (~3s)
│   └── demo02-concurrent.go    # Concurrent execution (~1s)
│
├── 02-goroutines/              # Part 2: Goroutine basics
│   ├── README.md
│   ├── demo03-lifecycle.go     # Goroutine lifecycle
│   └── demo04-closure-trap.go  # Loop variable capture bug
│
├── 03-synchronization/         # Part 3: Sync primitives
│   ├── README.md
│   ├── demo05-waitgroup.go     # WaitGroup patterns
│   ├── demo06-data-race.go     # Race condition example
│   ├── demo07-mutex.go         # Mutex solution
│   └── demo08-cond.go          # Condition variables
│
├── 04-channels/                # Part 4: Channels
│   ├── README.md
│   ├── demo09-unbuffered.go    # Unbuffered channels
│   ├── demo10-buffered.go      # Buffered channels
│   ├── demo11-closing.go       # Closing and range
│   └── demo12-worker-pool.go   # Worker pool pattern
│
└── homework/                   # Take-home assignments
    └── README.md               # 3 practical assignments
```

## 🚀 Quick Start

### Clone and Setup
```bash
git clone <repository-url>
cd go-concurrency-lesson
```

### Run Individual Demos
```bash
# Part 1: Motivation
cd 01-motivation
go run demo01-sequential.go
go run demo02-concurrent.go

# Part 2: Goroutines
cd ../02-goroutines
go run demo03-lifecycle.go
go run demo04-closure-trap.go

# Part 3: Synchronization
cd ../03-synchronization
go run demo05-waitgroup.go
go run demo06-data-race.go
go run -race demo06-data-race.go  # With race detector!
go run demo07-mutex.go
go run demo08-cond.go

# Part 4: Channels
cd ../04-channels
go run demo09-unbuffered.go
go run demo10-buffered.go
go run demo11-closing.go
go run demo12-worker-pool.go
```

### Run with Race Detector
**Always use the race detector during development:**
```bash
go run -race demo06-data-race.go
go test -race ./...
```

## 📝 Lesson Flow

### Part 1: Motivation (15 minutes)
**Goal:** See the performance benefit of concurrency

- **Demo 1:** Sequential delays waste time (~3 seconds)
- **Demo 2:** Concurrent execution is faster (~1 second)
- **Key Insight:** Goroutines eliminate waiting time

[→ Go to Part 1](01-motivation/)

---

### Part 2: Goroutines (15 minutes)
**Goal:** Understand goroutine lifecycle and common pitfalls

- **Demo 3:** Main doesn't wait for goroutines by default
- **Demo 4:** Loop variable capture bug (and fix)
- **Key Insight:** Always pass loop variables as parameters

[→ Go to Part 2](02-goroutines/)

---

### Part 3: Synchronization (20 minutes)
**Goal:** Master tools for coordinating goroutines

- **Demo 5:** Proper WaitGroup usage patterns
- **Demo 6:** Data race detection (use `-race` flag!)
- **Demo 7:** Mutex protects shared memory
- **Demo 8:** Condition variables for complex state
- **Key Insight:** Race detector is essential for development

[→ Go to Part 3](03-synchronization/)

---

### Part 4: Channels (25 minutes)
**Goal:** Master Go's idiomatic communication primitive

- **Demo 9:** Unbuffered channels = synchronization
- **Demo 10:** Buffered channels = decoupling
- **Demo 11:** Closing channels and range iteration
- **Demo 12:** Worker pool pattern (practical!)
- **Key Insight:** "Don't communicate by sharing memory; share memory by communicating"

[→ Go to Part 4](04-channels/)

---

## 🎓 For Teachers

### Teaching Materials
- **[Teacher's Guide](teacher-guide.md)** - Detailed lesson plan with timing, questions, and tips
- **[Slides](slides.md)** - Presentation slides with Mermaid diagrams
- **[Lesson Plan](LESSON_PLAN.md)** - Overall architecture and structure

### Recommended Approach
1. Start with motivation (show the problem, then the solution)
2. Live code each demo with student participation
3. Ask questions throughout (see teacher's guide)
4. Emphasize race detector usage
5. End with practical worker pool pattern

### Key Teaching Points
- Always show `-race` flag in action
- Emphasize "who closes which channel"
- Draw diagrams for worker pool
- Let students predict output before running
- Encourage experimentation

---

## 📚 For Students

### Study Path
1. Read each section's README before running demos
2. Run demos multiple times to see non-deterministic behavior
3. Experiment: modify code and predict outcomes
4. Complete homework assignments
5. Use `-race` flag for all your concurrent code

### Homework Assignments
Three practical assignments to solidify your learning:

1. **Rate Limiter** - Build concurrent rate limiter with semaphore pattern
2. **Pipeline** - Create 3-stage pipeline with context cancellation
3. **Trace Analysis** - Profile and optimize concurrent code

[→ View Homework](homework/)

---

## 🔑 Key Concepts

### Goroutines
```go
go function(args)  // Launch goroutine
```
- Lightweight (~2KB initial stack)
- Cheap to create (thousands easily)
- Scheduled by Go runtime

### Synchronization
```go
var wg sync.WaitGroup
wg.Add(1)
go func() {
    defer wg.Done()
    // work
}()
wg.Wait()
```

### Channels
```go
ch := make(chan int)      // Unbuffered
ch := make(chan int, 10)  // Buffered
ch <- value               // Send
value := <-ch             // Receive
close(ch)                 // Close (sender only!)
```

### Race Detection
```bash
go run -race main.go
go test -race ./...
```

---

## 🎯 Common Patterns

### Worker Pool (Fan-Out/Fan-In)
```go
jobs := make(chan int, 100)
results := make(chan int, 100)

// Start workers
for w := 1; w <= 3; w++ {
    go worker(jobs, results)
}

// Send jobs
for j := 1; j <= 10; j++ {
    jobs <- j
}
close(jobs)

// Collect results
for r := range results {
    fmt.Println(r)
}
```

### Pipeline
```go
gen := generate(nums...)
filtered := filter(gen)
squared := square(filtered)

for result := range squared {
    fmt.Println(result)
}
```

### Context Cancellation
```go
ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
defer cancel()

select {
case <-ctx.Done():
    return ctx.Err()
case result := <-ch:
    return result
}
```

---

## ⚠️ Common Pitfalls

### 1. Loop Variable Capture
```go
// ❌ WRONG
for i := 0; i < 5; i++ {
    go func() {
        fmt.Println(i)  // All print 5!
    }()
}

// ✅ CORRECT
for i := 0; i < 5; i++ {
    go func(i int) {
        fmt.Println(i)
    }(i)
}
```

### 2. Forgetting to Close Channels
```go
// ❌ WRONG - goroutine leak
go func() {
    for v := range ch {  // Hangs forever if ch never closed
        process(v)
    }
}()

// ✅ CORRECT
go func() {
    defer close(ch)
    for i := 0; i < 10; i++ {
        ch <- i
    }
}()
```

### 3. Closing from Wrong Side
```go
// ❌ WRONG - receiver closes
go func() {
    for v := range ch {
        process(v)
    }
    close(ch)  // Panic if sender still sending!
}()

// ✅ CORRECT - sender closes
go func() {
    defer close(ch)
    for i := 0; i < 10; i++ {
        ch <- i
    }
}()
```

---

## 📖 Resources

### Official Documentation
- [Effective Go - Concurrency](https://go.dev/doc/effective_go#concurrency)
- [Go Blog - Concurrency Patterns](https://go.dev/blog/pipelines)
- [Go Blog - Context](https://go.dev/blog/context)
- [Go Memory Model](https://go.dev/ref/mem)

### Tools
- [Race Detector](https://go.dev/blog/race-detector)
- [Execution Tracer](https://go.dev/blog/execution-tracer)
- [pprof Profiler](https://go.dev/blog/pprof)

### Videos
- "Concurrency is not Parallelism" by Rob Pike
- "Go Concurrency Patterns" by Rob Pike
- "Advanced Go Concurrency Patterns" by Sameer Ajmani

### Books
- "Concurrency in Go" by Katherine Cox-Buday
- "Go in Action" by William Kennedy
- "The Go Programming Language" by Donovan & Kernighan

---

## 🤝 Contributing

Found an issue or have a suggestion? Please open an issue or submit a pull request!

---

## 📄 License

This educational material is provided as-is for learning purposes.

---

## 🎓 About This Lesson

This lesson was designed to teach Go concurrency through progressive, hands-on examples. Each demo builds on the previous one, moving from basic concepts to practical patterns you'll use in real projects.

**Duration:** 90 minutes  
**Level:** Intermediate  
**Format:** Live coding with interactive discussions

---

## 🚀 Next Steps

1. Complete all demos in order
2. Experiment with modifications
3. Complete homework assignments
4. Build your own concurrent programs
5. Always use `-race` flag during development!

**Happy concurrent programming! 🎉**