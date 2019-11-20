package zero

/*
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
注意空字符串可被认为是有效字符串。

示例 1:

输入: "()"
输出: true
示例 2:

输入: "()[]{}"
输出: true
示例 3:

输入: "(]"
输出: false
示例 4:

输入: "([)]"
输出: false
示例 5:

输入: "{[]}"
输出: true
*/

// 用栈可解决
// 遇到左括号 (,[,{ 入栈，遇到右括号出栈，并比较出栈的括号是否匹配当前的右括号，不满足则是无效字符串
// 遍历完后，检查栈中是否剩余元素，剩则无效，反之就是有效字符串

func isValid(s string) bool {
	if s == "" {
		return true
	}
	pair := map[byte]byte{
		'(': ')',
		'[': ']',
		'{': '}',
	}
	cache := []byte{}
	for i := 0; i < len(s); i++ {
		c := s[i]
		// push to stack
		if _, ok := pair[c]; ok {
			cache = append(cache, c)
			continue
		} else if len(cache) == 0 {
			return false
		}
		cc := cache[len(cache)-1]
		cache = cache[:len(cache)-1]
		if pair[cc] != c {
			return false
		}
	}
	if len(cache) > 0 {
		return false
	}
	return true
}
