package three

/*
树是一个无向图，其中任何两个顶点只通过一条路径连接。 换句话说，一个任何没有简单环路的连通图都是一棵树。
给你一棵包含 n 个节点的树，标记为 0 到 n - 1 。给定数字 n 和一个有 n - 1 条无向边的 edges 列表（每一个边都是一对标签），其中 edges[i] = [ai, bi] 表示树中节点 ai 和 bi 之间存在一条无向边。
可选择树中任何一个节点作为根。当选择节点 x 作为根节点时，设结果树的高度为 h 。在所有可能的树中，具有最小高度的树（即，min(h)）被称为 最小高度树 。
请你找到所有的 最小高度树 并按 任意顺序 返回它们的根节点标签列表。

树的 高度 是指根节点和叶子节点之间最长向下路径上边的数量。

示例 1：
输入：n = 4, edges = [[1,0],[1,2],[1,3]]
输出：[1]
解释：如图所示，当根是标签为 1 的节点时，树的高度是 1 ，这是唯一的最小高度树。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/minimum-height-trees
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// 直观解法是，处理每个节点为根节点的树，计算其高度，然后取最小高度的就好了，但这样计算量很大超时了
// 另一种思路，想象一下可能拥有最小高度的树，肯定是一个围绕根节点成放射状展开的形态，所以我们可以从度为 1 的最外围节点开始考虑
// 先找到所有度为1的节点，然后把所有度为1的节点进队列，然后不断地bfs，最后找到的就是最里面的节点了
func findMinHeightTrees(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}
	table := make([][]int, n)
	degree := make([]int, n)
	for _, edge := range edges {
		i, j := edge[0], edge[1]
		degree[i]++
		degree[j]++
		if table[i] == nil {
			table[i] = make([]int, 0)
		}
		if table[j] == nil {
			table[j] = make([]int, 0)
		}
		table[i] = append(table[i], j)
		table[j] = append(table[j], i)
	}

	queue := make([]int, 0, n)
	for i, d := range degree {
		if d == 1 {
			queue = append(queue, i)
		}
	}
	ans := make([]int, 0)
	for len(queue) > 0 {
		ans = ans[:0]
		size := len(queue)
		for i := 0; i < size; i++ {
			cur := queue[i]
			ans = append(ans, cur)
			neighbors := table[cur]
			for _, neighbor := range neighbors {
				degree[neighbor]--
				if degree[neighbor] == 1 {
					queue = append(queue, neighbor)
				}
			}
		}
		queue = queue[size:]
	}
	return ans
}
