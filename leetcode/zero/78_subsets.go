/*
给定一组不含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。

说明：解集不能包含重复的子集。

示例:

输入: nums = [1,2,3]
输出:
[
  [3],
  [1],
  [2],
  [1,2,3],
  [1,3],
  [2,3],
  [1,2],
  []
]
*/

package zero

// 回溯法+剪枝，没啥好说的

func subsets(nums []int) [][]int {
	res := [][]int{{}}
	backtrackSubsets(nums, []int{}, &res)
	return res
}

func backtrackSubsets(nums []int, target []int, res *[][]int) {
	if len(nums) == 0 {
		return
	}
	for i := 0; i < len(nums); i++ {
		t := make([]int, len(target))
		copy(t, target)
		t = append(t, nums[i])
		*res = append(*res, t)
		backtrackSubsets(nums[i+1:], t, res)
	}
}
