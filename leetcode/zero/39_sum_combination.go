/*
给定一个无重复元素的数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。

candidates 中的数字可以无限制重复被选取。

说明：

所有数字（包括 target）都是正整数。
解集不能包含重复的组合。
示例 1:

输入: candidates = [2,3,6,7], target = 7,
所求解集为:
[
  [7],
  [2,2,3]
]
示例 2:

输入: candidates = [2,3,5], target = 8,
所求解集为:
[
  [2,2,2,2],
  [2,3,3],
  [3,5]
]
*/

package zero

var result = make([][]int, 0)

func combinationSum(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	backtrackSum(candidates, target, 0, []int{}, &res)
	return res
}

func backtrackSum(candidates []int, target, start int, maybe []int, res *[][]int) {
	// fmt.Println(maybe, target)
	if target < 0 {
		return
	}
	if target == 0 {
		*res = append(*res, maybe)
		// fmt.Printf("%v, %p\n", maybe, &maybe)
		return
	}
	for i := start; i < len(candidates); i++ {
		c := candidates[i]
		tmp := make([]int, len(maybe))
		copy(tmp, maybe)
		tmp = append(tmp, c)
		backtrackSum(candidates, target-c, i, tmp, res)
	}
}
