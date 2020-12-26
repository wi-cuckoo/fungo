package zero

/*
给定一个仅包含 0 和 1 、大小为 rows x cols 的二维二进制矩阵，找出只包含 1 的最大矩形，并返回其面积。

输入：matrix = [["1","0","1","0","0"],["1","0","1","1","1"],["1","1","1","1","1"],["1","0","0","1","0"]]
输出：6
解释：最大矩形如上图所示。

输入：matrix = []
输出：0

输入：matrix = [["0"]]
输出：0

输入：matrix = [["1"]]
输出：1

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/maximal-rectangle
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

import (
	"github.com/wi-cuckoo/fungo/util"
)

func maximalRectangle(matrix [][]byte) int {
	m, n := len(matrix), len(matrix[0])
	left := make([][]int, m)
	for i := 0; i < m; i++ {
		left[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if matrix[i][j] == '0' {
				continue
			}
			if j == 0 {
				left[i][j] = 0
			} else {
				left[i][j] = left[i][j-1] + 1
			}
		}
	}
	ans := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == '0' {
				continue
			}
			width := left[i][j]
			area := width * 1
			for k := i - 1; k >= 0; k-- {
				width = util.Min(width, left[k][j])
				area = util.Max(area, width*(i-k+1))
			}
			ans = util.Max(area, ans)
		}
	}
	return ans
}
