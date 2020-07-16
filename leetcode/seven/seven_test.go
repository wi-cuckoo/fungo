package seven

import (
	"fmt"
	"testing"
)

func TestRotateString(t *testing.T) {
	a, b := "gcmbf", "fgcmb"
	if !rotateString(a, b) {
		t.Error("should be true")
	}
}

func TestNumSubarrayProductLessThanK(t *testing.T) {
	fmt.Println(numSubarrayProductLessThanK([]int{10, 5, 2, 6}, 100))
	fmt.Println(numSubarrayProductLessThanK([]int{1, 2, 3}, 0))
	fmt.Println(numSubarrayProductLessThanK([]int{1, 1, 1}, 2))
	fmt.Println(numSubarrayProductLessThanK([]int{10, 9, 10, 4, 3, 8, 3, 3, 6, 2, 10, 10, 9, 3}, 19))
}

func TestIsBipartite(t *testing.T) {
	graph := [][]int{
		{1, 3}, {0, 2}, {1, 3}, {0, 2},
	}
	t.Log(isBipartite(graph))
	graph = [][]int{
		{1, 2, 3}, {0, 2}, {0, 1, 3}, {0, 2},
	}
	t.Log(isBipartite(graph))
	graph = [][]int{
		{}, {3}, {}, {1}, {},
	}
	t.Log(isBipartite(graph))
}
