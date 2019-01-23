/*
给定一个包含 m × n 个格子的面板，每一个格子都可以看成是一个细胞。每个细胞具有一个初始状态 live（1）即为活细胞， 或 dead（0）即为死细胞。每个细胞与其八个相邻位置（水平，垂直，对角线）的细胞都遵循以下四条生存定律：

如果活细胞周围八个位置的活细胞数少于两个，则该位置活细胞死亡；
如果活细胞周围八个位置有两个或三个活细胞，则该位置活细胞仍然存活；
如果活细胞周围八个位置有超过三个活细胞，则该位置活细胞死亡；
如果死细胞周围正好有三个活细胞，则该位置死细胞复活；
根据当前状态，写一个函数来计算面板上细胞的下一个（一次更新后的）状态。下一个状态是通过将上述规则同时应用于当前状态下的每个细胞所形成的，其中细胞的出生和死亡是同时发生的。

示例:

输入:
[
  [0,1,0],
  [0,0,1],
  [1,1,1],
  [0,0,0]
]
输出:
[
  [0,0,0],
  [1,0,1],
  [0,1,1],
  [0,1,0]
]
*/

package two

func gameOfLife(board [][]int) {
	if len(board) == 0 || len(board[0]) == 0 {
		return
	}
	m, n := len(board), len(board[0])
	cache := make([]int, m*n)
	// 设置cache
	x := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			cache[x] = nextStatus(i, j, board)
			x++
		}
	}
	// 通过cache更新board
	x = 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			board[i][j] = cache[x]
			x++
		}
	}
}

func nextStatus(i, j int, board [][]int) int {
	alive := board[i][j] == 1
	lives := 0
	for m := i - 1; m <= i+1 && m < len(board); m++ {
		for n := j - 1; n <= j+1 && n < len(board[0]); n++ {
			if m < 0 || n < 0 {
				continue
			}
			if board[m][n] == 1 {
				lives++
			}
		}
	}
	if alive {
		lives--
		switch {
		case lives == 2 || lives == 3:
			return 1
		default:
			return 0
		}
	}
	if lives == 3 {
		return 1
	}
	return 0
}
