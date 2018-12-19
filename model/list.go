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
