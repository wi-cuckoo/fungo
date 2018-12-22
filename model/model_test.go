package model

import (
	"container/heap"
	"testing"
)

func TestHeap(t *testing.T) {
	nums := []int{3, 2, 1, 5, 6, 4}
	h := NewHeap(nums)
	heap.Init(h)
	t.Log(h.List())
	val := heap.Pop(h)
	if val.(int) != 1 {
		t.Error("should be 1, not", val)
	}
}
