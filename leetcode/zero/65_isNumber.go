/*
验证给定的字符串是否可以解释为十进制数字。

例如:

"0" => true
" 0.1 " => true
"abc" => false
"1 a" => false
"2e10" => true
" -90e3   " => true
" 1e" => false
"e3" => false
" 6e-1" => true
" 99e2.5 " => false
"53.5e93" => true
" --6 " => false
"-+3" => false
"95a54e53" => false

说明: 我们有意将问题陈述地比较模糊。在实现代码之前，你应当事先思考所有可能的情况。这里给出一份可能存在于有效十进制数字中的字符列表：

数字 0-9
指数 - "e"
正/负号 - "+"/"-"
小数点 - "."
当然，在输入中，这些字符的上下文也很重要。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/valid-number
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package zero

// 状态机来解决
// 1. 先画出状态转换图，然后构造一张表，用查表方式来做状态跳转
// 注意：加一个 8 是为了处理空字符串
// state	blank	0-9		+/-	e	.	other
// 0		0		3		1	-1	2	-1
// 1		-1		3		-1	-1	2	-1
// 2		-1		4		-1	-1	-1	-1
// 3		8		3		-1	5	4	-1
// 4		8		4		-1	5	-1	-1
// 5		-1		7		6	-1	-1	-1
// 6		-1		7		-1	-1	-1	-1
// 7		8		7		-1	-1	-1	-1
// 8		8		-1		-1	-1	-1	-1

var trans = [][]int{
	{0, 3, 1, -1, 2, -1},
	{-1, 3, -1, -1, 2, -1},
	{-1, 4, -1, -1, -1, -1},
	{8, 3, -1, 5, 4, -1},
	{8, 4, -1, 5, -1, -1},
	{-1, 7, 6, -1, -1, -1},
	{-1, 7, -1, -1, -1, -1},
	{8, 7, -1, -1, -1, -1},
	{8, -1, -1, -1, -1, -1},
}
var finals = [9]int{0, 0, 0, 1, 1, 0, 0, 1, 1}

func choice(b byte) int {
	switch b {
	case ' ':
		return 0
	case '+', '-':
		return 2
	case 'e':
		return 3
	case '.':
		return 4
	default:
		if '0' <= b && b <= '9' {
			return 1
		}
		return 5
	}
}

func isNumber(s string) bool {
	state := 0
	for i := 0; i < len(s); i++ {
		state = trans[state][choice(s[i])]
		if state < 0 {
			return false
		}
	}
	return finals[state] == 1
}
