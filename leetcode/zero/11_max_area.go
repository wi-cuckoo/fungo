package zero

/*
给定 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。
在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i, ai) 和 (i, 0)。
找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。

说明：你不能倾斜容器，且 n 的值至少为 2。

![image](https://aliyun-lc-upload.oss-cn-hangzhou.aliyuncs.com/aliyun-lc-upload/uploads/2018/07/25/question_11.jpg)

图中垂直线代表输入数组 [1,8,6,2,5,4,8,3,7]。在此情况下，容器能够容纳水（表示为蓝色部分）的最大值为 49。

示例:

输入: [1,8,6,2,5,4,8,3,7]
输出: 49
*/

// 众所周知，盛水多少取决于最短的那个边。
// 使用滑动双指针，对于 A[i,j] 围成的容器，其容量为 min(A[i],A[j])*(j-i)
// 此时如果 i < j, 则中间可能存在容量更大的容器, 更新原则是移动最短的那一边
// 为什么移动短的那一边的呢, 咱们用一下反证法
// 假设 A[i] < A[j], 并令 d = j-i, 则容量 V = A[i]*d
// 如果我们移动高的那一边，考虑最好的情况: A[j-1]=A[i]，则 V' = A[i]*(d-1)
// 显然 V' < V, 所以移动高的那一边没有意义，移动短的那一边则可能会有惊喜

func maxArea(height []int) int {
	l, r := 0, len(height)-1
	ms := 0
	for l < r {
		s := min(height[l], height[r]) * (r - l)
		if s > ms {
			ms = s
		}
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return ms
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
