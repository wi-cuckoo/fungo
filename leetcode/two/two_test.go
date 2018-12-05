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
