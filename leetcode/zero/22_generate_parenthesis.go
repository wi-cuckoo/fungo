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

// 回溯法
// 用变量 l, r 分别记录左右括号的数量
// 当 l < n 的时候可以给待选字符串加一个左括号；当 r < n 的时候可以给其加一个右括号
// 只要在序列保持有效的情况下，添加的括号都是成立的

func generateParenthesis(n int) []string {
	res := make([]string, 0)
	backtrackGen("", 0, 0, n, &res)
	return res
}

func backtrackGen(s string, l, r, n int, res *[]string) {
	if len(s) == n*2 {
		*res = append(*res, s)
		return
	}
	if l < n {
		backtrackGen(s+"(", l+1, r, n, res)
	}
	if r < l {
		backtrackGen(s+")", l, r+1, n, res)
	}
}
