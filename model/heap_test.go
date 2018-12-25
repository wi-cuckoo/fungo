package model

import (
	"testing"
)

func TestHeap(t *testing.T) {
	nums := []int{3, 2, 1, 5, 8, 4}
	h := NewHeap(nums)
	t.Log(h.List())
	val := h.Pop()
	if val != 1 {
		t.Error("should be 1, not", val)
	}
	t.Log(h.List())
}
