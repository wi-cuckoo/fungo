package zero

/*
给定一个没有重复数字的序列，返回其所有可能的全排列。

示例:

输入: [1,2,3]
输出:
[
  [1,2,3],
  [1,3,2],
  [2,1,3],
  [2,3,1],
  [3,1,2],
  [3,2,1]
]
*/

// 回溯法，每次只遍历未使用过的元素

func permute(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}
	res := make([][]int, 0)
	visited := make([]int, len(nums))
	target := make([]int, 0, len(nums))
	backtrackPermute(nums, visited, target, &res)
	return res
}

func backtrackPermute(nums, visited, target []int, res *[][]int) {
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
		target = append(target, nums[i])
		visited[i] = 1
		backtrackPermute(nums, visited, target, res)
		target = target[:len(target)-1]
		visited[i] = 0
	}
}
