package main

import (
	"fmt"
)

func main() {
	fmt.Println(oddOne(1 << 32))
}

// cacl count of 1 is odd or not
func oddOne(x int) int {
	fmt.Printf("%b\n", x)
	x ^= x >> 1
	fmt.Printf("%b\n", x)
	x ^= x >> 2
	fmt.Printf("%b\n", x)
	x ^= x >> 4
	fmt.Printf("%b\n", x)
	x ^= x >> 8
	fmt.Printf("%b\n", x)
	x ^= x >> 16
	fmt.Printf("%b\n", x)
	x ^= x >> 32
	fmt.Printf("%b\n", x)
	return x & 1
}

// cacl count of 1
func countOne(x int) int {
	return 0
}
