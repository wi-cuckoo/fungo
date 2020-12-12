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

// 分冶法
func maxSubArrayV2(nums []int) int {
	return get(nums, 0, len(nums)-1).mSum
}

func pushUp(l, r Status) Status {
	iSum := l.iSum + r.iSum
	lSum := max(l.lSum, l.iSum+r.lSum)
	rSum := max(r.rSum, r.iSum+l.rSum)
	mSum := max(max(l.mSum, r.mSum), l.rSum+r.lSum)
	return Status{lSum, rSum, mSum, iSum}
}

func get(nums []int, l, r int) Status {
	if l == r {
		return Status{nums[l], nums[l], nums[l], nums[l]}
	}
	m := (l + r) >> 1
	lSub := get(nums, l, m)
	rSub := get(nums, m+1, r)
	return pushUp(lSub, rSub)
}

type Status struct {
	lSum, rSum, mSum, iSum int
}
