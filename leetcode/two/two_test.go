package two

import (
	"fmt"
	"testing"

	"github.com/wi-cuckoo/fungo/model"
)

func TestReverseList(t *testing.T) {
	l1 := &model.ListNode{
		Val: 1,
		Next: &model.ListNode{
			Val: 1,
			Next: &model.ListNode{
				Val: 2,
				Next: &model.ListNode{
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
	t.Log(nums)
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

func TestProductExceptSelf(t *testing.T) {
	nums := make([]int, 74989)
	for i := range nums {
		if i%2 == 0 {
			nums[i] = 1
			continue
		}
		nums[i] = -1
	}
	x := productExceptSelf(nums)
	t.Log(len(x))
}

func TestCalculate(t *testing.T) {
	s := " 3+5 / 2 + 23 - 100*23"
	if calculateV2(s) != -2272 {
		t.Error("fuck it")
	}
}

func TestMaxSlidingWindow(t *testing.T) {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k := 3
	res := maxSlidingWindow(nums, k)
	t.Log(res)
}

func TestMaximalSquare(t *testing.T) {
	matrix := [][]byte{
		{'1', '0', '1', '0', '0'},
		{'1', '0', '0', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '0', '1', '1', '1'},
	}
	if m := maximalSquare(matrix); m != 9 {
		t.Errorf("the maximalSquare should be 4, not %d", m)
	}
}

func TestGetSkyline(t *testing.T) {
	bs := [][]int{{2, 9, 10}, {3, 7, 15}, {5, 12, 12}, {15, 20, 10}, {19, 24, 8}}
	t.Log(getSkyline(bs))
}

func TestRemoveElements(t *testing.T) {
	head := model.NewListNode([]int{1, 2, 6, 3, 4, 6, 6, 7})
	removeElementsV2(head, 6).Print()
}
