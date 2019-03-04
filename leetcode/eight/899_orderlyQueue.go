/*
 给出了一个由小写字母组成的字符串 S。然后，我们可以进行任意次数的移动。

 在每次移动中，我们选择前 K 个字母中的一个（从左侧开始），将其从原位置移除，并放置在字符串的末尾。

 返回我们在任意次数的移动之后可以拥有的按字典顺序排列的最小字符串。



 示例 1：

 输入：S = "cba", K = 1
 输出："acb"
 解释：
 在第一步中，我们将第一个字符（“c”）移动到最后，获得字符串 “bac”。
 在第二步中，我们将第一个字符（“b”）移动到最后，获得最终结果 “acb”。
 示例 2：

 输入：S = "baaca", K = 3
 输出："aaabc"
 解释：
 在第一步中，我们将第一个字符（“b”）移动到最后，获得字符串 “aacab”。
 在第二步中，我们将第三个字符（“c”）移动到最后，获得最终结果 “aaabc”。


 提示：

 1 <= K <= S.length <= 1000
 S 只由小写字母组成。
*/

package eight

// 当 k=1 时，相当于从 i 处把字符串分隔再交换顺序拼接起来，因而选择最小字符出现的地方分隔
// 当 k>1 时，字符串可以通过操作得到任意序列，因为按字符从小到大排列即可
func orderlyQueue(S string, K int) string {
	min := 'z' + 1
	chars := [26]int{}
	for _, c := range S {
		if c < min {
			min = c
		}
		chars[int(c-'a')]++
	}
	if K == 1 {
		target := ""
		for i, c := range S {
			if c != min {
				continue
			}
			m := S[i:] + S[:i]
			if target == "" {
				target = m
			} else if m < target {
				target = m
			}
		}
		return target
	}
	str, index := make([]byte, len(S)), 0
	for c, n := range chars {
		char := byte(c + 'a')
		for i := 0; i < n; i++ {
			str[index] = char
			index++
		}
	}
	return string(str)
}
