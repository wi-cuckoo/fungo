package zero

/*
给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

示例:

输入: [-2,1,-3,4,-1,2,1,-5,4],
输出: 6
解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。
进阶:

如果你已经实现复杂度为 O(n) 的解法，尝试使用更为精妙的分治法求解。
*/

import "math"

// 1. 动态规划解
// * 只有一个数时，最大和显而易见
// * 当以第 i 个数结尾的连续子数组最大和为 sum[i], 则 第 i+1 数结尾时
//	 取 nums[i+1] 与 sum[i]+nums[i+1] 之间的最大者
// * 最后遍历sum数组，取最大值即可

func maxSubArray(nums []int) int {
	// dp[i] means the max sum of sub array that end with nums[i]
	sum := make([]int, len(nums))
	sum[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		sum[i] = max(sum[i-1]+nums[i], nums[i])
	}
	return max(sum...)
}

func max(nums ...int) int {
	n := nums[0]
	for _, v := range nums {
		if n < v {
			n = v
		}
	}
	return n
}

// 2. 分治法求解
// solution by divide-and-conquer
func maxSubArrayV2(nums []int) int {
	if len(nums) < 1 {
		return math.MinInt32
	}
	if len(nums) == 1 {
		return nums[0]
	}
	mid := (len(nums) - 1) / 2
	left := maxSubArrayV2(nums[:mid])
	right := maxSubArrayV2(nums[mid+1:])
	midSum := maxSubArrayCrossMid(mid, nums)
	return max(left, right, midSum)
}

func maxSubArrayCrossMid(mid int, nums []int) int {
	lsum, rsum := math.MinInt32, math.MinInt32
	sum := 0
	for i := mid; i >= 0; i-- {
		sum += nums[i]
		if lsum < sum {
			lsum = sum
		}
	}
	sum = 0
	for i := mid + 1; i < len(nums); i++ {
		sum += nums[i]
		if rsum < sum {
			rsum = sum
		}
	}
	if rsum < 0 {
		return lsum
	}
	return lsum + rsum
}
