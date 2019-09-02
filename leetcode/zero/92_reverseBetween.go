/*
反转从位置 m 到 n 的链表。请使用一趟扫描完成反转。

说明:
1 ≤ m ≤ n ≤ 链表长度。

示例:

输入: 1->2->3->4->5->NULL, m = 2, n = 4
输出: 1->4->3->2->5->NULL
*/

package zero

import (
	"github.com/wi-cuckoo/fungo/model"
)

// 与整体旋转类似，只是把其中一段当成整体来旋转。注意边界即当从头结点开始旋转（m=1）
// 1. 先将 p1 执行第 m-1 个节点，比如从第二个开始旋转，则 p1 就指向第一个节点
// 2. 将 q 节点指向 p1 下一个节点，如果是从头结点开始旋转，则将 q 节点指向头结点
// 3. 设定旋转需要遍历的范围 [0, m-n], 然后旋转该子段链表，完事后，p2指向第n个节点，q指向第n+1个节点，也可能是nil
// 4. 当是头结点开始旋转，则将 p1.Next 指向 q，返回 p2 即可；否则将 p1.Next.Next 指向 q后再将 p1.Next 指向 p2，返回 head
func reverseBetween(head *model.ListNode, m int, n int) *model.ListNode {
	p1 := head
	for i := 1; p1 != nil && i < m-1; i++ {
		p1 = p1.Next
	}
	var p2 *model.ListNode
	q := p1.Next
	// 如果是从第一个数开始旋转，则q取头结点
	if m == 1 {
		q = head
	}
	for i := 0; q != nil && i <= n-m; i++ {
		n := q.Next
		q.Next = p2
		p2, q = q, n
	}
	if m == 1 {
		p1.Next = q
		return p2
	}
	p1.Next.Next = q
	p1.Next = p2
	return head
}
