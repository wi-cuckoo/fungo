package one

import "github.com/wi-cuckoo/fungo/util"

/*
老师想给孩子们分发糖果，有 N 个孩子站成了一条直线，老师会根据每个孩子的表现，预先给他们评分。

你需要按照以下要求，帮助老师给这些孩子分发糖果：

每个孩子至少分配到 1 个糖果。
评分更高的孩子必须比他两侧的邻位孩子获得更多的糖果。
那么这样下来，老师至少需要准备多少颗糖果呢？



示例 1：

输入：[1,0,2]
输出：5
解释：你可以分别给这三个孩子分发 2、1、2 颗糖果。
示例 2：

输入：[1,2,2]
输出：4
解释：你可以分别给这三个孩子分发 1、2、1 颗糖果。
     第三个孩子只得到 1 颗糖果，这已满足上述两个条件。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/candy
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// 思路，分两种情况：
// 1. 从左往右，当 rating[i] > rating[i-1] 时，i 同学应该比 i-1 多一颗，否则为 1 颗
// 2. 从右往左，当 rating[i] > rating[i+1] 时，i 同学应该比 i+1 多一颗，否则为 1 颗
// 综合左右两种情况，对 i 点来说取最高的即可

func candy(ratings []int) int {
	left := make([]int, len(ratings))
	for i := 0; i < len(ratings); i++ {
		if i > 0 && ratings[i] > ratings[i-1] {
			left[i] = left[i-1] + 1
		} else {
			left[i] = 1
		}
	}
	n, right := len(ratings), 0
	ans := 0
	for i := n - 1; i >= 0; i-- {
		if i < n-1 && ratings[i] > ratings[i+1] {
			right++
		} else {
			right = 1
		}
		ans += util.Max(left[i], right)
	}
	return ans
}
