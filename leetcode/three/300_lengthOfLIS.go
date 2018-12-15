/*
给定一个无序的整数数组，找到其中最长上升子序列的长度。

示例:

输入: [10,9,2,5,3,7,101,18]
输出: 4
解释: 最长的上升子序列是 [2,3,7,101]，它的长度是 4。
说明:

可能会有多种最长上升子序列的组合，你只需要输出对应的长度即可。
你算法的时间复杂度应该为 O(n2) 。
进阶: 你能将算法的时间复杂度降低到 O(n log n) 吗?
*/

package three

import "fmt"

func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp := make([]int, len(nums))
	dp[0] = 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			dp[i] = dp[i-1]
			continue
		}
		dp[i] = 1
		for j := i - 1; j >= 0; j-- {
			if nums[i] == nums[j] && dp[j] > dp[i] {
				dp[i] = dp[j]
			}
			if nums[i] > nums[j] && dp[j]+1 > dp[i] {
				dp[i] = dp[j] + 1
			}
		}
	}
	fmt.Println(dp)
	for i := 1; i < len(dp); i++ {
		if dp[i] > dp[0] {
			dp[0] = dp[i]
		}
	}
	return dp[0]
}

// fuck 动态规划是个骗局
// 维护一个上升队列
func lengthOfLISV2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	queue := []int{nums[0]}
	for i := 1; i < len(nums); i++ {
		left, right := 0, len(queue)-1
		if nums[i] > queue[right] {
			queue = append(queue, nums[i])
			continue
		}
		// 二分查找寻找插入点
		for left < right {
			mid := (right + left) / 2
			if nums[i] > queue[mid] {
				left = mid + 1
			} else {
				right = mid
			}
		}
		queue[left] = nums[i]
	}
	return len(queue)
}
