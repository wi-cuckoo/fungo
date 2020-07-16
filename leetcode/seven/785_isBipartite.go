/*
给定一个无向图graph，当这个图为二分图时返回true。
如果我们能将一个图的节点集合分割成两个独立的子集A和B，并使图中的每一条边的两个节点一个来自A集合，
一个来自B集合，我们就将这个图称为二分图。graph将会以邻接表方式给出，graph[i]表示图中与节点i相连
的所有节点。每个节点都是一个在0到graph.length-1之间的整数。这图中没有自环和平行边： graph[i] 中
不存在i，并且graph[i]中没有重复的值。

示例 1:
输入: [[1,3], [0,2], [1,3], [0,2]]
输出: true
解释:
无向图如下:
0----1
|    |
|    |
3----2
我们可以将节点分成两组: {0, 2} 和 {1, 3}。

示例 2:
输入: [[1,2,3], [0,2], [0,1,3], [0,2]]
输出: false
解释:
无向图如下:
0----1
| \  |
|  \ |
3----2
我们不能将节点分割成两个独立的子集。
注意:

graph 的长度范围为 [1, 100]。
graph[i] 中的元素的范围为 [0, graph.length - 1]。
graph[i] 不会包含 i 或者有重复的值。
图是无向的: 如果j 在 graph[i]里边, 那么 i 也会在 graph[j]里边。
通过次数20,546提交次数42,449

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/is-graph-bipartite
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package seven

// 对于一条边的两个端点，如果要满足二分图必然划分到两个集合中去，因此可以采用染色的方法
// 从任意点 i 出发，如果点 i 未染色，则染成红色，然后 graph[i] 中的节点需要对应染成
// 蓝色，假如其中有的点已经染色并且颜色与 i 点相同则无法满足划分。如果顺利染色完成则
// 可以满足划分

const (
	_ = iota
	red
	blue
)

var valid bool

func isBipartite(graph [][]int) bool {
	valid = true
	nodes := make([]int, len(graph))
	for i := 0; i < len(graph); i++ {
		if nodes[i] > 0 {
			continue
		}
		nodes[i] = red
		painter(i, graph, nodes)
	}
	return valid
}

func painter(idx int, graph [][]int, nodes []int) {
	neighbors := graph[idx]
	cur := nodes[idx]
	cNei := red
	if cur == red {
		cNei = blue
	}
	for i := 0; i < len(neighbors); i++ {
		k := neighbors[i]
		if nodes[k] == cur {
			valid = false
			return
		}
		if nodes[k] > 0 {
			continue
		}
		nodes[k] = cNei
		painter(k, graph, nodes)
		if !valid {
			return
		}
	}
}
