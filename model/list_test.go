package model

import "testing"

func TestNewListNode(t *testing.T) {
	nums := []int{1, 2, 3, 4}
	head := NewListNode(nums)
	head.Print()
	i := 0
	for p := head; p != nil && i < len(nums); p = p.Next {
		if p.Val != nums[i] {
			t.Errorf("node value should be %d, not %d", nums[i], p.Val)
		}
		i++
	}
	if i != len(nums) {
		t.Errorf("invalid linked nodes length, should be %d, not %d", len(nums), i)
	}
}
