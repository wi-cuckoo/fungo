/*
给定一个二叉树，判断它是否是高度平衡的二叉树。

本题中，一棵高度平衡二叉树定义为：

一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过1。
*/

package one

import (
	"github.com/wi-cuckoo/fungo/model"
	"github.com/wi-cuckoo/fungo/util"
)

func isBalanced(root *model.TreeNode) bool {
	return height(root) != -1
}

func height(root *model.TreeNode) int {
	if root == nil {
		return 0
	}
	l := height(root.Left)
	r := height(root.Right)
	if l == -1 || r == -1 || util.Abs(l-r) > 1 {
		return -1
	}
	return util.Max(l, r) + 1
}
