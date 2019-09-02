/*
一条包含字母 A-Z 的消息通过以下方式进行了编码：

'A' -> 1
'B' -> 2
...
'Z' -> 26
给定一个只包含数字的非空字符串，请计算解码方法的总数。

示例 1:

输入: "12"
输出: 2
解释: 它可以解码为 "AB"（1 2）或者 "L"（12）。
示例 2:

输入: "226"
输出: 3
解释: 它可以解码为 "BZ" (2 26), "VF" (22 6), 或者 "BBF" (2 2 6) 。
*/
package zero

// be careful of zero and zero with zero, fuck
func numDecodings(s string) int {
	if s == "" {
		return 0
	}
	if s[0] == '0' {
		return 0
	}
	dp := make([]int, len(s)+1)
	dp[0], dp[1] = 1, 1

	for i := 1; i < len(s); i++ {
		if s[i] == '0' && (s[i-1] > '2' || s[i-1] == '0') {
			return 0
		}
		if s[i] == '0' || s[i-1] == '0' {
			dp[i+1] = dp[i-1]
			continue
		}
		if s[i-1:i+1] <= "26" {
			dp[i+1] = dp[i-1] + dp[i]
			continue
		}
		dp[i+1] = dp[i]
	}
	return dp[len(s)]
}
