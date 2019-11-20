// +build ignore,OMIT

package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 4, 5, 8, 8, 10}
	fmt.Println(search(nums, 1))
	fmt.Println(searchV2(nums, 2, 0, len(nums)-1))
}

// for version
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

// 递归版本
// 1. 基线条件：只剩下一个元素，直接比较即可
// 2. 确定如何缩小问题的规模，使其符合基线条件
func searchV2(nums []int, target, l, r int) int {
	if l > r {
		return 0
	}
	mid := (l + r) / 2
	switch {
	case target < nums[mid]:
		return searchV2(nums, target, l, mid-1)
	case target > nums[mid]:
		return searchV2(nums, target, mid+1, r)
	default:
		return mid
	}
}
