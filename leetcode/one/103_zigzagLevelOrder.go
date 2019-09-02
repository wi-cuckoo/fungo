// https://leetcode-cn.com/problems/binary-tree-zigzag-level-order-traversal/

package one

import "github.com/wi-cuckoo/fungo/model"

/*
给定一个二叉树，返回其节点值的锯齿形层次遍历。（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。

例如：
给定二叉树 [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回锯齿形层次遍历如下：

[
  [3],
  [20,9],
  [15,7]
]
*/

// 还是借助队列进行广度优先遍历，即层级遍历
// 需要注意的是录入结果集的时候需要变换顺序，顺序->逆序->顺序这样交替进行

func zigzagLevelOrder(root *model.TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	nums, flag := [][]int{}, false // flag 用于标识是否逆序录入节点数据
	queue := []*model.TreeNode{root}
	for len(queue) != 0 {
		ql := len(queue)
		levelNums := make([]int, ql)
		for i := 0; i < ql; i++ {
			node := queue[i]
			idx := i
			if flag {
				idx = ql - i - 1
			}
			levelNums[idx] = node.Val
			// 入队
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		nums = append(nums, levelNums)
		queue, flag = queue[ql:], !flag // 反转标识，截取队列
	}
	return nums
}
