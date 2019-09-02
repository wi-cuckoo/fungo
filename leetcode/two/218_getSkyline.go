/*
 天际线问题
 如果按照一个矩形一个矩形处理会非常麻烦，我们把这些矩形拆成两个点，一个左上角顶点，一个右上角顶点。
 将所有顶点按照横坐标进行排序然后开始遍历，遍历时通过一个堆来得知当前图形的最高位置，堆顶是所有顶点
 中最高的点，只要这个点没被移出堆，就说明这个最高的矩形还没结束。对于左顶点，我们将其加入堆中，对于
 右顶点，我们找出堆中相应的最顶点，然后移出左顶点，同时也意味着这个矩形的结束，为了区分左右顶点，我
 们以负数作为左顶点，正数作为右顶点
*/
package two

import (
	"sort"
)

func getSkyline(buildings [][]int) [][]int {
	dots := make([][]int, len(buildings)*2)
	for i, b := range buildings {
		dots[2*i] = []int{b[0], -b[2]}
		dots[2*i+1] = []int{b[1], b[2]}
	}
	sort.Slice(dots, func(i, j int) bool {
		a, b := dots[i], dots[j]
		if a[0] == b[0] {
			return a[1] < b[1]
		}
		return a[0] < b[0]
	})
	h, pre := NewPQ([]int{0}), 0
	res := make([][]int, 0)
	for _, d := range dots {
		x, y := d[0], d[1]
		if y < 0 {
			h.Push(-y)
		} else {
			// remove y
			h.Remove(y)
		}
		cur := h.Peek()
		if pre != cur {
			res = append(res, []int{x, cur})
			pre = cur
		}
	}
	return res
}

type pq struct {
	nums []int
}

// NewPQ ...
func NewPQ(nums []int) *pq {
	return &pq{
		nums: nums,
	}
}

func (p *pq) Push(n int) {
	p.nums = append(p.nums, n)
	for i := len(p.nums) - 2; i >= 0; i-- {
		if p.nums[i] < p.nums[i+1] {
			p.nums[i], p.nums[i+1] = p.nums[i+1], p.nums[i]
		}
	}
}

func (p *pq) Peek() int {
	return p.nums[0]
}

func (p *pq) Remove(n int) {
	l, r, mid := 0, len(p.nums)-1, 0
	for l <= r {
		mid = (l + r) / 2
		if p.nums[mid] == n {
			break
		}
		if p.nums[mid] > n {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	for i := mid; i > 0; i-- {
		p.nums[i] = p.nums[i-1]
	}
	p.nums = p.nums[1:]
}
