package one

import "testing"

func TestRob1(t *testing.T) {
	nums := []int{2, 7, 9, 3, 1}
	val := rob(nums)
	if val != 12 {
		t.Error("rob result should be 12")
	}

	nums = []int{1, 2, 3, 1}
	val = rob(nums)
	if val != 4 {
		t.Error("rob result should be 4")
	}
}
