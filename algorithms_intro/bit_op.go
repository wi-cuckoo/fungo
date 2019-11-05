package main

import (
	"fmt"
)

func main() {
	bits()
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

func bits() {
	bs := [125e5]byte{}
	nums := []int{124e5, 133e4 + 1, 133e4 + 2, 10001}
	for _, n := range nums {
		idx, m := n/8, n%8
		b := bs[idx]
		bs[idx] = (1 << uint(m)) | b
		// fmt.Println(idx, bs[idx])
	}
	for i, b := range bs {
		if b == 0 {
			continue
		}
		// fmt.Println(i, b)
		for j := 0; j < 8; j++ {
			if m := b >> uint(j); m&1 == 1 {
				fmt.Println(i*8 + j)
			}
		}
	}
}
