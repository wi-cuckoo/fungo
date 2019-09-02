package one

// 链表相交问题 —— 通过长度来寻找相交位置
// 分别遍历两链表，得出长度分别为 l1, l2
// 若 l2 > l1，则将 B 链表头节点指向其第 l2-l1 个节点，反之将 A 链表指向其第 l1-l2 个节点
// 此时两条链表长度相同，同时将头指针以一个单位移动，判断两链表头指针相等不，相等则是相交点
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	l1, l2 := 0, 0
	for p := headA; p != nil; p = p.Next {
		l1++
	}
	for p := headB; p != nil; p = p.Next {
		l2++
	}
	if l2 > l1 {
		for ; l2 > l1; headB = headB.Next {
			l2--
		}
	} else {
		for ; l2 < l1; headA = headA.Next {
			l1--
		}
	}
	for headA != headB {
		headA = headA.Next
		headB = headB.Next
	}
	return headA
}
