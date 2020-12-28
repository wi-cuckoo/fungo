package one

import (
	"math"

	"github.com/wi-cuckoo/fungo/util"
)

/*
给定一个整数数组 prices ，它的第 i 个元素 prices[i] 是一支给定的股票在第 i 天的价格。
设计一个算法来计算你所能获取的最大利润。你最多可以完成 k 笔交易。
注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。

示例 1：
输入：k = 2, prices = [2,4,1]
输出：2
解释：在第 1 天 (股票价格 = 2) 的时候买入，在第 2 天 (股票价格 = 4) 的时候卖出，这笔交易所能获得利润 = 4-2 = 2 。

示例 2：
输入：k = 2, prices = [3,2,6,5,0,3]
输出：7
解释：在第 2 天 (股票价格 = 2) 的时候买入，在第 3 天 (股票价格 = 6) 的时候卖出, 这笔交易所能获得利润 = 6-2 = 4 。
     随后，在第 5 天 (股票价格 = 0) 的时候买入，在第 6 天 (股票价格 = 3) 的时候卖出, 这笔交易所能获得利润 = 3-0 = 3 。


提示：
0 <= k <= 109
0 <= prices.length <= 1000
0 <= prices[i] <= 1000

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iv
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func maxProfit4(k int, prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	n := len(prices)
	if n/2 < k { // no more than n/2 times trade
		k = n / 2
	}
	// sell[i][j] 表示第 i 天收盘后进行了 j 笔交易，并且手上不持有该股票时获利最大
	// buy[i][j] 表示第 i 天收盘后进行了 j 笔交易，并且手上持有该股票时获利最大
	sell, buy := make([][]int, n), make([][]int, n)
	for i := 0; i < n; i++ {
		sell[i], buy[i] = make([]int, k+1), make([]int, k+1)
	}
	buy[0][0] = -prices[0]
	// 当 i=0时，只有prices[0]的股价，不能进行任何交易
	// 因此将buy[0][1..k]， sell[0][1..k]设很小，表示不合法
	for j := 1; j < k+1; j++ {
		sell[0][j] = math.MinInt32
		buy[0][j] = math.MinInt32
	}
	for i := 1; i < n; i++ {
		buy[i][0] = util.Max(buy[i-1][0], sell[i-1][0]-prices[i])
		for j := 1; j <= k; j++ {
			buy[i][j] = util.Max(buy[i-1][j], sell[i-1][j]-prices[i])
			sell[i][j] = util.Max(sell[i-1][j], buy[i-1][j-1]+prices[i])
		}
	}
	return util.Max(0, sell[n-1]...)
}
