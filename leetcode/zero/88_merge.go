/*
给定两个有序整数数组 nums1 和 nums2，将 nums2 合并到 nums1 中，使得 num1 成为一个有序数组。

说明:

初始化 nums1 和 nums2 的元素数量分别为 m 和 n。
你可以假设 nums1 有足够的空间（空间大小大于或等于 m + n）来保存 nums2 中的元素。
示例:

输入:
nums1 = [1,2,3,0,0,0], m = 3
nums2 = [2,5,6],       n = 3

输出: [1,2,2,3,5,6]
*/

package zero

import (
	"math"
	"sort"
)

// simple version
func mergeArray(nums1 []int, m int, nums2 []int, n int) {
	for i := m; i < m+n; i++ {
		nums1[i] = nums2[i-m]
	}
	sort.Ints(nums1)
}

// 双指针可解决，分别指向两个数组最末尾，每次取两个数组中最大值放到第一个数组最后面，然后指针前移
func mergeArrayV2(nums1 []int, m int, nums2 []int, n int) {
	t := m + n - 1
	m--
	n--
	for t > -1 {
		a, b := math.MinInt32, math.MinInt32
		if m > -1 {
			a = nums1[m]
		}
		if n > -1 {
			b = nums2[n]
		}
		max := a
		if b > max {
			max = b
			n--
		} else {
			m--
		}
		nums1[t] = max
		t--
	}
}
