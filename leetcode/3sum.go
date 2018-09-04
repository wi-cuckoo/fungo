package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum(nums))
}

// sort then find in both left and right
func threeSum(nums []int) [][]int {
	cache := map[string]bool{}
	res := make([][]int, 0)
	for i, a := range nums {
		for j, b := range nums[i+1:] {
			for _, c := range nums[i+j+2:] {
				if a+b+c == 0 {
					item := []int{a, b, c}
					sort.Ints(item)
					key := fmt.Sprintf("%d$%d$%d", item[0], item[1], item[2])
					if _, ok := cache[key]; ok {
						continue
					}
					cache[key] = true
					res = append(res, item)
				}
			}
		}
	}
	return res
}
