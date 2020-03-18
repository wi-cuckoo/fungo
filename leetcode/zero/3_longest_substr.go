/*

给定一个字符串，找出不含有重复字符的最长子串的长度。

示例 1:

输入: "abcabcbb"
输出: 3
解释: 无重复字符的最长子串是 "abc"，其长度为 3。
示例 2:

输入: "bbbbb"
输出: 1
解释: 无重复字符的最长子串是 "b"，其长度为 1。
示例 3:

输入: "pwwkew"
输出: 3
解释: 无重复字符的最长子串是 "wke"，其长度为 3。
     请注意，答案必须是一个子串，"pwke" 是一个子序列 而不是子串。
*/

package zero

// slide window
func lengthOfLongestSubstring(s string) int {
	table := [128]int{}
	fmax := func(i, j int) int {
		if i > j {
			return i
		}
		return j
	}
	l, r, max := 0, 0, 0
	for r = 0; r < len(s); r++ {
		idx := s[r] - '\000'
		l = fmax(table[idx], l)
		max = fmax(max, r-l+1)
		table[idx] = r + 1 // 记录下该字母最后一次出现的位置索引
	}
	return max
}
