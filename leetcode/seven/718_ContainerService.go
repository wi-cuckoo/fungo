/*
给两个整数数组 A 和 B ，返回两个数组中公共的、长度最长的子数组的长度。

示例：

输入：
A: [1,2,3,2,1]
B: [3,2,1,4,7]
输出：3
解释：
长度最长的公共子数组是 [3, 2, 1] 。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/maximum-length-of-repeated-subarray
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package seven

// 解法1：利用动态规划，考虑最长公共后缀，即 A[:i] 与 B[:j] 两字符串
// 当 A[i] == B[j] 时，两者有最长公共后缀，且等于 A[:i-1]B[:j-1]
// 的最长公共后缀加一，不等时为零

func findLength(A []int, B []int) int {
	max := 0
	dp := make([][]int, len(A))
	for i := 0; i < len(A); i++ {
		dp[i] = make([]int, len(B))
		for j := 0; j < len(B); j++ {
			if i == 0 || j == 0 {
				if A[i] == B[j] {
					dp[i][j] = 1
				}
				continue
			}
			if A[i] == B[j] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = 0
			}
		}
	}
	for i := 0; i < len(dp); i++ {
		for j := 0; j < len(dp[i]); j++ {
			if dp[i][j] > max {
				max = dp[i][j]
			}
		}
	}
	return max

}
