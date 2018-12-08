/*
给定一个二叉树，检查它是否是镜像对称的。

例如，二叉树 [1,2,2,3,4,4,3] 是对称的。

    1
   / \
  2   2
 / \ / \
3  4 4  3
但是下面这个 [1,2,2,null,3,null,3] 则不是镜像对称的:

    1
   / \
  2   2
   \   \
   3    3

说明:

如果你可以运用递归和迭代两种方法解决这个问题，会很加分。
*/

package one

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	left, right := []int{}, []int{}
	n1 := leftEach(root.Left, left)
	n2 := rightEach(root.Right, right)
	// fmt.Println(n1, n2)
	if len(n1) != len(n2) {
		return false
	}
	for i, v := range n1 {
		if v != n2[i] {
			return false
		}
	}
	return true
}

// 对左子树采用前序遍历
func leftEach(root *TreeNode, nums []int) []int {
	if root == nil {
		nums = append(nums, 0)
		return nums
	}
	nums = append(nums, root.Val)
	nums = leftEach(root.Left, nums)
	nums = leftEach(root.Right, nums)
	return nums
}

// 对右子树采用变形的前序遍历
func rightEach(root *TreeNode, nums []int) []int {
	if root == nil {
		nums = append(nums, 0)
		return nums
	}
	nums = append(nums, root.Val)
	nums = rightEach(root.Right, nums)
	nums = rightEach(root.Left, nums)
	return nums
}
