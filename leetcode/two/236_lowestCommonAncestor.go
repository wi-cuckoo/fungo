/*
给定一个二叉树, 找到该树中两个指定节点的最近公共祖先。

百度百科中最近公共祖先的定义为：“对于有根树 T 的两个结点 p、q，最近公共祖先表示为一个结点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。”

例如，给定如下二叉搜索树:  root = [6,2,8,0,4,7,9,null,null,3,5]
![](https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2018/12/14/binarysearchtree_improved.png)


示例 1:
输入: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 1
输出: 3
解释: 节点 5 和节点 1 的最近公共祖先是节点 3。

示例 2:
输入: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 4
输出: 5
解释: 节点 5 和节点 4 的最近公共祖先是节点 5。因为根据定义最近公共祖先节点可以为节点本身。


说明:

所有节点的值都是唯一的。
p、q 为不同节点且均存在于给定的二叉树中。
*/

package two

import (
	"github.com/wi-cuckoo/fungo/model"
)

// 反转所有节点的左指针指向自己的父节点，则 p， q两点到根节点的路径可看着单链表
// 于是问题就转化为求两条单链表的相交点，真tm机智
// 缺点是改变了树的根本结构，用完就成了一棵废树
func lowestCommonAncestor2(root, p, q *model.TreeNode) *model.TreeNode {
	// 反转节点左指针指向自个的父节点
	reverse(root, nil)
	cache := make(map[int]*model.TreeNode)
	for n := p; n != nil; n = n.Left {
		cache[n.Val] = n
	}
	for n := q; n != nil; n = n.Left {
		if target, ok := cache[n.Val]; ok {
			return target
		}
	}
	return nil
}

func reverse(node, parent *model.TreeNode) {
	if node == nil {
		return
	}
	reverse(node.Left, node)
	node.Left = parent
	reverse(node.Right, node)
}

// 第二种方案：利用递归的思想，
// p, q 一定存在，并且分散在某个节点的左右子树中，特殊情况是该节点为 p 或者 q
// 那么可以递归从左右子树中寻找
func lowestCommonAncestorV2(root, p, q *model.TreeNode) *model.TreeNode {
	if root == nil || root == q || root == p {
		return root
	}
	left := lowestCommonAncestorV2(root.Left, p, q)
	right := lowestCommonAncestorV2(root.Right, p, q)
	if left != nil && right != nil {
		return root
	}
	if left != nil {
		return left
	}
	return right
}
