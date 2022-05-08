package four

/*
给你一个长度为 n 的整数数组 nums ，其中 nums 的所有整数都在范围 [1, n] 内，且每个整数出现 一次 或 两次 。请你找出所有出现 两次 的整数，并以数组形式返回。
你必须设计并实现一个时间复杂度为 O(n) 且仅使用常量额外空间的算法解决此问题。

示例 1：

输入：nums = [4,3,2,7,8,2,3,1]
输出：[2,3]
示例 2：

输入：nums = [1,1,2]
输出：[1]
示例 3：

输入：nums = [1]
输出：[]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/find-all-duplicates-in-an-array
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// 由于数组元素取值范围均在 [1, n]，因此可以考虑使用计数排序，如元素 4，将其排到索引为 3 的位置
// 在排序过程中，如果目标位置的元素等于当前元素，且两者索引不一样，则可以判断是重复元素
// 需要注意的是，发现当前元素为重复元素后，需要将其赋一个特殊值标记，之后遇到该值的元素直接跳过
func findDuplicates(nums []int) []int {
	n := len(nums)
	var ans []int
	for i := 0; i < n; i++ {
		if nums[i] == -1 {
			continue
		}
		// 按顺序排列，nums[i] 应该出现的索引为 idx
		idx := nums[i] - 1
		// fmt.Println(nums, i, idx)
		if i == idx {
			continue
		}
		if nums[i] == nums[idx] {
			ans = append(ans, nums[i])
			nums[i] = -1
			continue
		}
		nums[i], nums[idx] = nums[idx], nums[i]
		i--
	}
	return ans
}
