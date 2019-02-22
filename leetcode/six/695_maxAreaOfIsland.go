/*
给定一个包含了一些 0 和 1的非空二维数组 grid , 一个 岛屿 是由四个方向 (水平或垂直) 的 1 (代表土地) 构成的组合。你可以假设二维矩阵的四个边缘都被水包围着。

找到给定的二维数组中最大的岛屿面积。(如果没有岛屿，则返回面积为0。)

示例 1:

[[0,0,1,0,0,0,0,1,0,0,0,0,0],
 [0,0,0,0,0,0,0,1,1,1,0,0,0],
 [0,1,1,0,1,0,0,0,0,0,0,0,0],
 [0,1,0,0,1,1,0,0,1,0,1,0,0],
 [0,1,0,0,1,1,0,0,1,1,1,0,0],
 [0,0,0,0,0,0,0,0,0,0,1,0,0],
 [0,0,0,0,0,0,0,1,1,1,0,0,0],
 [0,0,0,0,0,0,0,1,1,0,0,0,0]]
对于上面这个给定矩阵应返回 6。注意答案不应该是11，因为岛屿只能包含水平或垂直的四个方向的‘1’。

示例 2:

[[0,0,0,0,0,0,0,0]]
对于上面这个给定的矩阵, 返回 0。

注意: 给定的矩阵grid 的长度和宽度都不超过 50。
*/

package six

// 深度优先搜索搞定吧
// 很简单的一道题
func maxAreaOfIsland(grid [][]int) int {
	max := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 0 {
				continue
			}
			area := areaSearch(grid, i, j)
			if area > max {
				max = area
			}
		}
	}
	return max
}

func areaSearch(grid [][]int, i, j int) int {
	if grid[i][j] == 0 {
		return 0
	}
	grid[i][j] = 0 // 将该位置置为 0，避免重复计算
	s := 1
	if i > 0 {
		s += areaSearch(grid, i-1, j)
	}
	if i < len(grid)-1 {
		s += areaSearch(grid, i+1, j)
	}
	if j > 0 {
		s += areaSearch(grid, i, j-1)
	}
	if j < len(grid[0])-1 {
		s += areaSearch(grid, i, j+1)
	}
	return s
}
