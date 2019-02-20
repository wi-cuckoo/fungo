/*
给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。

示例 1：

输入: "babad"
输出: "bab"
注意: "aba" 也是一个有效答案。
示例 2：

输入: "cbbd"
输出: "bb"
*/

package zero

// 中心扩散算法
// 记录下其长度并与之前得到的最大长度比较，最终值为所需解
// 找到形如 bb 型或 aba 型字符串，向两边扩展
// 时间复杂度 O(N^2)，空间复杂度 O(1)

func longestPalindrome(s string) string {
	l := len(s)
	if l < 2 {
		return s
	}
	r := s[0:1]
	for i := 0; i < l-1; i++ {
		// ..bb..
		if s[i] == s[i+1] {
			left, right := i-1, i+2
			for left >= 0 && right < l {
				if s[left] != s[right] {
					break
				}
				left--
				right++
			}
			if right-left-1 > len(r) {
				r = s[left+1 : right]
			}
		}
		// ...aba...
		if i > 0 && s[i-1] == s[i+1] {
			left, right := i-2, i+2
			for left >= 0 && right < l {
				if s[left] != s[right] {
					break
				}
				left--
				right++
			}
			if right-left-1 > len(r) {
				r = s[left+1 : right]
			}
		}
	}
	return r
}
