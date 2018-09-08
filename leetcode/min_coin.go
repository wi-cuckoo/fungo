/*
硬币问题
我们有面值为1元3元5元的硬币若干枚，如何用最少的硬币凑够11元？

分析：
1 求问题的最优解：最小的硬币数
2 是否有子问题：f(n)表示的最少硬币数是是上一次拿时候的硬币数最少。
注意：f(n)是n元的最小硬币数，最后一次可拿的硬币数为1,3,5 则下一步
的最小硬币数为 f(n-vi) 它的状态变更不是按元数的，是按照上次拿的硬币钱目
3 状态转移方程为 f(n)= min(f(n-vi)+1)
4 边界问题(找到最后一个重复的问题) 这里
f(1)=1 ,f(2)=f(1)+f(1)=2 f(3)=min(1,f(2)+1) f(4)=f(3)+1 f(5)=1
*/

package main

import (
	"fmt"
)

func main() {
	fmt.Println(minCoin(11))
}

func minCoin(n int) int {
	dp := []int{0, 1, 2, 1, 2, 1}
	if n <= 5 {
		return dp[n]
	}
	coins := []int{1, 3, 5}
	for i := 6; i <= n; i++ {
		min := i
		for _, c := range coins {
			v := dp[i-c] + 1
			if min > v {
				min = v
			}
		}
		fmt.Println(i, min)
		dp = append(dp, min)
	}
	return dp[n]
}
