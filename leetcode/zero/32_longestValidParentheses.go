/*
给定一个只包含 '(' 和 ')' 的字符串，找出最长的包含有效括号的子串的长度。
示例 1:

输入: "(()"
输出: 2
解释: 最长有效括号子串为 "()"
示例 2:

输入: ")()())"
输出: 4
解释: 最长有效括号子串为 "()()"

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-valid-parentheses
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package zero

import "fmt"

func longestValidParentheses(s string) int {
	if len(s) < 2 {
		return 0
	}
	// dp[i] means the max valid parenthesis which ends with s[i]
	dp := make([]int, len(s))
	dp[0] = 0
	if s[0] == '(' && s[1] == ')' {
		dp[1] = 2
	}
	max := dp[1]
	for i := 2; i < len(s); i++ {
		if s[i] == '(' {
			dp[i] = 0
			continue
		}
		l := i - dp[i-1] - 1
		if l == -1 {
			continue
		}
		if s[l] == '(' {
			dp[i] = dp[i-1] + 2
		}
		if l > 0 && dp[i] != 0 {
			dp[i] += dp[l-1]
		}
		if dp[i] > max {
			max = dp[i]
		}
	}
	fmt.Println(dp)
	return max
}
