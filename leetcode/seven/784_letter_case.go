/*

给定一个字符串S，通过将字符串S中的每个字母转变大小写，我们可以获得一个新的字符串。返回所有可能得到的字符串集合。

示例:
输入: S = "a1b2"
输出: ["a1b2", "a1B2", "A1b2", "A1B2"]

输入: S = "3z4"
输出: ["3z4", "3Z4"]

输入: S = "12345"
输出: ["12345"]
注意：

S 的长度不超过12。
S 仅由数字和字母组成。
*/

package seven

func letterCasePermutation(S string) []string {
	res := make([]string, 0)
	backtrack(S, "", &res)
	return res
}

func backtrack(s, t string, res *[]string) {
	if len(s) == 0 {
		*res = append(*res, t)
		return
	}
	c := s[0]
	if c >= '0' && c <= '9' {
		backtrack(s[1:], t+s[0:1], res)
		return
	}
	c1 := c - 32
	if c >= 'A' && c <= 'Z' {
		c1 = c + 32
	}
	backtrack(s[1:], t+string(c), res)
	backtrack(s[1:], t+string(c1), res)
}
