/*
给定一个链表，返回链表开始入环的第一个节点。 如果链表无环，则返回 null。

为了表示给定链表中的环，我们使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。 如果 pos 是 -1，则在该链表中没有环。

说明：不允许修改给定的链表。
*/

package one

import (
	"github.com/wi-cuckoo/fungo/model"
)

func detectCycle(head *model.ListNode) *model.ListNode {
	if head == nil {
		return nil
	}
	p, q := head, head.Next
	for p != q {
		if q == nil || q.Next == nil {
			return nil
		}
		p, q = p.Next, q.Next.Next
	}
	// 此时 p, q 为相遇点
	// 此时将p放到头节点，q从相交节点往前走一步
	// 注：为啥要往前走一步哩，因为最开始就抢跑了一步。。。。
	p, q = head, q.Next
	for p != q {
		p, q = p.Next, q.Next
	}
	return p
}
