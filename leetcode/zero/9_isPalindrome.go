/*
判断一个整数是否是回文数。回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。

示例 1:

输入: 121
输出: true
示例 2:

输入: -121
输出: false
解释: 从左向右读, 为 -121 。 从右向左读, 为 121- 。因此它不是一个回文数。
示例 3:

输入: 10
输出: false
解释: 从右向左读, 为 01 。因此它不是一个回文数。
*/
package zero

func isPalindrome(num int) bool {
	if num < 0 {
		return false
	}
	if num == 0 {
		return true
	}
	cache := make([]int, 0)
	for {
		n := num % 10
		num /= 10
		cache = append(cache, n)
		if num == 0 {
			break
		}
	}
	for i := 0; i < len(cache)/2; i++ {
		if cache[i] != cache[len(cache)-i-1] {
			return false
		}
	}
	return true
}
