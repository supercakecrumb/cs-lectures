// Demo 6: Data race on increment
// Purpose: Demonstrate race conditions with unprotected shared memory
// Expected output: x is less than 1000 (varies each run)
// Run with: go run -race demo06-data-race.go
package main

import (
	"fmt"
	"time"
)

func main() {
	var x int
	for i := 0; i < 1000; i++ {
		go func() {
			x++
		}()
	}
	time.Sleep(300 * time.Millisecond)
	fmt.Println("x =", x)
}

// What's happening:
// 1. Goroutine A reads x (e.g., x = 5)
// 2. Goroutine B reads x (still 5)
// 3. Goroutine A increments and writes 6
// 4. Goroutine B increments and writes 6 (lost update!)
//
// Run with race detector:
//   go run -race demo06-data-race.go
//
// Output will show:
//   WARNING: DATA RACE
//   Write at 0x... by goroutine X
//   Previous write at 0x... by goroutine Y
