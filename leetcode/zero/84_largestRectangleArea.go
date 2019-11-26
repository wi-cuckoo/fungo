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
	return 0
}
