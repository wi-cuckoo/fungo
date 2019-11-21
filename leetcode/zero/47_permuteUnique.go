package zero

import (
	"sort"
)

/*
给定一个可包含重复数字的序列，返回所有不重复的全排列。

示例:

输入: [1,1,2]
输出:
[
  [1,1,2],
  [1,2,1],
  [2,1,1]
]
*/

func permuteUnique(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}
	sort.Ints(nums)
	res := make([][]int, 0)
	visited := make([]int, len(nums))
	target := make([]int, 0, len(nums))
	backtrackPermuteUnique(nums, visited, target, &res)
	return res
}

func backtrackPermuteUnique(nums, visited, target []int, res *[][]int) {
	if len(nums) == len(target) {
		cp := make([]int, len(target))
		copy(cp, target)
		*res = append(*res, cp)
		return
	}
	for i := 0; i < len(nums); i++ {
		if visited[i] == 1 {
			continue
		}
		// 和之前的数相等，并且之前的数还未使用过，只有出现这种情况，才会出现相同分支
		// 这种情况跳过即可
		if i > 0 && nums[i] == nums[i-1] && visited[i-1] == 0 {
			continue
		}
		target = append(target, nums[i])
		visited[i] = 1
		backtrackPermuteUnique(nums, visited, target, res)
		visited[i] = 0
		target = target[:len(target)-1]
	}
}
