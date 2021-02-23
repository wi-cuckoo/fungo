package ten

/*
今天，书店老板有一家店打算试营业 customers.length 分钟。每分钟都有一些顾客（customers[i]）会进入书店，
所有这些顾客都会在那一分钟结束后离开。在某些时候，书店老板会生气。 如果书店老板在第 i 分钟生气，那么
grumpy[i] = 1，否则 grumpy[i] = 0。 当书店老板生气时，那一分钟的顾客就会不满意，不生气则他们是满意的。

书店老板知道一个秘密技巧，能抑制自己的情绪，可以让自己连续 X 分钟不生气，但却只能使用一次。
请你返回这一天营业下来，最多有多少客户能够感到满意的数量。


示例：
输入：customers = [1,0,1,2,1,1,7,5], grumpy = [0,1,0,1,0,1,0,1], X = 3
输出：16
解释：
书店老板在最后 3 分钟保持冷静。
感到满意的最大客户数量 = 1 + 1 + 1 + 1 + 7 + 5 = 16.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/grumpy-bookstore-owner
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// 两钟思路
// 1. 假设在 [i, i+X] 分钟内使用了该技巧，则满意人数为 i 分钟前确定满意的人数，i+X 分钟之后确定满意的人数 及 [i, i+X]之
// 间满意的人数之和，由此对每分钟都计算一次该和取最大值则是结果
// 2. 对 [i, i+X] 分钟内使用该技巧，只是增加了该范围内原本不满意的人数，那么其范围前后该满意的还是不变，所以可以考虑求
// 该滑动窗口内增加的那部分人数的最大值，然后加上所有确定满意的人数
// 下面代码为第二种思路实现
func maxSatisfied(customers []int, grumpy []int, X int) int {
	n := len(grumpy)
	confirm, inc, win := 0, 0, 0
	l := 0
	for r := 0; r < n; r++ {
		confirm += (1 - grumpy[r]) * customers[r]
		win += grumpy[r] * customers[r]
		if win > inc {
			inc = win
		}
		if r-l+1 == X {
			win -= grumpy[l] * customers[l]
			l++
		}
	}
	return confirm + inc
}
