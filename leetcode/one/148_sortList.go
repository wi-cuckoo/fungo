/*
在 O(n log n) 时间复杂度和常数级空间复杂度下，对链表进行排序。

示例 1:

输入: 4->2->1->3
输出: 1->2->3->4
示例 2:

输入: -1->5->3->4->0
输出: -1->0->3->4->5
*/

package one

import (
	"sort"

	"github.com/wi-cuckoo/fungo/model"
)

// 时间复杂度 N*Log(N), 空间复杂度 N
func sortList(head *model.ListNode) *model.ListNode {
	nodes := make([]*model.ListNode, 0)
	for p := head; p != nil; p = p.Next {
		nodes = append(nodes, p)
	}
	if len(nodes) == 0 {
		return nil
	}
	sort.Slice(nodes, func(i, j int) bool {
		return nodes[i].Val < nodes[j].Val
	})
	for i := 0; i < len(nodes)-1; i++ {
		nodes[i].Next = nodes[i+1]
	}
	nodes[len(nodes)-1].Next = nil
	return nodes[0]
}

// 使用变形后的快排, 空间复杂度可降到常数级别，沃日，好像提交后慢了不少，评论区说用归并，等会搞个归并的V3版
// 1. 快排的话取头节点为基准节点，将后续的节点分为两条链表，一条所有值大于基准节点，一条小于基准节点
// 2. 分别对两条链表递归调用排序，然后将基准节点与两条链表链接起来
//  需要注意的是，较小链表排序后返回的是头节点，需要循环得打其尾节点，然后尾节点下一个节点指向基准节点

// 1. 归并排序的话，使用快慢指针将链表一分为二，分别递归调用
// 2. 调用完后，两条链表变为有序，则问题转化为合并两条有序链表
func sortListV2(head *model.ListNode) *model.ListNode {
	if head == nil {
		return head
	}
	pivot := head
	more := &model.ListNode{}
	less := &model.ListNode{}
	p1, p2 := more, less
	for p := head.Next; p != nil; {
		// 保存当前节点指向的下一个节点
		cur, next := p, p.Next
		cur.Next = nil
		if cur.Val >= pivot.Val {
			p1.Next = cur
			p1 = p1.Next
		} else {
			p2.Next = cur
			p2 = p2.Next
		}
		p = next
	}
	pivot.Next = sortListV2(more.Next)
	less = sortListV2(less.Next)
	for p := less; p != nil; p = p.Next {
		if p.Next == nil {
			p.Next = pivot
			break
		}
	}

	if less == nil {
		return pivot
	}

	return less
}
