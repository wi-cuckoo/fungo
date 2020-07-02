package four

/*
给定一个包含大写字母和小写字母的字符串，找到通过这些字母构造成的最长的回文串。

在构造过程中，请注意区分大小写。比如 "Aa" 不能当做一个回文字符串。

注意:
假设字符串的长度不会超过 1010。

示例 1:

输入:
"abccccdd"

输出:
7

解释:
我们可以构造的最长的回文串是"dccaccd", 它的长度是 7。
*/

func longestPalindrome(s string) int {
	alpha := [26 * 2]int{}
	for i := 0; i < len(s); i++ {
		if 'A' <= s[i] && s[i] <= 'Z' {
			alpha[s[i]-'A'+26]++
			continue
		}
		alpha[s[i]-'a']++
	}
	cnt := 0
	var odd bool
	for i := 0; i < len(alpha); i++ {
		n := alpha[i]
		if n%2 == 0 {
			cnt += n
			alpha[i] = 0
		} else {
			cnt += n - 1
			odd = true
		}
	}
	if odd {
		cnt++
	}
	return cnt
}
