package one

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
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
