/*

给定一个二维网格和一个单词，找出该单词是否存在于网格中。

单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。

示例:

board =
[
  ['A','B','C','E'],
  ['S','F','C','S'],
  ['A','D','E','E']
]

给定 word = "ABCCED", 返回 true.
给定 word = "SEE", 返回 true.
给定 word = "ABCB", 返回 false.
*/

package zero

// 还是回溯法去找，不放过任何一种可能

func exist(board [][]byte, word string) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == word[0] {
				if backtrackExist(board, i, j, word[1:]) {
					return true
				}
			}
		}
	}
	return false
}

func backtrackExist(board [][]byte, i, j int, w string) bool {
	if w == "" {
		return true
	}
	// set board[i][j] empty
	c := board[i][j]
	board[i][j] = '1'
	// check (i-1, j)
	if i > 0 && board[i-1][j] == w[0] {
		if backtrackExist(board, i-1, j, w[1:]) {
			return true
		}
	}
	// check (i+1, j)
	if i < len(board)-1 && board[i+1][j] == w[0] {
		if backtrackExist(board, i+1, j, w[1:]) {
			return true
		}
	}
	// check (i, j-1)
	if j > 0 && board[i][j-1] == w[0] {
		if backtrackExist(board, i, j-1, w[1:]) {
			return true
		}
	}
	if j < len(board[0])-1 && board[i][j+1] == w[0] {
		if backtrackExist(board, i, j+1, w[1:]) {
			return true
		}
	}
	board[i][j] = c
	return false
}
