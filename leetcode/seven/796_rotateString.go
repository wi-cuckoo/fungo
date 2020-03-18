/*
给定两个字符串, A 和 B。

A 的旋转操作就是将 A 最左边的字符移动到最右边。 例如, 若 A = 'abcde'，在移动一次之后结果就是'bcdea' 。如果在若干次旋转操作之后，A 能变成B，那么返回True。

示例 1:
输入: A = 'abcde', B = 'cdeab'
输出: true

示例 2:
输入: A = 'abcde', B = 'abced'
输出: false
注意：

A 和 B 长度不超过 100。
*/

package seven

import (
	"strings"
)

func rotateString(A string, B string) bool {
	for i := 0; i < len(A); i++ {
		if A[:i] == B[len(B)-i:] && A[i:] == B[:len(B)-i] {
			return true
		}
	}
	return false
}

// 评论区看到一个python答案，简直想找个地缝钻进去
// return len(A) == len(B) and B in A+A
func rotateStringV2(A string, B string) bool {
	return len(A) == len(B) && strings.Contains(A+A, B)
}
