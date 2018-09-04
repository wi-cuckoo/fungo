/*
给定一个排序链表，删除所有重复的元素，使得每个元素只出现一次。

示例 1:

输入: 1->1->2
输出: 1->2
示例 2:

输入: 1->1->2->3->3
输出: 1->2->3
*/

package main

import (
	"fmt"
)

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) Print() {
	if l == nil {
		return
	}
	fmt.Println(l.Val)
	l.Next.Print()
}

func main() {
	l1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val:  3,
						Next: nil,
					},
				},
			},
		},
	}
	deleteDuplicates(l1).Print()
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	p1, p2 := head, head.Next
	for p2 != nil {
		if p1.Val == p2.Val {
			p2 = p2.Next
			continue
		}
		p1.Next = p2
		p1, p2 = p2, p2.Next
	}
	if p1.Next != p2 {
		p1.Next = nil
	}
	return head
}
