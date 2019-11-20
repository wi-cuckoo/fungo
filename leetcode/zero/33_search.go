package zero

/*
假设按照升序排序的数组在预先未知的某个点上进行了旋转。

( 例如，数组 [0,1,2,4,5,6,7] 可能变为 [4,5,6,7,0,1,2] )。

搜索一个给定的目标值，如果数组中存在这个目标值，则返回它的索引，否则返回 -1 。

你可以假设数组中不存在重复的元素。

你的算法时间复杂度必须是 O(log n) 级别。

示例 1:

输入: nums = [4,5,6,7,0,1,2], target = 0
输出: 4
示例 2:

输入: nums = [4,5,6,7,0,1,2], target = 3
输出: -1
*/

// 把数组一分为二，其中一个有序另一个可能有序也可能部分有序，分别为 L1， L2
// 对有序数组L1判定目标值是否在序列中，如果在则用二分查找之
// 如不在, 则继续将L2重复上述步骤

func search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	return searchWidth(nums, l, r, target)
}

func searchWidth(nums []int, l, r, target int) int {
	// fmt.Println(l, r, nums)
	if l > r {
		return -1
	}
	mid := (l + r) / 2
	if nums[mid] == target {
		return mid
	}
	// nums[mid:high] 是有序的
	if nums[mid] < nums[r] {
		if nums[mid] < target && target <= nums[r] {
			return searchWidth(nums, mid+1, r, target)
		}
		return searchWidth(nums, l, mid-1, target)
	}
	// nums[l:mid]是有序的
	if nums[l] <= target && target < nums[mid] {
		return searchWidth(nums, l, mid-1, target)
	}
	return searchWidth(nums, mid+1, r, target)
}
