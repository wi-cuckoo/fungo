package main

import (
	"fmt"
)

func main() {
	fmt.Println(1 / 3)
}

func push(nums []int) {
	fmt.Printf("push nums: %v, len: %d, cap: %d, addr: %p \n", nums, len(nums), cap(nums), nums)
	nums = append(nums, 1)
	fmt.Printf("push nums: %v, len: %d, cap: %d, addr: %p \n", nums, len(nums), cap(nums), nums)
}
