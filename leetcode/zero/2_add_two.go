/*
给定两个非空链表来表示两个非负整数。位数按照逆序方式存储，它们的每个节点只存储单个数字。将两数相加返回一个新的链表。

你可以假设除了数字 0 之外，这两个数字都不会以零开头。

示例：
```
输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 0 -> 8
原因：342 + 465 = 807
```
*/

package zero

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) Print() {
	if l == nil {
		return
	}
	fmt.Println(l.Val)
	l.Next.Print()
}

func main() {
	l1 := &ListNode{
		Val: 7,
		Next: &ListNode{
			Val:  4,
			Next: nil,
		},
	}
	l2 := &ListNode{
		Val:  4,
		Next: nil,
	}
	addTwoNumbersV2(l1, l2).Print()
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	res := &ListNode{}
	p, p1, p2 := res, l1, l2
	last := &ListNode{}
	e := 0 // for sum > 10
	for p1 != nil && p2 != nil {
		v := p1.Val + p2.Val + e
		e = 0 // reset e to zero
		if v >= 10 {
			e = 1
			v -= 10
		}
		p.Val, p.Next = v, &ListNode{}
		p1, p2 = p1.Next, p2.Next
		last, p = p, p.Next
	}
	more := p1
	if p2 != nil {
		more = p2
	}
	for more != nil {
		v := more.Val + e
		e = 0 // reset e to zero
		if v >= 10 {
			e = 1
			v -= 10
		}
		p.Val, p.Next = v, &ListNode{}
		more = more.Next
		last, p = p, p.Next
	}
	last.Next = nil
	if e == 1 {
		last.Next = &ListNode{Val: 1}
	}
	return res
}

func addTwoNumbersV2(l1 *ListNode, l2 *ListNode) *ListNode {
	res := &ListNode{}
	p, last := res, &ListNode{}
	e := 0
	for l1 != nil || l2 != nil || e != 0 {
		v := e
		if l1 != nil {
			v += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			v += l2.Val
			l2 = l2.Next
		}
		p.Val, e = v%10, v/10
		p.Next = &ListNode{}
		last, p = p, p.Next
	}
	last.Next = nil
	return res
}
