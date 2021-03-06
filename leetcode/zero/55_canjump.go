package zero

/*
给定一个非负整数数组，你最初位于数组的第一个位置。

数组中的每个元素代表你在该位置可以跳跃的最大长度。

判断你是否能够到达最后一个位置。

示例 1:

输入: [2,3,1,1,4]
输出: true
解释: 从位置 0 到 1 跳 1 步, 然后跳 3 步到达最后一个位置。
示例 2:

输入: [3,2,1,0,4]
输出: false
解释: 无论怎样，你总会到达索引为 3 的位置。但该位置的最大跳跃长度是 0 ， 所以你永远不可能到达最后一个位置。
*/

// 动态规划
// 如果只有一个节点，肯定能到
// 如果

func canJump(nums []int) bool {
	if len(nums) == 0 {
		return true
	}
	dp := make([]int, len(nums)) // 记录到节点 n 时还能向前跳跃多远
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			dp[i] = nums[i]
		} else {
			if dp[i-1]-1 > nums[i] {
				dp[i] = dp[i-1] - 1
				continue
			}
			dp[i] = nums[i]
		}

		if dp[i] == 0 && i != len(nums)-1 {
			return false
		}
	}
	return true
}
