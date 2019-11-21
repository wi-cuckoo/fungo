package zero

/*
给定一个字符串数组，将字母异位词组合在一起。字母异位词指字母相同，但排列不同的字符串。

示例:

输入: ["eat", "tea", "tan", "ate", "nat", "bat"],
输出:
[
  ["ate","eat","tea"],
  ["nat","tan"],
  ["bat"]
]
说明：

所有输入均为小写字母。
不考虑答案输出的顺序。
*/

import "fmt"

// 主要点在于为每个单词计算出唯一的key来提供分组依据
// key的计算需要满足字母异位词的key是一致的

func groupAnagrams(strs []string) [][]string {
	cache := make(map[string][]string, 0)
	for _, s := range strs {
		k := key(s)
		if _, ok := cache[k]; !ok {
			cache[k] = []string{}
		}
		cache[k] = append(cache[k], s)
	}
	res := make([][]string, len(cache))
	i := 0
	for _, v := range cache {
		res[i] = v
		i++
	}
	return res
}

func key(s string) string {
	cache := [26]int{}
	for _, c := range s {
		cache[c-'a']++
	}
	key := ""
	for i, v := range cache {
		if v == 0 {
			continue
		}
		key += fmt.Sprintf("%d%d", v, i)
	}
	// fmt.Println(key)
	return key
}
