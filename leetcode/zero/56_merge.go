package zero

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

import "sort"

import "github.com/wi-cuckoo/fungo/model"

// 排序就完事了
// 1. 根据区间的 Start 值对整个 intervals 排序
// 2. 考虑第一个区间 p 和第二个区间 q, 首先有 p.Start <= q.Start
//    - 当 p.End < q.Start 时，说明两区间没有重叠，则将 p 区间加入解集
//    - 当 p.End >= q.Start 时，说明区间重叠，则更新 p.End = q.End 合并两区间为
//    - 重复上述比较

func merge(intervals []model.Interval) []model.Interval {
	if len(intervals) < 2 {
		return intervals
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].Start < intervals[j].Start
	})
	result := make([]model.Interval, 0)
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
