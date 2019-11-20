package zero

/*

给定一个数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。

candidates 中的每个数字在每个组合中只能使用一次。

说明：

所有数字（包括目标数）都是正整数。
解集不能包含重复的组合。
示例 1:

输入: candidates = [10,1,2,7,6,1,5], target = 8,
所求解集为:
[
  [1, 7],
  [1, 2, 5],
  [2, 6],
  [1, 1, 6]
]
示例 2:

输入: candidates = [2,5,2,1,2], target = 5,
所求解集为:
[
  [1,2,2],
  [5]
]
*/

import "sort"

func combinationSumX(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	sort.Ints(candidates)
	backtrackS(candidates, target, 0, []int{}, &res)
	return res
}

func backtrackS(candidates []int, target, start int, maybe []int, res *[][]int) {
	// fmt.Println(maybe, target)
	if target < 0 {
		return
	}
	if target == 0 {
		*res = append(*res, maybe)
		return
	}
	for i := start; i < len(candidates); i++ {
		// 避免重复解
		if i > start && candidates[i] == candidates[i-1] {
			continue
		}
		c := candidates[i]
		tmp := make([]int, len(maybe))
		copy(tmp, maybe)
		tmp = append(tmp, c)
		backtrackS(candidates, target-c, i+1, tmp, res)
	}
}
