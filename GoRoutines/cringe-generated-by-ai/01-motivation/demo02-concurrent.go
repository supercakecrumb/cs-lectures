// Demo 2: Concurrent execution with WaitGroup
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	var wg sync.WaitGroup
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(i int) {
			time.Sleep(time.Second)
			fmt.Println("done", i)
			wg.Done()
		}(i)
	}
	// горутина стопорится пока wg > 0
	wg.Wait()
	fmt.Println("total:", time.Since(start))
}
