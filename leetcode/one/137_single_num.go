/*
给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现了三次。找出那个只出现了一次的元素。

说明：

你的算法应该具有线性时间复杂度。 你可以不使用额外空间来实现吗？

示例 1:

输入: [2,2,3,2]
输出: 3
示例 2:

输入: [0,1,0,1,0,1,99]
输出: 99
*/

package main

import (
	"fmt"
)

func init() {
	a, b := 1, 19
	fmt.Println(a, b)
	a = a ^ b
	b = a ^ b
	a = a ^ b
	fmt.Println(a, b)
}

func main() {
	nums := []int{-2, -2, 1, 1, -3, 1, -3, -3, -4, -2}
	fmt.Println(singleNumber(nums))
}

func singleNumber(nums []int) int {
	m := 0
	for i := 0; i < 32; i++ {
		x := 0
		for _, n := range nums {
			x += (n >> uint(i)) & 1
		}
		m += (x % 3) << uint(i)
	}

	exist := false
	for _, n := range nums {
		if n == m {
			exist = true
			break
		}
	}

	if exist == false {
		return m - (1 << 32)
	}

	return m
}
