# Part 4: Channels - Idiomatic Communication

## Goal
Master Go's primary concurrency primitive: channels for safe communication between goroutines.

## Go Proverb
> "Don't communicate by sharing memory; share memory by communicating."

Channels are Go's way of implementing this philosophy.

---

## Demos

### Demo 9: Unbuffered Channels
**File:** [`demo09-unbuffered.go`](demo09-unbuffered.go:1)

**Run:**
```bash
go run demo09-unbuffered.go
```

**Expected Output:**
```
Main: about to receive
Goroutine: about to send
Goroutine: sent successfully
Main: received 42
```

**Teaching Points:**
- Unbuffered channel: `make(chan T)` or `make(chan T, 0)`
- **Synchronization point:** both sender and receiver must be ready
- Send blocks until receive is ready
- Receive blocks until send is ready
- This is a "rendezvous" pattern

**Visualization:**
```
Sender                Channel              Receiver
------                -------              --------
ch <- 42  ------>    [empty]  ------>     v := <-ch
(blocks)              (no buffer)          (blocks)

Both must be ready simultaneously!
```

**Deadlock Example:**
```go
ch := make(chan int)
ch <- 42  // Deadlock! No receiver
```

---

### Demo 10: Buffered Channels
**File:** [`demo10-buffered.go`](demo10-buffered.go:1)

**Run:**
```bash
go run demo10-buffered.go
```

**Expected Output:**
```
a
b
```

**Teaching Points:**
- Buffered channel: `make(chan T, capacity)`
- Sender doesn't block until buffer is full
- Receiver doesn't block until buffer is empty
- Decouples sender and receiver timing

**Buffer Size Guidelines:**
- **Size 0** (unbuffered): Tight synchronization, rendezvous
- **Size 1**: Simple async handoff, "mailbox" pattern
- **Size N**: Batch processing, rate limiting, burst handling

**Common Pattern - Semaphore:**
```go
sem := make(chan struct{}, 3) // Max 3 concurrent

for i := 0; i < 10; i++ {
    sem <- struct{}{}  // Acquire (blocks if 3 running)
    go func(i int) {
        defer func() { <-sem }()  // Release
        // Do work with max 3 concurrent
    }(i)
}
```

---

### Demo 11: Closing Channels
**File:** [`demo11-closing.go`](demo11-closing.go:1)

**Run:**
```bash
go run demo11-closing.go
```

**Expected Output:**
```
1
2
3
After close: v=0, ok=false
```

**Teaching Points:**
- [`close(ch)`](demo11-closing.go:14) signals "no more values"
- [`range`](demo11-closing.go:17) reads until channel is closed
- Reading from closed channel returns zero value + `ok=false`
- **Only sender should close** (receiver doesn't know if more senders exist)

**Channel Closing Rules:**
| Operation | Closed Channel | Result |
|-----------|----------------|--------|
| Send | `ch <- v` | **PANIC** |
| Receive | `v := <-ch` | Zero value, ok=false |
| Close | `close(ch)` | **PANIC** |
| Range | `for v := range ch` | Exits loop |

**Pattern:**
```go
func producer(ch chan int) {
    defer close(ch)  // Ensure close on return
    for i := 0; i < 10; i++ {
        ch <- i
    }
}

func consumer(ch chan int) {
    for v := range ch {  // Exits when closed
        fmt.Println(v)
    }
}
```

**When NOT to close:**
- Multiple senders (who closes?)
- Unclear ownership
- Channel is just for signaling (use `context` instead)

---

### Demo 12: Worker Pool Pattern
**File:** [`demo12-worker-pool.go`](demo12-worker-pool.go:1)

**Run:**
```bash
go run demo12-worker-pool.go
```

**Expected Output:**
```
worker 1 processing job 1
worker 2 processing job 2
worker 3 processing job 3
result: 2
worker 1 processing job 4
result: 4
worker 2 processing job 5
result: 6
worker 3 done
result: 8
worker 1 done
result: 10
worker 2 done
```

**Teaching Points:**
- **Fan-out:** Multiple workers read from jobs channel
- **Fan-in:** Multiple workers write to results channel
- [`WaitGroup`](demo12-worker-pool.go:15) tracks worker completion
- Separate goroutine closes results after all workers finish

**Architecture:**
```
Producer
   |
   v
[jobs chan] ---> Worker 1 ---> [results chan]
            ---> Worker 2 ---> 
            ---> Worker 3 --->
                                    |
                                    v
                                Consumer
```

**Critical Flow:**
1. Create jobs and results channels
2. Start N workers (each reads from jobs, writes to results)
3. Producer sends jobs and closes jobs channel
4. Workers process until jobs is closed
5. Separate goroutine waits for all workers, then closes results
6. Consumer reads results until closed

**Common Mistakes:**
```go
// ❌ WRONG: Close results in main (before workers finish)
close(results)
wg.Wait()

// ❌ WRONG: Don't close jobs (workers hang forever)
for j := 1; j <= 5; j++ {
    jobs <- j
}
// Missing: close(jobs)

// ❌ WRONG: Close results in worker (multiple closes = panic)
func worker(...) {
    defer close(results)  // Multiple workers = multiple closes!
    ...
}
```

**Real-World Applications:**
- Parallel HTTP requests to multiple APIs
- Batch database operations
- Image/video processing pipeline
- Log aggregation and processing
- Web scraping with rate limiting

---

## Channel Patterns Summary

### 1. Pipeline
Chain of stages connected by channels:
```go
func gen(nums ...int) <-chan int {
    out := make(chan int)
    go func() {
        defer close(out)
        for _, n := range nums {
            out <- n
        }
    }()
    return out
}

func sq(in <-chan int) <-chan int {
    out := make(chan int)
    go func() {
        defer close(out)
        for n := range in {
            out <- n * n
        }
    }()
    return out
}

// Usage: gen -> sq -> print
for n := range sq(gen(2, 3, 4)) {
    fmt.Println(n)  // 4, 9, 16
}
```

### 2. Fan-Out, Fan-In
Multiple goroutines read from same channel (fan-out), multiple write to same channel (fan-in).
See [`demo12-worker-pool.go`](demo12-worker-pool.go:1).

### 3. Select for Multiplexing
```go
select {
case v := <-ch1:
    fmt.Println("received from ch1:", v)
case v := <-ch2:
    fmt.Println("received from ch2:", v)
case <-time.After(time.Second):
    fmt.Println("timeout")
}
```

### 4. Quit Channel
```go
quit := make(chan struct{})
go func() {
    for {
        select {
        case <-quit:
            return
        default:
            // Do work
        }
    }
}()
// Later: close(quit)
```

---

## Channel Direction Annotations

Improve type safety by specifying direction:

```go
// Send-only channel
func producer(ch chan<- int) {
    ch <- 42
    // v := <-ch  // Compile error!
}

// Receive-only channel
func consumer(ch <-chan int) {
    v := <-ch
    // ch <- 42  // Compile error!
}

// Bidirectional (can be passed to either)
ch := make(chan int)
go producer(ch)
go consumer(ch)
```

---

## Best Practices

### ✅ DO:
- Use channels for communication between goroutines
- Close channels from sender side
- Use `range` to read until close
- Use buffered channels for known capacity
- Use channel direction annotations for clarity

### ❌ DON'T:
- Share channels between unrelated components
- Close from receiver side
- Close channels with multiple senders (use sync primitives)
- Forget to close channels in pipelines (goroutine leak)
- Over-buffer (wastes memory)

---

## Transition to Homework
"You now have all the tools! Time to build something real..."