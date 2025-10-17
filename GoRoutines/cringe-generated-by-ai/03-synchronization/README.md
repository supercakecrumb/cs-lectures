# Part 3: Synchronization Primitives

## Goal
Master the tools for coordinating goroutines and protecting shared state.

## Demos

### Demo 5: WaitGroup Best Practices
**File:** [`demo05-waitgroup.go`](demo05-waitgroup.go:1)

**Run:**
```bash
go run demo05-waitgroup.go
```

**Teaching Points:**
- ✅ Call [`wg.Add(1)`](demo05-waitgroup.go:15) **before** starting goroutine
- ✅ Use [`defer wg.Done()`](demo05-waitgroup.go:17) for guaranteed cleanup
- ✅ Call [`wg.Wait()`](demo05-waitgroup.go:22) in main to block until completion

**Common Mistakes:**
```go
// ❌ WRONG: Add inside goroutine (race condition)
go func() {
    wg.Add(1)  // Might call Wait() before this executes
    defer wg.Done()
}()

// ❌ WRONG: Forgot Done() (deadlock)
wg.Add(1)
go func() {
    // Missing wg.Done()
}()
wg.Wait()  // Hangs forever

// ❌ WRONG: Double Done() (panic)
wg.Add(1)
go func() {
    wg.Done()
    wg.Done()  // panic: negative WaitGroup counter
}()
```

---

### Demo 6: Data Race Detection
**File:** [`demo06-data-race.go`](demo06-data-race.go:1)

**Run:**
```bash
go run demo06-data-race.go        # Shows incorrect result
go run -race demo06-data-race.go  # Detects the race
```

**Expected Output (without -race):**
```
x = 987  # Or some other value < 1000
```

**With Race Detector:**
```
==================
WARNING: DATA RACE
Write at 0x00c000014098 by goroutine 7:
  main.main.func1()
      demo06-data-race.go:14 +0x3c

Previous write at 0x00c000014098 by goroutine 6:
  main.main.func1()
      demo06-data-race.go:14 +0x3c
==================
```

**Teaching Points:**
- `x++` is **not atomic** (read-modify-write)
- Multiple goroutines can read same value
- Updates get lost (race condition)
- **Always use `-race` flag during development!**

**The Problem:**
```
Time  Goroutine A    Goroutine B    x value
----  -----------    -----------    -------
t1    read x (5)                    5
t2                   read x (5)     5
t3    x = 6                         5
t4    write 6                       6
t5                   x = 6          6
t6                   write 6        6  ← Lost update!
```

---

### Demo 7: Mutex Solution
**File:** [`demo07-mutex.go`](demo07-mutex.go:1)

**Run:**
```bash
go run -race demo07-mutex.go  # No race detected
```

**Expected Output:**
```
x = 1000  # Always correct
```

**Teaching Points:**
- [`sync.Mutex`](demo07-mutex.go:16) ensures mutual exclusion
- Only one goroutine in critical section at a time
- [`Lock()`](demo07-mutex.go:21) acquires, [`Unlock()`](demo07-mutex.go:23) releases
- Prefer `defer mu.Unlock()` for safety

**Pattern:**
```go
mu.Lock()
defer mu.Unlock()
// Critical section - only one goroutine here
x++
```

**When to use `sync.RWMutex`:**
- Many readers, few writers
- Readers don't block each other
- Writers still get exclusive access

```go
var mu sync.RWMutex

// Multiple readers can hold RLock simultaneously
mu.RLock()
value := sharedData
mu.RUnlock()

// Writer needs exclusive Lock
mu.Lock()
sharedData = newValue
mu.Unlock()
```

---

### Demo 8: Condition Variables
**File:** [`demo08-cond.go`](demo08-cond.go:1)

**Run:**
```bash
go run demo08-cond.go
```

**Expected Output:**
```
ready!  # After ~1 second
```

**Teaching Points:**
- [`sync.Cond`](demo08-cond.go:14) for waiting on complex conditions
- [`cond.Wait()`](demo08-cond.go:29) atomically unlocks and blocks
- [`cond.Signal()`](demo08-cond.go:23) wakes one waiter
- [`cond.Broadcast()`](demo08-cond.go:23) wakes all waiters
- **Always check condition in a loop** (spurious wakeups)

**Pattern:**
```go
mu.Lock()
for !condition {
    cond.Wait()  // Unlocks mu, waits, then relocks
}
// Condition is true and we hold the lock
mu.Unlock()
```

**When to use:**
- Complex state conditions
- Multiple goroutines waiting for same event
- Need fine-grained control over wakeups

**Note:** For most cases, **channels are simpler and more idiomatic!**

---

## Synchronization Primitives Comparison

| Primitive | Use Case | Example |
|-----------|----------|---------|
| **WaitGroup** | Wait for N goroutines to complete | Parallel batch processing |
| **Mutex** | Protect shared memory | Concurrent counter, cache |
| **RWMutex** | Many readers, few writers | Configuration, read-heavy cache |
| **Cond** | Complex state conditions | Producer-consumer with multiple conditions |
| **Channel** | Communication + sync | Pipeline, worker pool (see next section) |

---

## Key Lessons

### Race Detector is Essential
```bash
go test -race ./...
go run -race main.go
go build -race
```

### Mutex Best Practices
✅ **DO:**
- Keep critical sections small
- Use `defer mu.Unlock()`
- Consider RWMutex for read-heavy workloads

❌ **DON'T:**
- Hold locks during I/O or long operations
- Lock in one function, unlock in another
- Forget to unlock (deadlock)

### Go Proverb
> "Don't communicate by sharing memory; share memory by communicating."

Prefer channels over shared memory when possible!

---

## Transition to Next Section
"We've seen low-level synchronization. Now let's explore Go's idiomatic approach: channels..."