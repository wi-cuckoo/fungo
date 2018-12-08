package two

import (
	"fmt"
	"testing"
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
