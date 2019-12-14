/*
给定一个二叉树，确定它是否是一个完全二叉树。

百度百科中对完全二叉树的定义如下：

若设二叉树的深度为 h，除第 h 层外，其它各层 (1～h-1) 的结点数都达到最大个数，第 h 层所有的结点都连续集中在最左边，这就是完全二叉树。（注：第 h 层可能包含 1~ 2h 个节点。）



示例 1：



输入：[1,2,3,4,5,6]
输出：true
解释：最后一层前的每一层都是满的（即，结点值为 {1} 和 {2,3} 的两层），且最后一层中的所有结点（{4,5,6}）都尽可能地向左。
示例 2：



输入：[1,2,3,4,5,null,7]
输出：false
解释：值为 7 的结点没有尽可能靠向左侧。
*/

package nine

import "github.com/wi-cuckoo/fungo/model"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isCompleteTree(root *model.TreeNode) bool {
	if root == nil {
		return false
	}

	queue := []*model.TreeNode{root}
	for len(queue) != 0 {
		node := queue[0]
		queue = queue[1:]
		if node.Left != nil && node.Right != nil {
			queue = append(queue, node.Left)
			queue = append(queue, node.Right)
			continue
		}
		if node.Right != nil && node.Left == nil {
			return false
		}
		if node.Right == nil {
			for i := 0; i < len(queue); i++ {
				if queue[i].Left == nil && queue[i].Right == nil {
					continue
				}
				return false
			}
			queue = queue[len(queue):]
		}

	}
	return true
}
