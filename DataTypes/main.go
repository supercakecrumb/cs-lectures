package main

import "fmt"

const (
	Badzim = 1 << 3 // 00001000₂
	Read   = 1 << 2 // 00000100₂
	Write  = 1 << 1 // 00000010₂
	Exec   = 1 << 0 // 00000001₂
)

func main() {
	var flags uint8 = 0b10101100
	hasRead := flags&Read != 0
	hasWrite := flags&Write != 0
	hasExec := flags&Exec != 0
	fmt.Printf("flags=%08b read=%t write=%t exec=%t\n", flags, hasRead, hasWrite, hasExec)

	// true  = 00000001
	// false = 00000000
}
