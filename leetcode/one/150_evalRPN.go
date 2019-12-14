/*
输入: ["10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"]
输出: 22
解释:
  ((10 * (6 / ((9 + 3) * -11))) + 17) + 5
= ((10 * (6 / (12 * -11))) + 17) + 5
= ((10 * (6 / -132)) + 17) + 5
= ((10 * 0) + 17) + 5
= (0 + 17) + 5
= 17 + 5
= 22
*/

package one

import (
	"fmt"
	"strconv"
)

var ops = map[string]int{
	"+": 1,
	"-": 1,
	"*": 1,
	"/": 1,
}

func evalRPN(tokens []string) int {
	stack := make([]int, 0, len(tokens))
	for _, t := range tokens {
		// 如果是0-9的数字，则直接入栈
		if _, ok := ops[t]; !ok {
			i, _ := strconv.Atoi(t)
			stack = append(stack, i)
			continue
		}
		fmt.Println(stack)
		l := len(stack)
		n1, n2 := stack[l-1], stack[l-2]
		switch t {
		case "+":
			stack[l-2] = n1 + n2
		case "-":
			stack[l-2] = n2 - n1
		case "*":
			stack[l-2] = n2 * n1
		case "/":
			stack[l-2] = n2 / n1
		}
		stack = stack[:l-1]
	}
	return stack[0]
}
