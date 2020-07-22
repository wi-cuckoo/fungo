/*
给定一个整数 n，生成所有由 1 ... n 为节点所组成的 二叉搜索树

示例：
输入：3
输出：
[
  [1,null,3,2],
  [3,2,null,1],
  [3,1,null,null,2],
  [2,1,3],
  [1,null,2,null,3]
]
解释：
以上的输出对应以下 5 种不同结构的二叉搜索树：

   1         3     3      2      1
    \       /     /      / \      \
     3     2     1      1   3      2
    /     /       \                 \
   2     1         2                 3


提示：

0 <= n <= 8

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/unique-binary-search-trees-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package zero

// 参照 96 题
// 以 [1,n] 的每一个点为根节点，以该点左侧为左子树，右侧为右子树
// 递归的建立所有可能的树

import "github.com/wi-cuckoo/fungo/model"

func generateTrees(n int) []*model.TreeNode {
	if n <= 0 {
		return []*model.TreeNode{}
	}
	return nil
}

func generate(i, n int) []*model.TreeNode {
	if i > n {
		return []*model.TreeNode{nil}
	}
	if i == n {
		return []*model.TreeNode{{Val: i}}
	}
	ans := make([]*model.TreeNode, 0, 64)
	for k := i; k <= n; k++ {
		lefts := generate(i, k-1)
		rights := generate(k+1, n)
		for _, l := range lefts {
			for _, r := range rights {
				ans = append(ans, &model.TreeNode{
					Val:   k,
					Left:  l,
					Right: r,
				})
			}
		}
	}
	return ans
}
