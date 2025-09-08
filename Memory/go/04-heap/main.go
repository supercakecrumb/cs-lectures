package main

import (
	"fmt"
	"runtime"
)

// Purpose: show heap allocation that outlives its function, slice growth (base may move),
// and Go's "memory clean" = garbage collection (GC).
// Run: go run ./Memory/go/04-heap

// NewInt returns a pointer that outlives this function.
// It escapes, so Go allocates it on the heap automatically.
func NewInt(v int) *int {
	x := v
	return &x // safe in Go; compiler keeps x alive on the heap
}

func main() {
	// 1) Heap object outlives creator
	p := NewInt(7)
	fmt.Printf("p=%p *p=%d\n", p, *p)

	// 2) Slice growth: underlying array can move as capacity increases
	s := make([]int, 0, 2)
	for i := 0; i < 6; i++ {
		s = append(s, i)
		base := "nil"
		if len(s) > 0 {
			base = fmt.Sprintf("%p", &s[0])
		}
		fmt.Printf("append %d -> len=%d cap=%d base=%s\n", i, len(s), cap(s), base)
	}

	// 3) Memory "cleaning" in Go = Garbage Collection (GC)
	//    GC automatically reclaims unreachable heap objects. You normally just drop references.
	var before, after runtime.MemStats
	runtime.ReadMemStats(&before)
	fmt.Printf("GC before: NumGC=%d, Alloc=%d bytes\n", before.NumGC, before.Alloc)

	// Create lots of short-lived allocations and then drop them.
	for i := 0; i < 1_000; i++ {
		b := make([]byte, 64*1024) // 64 KiB each
		_ = b                      // goes out of scope each loop
	}
	// Hint the collector for demo purposes (not for production usage).
	runtime.GC()

	runtime.ReadMemStats(&after)
	fmt.Printf("GC after:  NumGC=%d, Alloc=%d bytes\n", after.NumGC, after.Alloc)
}
