/*
删除链表中等于给定值 val 的所有节点。

示例:

输入: 1->2->6->3->4->5->6, val = 6
输出: 1->2->3->4->5
*/

package two

import (
	"github.com/wi-cuckoo/fungo/model"
)

// 递归版本，比较简单
func removeElements(head *model.ListNode, val int) *model.ListNode {
	if head == nil {
		return head
	}
	if head.Val == val {
		return removeElements(head.Next, val)
	}
	head.Next = removeElements(head.Next, val)
	return head
}

// 循环版本
func removeElementsV2(head *model.ListNode, val int) *model.ListNode {
	l := &model.ListNode{}
	p := l
	for r := head; r != nil; r = r.Next {
		if r.Val == val {
			continue
		}
		l.Next = r
		l = r
	}
	l.Next = nil
	return p.Next
}
