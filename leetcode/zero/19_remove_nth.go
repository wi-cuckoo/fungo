/*
给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点。

示例：

给定一个链表: 1->2->3->4->5, 和 n = 2.

当删除了倒数第二个节点后，链表变为 1->2->3->5.
说明：

给定的 n 保证是有效的。

进阶：

你能尝试使用一趟扫描实现吗？
*/

package zero

import "fmt"

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
			Val:  2,
			Next: nil,
		},
	}
	removeNthFromEnd(l1, 1).Print()
	fmt.Println("")
	removeNthFromEnd(l2, 2).Print()
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	p1, p2 := head, head
	if p1 == nil {
		return p1
	}
	for p2.Next != nil {
		if n > 0 {
			n--
			p2 = p2.Next
			continue
		}
		p1, p2 = p1.Next, p2.Next
	}
	fmt.Println(p1.Val, p2.Val, n)
	if n != 0 {
		return head.Next
	}
	p1.Next = p1.Next.Next
	return head
}
