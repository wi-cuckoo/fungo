/*
给定两个二进制字符串，返回他们的和（用二进制表示）。

输入为非空字符串且只包含数字 1 和 0。

示例 1:

输入: a = "11", b = "1"
输出: "100"
示例 2:

输入: a = "1010", b = "1011"
输出: "10101"
*/
package zero

func addBinary(a string, b string) string {
	long, short := a, b
	if len(a) < len(b) {
		long, short = b, a
	}
	for len(short) < len(long) {
		short = "0" + short
	}
	var e byte = '0'
	s := ""
	for i := len(long) - 1; i >= 0; i-- {
		sum := short[i] + long[i] + e
		switch sum {
		case 144:
			s = "0" + s
			e = '0'
		case 145:
			s = "1" + s
			e = '0'
		case 146:
			s = "0" + s
			e = '1'
		case 147:
			s = "1" + s
			e = '1'
		}
	}
	if e == '1' {
		s = "1" + s
	}
	return s
}
