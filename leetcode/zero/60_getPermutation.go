/*
给出集合 [1,2,3,…,n]，其所有元素共有 n! 种排列。

按大小顺序列出所有排列情况，并一一标记，当 n = 3 时, 所有排列如下：

"123"
"132"
"213"
"231"
"312"
"321"
给定 n 和 k，返回第 k 个排列。

说明：

给定 n 的范围是 [1, 9]。
给定 k 的范围是[1,  n!]。
示例 1:

输入: n = 3, k = 3
输出: "213"
示例 2:

输入: n = 4, k = 9
输出: "2314"
*/

package zero

func getPermutation(n int, k int) string {
	nums, total := make([]byte, 0, n), 1
	for i := 1; i <= n; i++ {
		nums = append(nums, byte(i)+'0')
		total *= i
	}
	// fmt.Println(total, string(nums))
	res := make([]byte, 0, n)
	for i := n; i > 0; i-- {
		t := total / i
		idx := k/t - 1
		if k%t > 0 {
			idx++
		}
		if idx < 0 {
			idx = 0
		}
		// fmt.Printf("total=%d k=%d i=%d t=%d idx=%d\n ", total, k, i, t, idx)
		res = append(res, nums[idx])
		total, k = t, k-t*idx
		nums = append(nums[0:idx], nums[idx+1:]...)
	}
	return string(res)
}
