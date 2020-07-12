/*
给定两个大小为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。
请你找出这两个正序数组的中位数，并且要求算法的时间复杂度为 O(log(m + n))。
你可以假设 nums1 和 nums2 不会同时为空。

示例 1:

nums1 = [1, 3]
nums2 = [2]

则中位数是 2.0
示例 2:

nums1 = [1, 2]
nums2 = [3, 4]

则中位数是 (2 + 3)/2 = 2.5

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/median-of-two-sorted-arrays
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package zero

// 由中位数的定义知道当数组长度为奇数时中位数为中间那个数，偶数时为中间两数均值
// 设数组长度为 n，那么中位数即是寻找数组第 k 小的数，k = n/2 或 k = n/2+1
// 回到这个问题，也就转化为寻找两个有序数组中第 k 小的数，设两数组长度分别为
// m, n => k = (m+n)/2 或者 k = (m+n)/2+1

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	l1, l2 := len(nums1), len(nums2)
	k := (l1 + l2 + 1) / 2
	return float64(findKthNum(nums1, nums2, k))
}

func findKthNum(nums1 []int, nums2 []int, k int) int {
	if len(nums1) == 0 {
		return nums2[k-1]
	}
	if len(nums2) == 0 {
		return nums1[k-1]
	}
	mid := k / 2
	m1, m2 := nums1[len(nums1)-1], nums2[len(nums2)-1]
	if len(nums1) > mid {
		m1 = nums1[mid]
	}
	if len(nums2) > mid {
		m2 = nums2[mid]
	}
	return m1 - m2
}

// 方法二：
// 在 [0, m][0,m] 中找到最大的 i，使得：A[i−1]≤B[j]，其中 j = (m + n + 1)/2 - i
