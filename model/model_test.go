package model

import (
	"container/heap"
	"testing"
)

func TestHeap(t *testing.T) {
	nums := []int{3, 2, 1, 5, 6, 4}
	h := NewHeap(nums)
	heap.Init(h)
	val := heap.Pop(h)
	if val.(int) != 6 {
		t.Error("should be 6, not", val)
	}
}
