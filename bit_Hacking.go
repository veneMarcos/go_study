package main

import (
	"fmt"
)

// iota consecutives numbers
// _ numbers are ignored
const (
	// you can do it
	// _  = iota             // 0
	// KB = 1 << (iota * 10) // 1 << (1 * 10)
	// MB = 1 << (iota * 10) // 1 << (2 * 10)
	// GB = 1 << (iota * 10) // 1 << (3 * 10)
	// TB = 1 << (iota * 10) // 1 << (4 * 10)

	// or you can do like this
	_  = iota             // 0
	KB = 1 << (iota * 10) // 1 << (1 * 10)
	MB
	GB
	TB
)

func main() {
	fmt.Println("binary\t\t\t\tdecimal")
	fmt.Printf("%b\t\t\t", KB)
	fmt.Printf("%d\n", KB)
	fmt.Printf("%b\t\t", MB)
	fmt.Printf("%d\n", MB)
	fmt.Printf("%b\t", GB)
	fmt.Printf("%d\n", GB)
	fmt.Printf("%b\t", TB)
	fmt.Printf("%d\n", TB)
}
