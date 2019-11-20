package zero

/*
将两个有序链表合并为一个新的有序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。

示例：

输入：1->2->4, 1->3->4
输出：1->1->2->3->4->4
*/

import (
	"github.com/wi-cuckoo/fungo/model"
)

// 很直观的链表操作
// 创建一个呆逼节点（dummy), 将呆逼节点赋给一个p指针
// 开始遍历，比较l1和l2当前节点大小，将p指向较小的那个
// 染后将其后移一个单位，并将p指向自身下一个节点，开始下一轮循环，直到有一个链表结束为止

// 8ms ?
func mergeTwoLists(l1 *model.ListNode, l2 *model.ListNode) *model.ListNode {
	dummy := &model.ListNode{}
	p := dummy
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
	return dummy.Next
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
