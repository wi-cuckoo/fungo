/*
给定一个未排序的整数数组，找出其中没有出现的最小的正整数。

示例 1:

输入: [1,2,0]
输出: 3
示例 2:

输入: [3,4,-1,1]
输出: 2
示例 3:

输入: [7,8,9,11,12]
输出: 1
说明:

你的算法的时间复杂度应为O(n)，并且只能使用常数级别的空间。
*/

package zero

func firstMissingPositive(nums []int) int {
	oneIdx, l := -1, len(nums) // l 表示正整数个数
	for i := 0; i < l; i++ {
		if nums[i] <= 0 {
			l--
			nums[i], nums[l] = nums[l], nums[i] // 将负数推至末尾
		}
		if nums[i] <= 0 {
			i--
			continue
		}
		if nums[i] == 1 {
			oneIdx = i
		}
	}
	// 如果没有1出现，则最小正数肯定是1，直接返回
	if oneIdx < 0 {
		return 1
	}
	// 如果存在1，考虑一种极端情况，整个数组元素为 1,2,3....n，这个极限情况就是解题的关键
	// 先交换 1 到首位
	nums[0], nums[oneIdx] = nums[oneIdx], nums[0]
	for i := 1; i < l; i++ {
		n := nums[i]
		if i == n-1 {
			continue
		}
		if n == 0 || n > l || n == nums[n-1] {
			nums[i] = 0
			continue
		}
		// 如果是该数字小于 n，则需要交换
		nums[i], nums[n-1] = nums[n-1], nums[i]
		i--
	}
	for i := 1; i < l; i++ {
		if nums[i] == 0 {
			return i + 1
		}
	}
	return l + 1
}

// 我真是傻逼，明明想到了与极限情况对比，还做了这么麻烦
func firstMissingPositiveV2(nums []int) int {
	return 1
}
