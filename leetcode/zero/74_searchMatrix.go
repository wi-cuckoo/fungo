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

import "fmt"

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) < 1 || len(matrix[0]) < 1 {
		return false
	}
	// 先二分搜索行，找到位于哪一行
	l, r := 0, len(matrix)-1
	for l <= r {
		mid := (l + r) / 2
		if matrix[mid][0] < target {
			l = mid + 1
			continue
		}
		if matrix[mid][0] > target {
			r = mid - 1
			continue
		}
		return true
	}
	if l > 0 {
		l--
	}
	fmt.Println(l, r)
	m, n := 0, len(matrix[l])-1
	for m <= n {
		mid := (m + n) / 2
		if matrix[l][mid] < target {
			m = mid + 1
			continue
		}
		if matrix[l][mid] > target {
			n = mid - 1
			continue
		}
		return true
	}
	return false
}

/*
func searchMatrix(matrix [][]int, target int) bool {
	// 处理一下异常情况
	h := len(matrix)
	if h == 0 {
		return false
	}
	w := len(matrix[0])
	if w == 0 {
		return false
	}

	// 先用线性查找查找target可能存在的那一行，可以优化成二分查找
	// 该行具有以下特征，matrix[line][0] <= target <= matrix[line][w-1]
	line := 0
	for i := 0; i < h; i++ {
		if matrix[i][w-1] == target {
			return true
		}
		if matrix[i][w-1] > target {
			line = i
			break
		}
	}
	//
	for i := 0; i < w; i++ {
		if matrix[line][i] == target {
			return true
		}
		if matrix[line][i] > target {
			break
		}
	}

	return false
}
*/
