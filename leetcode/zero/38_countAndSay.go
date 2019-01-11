/*
报数序列是指一个整数序列，按照其中的整数的顺序进行报数，得到下一个数。其前五项如下：

1.     1
2.     11
3.     21
4.     1211
5.     111221
1 被读作  "one 1"  ("一个一") , 即 11。
11 被读作 "two 1s" ("两个一"）, 即 21。
21 被读作 "one 2",  "one 1" （"一个二" ,  "一个一") , 即 1211。

给定一个正整数 n ，输出报数序列的第 n 项。

注意：整数顺序将表示为一个字符串。

示例 1:

输入: 1
输出: "1"
示例 2:

输入: 6
输出: "312211"

n per n
*/

package zero

import (
	"fmt"
)

// for version
func countAndSay(n int) string {
	trans := map[string]string{
		"1":  "11",
		"11": "21",
	}
	res := "1"
	for i := 1; i < n; i++ {
		if v, ok := trans[res]; ok {
			res = v
			continue
		}
		var n byte
		var m int
		next := ""
		for j := 0; j < len(res); j++ {
			if m == 0 {
				n, m = res[j], 1
				continue
			}
			if n == res[j] {
				m++
				if j == len(res)-1 {
					next += fmt.Sprintf("%d%c", m, n)
					break
				}
				continue
			}
			next += fmt.Sprintf("%d%c", m, n)
			n = res[j]
			m = 1
			if j == len(res)-1 {
				next += fmt.Sprintf("%d%c", m, n)
				break
			}
		}
		res = next
	}
	return res
}

// recursive version
func countAndSayV2(n int) string {
	if n == 1 {
		return "1"
	}
	if n == 2 {
		return "11"
	}

	res := countAndSayV2(n - 1)
	var cur byte
	var m int
	next := ""
	for j := 0; j < len(res); j++ {
		if m == 0 {
			cur, m = res[j], 1
			continue
		}
		if cur == res[j] {
			m++
			if j == len(res)-1 {
				next += fmt.Sprintf("%d%c", m, cur)
				break
			}
			continue
		}
		next += fmt.Sprintf("%d%c", m, cur)
		cur = res[j]
		m = 1
		if j == len(res)-1 {
			next += fmt.Sprintf("%d%c", m, cur)
			break
		}
	}
	return next
}
