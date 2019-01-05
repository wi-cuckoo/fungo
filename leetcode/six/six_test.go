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

func TestLeastInterval(t *testing.T) {
	tasks := []byte{'A', 'A', 'A', 'B', 'B', 'B'}
	x := leastInterval(tasks, 2)
	if x != 8 {
		t.Errorf("leastInterval should be 8, not %d", x)
	}

	tasks = []byte{'A', 'A', 'A', 'A', 'A', 'A', 'B', 'C', 'D', 'E', 'F', 'G'}
	x = leastInterval(tasks, 2)
	if x != 16 {
		t.Errorf("leastInterval should be 16, not %d", x)
	}
}
