/*
给出一个区间的集合，请合并所有重叠的区间。

示例 1:

输入: [[1,3],[2,6],[8,10],[15,18]]
输出: [[1,6],[8,10],[15,18]]
解释: 区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
示例 2:

输入: [[1,4],[4,5]]
输出: [[1,5]]
解释: 区间 [1,4] 和 [4,5] 可被视为重叠区间。
*/

package zero

import "sort"

func merge(intervals []Interval) []Interval {
	if len(intervals) < 2 {
		return intervals
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].Start < intervals[j].Start
	})
	result := make([]Interval, 0)
	p := intervals[0]
	for i := 1; i < len(intervals); i++ {
		q := intervals[i]
		if q.Start > p.End {
			result = append(result, p)
			p = q
			continue
		}
		if q.End > p.End {
			p.End = q.End
		}
	}
	result = append(result, p)
	return result
}