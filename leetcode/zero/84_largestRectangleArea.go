package zero

import "github.com/wi-cuckoo/fungo/util"

/*
给定 n 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1 。

求在该柱状图中，能够勾勒出来的矩形的最大面积。
![](https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2018/10/12/histogram.png)

以上是柱状图的示例，其中每个柱子的宽度为 1，给定的高度为 [2,1,5,6,2,3]。
![](https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2018/10/12/histogram_area.png)

图中阴影部分为所能勾勒出的最大矩形面积，其面积为 10 个单位。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/largest-rectangle-in-histogram
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// 首先分冶法其实和暴力枚举没差，基本枚举了以每个柱子为高的最大矩形
// 这里最优解法是用单调栈来枚举高，我们的目的是对于每个柱子，以其高为矩形高
// 枚举出左右两边能最远到达且其高度不低于该柱子高度的位置

// 单调栈原理

func largestRectangleArea(heights []int) int {
	n := len(heights)
	left, right := make([]int, n), make([]int, n)
	stack := make([]int, 0, n/2)
	for i := 0; i < n; i++ {
		// update stack
		for len(stack) > 0 {
			ns := len(stack)
			if heights[stack[ns-1]] < heights[i] {
				break
			}
			stack = stack[:ns-1]
		}
		if len(stack) == 0 {
			left[i] = -1
		} else {
			left[i] = stack[len(stack)-1]
		}
		stack = append(stack, i)
	}
	stack = []int{}
	for i := n - 1; i >= 0; i-- {
		// update stack
		for len(stack) > 0 {
			ns := len(stack)
			if heights[stack[ns-1]] < heights[i] {
				break
			}
			stack = stack[:ns-1]
		}
		if len(stack) == 0 {
			right[i] = n
		} else {
			right[i] = stack[len(stack)-1]
		}
		stack = append(stack, i)
	}
	ans := 0
	for i := 0; i < n; i++ {
		ans = util.Max(ans, (right[i]-left[i]-1)*heights[i])
	}
	return ans
}
