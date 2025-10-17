# Part 2: Goroutines - Basic Mechanics

## Goal
Understand goroutine lifecycle and common pitfalls.

## Demos

### Demo 3: Goroutine Lifecycle
**File:** [`demo03-lifecycle.go`](demo03-lifecycle.go:1)

**Run:**
```bash
go run demo03-lifecycle.go
```

**Possible Outputs:**
```
hello from main
```
or
```
hello from main
hello from goroutine
```
or
```
hello from goroutine
hello from main
```

**Teaching Points:**
- `main` function doesn't wait for goroutines
- When `main` exits, all goroutines are terminated
- Output order is non-deterministic
- Goroutine may not even start before `main` exits

**Question for Students:**
> "Run this 10 times. How many times do you see the goroutine message?"

**Key Lesson:**
We need explicit synchronization (like [`sync.WaitGroup`](../01-motivation/demo02-concurrent.go:13)) to control goroutine lifecycle.

---

### Demo 4: Closure Variable Capture Bug
**File:** [`demo04-closure-trap.go`](demo04-closure-trap.go:1)

**Run:**
```bash
go run demo04-closure-trap.go
```

**Expected Output (Buggy Version):**
```
=== BUGGY VERSION ===
i = 5
i = 5
i = 5
i = 5
i = 5

=== FIXED VERSION ===
i = 0
i = 1
i = 2
i = 3
i = 4
```

**Teaching Points:**
- **The Bug:** Goroutine closure captures loop variable `i` by reference
- All goroutines share the same `i` variable
- By the time goroutines execute, loop has finished and `i = 5`
- This is one of the most common Go concurrency bugs

**The Fix:**
```go
go func(i int) {  // Pass i as parameter
    defer wg.Done()
    fmt.Println("i =", i)
}(i)  // Creates a copy for each goroutine
```

**Alternative Fix (Go 1.22+):**
```go
for i := 0; i < 5; i++ {
    i := i  // Create loop-scoped copy
    go func() {
        defer wg.Done()
        fmt.Println("i =", i)
    }()
}
```

**Discussion Questions:**
> "Why does the buggy version sometimes print different values?"
> "What would happen if we added a sleep before starting goroutines?"

---

## Key Concepts

### Goroutine Properties
- **Lightweight:** ~2KB initial stack (vs ~1MB for OS threads)
- **Cheap to create:** Can easily spawn thousands
- **Scheduled by Go runtime:** M:N scheduling (M goroutines on N OS threads)
- **Non-preemptive at function boundaries:** Cooperative scheduling

### Common Patterns
✅ **DO:**
- Pass loop variables as parameters
- Use `defer wg.Done()` for cleanup
- Add to WaitGroup before starting goroutine

❌ **DON'T:**
- Capture loop variables in closures
- Forget to call `Done()`
- Call `Add()` inside the goroutine

---

## Transition to Next Section
"Now we understand goroutines, but what about shared data? Let's explore synchronization primitives..."