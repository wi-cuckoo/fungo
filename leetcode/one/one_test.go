package one

import "testing"

func TestListToArray(t *testing.T) {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val:  2,
			Next: nil,
		},
	}
	nums := head.Array()
	for i, v := range []int{1, 2} {
		if nums[i] != v {
			t.Errorf("nums[%d] shoudl be %d, not %d", i, v, nums[i])
		}
	}
}

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

func TestBuildTree1(t *testing.T) {
	pre := []int{1, 2}
	in := []int{1, 2}
	tree := buildTree(pre, in)
	for i, v := range tree.PreOrderTraversal() {
		if pre[i] != v {
			t.Errorf("pre[%d] should be %d, not %d", i, pre[i], v)
		}
	}
	t.Log(tree.PreOrderTraversal())
}
