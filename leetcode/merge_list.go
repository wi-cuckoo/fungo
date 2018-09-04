/*
将两个有序链表合并为一个新的有序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。

示例：

输入：1->2->4, 1->3->4
输出：1->1->2->3->4->4
*/
package main

import (
	"fmt"
)

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
			Val: 2,
			Next: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val:  5,
					Next: nil,
				},
			},
		},
	}
	l2 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}
	mergeTwoLists(l1, l2).Print()
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	cur := &ListNode{}
	head := &ListNode{Next: cur}
	for {
		p := cur
		if l1 == nil && l2 == nil {
			break
		}

		if l1 == nil && l2 != nil {
			p.Next = l2
			l2 = l2.Next
		} else if l1 != nil && l2 == nil {
			p.Next = l1
			l1 = l1.Next
		} else if l1.Val >= l2.Val {
			p.Next = l2
			l2 = l2.Next
		} else {
			p.Next = l1
			l1 = l1.Next
		}
		cur = p.Next
	}
	return head.Next.Next
}
