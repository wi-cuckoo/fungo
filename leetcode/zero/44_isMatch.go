package zero

import "fmt"

/*
给定一个字符串 (s) 和一个字符模式 (p) ，实现一个支持 '?' 和 '*' 的通配符匹配。

'?' 可以匹配任何单个字符。
'*' 可以匹配任意字符串（包括空字符串）。
两个字符串完全匹配才算匹配成功。

说明:

s 可能为空，且只包含从 a-z 的小写字母。
p 可能为空，且只包含从 a-z 的小写字母，以及字符 ? 和 *。
示例 1:

输入:
s = "aa"
p = "a"
输出: false
解释: "a" 无法匹配 "aa" 整个字符串。
示例 2:

输入:
s = "aa"
p = "*"
输出: true
解释: '*' 可以匹配任意字符串。
示例 3:

输入:
s = "cb"
p = "?a"
输出: false
解释: '?' 可以匹配 'c', 但第二个 'a' 无法匹配 'b'。
示例 4:

输入:
s = "adceb"
p = "*a*b"
输出: true
解释: 第一个 '*' 可以匹配空字符串, 第二个 '*' 可以匹配字符串 "dce".
示例 5:

输入:
s = "acdcb"
p = "a*c?b"
输入: false

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/wildcard-matching
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// 用动态规划解
// 考虑 s 的前 i 个字符能否匹配上 p 的前 j 个字符，用一个二维数组来记录
// 有几种情况需要说明，初始的 s 为空或 p 为空以及两者都为空是比较简单
// 当 s 与 p 都不为空的时候，需要计算 dp[i][j] 分三种情况
// 其中 s[i-1] 表示 s 的第 i 个字符，p[j-1] 表示 p 的第 j 个字符
// * s[i-1] == p[j-1] || p[j-1] == '?'，则 dp[i][j] = dp[i-1][j-1]
// * p[j-1] == '*'，则 dp[i][j] 取决于 dp[i-1][j-1] || dp[i-1][j] || dp[i][j-1]
// * 其他，则 dp[i][j] 无法匹配，置为 false
func isMatch(s string, p string) bool {
	// dp[i][j] 表示 s 的前 i 个字符能否匹配上 p 的前 j 个字符
	dp := make([][]bool, len(s)+1)
	for i := 0; i < len(s)+1; i++ {
		dp[i] = make([]bool, len(p)+1)
		for j := 0; j < len(p)+1; j++ {
			// 都为空字符的时候
			if i == 0 && j == 0 {
				dp[i][j] = true
				continue
			}
			// s 为空，而 p 不为空的时候
			if i == 0 {
				dp[i][j] = dp[i][j-1] && p[j-1] == '*'
				continue
			}
			// s 不为空，而 p 为空的情况
			if j == 0 {
				dp[i][j] = false
				continue
			}
			// s 不为空，p 也不为空
			switch {
			case s[i-1] == p[j-1] || p[j-1] == '?':
				dp[i][j] = dp[i-1][j-1]
			case p[j-1] == '*':
				dp[i][j] = dp[i-1][j-1] || dp[i-1][j] || dp[i][j-1]
			default:
				dp[i][j] = false
			}
		}
	}
	fmt.Println(dp)
	return dp[len(s)][len(p)]
}
