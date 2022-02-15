package thirteen


/*
给你一个 m * n 的矩阵，矩阵中的数字 各不相同 。请你按 任意 顺序返回矩阵中的所有幸运数。

幸运数是指矩阵中满足同时下列两个条件的元素：

在同一行的所有元素中最小
在同一列的所有元素中最大
 

示例 1：

输入：matrix = [[3,7,8],[9,11,13],[15,16,17]]
输出：[15]
解释：15 是唯一的幸运数，因为它是其所在行中的最小值，也是所在列中的最大值。
示例 2：

输入：matrix = [[1,10,4,2],[9,3,8,7],[15,16,17,12]]
输出：[12]
解释：12 是唯一的幸运数，因为它是其所在行中的最小值，也是所在列中的最大值。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/lucky-numbers-in-a-matrix
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// calculate the minimum number for every row, and the maximum number for every column
// then if matrix[i][j] == rowMin[i] && matrix[i][j] == colMax[j], that is it
func luckyNumbers(matrix [][]int) []int {
	m, n := len(matrix), len(matrix[0]) // m row, n column
	rowMin := make([]int, m)
	for i := 0; i < m; i++ {
		min := math.MaxInt32
		for j := 0; j < n; j++ {
			if cur := matrix[i][j]; cur < min {
				min = cur
			}
		}
		rowMin[i] = min
	}
	ans := make([]int, 0, m)
	for j := 0; j < n; j++ {
		max := 0
		for i := 1; i < m; i++ {
			if matrix[i][j] > matrix[max][j] {
				max = i
			}
		}
		if colMax := matrix[max][j]; colMax == rowMin[max] {
			ans = append(ans, colMax)
		}
	}
	return ans
}