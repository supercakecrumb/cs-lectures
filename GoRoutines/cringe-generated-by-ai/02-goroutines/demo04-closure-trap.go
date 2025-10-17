// Demo 4: Loop variable capture bug
// Purpose: Show the common mistake of capturing loop variables in closures
// Expected output: Often prints repeated values (like "i = 5" multiple times)
package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("=== BUGGY VERSION ===")
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println("i =", i) // BUG: captures loop variable by reference
		}()
	}
	wg.Wait()

	fmt.Println("\n=== FIXED VERSION ===")
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) { // FIX: pass i as parameter
			defer wg.Done()
			fmt.Println("i =", i)
		}(i)
	}
	wg.Wait()
}
