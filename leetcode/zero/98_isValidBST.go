/*
给定一个二叉树，判断其是否是一个有效的二叉搜索树。

假设一个二叉搜索树具有如下特征：

节点的左子树只包含小于当前节点的数。
节点的右子树只包含大于当前节点的数。
所有左子树和右子树自身必须也是二叉搜索树。
示例 1:

输入:
    2
   / \
  1   3
输出: true
示例 2:

输入:
    5
   / \
  1   4
     / \
    3   6
输出: false
解释: 输入为: [5,1,4,null,null,3,6]。
     根节点的值为 5 ，但是其右子节点值为 4 。
*/

package zero

import "github.com/wi-cuckoo/fungo/model"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *model.TreeNode) bool {
	nums := make([]int, 0)
	nums = midEach(root, nums)
	for i := 1; i < len(nums); i++ {
		if nums[i-1] >= nums[i] {
			return false
		}
	}
	return true
}

// 判断是不是二叉搜索树使用中序遍历，判断遍历后的数组是不是有序就行了，我真的傻想了半天，忘了这茬
// 使用递归的时候需要保存每个节点对应取值的上界和下界，因为右子树的所有节点值要大于左子树的节点值
// 并不仅仅是左节点小于父节点，右节点大于父节点
func midEach(root *model.TreeNode, nums []int) []int {
	if root == nil {
		return nums
	}
	nums = midEach(root.Left, nums)
	nums = append(nums, root.Val)
	nums = midEach(root.Right, nums)
	return nums
}
