/*
给定一个数组，将数组中的元素向右移动 k 个位置，其中 k 是非负数。

示例 1:

输入: [1,2,3,4,5,6,7] 和 k = 3
输出: [5,6,7,1,2,3,4]
解释:
向右旋转 1 步: [7,1,2,3,4,5,6]
向右旋转 2 步: [6,7,1,2,3,4,5]
向右旋转 3 步: [5,6,7,1,2,3,4]
示例 2:

输入: [-1,-100,3,99] 和 k = 2
输出: [3,99,-1,-100]
解释:
向右旋转 1 步: [99,-1,-100,3]
向右旋转 2 步: [3,99,-1,-100]
说明:

尽可能想出更多的解决方案，至少有三种不同的方法可以解决这个问题。
要求使用空间复杂度为 O(1) 的原地算法。
*/

package one

// stupid way
func rotate(nums []int, k int) {
	l := len(nums)
	if k == 0 || l == 1 {
		return
	}
	if k > l {
		k = k - l
	}
	for i := 0; i < k; i++ {
		for j := l - 1; j > 0; j-- {
			nums[j], nums[j-1] = nums[j-1], nums[j]
		}
	}
}

func rotateV2(nums []int, k int) {
	if len(nums) < 2 {
		return
	}
	k %= len(nums)
	if k == 0 {
		return
	}
	// reverse all
	for i, j := 0, len(nums)-1; i < j; {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
	// reverse 0 ~ k
	for i, j := 0, k-1; i < j; {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
	for i, j := k, len(nums)-1; i < j; {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}
