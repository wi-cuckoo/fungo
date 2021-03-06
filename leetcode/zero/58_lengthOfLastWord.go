/*
给定一个仅包含大小写字母和空格 ' ' 的字符串，返回其最后一个单词的长度。

如果不存在最后一个单词，请返回 0 。

说明：一个单词是指由字母组成，但不包含任何空格的字符串。

示例:

输入: "Hello World"
输出: 5
*/

package zero

import (
	"strings"
)

func lengthOfLastWord(s string) int {
	arr := strings.Split(strings.TrimRight(s, " "), " ")
	return len(arr[len(arr)-1])
}

// no package
func lengthOfLastWordV2(s string) int {
	n := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' {
			if n != 0 {
				break
			}
			continue
		}
		n++
	}
	return n
}
