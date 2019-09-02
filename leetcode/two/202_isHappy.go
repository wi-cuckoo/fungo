/*
编写一个算法来判断一个数是不是“快乐数”。

一个“快乐数”定义为：对于一个正整数，每一次将该数替换为它每个位置上的数字的平方和，然后重复这个过程直到这个数变为 1，也可能是无限循环但始终变不到 1。如果可以变为 1，那么这个数就是快乐数。

示例:

输入: 19
输出: true
解释:
12 + 92 = 82
82 + 22 = 68
62 + 82 = 100
12 + 02 + 02 = 1
*/

package two

// 分析得非快乐数都会进入以下数字循环之中
var cache = map[int]int{
	4:   1,
	16:  1,
	37:  1,
	58:  1,
	89:  1,
	145: 1,
	42:  1,
	20:  1,
}

func isHappy(n int) bool {
	for {
		x := epow(n)
		if x == 1 {
			return true
		}
		if _, ok := cache[x]; ok {
			return false
		}
		n = x
	}
}

func epow(n int) int {
	x := 0
	for n != 0 {
		m := n % 10
		x += m * m
		n /= 10
	}
	return x
}
