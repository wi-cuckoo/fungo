package ten

/*
给定一个由若干 0 和 1 组成的数组 A，我们最多可以将 K 个值从 0 变成 1 。
返回仅包含 1 的最长（连续）子数组的长度。

示例 1：
输入：A = [1,1,1,0,0,0,1,1,1,1,0], K = 2
输出：6
解释：
[1,1,1,0,0,1,1,1,1,1,1]
粗体数字从 0 翻转到 1，最长的子数组长度为 6。

示例 2：
输入：A = [0,0,1,1,0,0,1,1,1,0,1,1,0,0,0,1,1,1,1], K = 3
输出：10
解释：
[0,0,1,1,1,1,1,1,1,1,1,1,0,0,0,1,1,1,1]
粗体数字从 0 翻转到 1，最长的子数组长度为 10。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/max-consecutive-ones-iii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

import (
	"github.com/wi-cuckoo/fungo/util"
)

// 本题显然使用滑动窗口，
// 和[424题](https://leetcode-cn.com/problems/longest-repeating-character-replacement0)基本一致
// 本题更加简单，计算maxCnt只针对值为 1 的子串
func longestOnes(A []int, K int) int {
	cnt := [2]int{}
	maxCnt, left := 0, 0
	for right, n := range A {
		cnt[n]++
		maxCnt = util.Max(maxCnt, cnt[1])
		if right-left+1-maxCnt > K {
			cnt[n]--
			left++
		}
	}
	return len(A) - left
}
