/*
给定一个 n x n 矩阵，其中每行和每列元素均按升序排序，找到矩阵中第 k 小的元素。
请注意，它是排序后的第 k 小元素，而不是第 k 个不同的元素。
示例：

matrix = [
   [ 1,  5,  9],
   [10, 11, 13],
   [12, 13, 15]
],
k = 8,

返回 13。

提示：
你可以假设 k 的值永远是有效的，1 ≤ k ≤ n2 。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/kth-smallest-element-in-a-sorted-matrix
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package three

// 1. 可以使用暴力但比较傻逼，优化方法是用合并排序将数组两两合并，对暴力方法中排序部分做优化
// 2. 二分查找，对该矩阵，可以知左上角数最小，右上角数最大，那么对于中间值某个值就可以一次遍历
// 计算出大于它的数有多少，小于的有多少

func kthSmallest(matrix [][]int, k int) int {
	n := len(matrix)
	l, r := matrix[0][0], matrix[n-1][n-1]
	for l < r {
		mid := (l + r) / 2
		if fuckth(matrix, k, mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}

func fuckth(matrix [][]int, k int, mid int) bool {
	// from left-bottom to right-top
	n := len(matrix)
	i, j := n-1, 0
	cnt := 0
	for i >= 0 && j < n {
		// 如果当前值大于mid，则需要向上探索
		if matrix[i][j] > mid {
			i--
			continue
		}
		j++
		cnt += i + 1
	}
	return cnt >= k
}
