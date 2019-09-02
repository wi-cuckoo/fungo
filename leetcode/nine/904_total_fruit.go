/*
在一排树中，第 i 棵树产生 tree[i] 型的水果。
你可以从你选择的任何树开始，然后重复执行以下步骤：

把这棵树上的水果放进你的篮子里。如果你做不到，就停下来。
移动到当前树右侧的下一棵树。如果右边没有树，就停下来。
请注意，在选择一颗树后，你没有任何选择：你必须执行步骤 1，然后执行步骤 2，然后返回步骤 1，然后执行步骤 2，依此类推，直至停止。

你有两个篮子，每个篮子可以携带任何数量的水果，但你希望每个篮子只携带一种类型的水果。
用这个程序你能收集的水果总量是多少？



示例 1：

输入：[1,2,1]
输出：3
解释：我们可以收集 [1,2,1]。
示例 2：

输入：[0,1,2,2]
输出：3
解释：我们可以收集 [1,2,2].
如果我们从第一棵树开始，我们将只能收集到 [0, 1]。
示例 3：

输入：[1,2,3,2,2]
输出：4
解释：我们可以收集 [2,3,2,2].
如果我们从第一棵树开始，我们将只能收集到 [1, 2]。
示例 4：

输入：[3,3,3,1,2,1,1,2,3,3,4]
输出：5
解释：我们可以收集 [1,2,1,1,2].
如果我们从第一棵树或第八棵树开始，我们将只能收集到 4 个水果。


提示：

1 <= tree.length <= 40000
0 <= tree[i] < tree.length
*/
package nine

func totalFruit(tree []int) int {
	if len(tree) <= 2 {
		return len(tree)
	}
	dp := make([]int, len(tree))
	dp[0], dp[1] = 1, 2
	f1, f2 := tree[0], tree[1] // mean two kinds of fruit
	diff := 0
	if f1 != f2 {
		diff = 1
	}
	for i := 2; i < len(tree); i++ {
		//fmt.Println(i, tree[i], diff, f1, f2)
		if tree[i] == f1 || tree[i] == f2 {
			dp[i] = dp[i-1] + 1
			if tree[i] != tree[i-1] {
				diff = i
			}
			continue
		}
		f1, f2 = tree[diff], tree[i]
		dp[i] = i - diff + 1
		diff = i
	}
	maxSum := 0
	for _, n := range dp {
		if maxSum < n {
			maxSum = n
		}
	}
	return maxSum
}
