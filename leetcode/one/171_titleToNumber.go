/*
给定一个Excel表格中的列名称，返回其相应的列序号。

例如，

    A -> 1
    B -> 2
    C -> 3
    ...
    Z -> 26
    AA -> 27
    AB -> 28
    ...
示例 1:

输入: "A"
输出: 1
示例 2:

输入: "AB"
输出: 28
示例 3:

输入: "ZY"
输出: 701
*/

package one

func titleToNumber(s string) int {
	x := 0
	for i := 0; i < len(s); i++ {
		m := int(s[i]-'A') + 1
		x = 26*x + m
	}
	return x
}
