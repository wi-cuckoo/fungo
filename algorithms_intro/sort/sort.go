package sort

import (
	"github.com/wi-cuckoo/fungo/model"
)

// SelectSort sort nums with selection [3, 2, 4, 5, 1, 9]
func SelectSort(nums []int) {
	l := len(nums)
	for i := 0; i < l; i++ {
		if i == l-1 {
			break
		}
		n, min := i, nums[i] // index of max num in nums[i:]
		for j := i + 1; j < l; j++ {
			if nums[j] < min {
				min = nums[j]
				n = j
			}
		}
		nums[i], nums[n] = nums[n], nums[i]
	}
}

// QuickSort use quick sort
func QuickSort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	pivot := nums[0]
	less := make([]int, 0)
	greater := make([]int, 0)
	for _, i := range nums[1:] {
		if i > pivot {
			greater = append(greater, i)
		} else {
			less = append(less, i)
		}
	}
	result := append(QuickSort(less), pivot)
	return append(result, QuickSort(greater)...)
}

// QuickSortList for linked nodes
// 1. 以头节点为基准节点，将整条链表分成两段，一段所有值小于基准值，另一端大于等于基准值
// 2. 递归的将两段链表排序
// 3. 合并两段链表及基准节点
// TODO 优化点是可以把排序函数单独抽出来，并返回每次排完的首尾节点，方便合并用
func QuickSortList(head *model.ListNode) *model.ListNode {
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
	pivot.Next = QuickSortList(more.Next)
	less = QuickSortList(less.Next)
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
