/*
给定一个字符串 S 和一个字符串 T，请在 S 中找出包含 T 所有字母的最小子串。

示例：

输入: S = "ADOBECODEBANC", T = "ABC"
输出: "BANC"
说明：

如果 S 中不存这样的子串，则返回空字符串 ""。
如果 S 中存在这样的子串，我们保证它是唯一的答案。
*/

package zero

import "fmt"

func minWindow(s string, t string) string {
	chars, cache := [26]int{}, [26]int{}
	for _, c := range t {
		n := int(c - 'A')
		chars[n]++
		cache[n]++
	}
	l, r := -1, -1
	total := len(t)
	for i := 0; i < len(s) && total > 0; i++ {
		n := int(s[i] - 'A')
		if chars[n] == 0 {
			continue
		}
		chars[n]--
		if l == -1 {
			l = i
		}
		total--
		r = i
	}
	if r < 0 {
		return ""
	}
	// 经过第一次从左至右循环，我们找到了第一个包含所有T的子串，但并非答案
	min := s[l : r+1]
	// 接着，我们需要通过窗口滑动来确定最终的答案
	fmt.Println(min)

	return min
}
