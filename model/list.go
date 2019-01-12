package model

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// Array convert list to array
func (l *ListNode) Array() []int {
	nums := []int{}
	for p := l; p != nil; p = p.Next {
		nums = append(nums, p.Val)
	}
	return nums
}

// Print the list as array
func (l *ListNode) Print() {
	fmt.Println(l.Array())
}

// NewListNode help to build a linked nodes by array
func NewListNode(nums []int) *ListNode {
	head := &ListNode{}
	p1 := head
	for _, n := range nums {
		n := &ListNode{
			Val: n,
		}
		p1.Next = n
		p1 = p1.Next
	}
	return head.Next
}
