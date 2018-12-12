/*
给出 n 代表生成括号的对数，请你写出一个函数，使其能够生成所有可能的并且有效的括号组合。

例如，给出 n = 3，生成结果为：

[
  "((()))",
  "(()())",
  "(())()",
  "()(())",
  "()()()"
]
*/

package zero

func generateParenthesis(n int) []string {
	res := make([]string, 0)
	backtrack("", 0, 0, n, &res)
	return res
}

func backtrack(s string, l, r, n int, res *[]string) {
	if len(s) == n*2 {
		*res = append(*res, s)
		return
	}
	if l < n {
		backtrack(s+"(", l+1, r, n, res)
	}
	if r < l {
		backtrack(s+")", l, r+1, n, res)
	}
}
