package model

import (
	"testing"
)

func TestHeap(t *testing.T) {
	nums := []int{3, 2, 1, 5, 8, 4}
	h := NewHeap(nums)
	h.Push(0)
	t.Log(h.List())
	min := h.Pop()
	if min != 0 {
		t.Error("pop should be 0, not", min)
	}
	t.Log(h.List())
}
