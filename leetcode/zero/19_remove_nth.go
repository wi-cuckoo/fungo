package zero

import "github.com/wi-cuckoo/fungo/model"

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

// 双指针，开始时p1,p2都指向头节点
// 让 p2 指针先向前走 n+1 步，然后 p1, p2 同时开始走
// 当 p2 到达末尾节点时，p1节点的下一个节点就是需要删除的

func removeNthFromEnd(head *model.ListNode, n int) *model.ListNode {
	dummy := &model.ListNode{Next: head}
	p1, p2 := dummy, head
	for ; n > 0; n-- {
		p2 = p2.Next
	}
	for p2 != nil {
		p1, p2 = p1.Next, p2.Next
	}
	p1.Next = p1.Next.Next
	return dummy.Next
}
