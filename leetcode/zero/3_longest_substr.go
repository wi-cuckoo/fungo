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

import "fmt"

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	// cache 记录某个字符最近一次出现的位置
	cache := map[byte]int{
		s[0]: 0,
	}
	// dp 每一项都是一个两位数组，记录以 s[i] 字母结尾的最长无重复子串的首位索引
	dp := make([][]int, len(s))
	dp[0] = []int{0, 0}
	for i := 1; i < len(s); i++ {
		idx, ok := cache[s[i]]
		if !ok || idx < dp[i-1][0] {
			dp[i] = []int{dp[i-1][0], i}
		} else {
			dp[i] = []int{idx + 1, i}
		}
		cache[s[i]] = i
	}
	fmt.Println(dp)
	l := 0
	for _, d := range dp {
		dl := d[1] - d[0] + 1
		if l < dl {
			l = dl
		}
	}
	return l
}
