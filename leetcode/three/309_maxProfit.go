/*
给定一个整数数组，其中第 i 个元素代表了第 i 天的股票价格 。​

设计一个算法计算出最大利润。在满足以下约束条件下，你可以尽可能地完成更多的交易（多次买卖一支股票）:

你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
卖出股票后，你无法在第二天买入股票 (即冷冻期为 1 天)。
示例:

输入: [1,2,3,0,2]
输出: 3
解释: 对应的交易状态为: [买入, 卖出, 冷冻期, 买入, 卖出]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-with-cooldown
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。。
*/

package three

// 在买卖股票的问题中，买入股票意味着负收益，卖出股票意味着正收益
// 考虑第i天结束后所获得的累计最大收益，设为 f[i]，那么第i天有三种情况
// 1. 第 i 天结束后持有一只股票，记为 f[i][0]，那么要么是今天刚买入，要么就是第i-1天结束就持有了
// 2. 第 i 天结束后处于冻结期，记为 f[i][1]
// 3. 第 i 天结束后既不持有股票也不处于冻结期，记为 f[i][2]
// 分别思考 f[i][0], f[i][1], f[i][2] 的状态转移方程

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	f0, f1, f2 := -prices[0], 0, 0
	for i := 1; i < len(prices); i++ {
		tf0 := max(f0, f2-prices[i])
		tf1 := f0 + prices[i]
		f2 = max(f1, f2)
		f0, f1 = tf0, tf1
	}
	return max(f1, f2)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
