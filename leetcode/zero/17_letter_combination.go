/*
给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。

给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。



示例:

输入："23"
输出：["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].
说明:
尽管上面的答案是按字典序排列的，但是你可以任意选择答案输出的顺序。
*/

package zero

var letters = [8][]string{
	[]string{"a", "b", "c"},
	[]string{"d", "e", "f"},
	[]string{"g", "h", "i"},
	[]string{"j", "k", "l"},
	[]string{"m", "n", "o"},
	[]string{"p", "q", "r", "s"},
	[]string{"t", "u", "v"},
	[]string{"w", "x", "y", "z"},
}

// 回溯法
// 1. 建立一个数字到字母的映射关系 letters
// 2. 依次递归 digits 每个数字，找出其可能的字母
// 3. 当递归完后将字母搜集起来就行了
var combinations = make([]string, 0)

func letterCombinations(digits string) []string {
	combinations = []string{}
	if len(digits) == 0 {
		return combinations
	}
	backtrack(digits, "", len(digits))
	return combinations
}

func backtrack(s, t string, n int) {
	if len(t) == n {
		// fmt.Println(t)
		combinations = append(combinations, t)
		return
	}
	for _, x := range letters[int(s[0])-50] {
		backtrack(s[1:], t+x, n)
	}
}
