package sort

import (
	"testing"

	"github.com/wi-cuckoo/fungo/model"
)

func TestQuickTest(t *testing.T) {
	nums := []int{3, 2, 4, 5, 1, 9}
	QuickSortV2(nums)
	for i, v := range []int{1, 2, 3, 4, 5, 9} {
		if nums[i] != v {
			t.Errorf("index %d should be %d", i, v)
		}
	}
}

// 测一把两种快排的性能
func BenchmarkQuickSort(t *testing.B) {
	for i := 0; i < t.N; i++ {
		nums := []int{3, 2, 4, 5, 1, 9}
		QuickSortV2(nums)
	}
}

// TestSelectSort ...
func TestSelectSort(t *testing.T) {
	nums := []int{3, 2, 4, 8, 1, 9}
	SelectSort(nums)
	for i, v := range []int{1, 2, 3, 4, 8, 9} {
		if nums[i] != v {
			t.Errorf("index %d should be %d, not %d", i, v, nums[i])
		}
	}
}

// TestBubbleSort ...
func TestBubbleSort(t *testing.T) {
	nums := []int{3, 2, 4, 8, 1, 9}
	BubbleSort(nums)
	for i, v := range []int{1, 2, 3, 4, 8, 9} {
		if nums[i] != v {
			t.Errorf("index %d should be %d, not %d", i, v, nums[i])
		}
	}
}

func TestInsertSort(t *testing.T) {
	nums := []int{3, 2, 4, 8, 1, 9}
	InsertSort(nums)
	for i, v := range []int{1, 2, 3, 4, 8, 9} {
		if nums[i] != v {
			t.Errorf("index %d should be %d, not %d", i, v, nums[i])
		}
	}
}

func TestMerge(t *testing.T) {
	t.Log(merge([]int{1, 3, 6}, []int{2, 4, 5, 10}))
}

func TestMergeSort(t *testing.T) {
	nums := []int{3, 2, 4, 8, 1, 9}
	nums = MergeSort(nums)
	for i, v := range []int{1, 2, 3, 4, 8, 9} {
		if nums[i] != v {
			t.Errorf("index %d should be %d, not %d", i, v, nums[i])
		}
	}
}

func TestHeapSort(t *testing.T) {
	nums := []int{3, 2, 4, 8, 1, 9}
	nums = HeapSort(nums)
	for i, v := range []int{1, 2, 3, 4, 8, 9} {
		if nums[i] != v {
			t.Errorf("index %d should be %d, not %d", i, v, nums[i])
		}
	}
}
func TestShuffle(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	shuffle(nums)
	t.Log(nums)
}

func TestQuickSortList(t *testing.T) {
	head := &model.ListNode{
		Val: 3,
		Next: &model.ListNode{
			Val: 2,
			Next: &model.ListNode{
				Val: 5,
				Next: &model.ListNode{
					Val:  4,
					Next: nil,
				},
			},
		},
	}
	QuickSortList(head).Print()
}
