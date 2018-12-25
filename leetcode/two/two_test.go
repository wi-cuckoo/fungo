package two

import (
	"fmt"
	"testing"

	"github.com/wi-cuckoo/fungo/model"
)

func TestReverseList(t *testing.T) {
	l1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val:  1,
					Next: nil,
				},
			},
		},
	}
	// reverseList(l1).Print()
	// reverseList2(l1).Print()
	fmt.Println(isPalindrome(l1))
}

func TestRob(t *testing.T) {
	nums := []int{2, 3, 2}
	if rob(nums) != 3 {
		t.Error("rob result should be 3")
	}

	nums = []int{1, 2, 3, 1}
	if rob(nums) != 4 {
		t.Error("rob result should be 4")
	}
}

func TestFindKthLargest(t *testing.T) {
	nums := []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	k := findKthLargest(nums, 4)
	if k != 4 {
		t.Errorf("the 4th largest num should be 4, not %d", k)
	}

	k = findKthLargestV2(nums, 4)
	if k != 4 {
		t.Errorf("the 4th largest num should be 4, not %d", k)
	}
}

func TestSearchMatrix(t *testing.T) {
	matrix := [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}
	if !searchMatrix(matrix, 5) {
		t.Error("5 should be exist")
	}

	if !searchMatrixV2(matrix, 5) {
		t.Error("5 should be exist")
	}
}

func TestKthSmallest(t *testing.T) {
	var tree = &model.TreeNode{
		Val: 5,
		Left: &model.TreeNode{
			Val: 3,
			Left: &model.TreeNode{
				Val: 2,
				Left: &model.TreeNode{
					Val: 1,
				},
			},
			Right: &model.TreeNode{
				Val: 4,
			},
		},
		Right: &model.TreeNode{
			Val: 6,
		},
	}
	k := kthSmallest(tree, 3)
	if k != 3 {
		t.Error("k should be 3, not", k)
	}
	t.Log(k)
}
