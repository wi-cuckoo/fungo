package two

import "fmt"

// ListNode ...
type ListNode struct {
	Val  int
	Next *ListNode
}

// Print lll
func (l *ListNode) Print() {
	if l == nil {
		return
	}
	fmt.Print(l.Val, " ")
	l.Next.Print()
}
