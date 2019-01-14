/*
给定一个整数数组 nums ，找出一个序列中乘积最大的连续子序列（该序列至少包含一个数）。

示例 1:

输入: [2,3,-2,4]
输出: 6
解释: 子数组 [2,3] 有最大乘积 6。
示例 2:

输入: [-2,0,-1]
输出: 0
解释: 结果不能为 2, 因为 [-2,-1] 不是子数组。
*/

package one

func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		if dp[i-1]*nums[i] > dp[i-1] {
			dp[i] = dp[i-1] * nums[i]
			continue
		}
		max, tmp := nums[i], nums[i]
		for l := i - 1; l >= 0; l-- {
			tmp *= nums[l]
			if tmp > max {
				max = tmp
			}
		}
		dp[i] = max
	}
	for i := 1; i < len(dp); i++ {
		if dp[i] > dp[0] {
			dp[0] = dp[i]
		}
	}
	return dp[0]
}
