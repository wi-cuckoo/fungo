/*
给定两个字符串 s1 和 s2，写一个函数来判断 s2 是否包含 s1 的排列。

换句话说，第一个字符串的排列之一是第二个字符串的子串。

示例1:

输入: s1 = "ab" s2 = "eidbaooo"
输出: True
解释: s2 包含 s1 的排列之一 ("ba").


示例2:

输入: s1= "ab" s2 = "eidboaoo"
输出: False
*/

package five

// 利用hash算法比对，优化点：可以缩小待比较空间，比如用双指针减少不必比较的子串·
func checkInclusion(s1 string, s2 string) bool {
	table := [26]int{}
	for _, c := range s1 {
		table[int(c-'a')]++
	}
	h1 := hash(s1)
	for i := 0; i <= len(s2)-len(s1); i++ {
		if table[int(s2[i]-'a')] == 0 {
			continue
		}
		h2 := hash(s2[i : i+len(s1)])
		if h1 == h2 {
			return true
		}
	}
	return false
}

func hash(s string) string {
	table := [26]int{}
	for _, c := range s {
		table[int(c-'a')]++
	}
	bs := make([]byte, 0, 26)
	for i, t := range table {
		if t == 0 {
			continue
		}
		bs = append(bs, byte(i)+'a', byte(t)+'0')
	}
	return string(bs)
}
