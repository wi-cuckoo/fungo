package sort

import (
	"fmt"
)

// increasing order
func InsertSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		fmt.Println("current ele: ", arr[i])
		cur := arr[i]
		j = i - 1
		for j := i - 1; j >= 0; j-- {
			// if arr[i] less than arr[j], swap them
			if arr[i] < arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
				i = j
			}
		}
	}
}
