/*
给定一个以字符串表示的非负整数 num，移除这个数中的 k 位数字，使得剩下的数字最小。

注意:

num 的长度小于 10002 且 ≥ k。
num 不会包含任何前导零。
示例 1 :

输入: num = "1432219", k = 3
输出: "1219"
解释: 移除掉三个数字 4, 3, 和 2 形成一个新的最小的数字 1219。
*/

package four

import (
	"fmt"
)

func removeKdigits(num string, k int) string {
	// 啥都不用移除
	if k == 0 {
		return num
	}
	// 数字全部移除，damn it
	if k >= len(num) {
		return "0"
	}
	stack := make([]byte, 0, len(num))
	stack = append(stack, num[0])
	for i := 1; i < len(num); i++ {
		// 如果当前数字比上一个数字大，则入栈
		if num[i] >= num[i-1] {
			stack = append(stack, num[i])
			continue
		}
		// 否则，将栈中大于该数字的出栈，出栈的则是需要删除的
		var j int
		for j = len(stack) - 1; j >= 0; j-- {
			if stack[j] > num[i] {
				k--
				continue
			}
			break
		}
		fmt.Println(i, j, k, string(stack), string(stack[:j+1]))
		if k <= 0 {
			return trim0(string(stack[:j+1-k]) + num[i:])
		}
		stack = append(stack[:j+1], num[i])
	}
	return trim0(string(stack[:len(stack)-k]))
}

func trim0(s string) string {
	var i int
	for i = 0; i < len(s); i++ {
		if s[i] != '0' {
			break
		}
	}
	if i >= len(s) {
		return "0"
	}
	return s[i:]
}
