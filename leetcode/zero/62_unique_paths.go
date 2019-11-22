/*
一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为“Start” ）。

机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为“Finish”）。

问总共有多少条不同的路径？



例如，上图是一个7 x 3 的网格。有多少可能的路径？

说明：m 和 n 的值均不超过 100。

示例 1:

输入: m = 3, n = 2
输出: 3
解释:
从左上角开始，总共有 3 条路径可以到达右下角。
1. 向右 -> 向右 -> 向下
2. 向右 -> 向下 -> 向右
3. 向下 -> 向右 -> 向右
示例 2:

输入: m = 7, n = 3
输出: 28
*/

package zero

// 1. 递归，我他妈都佩服我自己，三行代码
// long time to say
func uniquePaths(m int, n int) int {
	if n == 1 || m == 1 {
		return 1
	}
	return uniquePaths(m, n-1) + uniquePaths(m-1, n)
}

// 2. 动态规划
// 对于点 (i, j)，到达该点的路径取决于到达 (i, j-1) 和 (i-1, j) 的路径
// 并且满足 f(i,j) = f(i,j-1)+f(i-1,j)
func uniquePathsV2(m int, n int) int {
	if n == 1 || m == 1 {
		return 1
	}
	dp := make([][]int, m) // 2-dimension arra
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		dp[i][0] = 1
		for j := 1; j < n; j++ {
			if i == 0 {
				dp[i][j] = 1
				continue
			}
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}
