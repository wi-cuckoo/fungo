/*
给定一个二叉树，返回其按层次遍历的节点值。 （即逐层地，从左到右访问所有节点）。

例如:
给定二叉树: [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回其层次遍历结果：

[
  [3],
  [9,20],
  [15,7]
]
*/

package one

// 使用广度优先搜索 BFS， 即可达到目的
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	nums := [][]int{}
	queue := []*TreeNode{root}
	for len(queue) != 0 {
		ql := len(queue)
		levelNums := make([]int, ql)
		for i, node := range queue[0:ql] {
			levelNums[i] = node.Val
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		nums = append(nums, levelNums)
		queue = queue[ql:]
	}
	return nums
}
