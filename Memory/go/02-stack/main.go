package main

import "fmt"

// Purpose: show stack frames via changing local addresses and the order of defer/panic/recover.
// Run: go run ./Memory/go/02-stack

func f2() {
	x := 200
	fmt.Printf("f2:   &x=%p\n", &x)

	defer fmt.Println("defer f2: runs first during unwind")
	panic("boom in f2")
}

func f1() {
	y := 100
	fmt.Printf("f1:   &y=%p\n", &y)
	defer fmt.Println("defer f1: runs after f2's defer")
	f2()
}

func main() {
	z := 42
	fmt.Printf("main: &z=%p\n", &z)

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover in main:", r)
		}
		fmt.Println("defer main: runs last")
	}()

	f1()
	// not reached (panic unwinds to main, where we recover)
}
