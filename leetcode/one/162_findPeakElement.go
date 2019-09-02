/*
峰值元素是指其值大于左右相邻值的元素。

给定一个输入数组 nums，其中 nums[i] ≠ nums[i+1]，找到峰值元素并返回其索引。

数组可能包含多个峰值，在这种情况下，返回任何一个峰值所在位置即可。

你可以假设 nums[-1] = nums[n] = -∞。

示例 1:

输入: nums = [1,2,3,1]
输出: 2
解释: 3 是峰值元素，你的函数应该返回其索引 2。
示例 2:

输入: nums = [1,2,1,3,5,6,4]
输出: 1 或 5
解释: 你的函数可以返回索引 1，其峰值元素为 2；
     或者返回索引 5， 其峰值元素为 6。
说明:

你的解法应该是 O(logN) 时间复杂度的。
*/

package one

// 时间复杂度O(N)
func findPeakElement(nums []int) int {
	if len(nums) == 1 || nums[0] > nums[1] {
		return 0
	}
	for i := 1; i < len(nums)-1; i++ {
		if nums[i] > nums[i-1] && nums[i] > nums[i+1] {
			return i
		}
	}
	if nums[len(nums)-1] > nums[len(nums)-2] {
		return len(nums) - 1
	}
	return -1
}

// 分析一下
/*
 * 由于 nums[i] != nums[i+1]，表明相邻的两数字不相等
 * 对 i, j 满足 0 <= i < j <= len(nums)-1 && nums[i] < nums[j] 的情况下，必然在 nums[i:j+1] 之间能找到峰值
 * 可以假定， 如果从 i 到 j 是一个上升序列，则 j 是一个峰值; 如果是一个下降序列，则 i 是一个峰值，其余情况均存在峰值
 * 所以算法的目的就是使用二分去逼近这个尖刺
 */
func findPeakElementV2(nums []int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (l + r) / 2
		if mid == l {
			if nums[mid] > nums[r] {
				return l
			}
			return r
		}
		if nums[mid] < nums[mid+1] {
			l = mid
		}
		r = mid
	}
	return 0
}
