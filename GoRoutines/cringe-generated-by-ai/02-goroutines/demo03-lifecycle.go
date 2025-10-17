// Demo 3: Goroutine lifecycle
package main

import (
	"fmt"
	"time"
)

func main() {
	go fmt.Println("hello from goroutine")
	time.Sleep(10 * time.Nanosecond)
	fmt.Println("hello from main")
}
