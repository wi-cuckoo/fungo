/*
给定一个包含 m x n 个元素的矩阵（m 行, n 列），请按照顺时针螺旋顺序，返回矩阵中的所有元素。

示例 1:

输入:
[
 [ 1, 2, 3 ],
 [ 4, 5, 6 ],
 [ 7, 8, 9 ]
]
输出: [1,2,3,6,9,8,7,4,5]
示例 2:

输入:
[
  [1, 2, 3, 4],
  [5, 6, 7, 8],
  [9,10,11,12]
]
输出: [1,2,3,4,8,12,11,10,9,5,6,7]
*/

package zero

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}
	m, n := len(matrix)-1, len(matrix[0])-1 // 终点
	x, y := 0, 0                            // 起点
	/*
	 * (x, y) ... (x, n)
	 *  ...   ...  ...
	 * (m, y) ... (m, n)
	 */
	res := make([]int, 0, (m+1)*(n+1))
	for x <= m && y <= n {
		// 横着走 左 -> 右
		for i := y; i <= n; i++ {
			res = append(res, matrix[x][i])
		}
		// 竖着走 上 -> 下
		for i := x + 1; i <= m; i++ {
			res = append(res, matrix[i][n])
		}
		// 倒着走 右 -> 左，注意只剩一行的情况
		for i := n - 1; i > y && m != x; i-- {
			res = append(res, matrix[m][i])
		}
		// 逆着走 下 -> 上，注意只剩下一列的情况
		for i := m; i > x && n != y; i-- {
			res = append(res, matrix[i][y])
		}
		// 更新 x, y, m, n
		x++
		y++
		m--
		n--
	}

	return res
}
