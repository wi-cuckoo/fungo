/*
在一个由 0 和 1 组成的二维矩阵内，找到只包含 1 的最大正方形，并返回其面积。

示例:

输入:

1 0 1 0 0
1 0 1 1 1
1 1 1 1 1
1 0 0 1 0

输出: 4
*/

package two

func maximalSquare(matrix [][]byte) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	m, n := len(matrix), len(matrix[0])
	dp := make([]int, m*n)
	// 注：f(x) 表以位置x为可能正方形的右下顶点时所对应的边长
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			idx := j + n*i
			// 如果是第一行或者第一列则该位置是0对应的f(x)就为0，是1则为1
			// 如果该位置数字为0，则对应的f(x)就为0，可以和上面的判断合并
			if i == 0 || j == 0 || matrix[i][j] == '0' {
				dp[idx] = int(matrix[i][j] - '0')
				continue
			}
			// 如果该位置非一行一列并且值为 1，则观其左上两位置f(x)值
			d1, d2 := dp[idx-1], dp[idx-n]
			if d1 > d2 {
				d1, d2 = d2, d1
			}
			// fmt.Println(i, j, idx, d1, d2)
			// 这里判断考虑情况较多
			// 1. 以较小正方形为准，边加1的情况下，如果左上顶点为1，则可以成功加上这条边
			// 2. 如果1中的边加不上，则只能在较小正方形和1中取最大
			if dp[idx-d1*(n+1)] >= 1 {
				dp[idx] = d1 + 1
			} else if d1 > 1 {
				dp[idx] = d1
			} else {
				dp[idx] = 1
			}
		}
	}
	// fmt.Println(dp)
	for i := 1; i < len(dp); i++ {
		if dp[i] > dp[0] {
			dp[0] = dp[i]
		}
	}
	return dp[0] * dp[0]
}
