// Demo 1: Sequential delays
package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	for i := 1; i <= 3; i++ {
		time.Sleep(time.Second)
		fmt.Println("done", i)
	}
	fmt.Println("total:", time.Since(start))
}
