package sort

import (
	"github.com/wi-cuckoo/fungo/model"
)

// SelectSort sort nums with selection [3, 2, 4, 5, 1, 9]
func SelectSort(nums []int) {
	l := len(nums)
	for i := 0; i < l-1; i++ {
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

// QuickSortV2 快速排序的原地算法，不申请额外存储空间
func QuickSortV2(nums []int) {
	if len(nums) < 2 {
		return
	}
	pivot, l := 0, len(nums)
	for i := 1; i < l; i++ {
		// 如果当前值小于基准值，放到基准左边去
		if nums[i] <= nums[pivot] {
			nums[i], nums[pivot] = nums[pivot], nums[i]
			pivot = i
			continue
		}
		// 如果当前值大于基准值，则与最后一个大佬交换
		nums[i], nums[l-1] = nums[l-1], nums[i]
		l-- // 在下一轮需要比较倒数第二个
		i-- // 这里很关键，需要继续比较当前值
	}
	QuickSortV2(nums[:pivot])
	QuickSortV2(nums[pivot+1:])
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

// BubbleSort 冒泡排序时间复杂度 N^2, 空间复杂度 1
// 1. 循环一遍，一次比较两个元素，如果顺序错误就交换
// 2. 重复第一步，直到无需交换为止
// 注意：每轮比较后，最大的已经安排好了，所以次轮可以少比较一个
func BubbleSort(nums []int) {
	for offset := 1; offset < len(nums); offset++ {
		// 使用 offset 将待排序范围缩小
		for i := 0; i < len(nums)-offset; i++ {
			if nums[i] > nums[i+1] {
				nums[i], nums[i+1] = nums[i+1], nums[i]
			}
		}
	}
}

// InsertSort 插入排序，时间复杂度 N^2, 空间复杂度 1
// 1. 从第二个元素开始遍历，认为此前所有元素已经排序完毕
// 2. 从此前的序列中找到合适的插入位置，该位置及之后的元素顺移
func InsertSort(nums []int) {
	for i := 1; i < len(nums); i++ {
		// find the pos to insert
		for j := i; j > 0; j-- {
			if nums[j] < nums[j-1] {
				nums[j], nums[j-1] = nums[j-1], nums[j]
				continue
			}
			break
		}
	}
}

// MergeSort 归并排序
func MergeSort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	mid := len(nums) / 2
	return merge(MergeSort(nums[0:mid]), MergeSort(nums[mid:]))
}

func merge(left []int, right []int) []int {
	res := make([]int, len(left)+len(right))
	offset, l, r := 0, 0, 0
	for l < len(left) && r < len(right) {
		if left[l] < right[r] {
			res[offset] = left[l]
			l++
		} else {
			res[offset] = right[r]
			r++
		}
		offset++
	}
	for ; l < len(left); l++ {
		res[offset] = left[l]
		offset++
	}
	for ; r < len(right); r++ {
		res[offset] = right[r]
		offset++
	}
	return res
}

// HeapSort 堆排序
func HeapSort(nums []int) []int {
	res := make([]int, len(nums))
	h := model.NewHeap(nums)
	for i := 0; i < len(res); i++ {
		res[i] = h.Pop()
	}
	return res
}
