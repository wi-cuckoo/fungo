/*
给定一棵二叉树，想象自己站在它的右侧，按照从顶部到底部的顺序，返回从右侧所能看到的节点值。

示例:

输入: [1,2,3,null,5,null,4]
输出: [1, 3, 4]
解释:

   1            <---
 /   \
2     3         <---
 \     \
  5     4       <---

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/binary-tree-right-side-view
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package on

import "github.com/wi-cuckoo/fungo/model"

// 采用 BFS 即可，每一层从右侧开始

func rightSideView(root *model.TreeNode) []int {
	if root == nil {
		return []int{}
	}
	ans := make([]int, 0, 8)
	queue := make([]*model.TreeNode, 0, 32)
	queue = append(queue, root)
	for len(queue) > 0 {
		// record the first view from right side
		ans = append(ans, queue[0].Val)
		// dequeue
		n := len(queue)
		for i := 0; i < n; i++ {
			q := queue[i]
			if q.Right != nil {
				queue = append(queue, q.Right)
			}
			if q.Left != nil {
				queue = append(queue, q.Left)
			}
		}
		queue = queue[n:]
	}
	return ans
}
