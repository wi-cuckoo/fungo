package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 4, 5, 8, 8, 10}
	fmt.Println(search(nums, 5))
}

func search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (l + r) / 2
		if target < nums[mid] {
			r = mid - 1
			continue
		}
		if target > nums[mid] {
			l = mid + 1
			continue
		}
		return mid
	}
	return 0
}
