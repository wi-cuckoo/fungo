package three

/*
给你一个字符串 s 和一个整数 k ，请你找出 s 中的最长子串， 要求该子串中的每一字符出现次数都不少于 k 。返回这一子串的长度。

示例 1：
输入：s = "aaabb", k = 3
输出：3
解释：最长子串为 "aaa" ，其中 'a' 重复了 3 次。

示例 2：
输入：s = "ababbc", k = 2
输出：5
解释：最长子串为 "ababb" ，其中 'a' 重复了 2 次， 'b' 重复了 3 次。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-substring-with-at-least-k-repeating-characters
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

import (
	"github.com/wi-cuckoo/fungo/util"
)

// 思路：计算出 s 中各字母出现的次数，如果某个字符出现次数低于 k 个，则该字符必然不会出现在符合条件的子串中
// 因此以这些字符去将 s 分割成若干段，再对这些段继续重复上述做法，得出最后答案
func longestSubstring(s string, k int) int {
	if len(s) < k {
		return 0
	}
	cnt := [26]int{}
	for i := 0; i < len(s); i++ {
		cnt[s[i]-'a']++
	}
	ans, l, r := 0, 0, 0
	for r = 0; r < len(s); r++ {
		if cnt[s[r]-'a'] >= k {
			continue
		}
		ans = util.Max(ans, longestSubstring(s[l:r], k))
		l = r + 1
	}
	if l == 0 {
		return len(s)
	}
	return util.Max(ans, longestSubstring(s[l:r], k))
}
