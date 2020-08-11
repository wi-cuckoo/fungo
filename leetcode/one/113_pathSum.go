/*
给定一个二叉树和一个目标和，找到所有从根节点到叶子节点路径总和等于给定目标和的路径。
说明: 叶子节点是指没有子节点的节点。

示例:
给定如下二叉树，以及目标和 sum = 22，

              5
             / \
            4   8
           /   / \
          11  13  4
         /  \    / \
        7    2  5   1
返回:

[
   [5,4,11,2],
   [5,8,4,5]
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/path-sum-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package one

import "github.com/wi-cuckoo/fungo/model"

// 思路就是回溯，没啥好说的

var ans [][]int

func pathSum(root *model.TreeNode, sum int) [][]int {
	ans = make([][]int, 0, 32)
	visited := make([]int, 0, 32)
	back(root, sum, visited)
	return ans
}

func back(root *model.TreeNode, sum int, visited []int) {
	if root == nil {
		return
	}
	visited = append(visited, root.Val)
	sum -= root.Val
	if root.Left == nil && root.Right == nil && sum == 0 {
		ans1 := make([]int, len(visited))
		copy(ans1, visited)
		ans = append(ans, ans1)
		return
	}
	back(root.Left, sum, visited)
	back(root.Right, sum, visited)
	visited = visited[:len(visited)-1]
}
