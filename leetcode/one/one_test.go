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

func TestTitleToNumber(t *testing.T) {
	s := "A"
	if titleToNumber(s) != 1 {
		t.Error(s + "should be 1")
	}

	s = "AB"
	if titleToNumber(s) != 28 {
		t.Error(s + "should be 28")
	}
}

func TestEvalRPN(t *testing.T) {
	tokens := map[int][]string{
		9:  {"2", "1", "+", "3", "*"},
		22: {"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"},
	}
	for k, token := range tokens {
		x := evalRPN(token)
		if x != k {
			t.Errorf("%v should be %d, not %d", token, k, x)
		}
	}
}
