package interview

import (
	"github.com/wi-cuckoo/fungo/model"
)

/*
题目描述：
在给定的网格中，每个单元格可以有以下三个值之一：

值 0 代表空单元格；
值 1 代表新鲜苹果
值 2 代表发霉到苹果。
每分钟，任何与发霉到苹果（在 4 个正方向上）相邻的新鲜苹果都会发霉。

返回直到单元格中没有新鲜苹果为止所必须经过的最小分钟数。如果不可能，返回 -1。
输入：[[2,1,1],[1,1,0],[0,1,1]]
输出：4

示例 2：
输入：[[2,1,1],[0,1,1],[1,0,1]]
输出：-1
解释：左下角的苹果（第 2 行， 第 0 列）永远不会发霉，因为发霉只会发生在 4 个正向上。

示例 3：
输入：[[0,2]]
输出：0
解释：因为 0 分钟时已经没有新鲜苹果了，所以答案就是 0 。
*/
var m, n int

var (
	posx = []int{0, -1, 0, 1}
	posy = []int{1, 0, -1, 0}
)

func appleRotting(apples [][]int) int {
	m, n = len(apples), len(apples[0])
	queue := make([][]int, 0, 8)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if apples[i][j] == 2 {
				queue = append(queue, []int{i, j})
			}
		}
	}
	round := -1
	for len(queue) > 0 {
		ql := len(queue)
		for p := 0; p < ql; p++ {
			dx, dy := queue[p][0], queue[p][1]
			for k := 0; k < len(posx); k++ {
				x := posx[k] + dx
				y := posy[k] + dy
				if x < 0 || y < 0 || x >= m || y >= n || apples[x][y] != 1 {
					continue
				}
				apples[x][y] = 2
				queue = append(queue, []int{x, y})
			}
		}
		queue = queue[ql:]
		round++
	}
	// check
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if apples[i][j] == 1 {
				return -1
			}
		}
	}
	return round
}

/*
题目描述：堆排序
*/
func heapSort(nums []int) []int {
	h := model.NewHeap(nums)
	ans := make([]int, 0, len(nums))
	for i := 0; i < len(nums); i++ {
		ans = append(ans, h.Pop())
	}
	return ans
}

/*
在一棵满二叉排序树深度为k，节点数为2^k-1;节点值为1至（2^k - 1）,给出k和任意三个节点的值，
输出包含该三个节点的最小子树的根节点。

样例输入：4 10 15 13
样例输出：12
*/
// 假定 i < j < h 有序
func findNodes(k int, i, j, h int) int {
	n := 1
	for i := 0; i < k; i++ {
		n = n << 1
	}
	l, r := 1, n-1
	for l <= r {
		mid := (l + r) / 2
		// fmt.Println(l, r, mid)
		if mid > h {
			r = mid - 1
			continue
		}
		if mid < i {
			l = mid + 1
			continue
		}
		return mid
	}
	return 0
}

// 题目描述：给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和
// O(n) f(i) = max(f(i-1)+nums[i], nums[i])
func maxSubArr(nums []int) int {
	if len(nums) < 1 {
		return 0
	}
	dp := make([]int, len(nums))
	m := nums[0]
	dp[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		if dp[i-1] < 0 {
			dp[i] = nums[i]
		} else {
			dp[i] = dp[i-1] + nums[i]
		}
		if m < dp[i] {
			m = dp[i]
		}
	}
	return m
}
