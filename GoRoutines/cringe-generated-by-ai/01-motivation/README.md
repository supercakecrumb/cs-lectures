# Part 1: Motivation - Why Concurrency?

## Goal
Demonstrate the performance difference between sequential and concurrent execution.

## Demos

### Demo 1: Sequential Delays
**File:** [`demo01-sequential.go`](demo01-sequential.go:1)

**Run:**
```bash
go run demo01-sequential.go
```

**Expected Output:**
```
done 1
done 2
done 3
total: ~3s
```

**Teaching Points:**
- Sequential code waits for each operation to complete
- Total time = sum of all delays
- CPU is idle during sleep/I-O operations

**Question for Students:**
> "Can we do better? Can we fit this work into ~1 second?"

---

### Demo 2: Concurrent Execution
**File:** [`demo02-concurrent.go`](demo02-concurrent.go:1)

**Run:**
```bash
go run demo02-concurrent.go
```

**Expected Output:**
```
done 2
done 1
done 3
total: ~1s
```

**Teaching Points:**
- Goroutines run concurrently
- Total time ≈ longest single operation
- Order of output may vary (non-deterministic)
- [`sync.WaitGroup`](demo02-concurrent.go:13) ensures main waits for all goroutines

**Key Concepts Introduced:**
- `go` keyword launches a goroutine
- [`sync.WaitGroup`](demo02-concurrent.go:13) for synchronization
- [`wg.Add(1)`](demo02-concurrent.go:15) before starting goroutine
- [`defer wg.Done()`](demo02-concurrent.go:17) to signal completion
- [`wg.Wait()`](demo02-concurrent.go:22) blocks until all goroutines finish

**Discussion:**
> "Where in real systems do we see similar patterns?"
> - Web servers handling multiple requests
> - Batch processing of files
> - Parallel API calls
> - Database query fan-out

---

## Transition to Next Section
"Now that we've seen the power of goroutines, let's understand how they work under the hood..."