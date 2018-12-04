/*
不使用运算符 + 和 - ​​​​​​​，计算两整数 ​​​​​​​a 、b ​​​​​​​之和。

示例 1:

输入: a = 1, b = 2
输出: 3
示例 2:

输入: a = -2, b = 3
输出: 1
*/
package main

import (
	"fmt"
)

func main() {
	fmt.Printf("%#x\n", -1>>1)
	// fmt.Println(getSum(-1, 3))
}

func getSum(a int, b int) int {
	for i := 0; i < 64; i++ {
		fmt.Printf("%b\n", a)
		a >>= 1
	}
	return 0
}
