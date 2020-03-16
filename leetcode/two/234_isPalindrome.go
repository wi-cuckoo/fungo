package two

import "github.com/wi-cuckoo/fungo/model"

/*
请判断一个链表是否为回文链表。

示例 1:

输入: 1->2
输出: false
示例 2:

输入: 1->2->2->1
输出: true
进阶：
你能否用 O(n) 时间复杂度和 O(1) 空间复杂度解决此题？

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/palindrome-linked-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func isPalindrome(head *model.ListNode) bool {
	n := 0
	for p := head; p != nil; {
		n++
		p = p.Next
	}
	// reverse half of list
	var q *model.ListNode // dummy node
	p := head
	for i := 0; i < n/2; i++ {
		n := p.Next
		p.Next = q
		q, p = p, n
	}
	if n%2 > 0 {
		p = p.Next
	}
	for p != nil && q != nil {
		if p.Val != q.Val {
			return false
		}
		p, q = p.Next, q.Next
	}
	return p == nil && q == nil
}
