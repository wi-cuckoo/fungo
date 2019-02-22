/*
给定一些标记了宽度和高度的信封，宽度和高度以整数对形式 (w, h) 出现。当另一个信封的宽度和高度都比这个信封大的时候，这个信封就可以放进另一个信封里，如同俄罗斯套娃一样。

请计算最多能有多少个信封能组成一组“俄罗斯套娃”信封（即可以把一个信封放到另一个信封里面）。

说明:
不允许旋转信封。

示例:

输入: envelopes = [[5,4],[6,4],[6,7],[2,3]]
输出: 3
解释: 最多信封的个数为 3, 组合为: [2,3] => [5,4] => [6,7]。
*/

package three

import (
	"sort"
)

// 把问题转化成最长上升子序列的长度
// 考虑将信封按照宽大小排序使得宽有序，只要前后信封宽不相等则满足套娃需求，问题则转化成求一个高组成的最长上升子序列
// 这里需要注意的是，如果信封中宽都不相等，则求高的最长上升子序列即可
// 如果有宽相等，则在排序时需要根据其高逆序排列，避免了最后得到的最长上升子序列有误
func maxEnvelopes(envelopes [][]int) int {
	if len(envelopes) < 2 {
		return len(envelopes)
	}
	sort.Slice(envelopes, func(i, j int) bool {
		a, b := envelopes[i], envelopes[j]
		if a[0] < b[0] {
			return true
		}
		if a[0] == b[0] {
			return a[1] > b[1]
		}
		return false
	})
	nums := make([]int, len(envelopes))
	for i := 0; i < len(envelopes); i++ {
		nums[i] = envelopes[i][1]
	}
	queue := []int{nums[0]}
	for i := 1; i < len(envelopes); i++ {
		left, right := 0, len(queue)-1
		if nums[i] > queue[right] {
			queue = append(queue, nums[i])
			continue
		}
		// 二分查找寻找插入点
		for left < right {
			mid := (right + left) / 2
			if nums[i] > queue[mid] {
				left = mid + 1
			} else {
				right = mid
			}
		}
		queue[left] = nums[i]
	}
	return len(queue)
}
