/*
给定两个整数，分别表示分数的分子 numerator 和分母 denominator，以字符串形式返回小数。

如果小数部分为循环小数，则将循环的部分括在括号内。

示例 1:
输入: numerator = 1, denominator = 2
输出: "0.5"


示例 3:
输入: numerator = 2, denominator = 3
输出: "0.(6)"
*/

package one

import "strconv"

func fractionToDecimal(x int, y int) string {
	if x == 0 {
		return "0"
	}
	neg := x*y < 0
	if y < 0 {
		y = 0 - y
	}
	if x < 0 {
		x = 0 - x
	}
	res, i := []int{x / y}, 1 // 小数从索引为 1 算起
	remain := x % y           // remain 为余数
	cache := map[int]int{
		remain: i,
	}
	for remain != 0 {
		remain *= 10 // 把余数乘以 10，继续干
		res = append(res, remain/y)
		remain %= y
		if _, ok := cache[remain]; ok {
			break
		}
		i++
		cache[remain] = i
	}
	s, flag := "", false
	for i, n := range res {
		if i == 1 {
			s += "."
		}
		if i == cache[remain] {
			s += "("
			flag = true
		}
		s += strconv.Itoa(n)
	}
	if flag {
		s += ")"
	}
	if neg {
		return "-" + s
	}
	return s
}
