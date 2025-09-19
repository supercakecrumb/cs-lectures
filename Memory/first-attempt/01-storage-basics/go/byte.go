package main

import "fmt"

var a = 1

const b = 2

func main() {
	fmt.Println("1")

	defer fmt.Println("defer1")

	fmt.Println("2")

	defer fmt.Println("defer2")

	execDefer()

	fmt.Println("3")
}

func execDefer() {
	defer fmt.Println("execDefer")
}
