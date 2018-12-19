package one

// TreeNode ...
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) Array() []int {
	nums := []int{}
	for p := l; p != nil; p = p.Next {
		nums = append(nums, p.Val)
	}
	return nums
}
