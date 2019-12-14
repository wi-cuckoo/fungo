package one

import (
	"testing"

	"github.com/wi-cuckoo/fungo/model"
)

func TestListToArray(t *testing.T) {
	head := &model.ListNode{
		Val: 1,
		Next: &model.ListNode{
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

func TestFractionToDecimal(t *testing.T) {
	t.Log(fractionToDecimal(4, 9))
	t.Log(fractionToDecimal(4, 333))
}

func TestSortList(t *testing.T) {
	head := model.NewListNode([]int{-1, 5, 3, 4, 0})
	sortListV2(head).Print()
}

func TestMaxProduct(t *testing.T) {
	nums := []int{7, -2, -4, 0, -19}
	max := maxProductV2(nums)
	if max != 56 {
		t.Errorf("max should be %d, not %d", 56, max)
	}
}

func TestCanCompleteCircuit(t *testing.T) {
	gas, cost := []int{5, 8, 2, 8}, []int{6, 5, 6, 6}
	t.Log(canCompleteCircuit(gas, cost) == 3)
}
