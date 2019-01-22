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
	l, r, total := -1, -1, len(t)
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
	// 重新计算第一个待选区间的目标字符个数
	for _, c := range s[l : r+1] {
		n := int(c - 'A')
		if cache[n] > 0 {
			chars[n]++
		}
	}
	// 接着，我们需要通过窗口滑动来确定最终的答案
	for valid := true; valid; {
		// 以min串左边开始向右滑动，检测到min不合法后停止
		var c byte
		for valid = true; valid; l++ {
			n := int(s[l] - 'A')
			if cache[n] == 0 {
				continue
			}
			chars[n]--
			// 当去掉该字符后，min串变得不合法
			if chars[n] < cache[n] {
				c, valid = s[l], false
			}
		}
		fmt.Println(s[l:r], l, r, c, chars)
		// 以min串右边向右滑动，使得min合法为止
		i := r
		for i++; i < len(s) && !valid; i++ {
			n := int(s[i] - 'A')
			if cache[n] == 0 {
				continue
			}
			chars[n]++
			if s[i] == c {
				valid = true
			}
		}
		if !valid {
			l--
		} else {
			r = i
		}
		fmt.Println(valid, s[l:r], l, r, chars)
		// 计算最新的min
		if len(s[l:r]) < len(min) {
			min = s[l:r]
		}
	}

	return min
}
