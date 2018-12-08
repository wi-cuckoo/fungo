package sort

import (
	"testing"
)

func TestQuickTest(t *testing.T) {
	nums := []int{3, 2, 4, 5, 1, 9}
	sorted := QuickSort(nums)
	t.Log("sorted: ", sorted)
	if len(sorted) != 6 {
		t.Fatal("sort result length incorrect")
	}
	for i, v := range []int{1, 2, 3, 4, 5, 9} {
		if sorted[i] != v {
			t.Errorf("index %d should be %d", i, v)
		}
	}
}

// TestSelectSort ...
func TestSelectSort(t *testing.T) {
	nums := []int{3, 2, 4, 8, 1, 9}
	SelectSort(nums)
	t.Log("sorted: ", nums)
	for i, v := range []int{1, 2, 3, 4, 8, 9} {
		if nums[i] != v {
			t.Errorf("index %d should be %d", i, v)
		}
	}
}

func TestShuffle(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	shuffle(nums)
	t.Log(nums)
}
