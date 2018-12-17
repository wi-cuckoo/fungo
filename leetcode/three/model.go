package three

import "fmt"

//ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// Print ...
func (l *ListNode) Print() {
	if l == nil {
		return
	}
	fmt.Print(l.Val, " ")
	l.Next.Print()
}

// NewLine ...
func (l *ListNode) NewLine() {
	fmt.Println()
}
