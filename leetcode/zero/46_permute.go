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

package zero

func permute(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}
	res := make([][]int, 0)
	backtrackPermute(nums, []int{}, &res)
	return res
}

func backtrackPermute(nums []int, target []int, res *[][]int) {
	if len(nums) == 0 {
		*res = append(*res, target)
		return
	}
	for i := 0; i < len(nums); i++ {
		newNums := make([]int, len(nums)-1)
		copy(newNums[0:i], nums[0:i])
		copy(newNums[i:], nums[i+1:])
		newTarget := append(target, nums[i])
		backtrackPermute(newNums, newTarget, res)
	}
}
