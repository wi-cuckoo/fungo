/*
给定一个二叉树，编写一个函数来获取这个树的最大宽度。树的宽度是所有层中的最大宽度。这个二叉树与满二叉树（full binary tree）结构相同，但一些节点为空。

每一层的宽度被定义为两个端点（该层最左和最右的非空节点，两端点间的null节点也计入长度）之间的长度。

示例 1:

输入:

           1
         /   \
        3     2
       / \     \
      5   3     9

输出: 4
解释: 最大值出现在树的第 3 层，宽度为 4 (5,3,null,9)。
*/

package six

import (
	"github.com/wi-cuckoo/fungo/model"
)

// 使用队列广度优先遍历方式真费时，主要耗时在于append操作
func widthOfBinaryTree(root *model.TreeNode) int {
	if root == nil {
		return 0
	}
	queue, max := []*model.TreeNode{root}, 1
	for len(queue) != 0 {
		l := len(queue)
		for i := 0; i < l; i++ {
			if queue[i] == nil {
				queue = append(queue, nil)
				queue = append(queue, nil)
				continue
			}
			queue = append(queue, queue[i].Left)
			queue = append(queue, queue[i].Right)
		}
		// 去掉 queue[l:] 队列首尾的 nil 节点
		ql := len(queue) - 1
		for ql >= l {
			if queue[l] == nil {
				l++
				continue
			}
			if queue[ql] == nil {
				ql--
				continue
			}
			break
		}
		if ql < l {
			break
		}
		queue = queue[l : ql+1]
		if m := len(queue); m > max {
			max = m
		}
	}
	return max
}

// 尝试使用深度优先搜索
/**
  假设满二叉树表示成数组序列, 根节点所在的位置为1, 则任意位于i节点的左右子节点的index为2*i, 2*i+1
  用一个List保存每层的左端点, 易知二叉树有多少层List的元素就有多少个. 那么可以在dfs的过程中记录每个
  节点的index及其所在的层level, 如果level > List.size()说明当前节点就是新的一层的最左节点, 将其
  加入List中, 否则判断当前节点的index减去List中对应层的最左节点的index的宽度是否大于最大宽度并更新
  **/
func fuckDFS(root *model.TreeNode, level, index int, left *[]int) int {
	if root == nil {
		return 0
	}
	if level == len(*left) {
		*left = append(*left, index)
	}
	width := index - (*left)[level] + 1
	l := fuckDFS(root.Left, level+1, index*2, left)
	r := fuckDFS(root.Right, level+1, index*2+1, left)
	if l > width {
		width = l
	}
	if r > width {
		width = r
	}
	return width
}
