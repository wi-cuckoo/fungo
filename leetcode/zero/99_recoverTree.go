/*
二叉搜索树中的两个节点被错误地交换。
请在不改变其结构的情况下，恢复这棵树。

示例 1:

输入: [1,3,null,null,2]

   1
  /
 3
  \
   2

输出: [3,1,null,null,2]

   3
  /
 1
  \
   2
示例 2:

输入: [3,1,4,null,null,2]

  3
 / \
1   4
   /
  2

输出: [2,1,4,null,null,3]

  2
 / \
1   4
   /
  3
进阶:

使用 O(n) 空间复杂度的解法很容易实现。
你能想出一个只使用常数空间的解决方案吗？

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/recover-binary-search-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package zero

import "github.com/wi-cuckoo/fungo/model"

var nodes []*model.TreeNode

func recoverTree(root *model.TreeNode) {
	nodes = make([]*model.TreeNode, 0, 64)
	inorder(root)
	i, j := 0, 0
	for ; i < len(nodes)-1; i++ {
		if nodes[i].Val > nodes[i+1].Val {
			break
		}
	}
	for j = len(nodes) - 1; j > 0; j-- {
		if nodes[j].Val < nodes[j-1].Val {
			break
		}
	}
	nodes[i].Val, nodes[j].Val = nodes[j].Val, nodes[i].Val
}

func inorder(root *model.TreeNode) {
	if root == nil {
		return
	}
	inorder(root.Left)
	nodes = append(nodes, root)
	inorder(root.Right)
}
