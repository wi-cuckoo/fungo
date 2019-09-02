/*
合并 k 个排序链表，返回合并后的排序链表。请分析和描述算法的复杂度。

示例:

输入:
[
  1->4->5,
  1->3->4,
  2->6
]
输出: 1->1->2->3->4->4->5->6
*/

package zero

import (
	"github.com/wi-cuckoo/fungo/model"
)

func mergeKLists(lists []*model.ListNode) *model.ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	if len(lists) == 2 {
		return mergeTwoLists(lists[0], lists[1])
	}

	l1 := mergeKLists(lists[:len(lists)/2])
	l2 := mergeKLists(lists[len(lists)/2:])
	return mergeTwoLists(l1, l2)
}
