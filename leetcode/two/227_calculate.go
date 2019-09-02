/*
实现一个基本的计算器来计算一个简单的字符串表达式的值。

字符串表达式仅包含非负整数，+， - ，*，/ 四种运算符和空格  。 整数除法仅保留整数部分。

示例 1:
输入: "3+2*2"
输出: 7

示例 2:
输入: " 3/2 "
输出: 1

示例 3:
输入: " 3+5 / 2 "
输出: 5

说明：
	你可以假设所给定的表达式都是有效的。
	请不要使用内置的库函数 eval。
*/

package two

import (
	"strconv"
)

var ops = map[byte]int{
	'+': 0,
	'-': 0,
	'*': 1,
	'/': 1,
}

// 转成逆波兰表达式再求值太慢了
func calculate(s string) int {
	// 运算符栈s2和储存中间结果的栈s1
	s1, s2 := make([]string, 0), make([]string, 0)
	for i := 0; i < len(s); i++ {
		// 空格直接跳过
		if s[i] == ' ' {
			continue
		}
		// 如果是数字, 提取数字压入 s1
		if s[i] >= '0' && s[i] <= '9' {
			k := i + 1
			for k < len(s) && s[k] >= '0' && s[k] <= '9' {
				k++
			}
			s1, i = append(s1, s[i:k]), k-1
			continue
		}
		// 如果是运算符，需要比较s2栈顶的云算法优先级
		for {
			if len(s2) == 0 {
				s2 = append(s2, s[i:i+1])
				break
			}
			op := s2[len(s2)-1]
			if ops[s[i]] > ops[op[0]] {
				s2 = append(s2, s[i:i+1])
				break
			}
			s1 = append(s1, op)
			s2 = s2[:len(s2)-1]
		}
	}
	for i := len(s2); i > 0; i-- {
		s1 = append(s1, s2[i-1])
	}
	res := make([]int, 0)
	for _, x := range s1 {
		if x[0] >= '0' && x[0] <= '9' {
			n, _ := strconv.Atoi(x)
			res = append(res, n)
			continue
		}
		// 如果是符号
		l, n := len(res), 0
		switch x {
		case "+":
			n = res[l-2] + res[l-1]
		case "-":
			n = res[l-2] - res[l-1]
		case "*":
			n = res[l-2] * res[l-1]
		case "/":
			n = res[l-2] / res[l-1]
		}
		res[l-2] = n
		res = res[:l-1]
	}
	return res[0]
}

// 4ms, good job
func calculateV2(s string) int {
	ops := make([]byte, 0, len(s))
	res := make([]int, 0, len(s))
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			continue
		}
		// 如果是运算符，直接入队
		if s[i] < '0' || s[i] > '9' {
			ops = append(ops, s[i])
			continue
		}
		// 如果是数字提取出来
		n := 0
		for ; i < len(s) && s[i] >= '0' && s[i] <= '9'; i++ {
			n = n*10 + int(s[i]-'0')
		}
		i--
		// 如果操作符栈顶是高优先级，则取出操作数进行计算
		lo := len(ops)
		if lo > 0 && ops[lo-1] == '*' {
			res[len(res)-1] *= n
			ops = ops[:lo-1]
		} else if lo > 0 && ops[lo-1] == '/' {
			res[len(res)-1] /= n
			ops = ops[:lo-1]
		} else {
			res = append(res, n)
		}
	}
	m := res[0]
	for i := 1; i < len(res); i++ {
		if ops[i-1] == '+' {
			m += res[i]
		} else {
			m -= res[i]
		}
	}
	return m
}
