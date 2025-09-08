package main

import "fmt"

// Purpose: show "static-like" lifetime using package-level vars (program-long)
// and a closure that persists across calls.
// Run: go run ./Memory/go/03-static

// Package-level variable: lives for the whole program (like a C static/global).
var counter int

func hit() {
	counter++
	fmt.Println("hit count =", counter)
}

// Closure-captured variable persists across calls to the returned function.
func makeCounter() func() int {
	c := 0
	return func() int { c++; return c }
}

func main() {
	hit()
	hit()
	hit() // 1, 2, 3

	next := makeCounter()
	fmt.Println("closure:", next(), next(), next()) // 1 2 3

	const version = "v1" // compile-time constant (read-only); here for contrast
	_ = version
}
