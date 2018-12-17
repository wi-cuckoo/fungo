/*
给定一个数组，它的第 i 个元素是一支给定股票第 i 天的价格。

如果你最多只允许完成一笔交易（即买入和卖出一支股票），设计一个算法来计算你所能获取的最大利润。

注意你不能在买入股票前卖出股票。

示例 1:

输入: [7,1,5,3,6,4]
输出: 5
解释: 在第 2 天（股票价格 = 1）的时候买入，在第 5 天（股票价格 = 6）的时候卖出，最大利润 = 6-1 = 5 。
     注意利润不能是 7-1 = 6, 因为卖出价格需要大于买入价格。
示例 2:

输入: [7,6,4,3,1]
输出: 0
解释: 在这种情况下, 没有交易完成, 所以最大利润为 0。
*/
package one

import (
	"fmt"
	"math"
)

func maxProfit1(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	fn := make([]int, 0, 2) // mean price of buy and sell
	for i := 1; i < len(prices); i++ {
		if len(fn) == 0 {
			if prices[i] > prices[i-1] {
				fn = append(fn, i-1, i)
			}
			continue
		}
		s1 := prices[fn[1]] - prices[fn[0]]
		s2 := prices[i] - prices[fn[0]]
		s3 := prices[i] - prices[fn[1]+1]
		l := fn[1] + 1
		for j := fn[1] + 1; j < i; j++ {
			if s := prices[i] - prices[j]; s > s3 {
				s3 = s
				l = j
			}
		}

		if max(s1, s2, s3) == s3 {
			fn[0], fn[1] = l, i
		}
		if max(s1, s2, s3) == s2 {
			fn[1] = i
		}

		fmt.Println(fn, i, l, s1, s2, s3)

	}
	if len(fn) == 0 {
		return 0
	}
	return prices[fn[1]] - prices[fn[0]]
}

func max(nums ...int) int {
	n := nums[0]
	for _, v := range nums {
		if n < v {
			n = v
		}
	}
	return n
}

func maxProfitV2(nums []int) int {
	minPrice, maxProfit := math.MaxInt32, 0
	for _, p := range nums {
		if p < minPrice {
			minPrice = p
		}
		if p-minPrice > maxProfit {
			maxProfit = p - minPrice
		}
	}
	return maxProfit
}
