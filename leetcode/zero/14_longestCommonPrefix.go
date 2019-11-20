package zero

/*
编写一个函数来查找字符串数组中的最长公共前缀。

如果不存在公共前缀，返回空字符串 ""。

示例 1:

输入: ["flower","flow","flight"]
输出: "fl"
示例 2:

输入: ["dog","racecar","car"]
输出: ""
解释: 输入不存在公共前缀。
说明:

所有输入只包含小写字母 a-z
*/

// 简单，找出最短的那个字符串，然后迭代它
// 每迭代一个字符，检查其他字符串在相应位置的字符是否一致，直到不一致或者迭代完为止

func longestCommonPrefix(strs []string) string {
	var res []byte
	// find the shortest word
	shortest := ""
	for _, s := range strs {
		if s == "" {
			shortest = s
			break
		}
		if len(shortest) == 0 {
			shortest = s
			continue
		}
		if len(shortest) > len(s) {
			shortest = s
		}
	}
	for i, c := range shortest {
		flag := true
		for _, s := range strs {
			if byte(c) != s[i] {
				flag = false
				break
			}
		}
		if !flag {
			break
		}
		res = append(res, byte(c))
	}
	return string(res)
}
