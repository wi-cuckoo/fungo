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

// 维护一个上升队列
// 当遇到比队尾元素大的元素时，进行入队操作
// 当遇到比队尾元素小的元素时，通过二分查找找到一个可以被替换的元素并替换掉
// 这里有一个容易想不清楚的地方，就是替换后明显导致该队列并非真正的最长上升
// 子序列排列了。确实如此，但我们求的长度，所以可以证明这个长度时不会错的
func lengthOfLIS(nums []int) int {
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
