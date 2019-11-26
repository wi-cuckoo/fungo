/*
 给定三个字符串 s1, s2, s3, 验证 s3 是否是由 s1 和 s2 交错组成的。

 示例 1:

 输入: s1 = "aabcc", s2 = "dbbca", s3 = "aadbbcbcac"
 输出: true
 示例 2:

 输入: s1 = "aabcc", s2 = "dbbca", s3 = "aadbbbaccc"
 输出: false
*/

package zero

// 动态规划
// 1. 考虑s1，s2的某个前缀能不能交错组成s3的一个前缀

func isInterleave(s1 string, s2 string, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}
	dp := make([][]bool, len(s1)+1)
	for i := 0; i <= len(s1); i++ {
		dp[i] = make([]bool, len(s2)+1)
		for j := 0; j <= len(s2); j++ {
			if i == 0 && j == 0 {
				dp[i][j] = true
				continue
			}
			if i == 0 {
				dp[i][j] = dp[i][j-1] && s3[i+j-1] == s2[j-1]
			} else if j == 0 {
				dp[i][j] = dp[i-1][j] && s3[i+j-1] == s1[i-1]
			} else {
				dp[i][j] = (dp[i-1][j] && s3[i+j-1] == s1[i-1]) || (dp[i][j-1] && s3[i+j-1] == s2[j-1])
			}
		}
	}
	return dp[len(s1)][len(s2)]
}
