/*

给定一个三角形，找出自顶向下的最小路径和。每一步只能移动到下一行中相邻的结点上。

例如，给定三角形：

[
     [2],
    [3,4],
   [6,5,7],
  [4,1,8,3]
]
自顶向下的最小路径和为 11（即，2 + 3 + 5 + 1 = 11）。

说明：

如果你可以只使用 O(n) 的额外空间（n 为三角形的总行数）来解决这个问题，那么你的算法会很加分。
*/

package main

import (
	"fmt"
)

func main() {
	// grid := [][]int{
	// 	[]int{2},
	// 	[]int{3, 4},
	// 	[]int{6, 5, 7},
	// 	[]int{4, 2, 1, 3},
	// }
	grid2 := [][]int{
		[]int{-1},
		[]int{-2, -3},
	}
	fmt.Println(minimumTotal(grid2))
}

func minimumTotal(triangle [][]int) int {
	dp := make([]int, len(triangle))
	dp[0] = triangle[0][0]
	if len(triangle) == 1 {
		return dp[0]
	}
	for i := 1; i < len(triangle); i++ {
		lt := len(triangle[i])
		p := dp[0]
		dp[lt-1] = dp[lt-2] + triangle[i][lt-1]
		dp[0] = dp[0] + triangle[i][0]
		for j := 1; j < len(triangle[i])-1; j++ {
			lj := dp[j] // store the last
			if lj > p {
				lj = p
			}
			// re-value p and dp[j]
			p, dp[j] = dp[j], lj+triangle[i][j]
		}
	}
	min := dp[0]
	for _, n := range dp {
		if min >= n {
			min = n
		}
	}
	return min
}
