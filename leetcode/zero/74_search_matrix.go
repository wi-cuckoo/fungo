package zero

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
