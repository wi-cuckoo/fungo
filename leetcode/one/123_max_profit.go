/*
给定一个数组，它的第 i 个元素是一支给定的股票在第 i 天的价格。

设计一个算法来计算你所能获取的最大利润。你最多可以完成 两笔 交易。

注意: 你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。

示例 1:

输入: [3,3,5,0,0,3,1,4]
输出: 6
解释: 在第 4 天（股票价格 = 0）的时候买入，在第 6 天（股票价格 = 3）的时候卖出，这笔交易所能获得利润 = 3-0 = 3 。
     随后，在第 7 天（股票价格 = 1）的时候买入，在第 8 天 （股票价格 = 4）的时候卖出，这笔交易所能获得利润 = 4-1 = 3 。
示例 2:

输入: [1,2,3,4,5]
输出: 4
解释: 在第 1 天（股票价格 = 1）的时候买入，在第 5 天 （股票价格 = 5）的时候卖出, 这笔交易所能获得利润 = 5-1 = 4 。
     注意你不能在第 1 天和第 2 天接连购买股票，之后再将它们卖出。
     因为这样属于同时参与了多笔交易，你必须在再次购买前出售掉之前的股票。
示例 3:

输入: [7,6,4,3,1]
输出: 0
解释: 在这个情况下, 没有交易完成, 所以最大利润为 0。
*/

package one

import "fmt"

// 一分为二，考虑第i天前后各进行一次交易能获得的最大利润
// 1. 取两个数组 left，right
//    * left 表示从左往右看第 i 天内进行一次交易能获得的最大利润
//    * right 表示从右往左看第 i 天后进行一次交易能获得的最大利润
// 2. 重新从第一天开始遍历，将每天前后两次交易利润加起来，最终得到最大的总利润
func maxProfit3(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	left, right := make([]int, len(prices)), make([]int, len(prices))

	// cacl left at i
	minPrice, maxProfit := 999, 0
	for i, p := range prices {
		if p < minPrice {
			minPrice = p
		}
		if p-minPrice > maxProfit {
			maxProfit = p - minPrice
		}
		left[i] = maxProfit
	}

	// cacl right at i
	maxPrice := -999
	maxProfit = 0
	for i := len(prices) - 1; i >= 0; i-- {
		if prices[i] > maxPrice {
			maxPrice = prices[i]
		}
		if maxPrice-prices[i] > maxProfit {
			maxProfit = maxPrice - prices[i]
		}
		right[i] = maxProfit
	}

	fmt.Println(left, right)

	prifit := 0
	for i := 0; i < len(prices); i++ {
		if left[i]+right[i] > prifit {
			prifit = left[i] + right[i]
		}
	}
	return prifit
}
