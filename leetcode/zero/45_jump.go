/*
给定一个非负整数数组，你最初位于数组的第一个位置。

数组中的每个元素代表你在该位置可以跳跃的最大长度。

你的目标是使用最少的跳跃次数到达数组的最后一个位置。

示例:

输入: [2,3,1,1,4]
输出: 2
解释: 跳到最后一个位置的最小跳跃数是 2。
     从下标为 0 跳到下标为 1 的位置，跳 1 步，然后跳 3 步到达数组的最后一个位置。
说明:

假设你总是可以到达数组的最后一个位置。
*/

package zero

// 贪心解法
func jump(nums []int) int {
	if len(nums) < 2 {
		return 0
	}

	start, end, steps := 0, 0, 0
	for end < len(nums)-1 {
		steps++
		farthest := end
		for i := start; i < end+1; i++ {
			if nums[i]+i > farthest {
				farthest = nums[i] + i
			}
		}
		start = end + 1
		end = farthest
	}

	return steps
}

// 动态规划解法
func jumpV2(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	step, pos := 0, 0
	for {
		step++
		next := nums[pos]
		if next >= len(nums)-pos-1 {
			break
		}
		nextPos := next + pos

		// 找出相比 nextPos 更优的下一个跳跃点
		longer := 0
		for i := pos + 1; i <= nextPos; i++ {
			dis := nums[i] - nums[nextPos] - (nextPos - i)
			if dis >= longer {
				longer = dis
				pos = i
			}
		}
	}
	return step
}
