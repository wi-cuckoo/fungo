package main

import (
	"fmt"
)

func main() {
	matrix := [][]int{
		[]int{1, 2, 3, 4},
		[]int{5, 6, 7, 8},
		[]int{9, 10, 11, 12},
		[]int{13, 14, 15, 16},
	}
	rotate(matrix)
	fmt.Println(matrix)
}

func rotate(matrix [][]int) {
	n := len(matrix)
	// 对角线交换值
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i+j == n-1 {
				break
			}
			matrix[i][j], matrix[n-j-1][n-i-1] = matrix[n-j-1][n-i-1], matrix[i][j]
		}
	}
	// 上下对称交换值
	for i := 0; i < n/2; i++ {
		matrix[i], matrix[n-i-1] = matrix[n-i-1], matrix[i]
	}
}
