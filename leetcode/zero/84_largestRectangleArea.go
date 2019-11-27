package zero

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

// divide-and-conquer
// 将柱子一分为二，则最大矩形存在于三种可能
// 1，存在左半边区域，则对左半边区域继续二分
// 2. 存在右半边区域，则对右半边区域继续二分
// 3. 跨两边区域，这种情况由中间向两边展开，找出最大面积所在即可
// 最后对比三种情况所获得的最大矩形面积，取最大值即可

func largestRectangleArea(heights []int) int {
	if len(heights) == 0 {
		return 0
	}
	if len(heights) == 1 {
		return heights[0]
	}
	mid := len(heights) / 2
	left := largestRectangleArea(heights[:mid])
	right := largestRectangleArea(heights[mid:])
	cross := largestRectangleAreaCross(mid, heights)
	for _, n := range [2]int{left, right} {
		if n > cross {
			cross = n
		}
	}
	return cross
}

func largestRectangleAreaCross(mid int, heights []int) int {
	// 如果只有两根柱子，而又要求是跨区域，则面积就是 2*min(heights)
	lh := heights[mid-1]
	max := 0
	for i := mid - 1; i >= 0; i-- {
		if heights[i] < lh {
			lh = heights[i]
		}
		rh := lh
		for j := mid; j < len(heights); j++ {
			if heights[j] < rh {
				rh = heights[j]
			}
			if area := (j - i + 1) * rh; area > max {
				max = area
			}
		}
	}
	return max
}

// 另一种分治思路，找到最矮的柱子，则最大面积也存在三种分布情况
// 1. 包含最矮柱子，则此时面积为 minHeight * len(heights)
// 2. 存在于最矮柱子左边区域
// 3. 存在于最矮柱子右边区域
