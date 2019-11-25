/*
假设你正在爬楼梯。需要 n 阶你才能到达楼顶。

每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？

注意：给定 n 是一个正整数。

示例 1：

输入： 2
输出： 2
解释： 有两种方法可以爬到楼顶。
1.  1 阶 + 1 阶
2.  2 阶
示例 2：

输入： 3
输出： 3
解释： 有三种方法可以爬到楼顶。
1.  1 阶 + 1 阶 + 1 阶
2.  1 阶 + 2 阶
3.  2 阶 + 1 阶
*/

package zero

var cache = map[int]int{}

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	if v, ok := cache[n]; ok {
		return v
	}
	v := climbStairs(n-1) + climbStairs(n-2)
	cache[n] = v
	return v
}

func climbStairsV2(n int) int {
	s1, s2 := 1, 2
	if n == 1 {
		return s1
	}
	if n == 2 {
		return s2
	}
	res := 0
	for i := 3; i <= n; i++ {
		res = s1 + s2
		s1, s2 = s2, res
	}
	return res
}
