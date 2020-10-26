/*
给定一个只包含正整数的非空数组。是否可以将这个数组分割成两个子集，使得两个子集的元素和相等。
注意:

每个数组中的元素不会超过 100, 数组的大小不会超过 200
示例 1:

输入: [1, 5, 11, 5]
输出: true

解释: 数组可以分割成 [1, 5, 5] 和 [11].


示例 2:

输入: [1, 2, 3, 5]
输出: false

解释: 数组不能分割成两个元素和相等的子集.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/partition-equal-subset-sum
*/

package four

func canPartition(nums []int) bool {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	if len(nums) < 2 || sum%2 != 0 {
		return false
	}
	target := sum / 2
	dp := make([][]bool, len(nums))
	for i := range dp {
		dp[i] = make([]bool, target+1)
	}
	for i := 0; i < len(nums); i++ {
		dp[i][0] = true
	}
	dp[0][nums[0]] = true
	for i := 1; i < len(nums); i++ {
		v := nums[i]
		for j := 1; j <= target; j++ {
			if j >= v {
				dp[i][j] = dp[i-1][j] || dp[i-1][j-v]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[len(nums)-1][target]
}
