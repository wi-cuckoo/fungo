/*
将两个有序链表合并为一个新的有序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。

示例：

输入：1->2->4, 1->3->4
输出：1->1->2->3->4->4
*/

package zero

import (
	"github.com/wi-cuckoo/fungo/model"
)

// 8ms ?
func mergeTwoLists(l1 *model.ListNode, l2 *model.ListNode) *model.ListNode {
	head := &model.ListNode{}
	p := head
	for ; l1 != nil && l2 != nil; p = p.Next {
		if l1.Val < l2.Val {
			p.Next = l1
			l1 = l1.Next
			continue
		}
		p.Next = l2
		l2 = l2.Next
	}
	if l1 != nil {
		p.Next = l1
	} else {
		p.Next = l2
	}
	return head.Next
}

// 4ms ?
func mergeTwoListsV2(l1 *model.ListNode, l2 *model.ListNode) *model.ListNode {
	head := &model.ListNode{}
	cur := head
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
	return head.Next
}
