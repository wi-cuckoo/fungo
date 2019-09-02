/*
编写一个高效的算法来搜索 m x n 矩阵 matrix 中的一个目标值 target。该矩阵具有以下特性：

每行的元素从左到右升序排列。
每列的元素从上到下升序排列。
示例:

现有矩阵 matrix 如下：

[
  [1,   4,  7, 11, 15],
  [2,   5,  8, 12, 19],
  [3,   6,  9, 16, 22],
  [10, 13, 14, 17, 24],
  [18, 21, 23, 26, 30]
]
给定 target = 5，返回 true。

给定 target = 20，返回 false。
*/

package two

func searchMatrix(matrix [][]int, target int) bool {
	// 找到第一个元素小于 target 的最后一行
	for i := 0; i < len(matrix); i++ {
		l := len(matrix[i])
		if l < 1 {
			return false
		}
		if matrix[i][0] > target {
			return false
		}
		if matrix[i][l-1] < target {
			continue
		}
		if searchLine(matrix[i], target) {
			return true
		}
	}

	return false
}

func searchLine(nums []int, target int) bool {
	l, r := 0, len(nums)
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] > target {
			r = mid - 1
			continue
		}
		if nums[mid] < target {
			l = mid + 1
			continue
		}
		return true
	}
	return false
}

/*
 * 注意审题，每一列的元素从上到下也是有序的，这样一来我们可以阶梯式的搜索
 */
func searchMatrixV2(matrix [][]int, target int) bool {
	if len(matrix) < 1 {
		return false
	}
	i, j := 0, len(matrix[0])-1
	for i < len(matrix) && j >= 0 {
		if matrix[i][j] == target {
			return true
		}
		if matrix[i][j] < target {
			i++
		} else {
			j--
		}
	}
	return false
}
