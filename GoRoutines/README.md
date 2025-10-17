# Go Goroutines and Concurrency Course

A comprehensive course on Go concurrency patterns, goroutines, channels, and synchronization primitives.

## 📁 Project Structure

```
.
├── demos/                  # Basic goroutine demonstrations
│   ├── 01-goroutines.go           # Basic goroutine usage
│   ├── 02-goroutines-anon.go      # Anonymous functions with goroutines
│   └── 03-mutex.go                # Mutex for synchronization
│
├── channels/              # Channel examples and patterns
│   ├── 01-basics.go              # Channel basics
│   ├── 02-buffered.go            # Buffered channels
│   ├── 03-directions.go          # Channel directions
│   ├── 04-ownership.go           # Channel ownership pattern
│   ├── 05-select.go              # Select statement
│   ├── 06-for-select.go          # For-select pattern
│   └── 07-range.go               # Range over channels
│
├── homework/              # Homework tasks and tests
│   ├── tasks.go                  # Task implementations (blank)
│   ├── tasks_test.go             # Tests and benchmarks
│   └── errors.go                 # Common errors to fix
│
└── Makefile              # Build and test automation
```

## 🚀 Getting Started

### Prerequisites
- Go 1.21 or higher
- Make (optional, for using Makefile commands)

### Running Examples

#### Demos
```bash
# Run all demos
make run-demos

# Or run individually
go run demos/01-goroutines.go
go run demos/02-goroutines-anon.go
go run demos/03-mutex.go
```

#### Channel Examples
```bash
# Run all channel examples
make run-channels

# Or run individually
go run channels/01-basics.go
go run channels/05-select.go
```

## 📚 Course Content

### 1. Demos (`demos/`)

#### [`01-goroutines.go`](demos/01-goroutines.go:1)
- Basic goroutine creation and execution
- Understanding concurrent execution

#### [`02-goroutines-anon.go`](demos/02-goroutines-anon.go:1)
- Anonymous functions with goroutines
- Good vs bad closure practices
- Loop variable capture issues

#### [`03-mutex.go`](demos/03-mutex.go:1)
- Race conditions
- Mutex for synchronization
- Safe concurrent access patterns

### 2. Channels (`channels/`)

#### [`01-basics.go`](channels/01-basics.go:1)
- Creating and using channels
- Sending and receiving
- Closing channels
- Broadcast signal pattern

#### [`02-buffered.go`](channels/02-buffered.go:1)
- Buffered vs unbuffered channels
- Deadlock examples
- Worker pool pattern
- Rate limiting
- Producer/Consumer pattern

#### [`03-directions.go`](channels/03-directions.go:1)
- Send-only channels (`chan<-`)
- Receive-only channels (`<-chan`)
- Pipeline pattern
- Type safety benefits

#### [`04-ownership.go`](channels/04-ownership.go:1)
- Channel ownership pattern
- Encapsulation with channels
- Safe channel lifecycle management

#### [`05-select.go`](channels/05-select.go:1)
- Select statement basics
- Non-blocking operations
- Timeout patterns
- Cancellation with done channels

#### [`06-for-select.go`](channels/06-for-select.go:1)
- Combining for and select
- Multiple channel handling
- Proper exit conditions

#### [`07-range.go`](channels/07-range.go:1)
- Ranging over channels
- Proper channel closing
- Early termination patterns

## 🎯 Homework Tasks

### Task 1: Parallel Computing
Implement `ParallelSum` and `SquareSum` functions that process data concurrently.

### Task 2: HTTP Concurrency
Implement `FetchURLs` to fetch multiple URLs concurrently with timeout support.

### Task 3: Pipeline Pattern
Create a 3-stage pipeline: generate → square → filter even numbers.

### Task 4: Worker Pool
Implement a worker pool pattern for concurrent job processing.

### Task 5: Rate Limiter
Build a rate-limited processor that respects throughput limits.

### Task 6: Fan-Out/Fan-In
Distribute work across workers and collect results.

### Task 7: Timeout Pattern
Process data with timeout constraints.

### Task 8: Semaphore Pattern
Limit concurrent operations using semaphore pattern.

## 🧪 Testing

### Run All Tests
```bash
make test
# or
go test ./homework -v
```

### Run Specific Tests
```bash
make test-task1    # Test ParallelSum
make test-task2    # Test FetchURLs
make test-task3    # Test ProcessPipeline
make test-task4    # Test WorkerPool
```

### Run with Race Detector
```bash
make race
# or
go test ./homework -race
```

### Generate Coverage Report
```bash
make coverage
# Opens coverage.html in browser
```

## 📊 Benchmarks

### Run All Benchmarks
```bash
make bench
# or
go test ./homework -bench=. -benchmem
```

### Run Specific Benchmarks
```bash
make bench-task1   # Benchmark ParallelSum
make bench-task4   # Benchmark WorkerPool
```

### Compare Worker Counts
```bash
go test ./homework -bench=BenchmarkParallelSum -benchmem
```

Expected output shows performance with different worker counts:
```
BenchmarkParallelSum/1worker-8     1000    1234567 ns/op
BenchmarkParallelSum/2workers-8    2000     654321 ns/op
BenchmarkParallelSum/4workers-8    3000     345678 ns/op
```

## 🐛 Common Errors (`homework/errors.go`)

The `errors.go` file contains 7 common concurrency bugs for students to fix:

1. **Easy**: Missing `WaitGroup.Wait()`
2. **Easy-Medium**: Loop variable capture bug
3. **Medium**: Channel not closed (deadlock)
4. **Medium-Hard**: Race condition on shared variable
5. **Hard**: Goroutine leak with blocked channel
6. **Hard**: Select without default in tight loop
7. **Hard**: Deadlock with mutual channel dependency

### How to Fix Errors

1. Read the code and comments
2. Identify the bug
3. Fix the issue
4. Test your fix

## 📖 Key Concepts

### Goroutines
- Lightweight threads managed by Go runtime
- Use `go` keyword to launch
- Always synchronize completion

### Channels
- Communication between goroutines
- Buffered vs unbuffered
- Always close when done sending
- Use directional channels for safety

### Synchronization
- `sync.WaitGroup` - Wait for goroutines to complete
- `sync.Mutex` - Protect shared data
- Channels - Communicate and synchronize

### Patterns
- **Worker Pool**: Fixed number of workers processing jobs
- **Pipeline**: Chain of processing stages
- **Fan-Out/Fan-In**: Distribute work, collect results
- **Timeout**: Limit operation duration
- **Rate Limiting**: Control throughput
- **Semaphore**: Limit concurrent operations

## 🎓 Learning Path

1. Start with `demos/` to understand basics
2. Study `channels/` examples in order
3. Implement homework tasks in `homework/tasks.go`
4. Run tests to verify implementations
5. Fix bugs in `homework/errors.go`
6. Run benchmarks to understand performance

## 🔧 Makefile Commands

```bash
make help          # Show all available commands
make test          # Run all tests
make test-verbose  # Run tests with verbose output
make bench         # Run all benchmarks
make race          # Run tests with race detector
make coverage      # Generate coverage report
make clean         # Clean test cache
make run-demos     # Run all demo files
make run-channels  # Run all channel examples
```

## 📝 Tips

1. **Always use WaitGroup** when launching goroutines
2. **Close channels** when done sending
3. **Pass loop variables** as parameters to goroutines
4. **Use mutex** for shared data access
5. **Test with race detector** (`-race` flag)
6. **Use buffered channels** to prevent blocking
7. **Handle timeouts** for long-running operations
8. **Use select** for multiple channel operations

## 🤝 Contributing

Feel free to submit issues and enhancement requests!

## 📄 License

This course material is provided for educational purposes.