/*
如果一个正整数自身是回文数，而且它也是一个回文数的平方，那么我们称这个数为超级回文数。

现在，给定两个正整数 L 和 R （以字符串形式表示），返回包含在范围 [L, R] 中的超级回文数的数目。



示例：

输入：L = "4", R = "1000"
输出：4
解释：
4，9，121，以及 484 是超级回文数。
注意 676 不是一个超级回文数： 26 * 26 = 676，但是 26 不是回文数。


提示：

1 <= len(L) <= 18
1 <= len(R) <= 18
L 和 R 是表示 [1, 10^18) 范围的整数的字符串。
int(L) <= int(R)
*/

package main

import (
	"fmt"
)

func main() {
	fmt.Println(checkPalin("2002"))
	fmt.Println(superpalindromesInRange("4", "100000000"))
}

func superpalindromesInRange(L string, R string) int {
	cnt := 0
	llen, rlen := len(L), len(R)
	for i := llen; i <= (rlen+1)/2; i++ {

	}
	return cnt
}

func checkPalin(s string) bool {
	l, r := 0, len(s)-1
	for l < r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}

func checkSum(n int) bool {
	sum := 0
	for n != 0 {
		i := n % 10
		n = n / 10
		sum += i * i
	}
	return sum < 10
}
