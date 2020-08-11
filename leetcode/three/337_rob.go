/*
在上次打劫完一条街道之后和一圈房屋后，小偷又发现了一个新的可行窃的地区。这个地区只有一个入口，
我们称之为“根”。 除了“根”之外，每栋房子有且只有一个“父“房子与之相连。一番侦察之后，聪明的小偷
意识到“这个地方的所有房屋的排列类似于一棵二叉树”。 如果两个直接相连的房子在同一天晚上被打劫，
房屋将自动报警。
计算在不触动警报的情况下，小偷一晚能够盗取的最高金额。

示例 1:

输入: [3,2,3,null,3,null,1]

     3
    / \
   2   3
    \   \
     3   1

输出: 7
解释: 小偷一晚能够盗取的最高金额 = 3 + 3 + 1 = 7.
示例 2:

输入: [3,4,5,1,3,null,1]

     3
    / \
   4   5
  / \   \
 1   3   1

输出: 9
解释: 小偷一晚能够盗取的最高金额 = 4 + 5 = 9.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/house-robber-iii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package three

import "github.com/wi-cuckoo/fungo/model"

func rob(root *model.TreeNode) int {
	if root == nil {
		return 0
	}
	v1, v2 := fuckTwoChild(root)
	if v1 > v2 {
		return v1
	}
	return v2
}

// 返回包含偷窃root节点与不偷窃root节点的两种情况收益值
func fuckTwoChild(root *model.TreeNode) (int, int) {
	if root == nil {
		return 0, 0
	}
	l1, l2 := fuckTwoChild(root.Left)  // l1: 偷了 root.Left, l2: 不偷 root.Left
	r1, r2 := fuckTwoChild(root.Right) // r1: 偷了 root.Right, r2: 不偷 root.Right
	s1 := root.Val + l2 + r2
	s2 := l1 + r1
	if s1 > s2 {
		return s1, s2
	}
	return s2, s2
}
