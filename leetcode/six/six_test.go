package six

import "testing"

func TestFindLengthOfLCIS(t *testing.T) {
	nums := []int{1, 3, 5, 4, 7}
	l := findLengthOfLCIS(nums)
	if l != 3 {
		t.Error("longth should be 3")
	}
	t.Log(l)

	nums = []int{2, 2, 2, 2, 2}
	l = findLengthOfLCIS(nums)
	if l != 1 {
		t.Error("longth should be 1")
	}
	t.Log(l)
}
