// +build ignore,OMIT

package main

import "fmt"

// 需求，将矩阵中相连的且值为W点标记为 S，相连为水平或垂直相邻

func main() {
	matrix := [][]byte{
		{'A', 'W', 'W', 'R'},
		{'E', 'E', 'E', 'W'},
		{'E', 'W', 'A', 'R'},
		{'E', 'A', 'W', 'R'},
	}
	fmt.Println(matrix)
	n, m := len(matrix), len(matrix[0])
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if matrix[i][j] == 'W' {
				dfs(matrix, i, j, n, m)
			}
		}
	}
	fmt.Println(matrix)
}

func dfs(matrix [][]byte, i, j, n, m int) {
	if i < 0 || i >= n || j < 0 || j >= m || matrix[i][j] != 'W' {
		return
	}
	matrix[i][j] = 'S'
	dfs(matrix, i-1, j, n, m) // up
	dfs(matrix, i+1, j, n, m) // down
	dfs(matrix, i, j+1, n, m) // right
	dfs(matrix, i, j-1, n, m) // left
}

var (
	dx = [4]int{1, -1, 0, 0}
	dy = [4]int{0, 0, 1, -1}
)

func bfs(matrix [][]byte, i, j, n, m int) {
	queue := [][]int{}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if matrix[i][j] == 'W' {
				queue = append(queue, []int{i, j})
			}
		}
	}
	for len(queue) > 0 {
		point := queue[0]
		queue = queue[1:]
		x, y := point[0], point[1]
		matrix[x][y] = 'S'
		for i := 0; i < 4; i++ {
			mx, my := x+dx[i], y+dy[i]
			if mx < 0 || my < 0 || mx >= n || my >= m || matrix[mx][my] != 'W' {
				continue
			}
			queue = append(queue, []int{mx, my})
		}
	}
}
