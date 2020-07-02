package five

/*

给定一个整数数组和一个整数 k，你需要找到该数组中和为 k 的连续的子数组的个数。

示例 1 :

输入:nums = [1,1,1], k = 2
输出: 2 , [1,1] 与 [1,1] 为两种不同的情况。
说明 :

数组的长度为 [1, 20,000]。
数组中元素的范围是 [-1000, 1000] ，且整数 k 的范围是 [-1e7, 1e7]。
*/

// 注意：滑动窗口不行，为啥不行，我想因为有负数，没法检测到所有子数组
// 思路： sum(i, j) = sum(0, j) - sum(0, i)
func subarraySum(nums []int, k int) int {
	freq := map[int]int{0: 1}
	sum, cnt := 0, 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		cnt += freq[sum-k]
		freq[sum]++ // 为了防止 k = 0 时引起查询混乱，该语句应该放循环体最后
	}
	return cnt
}
