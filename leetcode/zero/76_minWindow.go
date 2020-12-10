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

// 双指针+滑动窗口
// 1. 移动右指针 r，找到第一个包含 T 的子串为止，记为 min = S[l:r+1]
// 2. 移动左指针 l，直到某个位置导致 S[l:r+1] 不合法为止，更新 min
// 3. 接着移动右指针 r，重复步骤 1.

func minWindow(s string, t string) string {
	chars, cache := [128]int{}, [128]int{}
	for i := 0; i < len(t); i++ {
		cache[t[i]]++
	}
	check := func() bool {
		for i := 0; i < len(cache); i++ {
			if chars[i] < cache[i] {
				return false
			}
		}
		return true
	}
	min := len(s) + 1
	ansL, ansR := -1, -1
	for l, r := 0, 0; r < len(s); r++ {
		if cache[s[r]] > 0 {
			chars[s[r]]++
		}
		for check() && l <= r {
			if v := r - l + 1; v < min {
				min, ansL, ansR = v, l, l+v
			}
			if cache[s[l]] > 0 {
				chars[s[l]]--
			}
			l++
		}
	}
	if ansL == -1 {
		return ""
	}
	return s[ansL:ansR]
}
