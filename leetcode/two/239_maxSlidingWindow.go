/*
给定一个数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口 k 内的数字。滑动窗口每次只向右移动一位。

返回滑动窗口最大值。

示例:

输入: nums = [1,3,-1,-3,5,3,6,7], 和 k = 3
输出: [3,3,5,5,6,7]
解释:

  滑动窗口的位置                最大值
---------------               -----
[1  3  -1] -3  5  3  6  7       3
 1 [3  -1  -3] 5  3  6  7       3
 1  3 [-1  -3  5] 3  6  7       5
 1  3  -1 [-3  5  3] 6  7       5
 1  3  -1  -3 [5  3  6] 7       6
 1  3  -1  -3  5 [3  6  7]      7
注意：

你可以假设 k 总是有效的，1 ≤ k ≤ 输入数组的大小，且输入数组不为空。

进阶：

你能在线性时间复杂度内解决此题吗？
*/

package two

import "container/list"

// 纯手搓实现，略慢，时间复杂度为 K + (N-K)*K
func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 || k == 1 {
		return nums
	}
	res := make([]int, len(nums)-k+1)
	max := 0
	for i := 1; i < k; i++ {
		if nums[i] > nums[max] {
			max = i
		}
	}
	res[0] = nums[max]
	for i := k; i < len(nums); i++ {
		// 如果窗口新加入的数是比上一个窗口最大值还大，直接更新就好了
		if nums[i] >= nums[max] {
			res[i-k+1], max = nums[i], i
			continue
		}
		// 如果新加入的值没那么大，也就是小于上一个窗口最大值
		if max > i-k {
			res[i-k+1] = nums[max]
			continue
		}
		// 将最大值索引往右移一位，并尝试寻找新的最大值
		max++
		for j := max + 1; j <= i; j++ {
			if nums[j] > nums[max] {
				max = j
			}
		}
		res[i-k+1] = nums[max]
	}
	return res
}

func maxSlidingWindowV2(nums []int, k int) []int {
	if len(nums) == 0 {
		return nums
	}
	result := make([]int, 0)
	window := list.New()
	for i, x := range nums {
		if i >= k && window.Front().Value.(int) <= i-k { // 移除头元素
			window.Remove(window.Front())
		}
		for window.Len() > 0 && nums[window.Back().Value.(int)] <= x { // 移除尾元素
			window.Remove(window.Back())
		}
		window.PushBack(i)
		if i >= k-1 {
			result = append(result, nums[window.Front().Value.(int)])
		}
	}
	return result
}
