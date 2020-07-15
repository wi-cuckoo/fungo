/*
给定一个整数 n，求以 1 ... n 为节点组成的二叉搜索树有多少种？

示例:

输入: 3
输出: 5
解释:
给定 n = 3, 一共有 5 种不同结构的二叉搜索树:

   1         3     3      2      1
    \       /     /      / \      \
     3     2     1      1   3      2
    /     /       \                 \
   2     1         2                 3

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/unique-binary-search-trees
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package zero

// 注意是二叉搜索树
// 假设 G(n) 为1 ... n 为节点组成的二叉搜索树数量，f(i, n) 为以节点 i 为根节点组成的
// 二叉搜索树数量， 则 G(n) = f(1, n)+f(2,n)+...+f(n,n)
// 而对于 f(i,n), 即以i节点为根节点时，左子树组成的二叉搜索树数量为 G(i-1)，右子树的
// 为 G(n-i)，因此 f(i,n) = G(i-1)*G(n-i)，代入G(n)的等式即可推出最后结果

func numTrees(n int) int {
	gdp := make([]int, n+1)
	gdp[0], gdp[1] = 1, 1
	for i := 2; i <= n; i++ {
		g := 0
		for j := 1; j < i; j++ {
			g += gdp[j-1] * gdp[i-j]
		}
		gdp[n] = g
	}
	return gdp[n]
}
