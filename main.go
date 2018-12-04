package main

import (
	"fmt"
)

func main() {
	nums := make([]int, 0, 10)
	nums = append(nums, []int{3, 2, 4}...)
	fmt.Printf("nums: %v, cap: %d, addr: %p \n", nums, cap(nums), nums)
	push(nums)
	fmt.Printf("nums: %v, cap: %d, addr: %p \n", nums, cap(nums), nums)
}

func push(nums []int) {
	fmt.Printf("push nums: %v, len: %d, cap: %d, addr: %p \n", nums, len(nums), cap(nums), nums)
	nums = append(nums, 1)
	fmt.Printf("push nums: %v, len: %d, cap: %d, addr: %p \n", nums, len(nums), cap(nums), nums)
}
