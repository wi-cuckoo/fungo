/*
给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。

你的算法时间复杂度必须是 O(log n) 级别。

如果数组中不存在目标值，返回 [-1, -1]。

示例 1:

输入: nums = [5,7,7,8,8,10], target = 8
输出: [3,4]
示例 2:

输入: nums = [5,7,7,8,8,10], target = 6
输出: [-1,-1]
*/

package zero

// 二分法
// 先用普通二分查找找到target出现的位置，可能是好几个位置的其中一个，记为 p
// 从 p 位置分别向左和向右寻找相同的数，记最左与最右达到的位置为 p1, p2，答案即是 [p1, p2]

func searchRange(nums []int, target int) []int {
	p1, p2 := -1, -1
	for i := 0; i < len(nums); i++ {
		if nums[i] == target && p1 == -1 {
			p1, p2 = i, i
			continue
		}
		if nums[i] == target {
			p2 = i
		}
	}
	return []int{p1, p2}
}

// binary search
func searchRangeV2(nums []int, target int) []int {
	p1, p2 := -1, -1
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (right + left) / 2
		if target < nums[mid] {
			right = mid - 1
			continue
		}
		if target > nums[mid] {
			left = mid + 1
			continue
		}
		p1, p2 = mid, mid
		break
	}
	for j := p1 - 1; j >= 0; j-- {
		if target == nums[j] {
			p1 = j
			continue
		}
		break
	}
	for j := p2 + 1; j < len(nums); j++ {
		if target == nums[j] {
			p2 = j
			continue
		}
		break
	}
	return []int{p1, p2}
}
