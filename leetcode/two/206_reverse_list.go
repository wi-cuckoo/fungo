/*

反转一个单链表。

示例:

输入: 1->2->3->4->5->NULL
输出: 5->4->3->2->1->NULL
进阶:
你可以迭代或递归地反转链表。你能否用两种方法解决这道题？
*/

package two

func reverseList(head *ListNode) *ListNode {
	tail := &ListNode{
		Val:  head.Val,
		Next: nil,
	}
	for head.Next != nil {
		head = head.Next
		n := &ListNode{
			Val:  head.Val,
			Next: tail,
		}
		tail = n
	}
	return tail
}

func reverseList2(head *ListNode) *ListNode {
	var p *ListNode
	q := head
	for q != nil {
		n := q.Next
		q.Next = p
		p, q = q, n
	}
	return p
}
