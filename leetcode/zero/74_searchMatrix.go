/*
编写一个高效的算法来判断 m x n 矩阵中，是否存在一个目标值。该矩阵具有如下特性：

每行中的整数从左到右按升序排列。
每行的第一个整数大于前一行的最后一个整数。
示例 1:

输入:
matrix = [
  [1,   3,  5,  7],
  [10, 11, 16, 20],
  [23, 30, 34, 50]
]
target = 3
输出: true
示例 2:

输入:
matrix = [
  [1,   3,  5,  7],
  [10, 11, 16, 20],
  [23, 30, 34, 50]
]
target = 13
输出: false
*/

package zero

func searchMatrix(matrix [][]int, target int) bool {
	// 将二维数组考虑成有序一维数组
	if len(matrix) < 1 {
		return false
	}
	m, n := len(matrix), len(matrix[0])
	l, r := 0, m*n-1
	for l <= r {
		mid := (l + r) / 2
		i, j := mid/n, mid%n
		if matrix[i][j] == target {
			return true
		}
		if matrix[i][j] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return false
}
