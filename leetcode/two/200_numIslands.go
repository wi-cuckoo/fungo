/*
给定一个由 '1'（陆地）和 '0'（水）组成的的二维网格，计算岛屿的数量。一个岛被水包围，并且它是通过水平方向或垂直方向上相邻的陆地连接而成的。你可以假设网格的四个边均被水包围。

示例 1:

输入:
1111010
1101110
1100000
0000010

输出: 1
示例 2:

输入:
11000
11000
00100
00011

输出: 3
*/

package two

func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	// 用广度优先搜索试试
	count := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			// 如果是水域，则跳过
			if grid[i][j] != '1' {
				continue
			}
			// 如果是陆地，则摧毁这些岛屿并计数加一
			count++
			destory(i, j, grid)
		}
	}
	return count
}

func destory(i, j int, grid [][]byte) {
	grid[i][j] = '0'
	if i > 0 && grid[i-1][j] == '1' {
		destory(i-1, j, grid)
	}
	if i < len(grid)-1 && grid[i+1][j] == '1' {
		destory(i+1, j, grid)
	}
	if j > 0 && grid[i][j-1] == '1' {
		destory(i, j-1, grid)
	}
	if j < len(grid[0])-1 && grid[i][j+1] == '1' {
		destory(i, j+1, grid)
	}
}
