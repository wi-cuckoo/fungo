/*
给定一个非空二叉树，返回其最大路径和。
本题中，路径被定义为一条从树中任意节点出发，达到任意节点的序列。该路径至少包含一个节点，且不一定经过根节点。
示例 1:
输入: [1,2,3]

       1
      / \
     2   3

输出: 6

示例 2:
输入: [-10,9,20,null,null,15,7]

   -10
   / \
  9  20
    /  \
   15   7

输出: 42

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/binary-tree-maximum-path-sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package one

import (
	"math"

	"github.com/wi-cuckoo/fungo/model"
	"github.com/wi-cuckoo/fungo/util"
)

// 考虑某个节点 i，记包含该节点及其子树，且以该节点结束的路径最大和为 f(i)
// 则会有递推公式，f(i) = max(val, val+f(i.left), val+f(i.right))
// 解释：val 为节点 i 的值，f(i.left) 为节点 i 的左子节点的路径最大和
// 同时，对于节点 i，问题真正所求的最大路径和其实应该是
// max(val, val+f(i.left), val+f(i.right), f(i.left)+val+f(i.right))

func maxPathSum(root *model.TreeNode) int {
	maxPath := math.MinInt32
	fuckMaxPathSum(root, &maxPath)
	return maxPath
}

func fuckMaxPathSum(root *model.TreeNode, maxPath *int) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		if root.Val > *maxPath {
			*maxPath = root.Val
		}
		return root.Val
	}
	left, right := fuckMaxPathSum(root.Left, maxPath), fuckMaxPathSum(root.Right, maxPath)
	val := util.Max(
		root.Val,
		root.Val+left,
		root.Val+right,
		root.Val+right+left,
	)
	if val > *maxPath {
		*maxPath = val
	}
	return util.Max(root.Val, root.Val+left, root.Val+right)
}
