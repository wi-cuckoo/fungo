package four

/*
给你一个仅由大写英文字母组成的字符串，你可以将任意位置上的字符替换成另外的字符，总共可最多替换 k 次。在执行上述操作后，
找到包含重复字母的最长子串的长度。
注意：字符串长度 和 k 不会超过 104。

示例 1：
输入：s = "ABAB", k = 2
输出：4
解释：用两个'A'替换为两个'B',反之亦然。

示例 2：
输入：s = "AABABBA", k = 1
输出：4
解释：
将中间的一个'A'替换为'B',字符串变为 "AABBBBA"。
子串 "BBBB" 有最长重复字母, 答案为 4。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-repeating-character-replacement
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

import (
	"github.com/wi-cuckoo/fungo/util"
)

// 解题思路：使用双指针+滑动窗口，合法的窗口为窗口长度减去窗口内最多出现字符的个数，差值不超过 K
// 我的代码，比较笨
func characterReplacement(s string, k int) int {
	l, r := 0, 0
	ans := 0
	chars := [26]int{}
	for ; r < len(s); r++ {
		chars[s[r]-'A']++
		maxChar := maxfreq(chars)
		length := r - l + 1
		// same+k is less than the length of [l:r] means s[l:r] is valid
		if length-maxChar <= k {
			continue
		}
		// invalid
		// fmt.Println(ans, r-l, l, r)
		ans = util.Max(ans, r-l)
		// move l to make s[l:r] invalid
		for {
			chars[s[l]-'A']--
			l++
			length, maxChar = r-l+1, maxfreq(chars)
			if length-maxChar <= k {
				break
			}
		}
	}
	return util.Max(ans, r-l)
}

func maxfreq(chars [26]int) int {
	ans := 0
	for _, n := range chars {
		if n > ans {
			ans = n
		}
	}
	return ans
}

// 很明显，上述代码为了获得窗口内出现次数最多的字符的数量，不得不每次循环查找，这是需要优化的点，可以用大顶堆
// 来优化，但代码会显得很臃肿，其实仔细分析，对于一个已经找到的最大窗口 w1 来说，之后遇到的长度小于 w1 的窗口
// 我们都可以不用考虑啦，因而可以写出更简单的代码，如下；
func characterReplacement2(s string, k int) int {
	cnt := [26]int{}
	maxCnt, left := 0, 0
	for right, ch := range s {
		cnt[ch-'A']++
		maxCnt = util.Max(maxCnt, cnt[ch-'A'])
		if right-left+1-maxCnt > k {
			cnt[s[left]-'A']--
			left++
		}
	}
	return len(s) - left
}
