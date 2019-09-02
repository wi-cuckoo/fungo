/*
给定一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？找出所有满足条件且不重复的三元组。

注意：答案中不可以包含重复的三元组。

例如, 给定数组 nums = [-1, 0, 1, 2, -1, -4]，

满足要求的三元组集合为：
[
  [-1, 0, 1],
  [-1, -1, 2]
]
*/

package zero

import (
	"sort"
)

// # 双指针法
// 1. 先将 nums 排序
// 2. 遍历最小数到值不超过0的所有数字，对每个数字 nums[i] 寻找 [i+1:] 之间两数之和等于 -nums[i] 的解
//    - 取两个指针 l = i+1, r = len(nums)-1，然后遍历 [l:r]之间的数字找出满足条件的数字
//    - 如果找到的话，将两指针同时向内移动，并且需要跳过与其相同的数
func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return [][]int{}
	}
	sort.Ints(nums)
	set := make([][]int, 0)
	for i := 0; i < len(nums) && nums[i] <= 0; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l, r := i+1, len(nums)-1
		for l < r {
			if nums[i]+nums[l]+nums[r] == 0 {
				set = append(set, []int{nums[i], nums[l], nums[r]})
				for l < r && nums[l] == nums[l+1] {
					l++
				}
				for l < r && nums[r] == nums[r-1] {
					r--
				}
				l++
				r--
				continue
			}
			if nums[l]+nums[r] > -nums[i] {
				r--
			} else {
				l++
			}
		}
	}

	return set
}

// deprecated code
func de(nums []int) [][]int {
	if len(nums) < 3 {
		return [][]int{}
	}
	sort.Ints(nums)
	set := make([][]int, 0)
	l, r := 0, len(nums)-1
	for l < r {
		x, y := nums[l], nums[r]
		// fmt.Println(x, y)
		for i := l + 1; i < r; i++ {
			if x+y+nums[i] == 0 {
				set = append(set, []int{x, y, nums[i]})
				break
			}
		}
		if x+y >= 0 {
			for y == nums[r] && l < r {
				r--
			}
		} else {
			for x == nums[l] && l < r {
				l++
			}
		}
	}

	return set
}
