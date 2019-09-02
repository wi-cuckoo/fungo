/*
给定一个整数 n，返回 n! 结果尾数中零的数量。

示例 1:

输入: 3
输出: 0
解释: 3! = 6, 尾数中没有零。
示例 2:

输入: 5
输出: 1
解释: 5! = 120, 尾数中有 1 个零.
说明: 你算法的时间复杂度应为 O(log n) 。
*/

package one

// cacl count of five
func trailingZeroes(n int) int {
	cnt := 0
	for i := 0; i <= n; i += 5 {
		if i >= 5 {
			for x := i / 5; x > 0; x = x / 5 {
				cnt++
				if x%5 != 0 {
					break
				}
			}
		}
	}
	return cnt
}
