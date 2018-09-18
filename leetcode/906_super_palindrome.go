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
	"strconv"
)

func main() {
	fmt.Println(mirror("ab", false))
	// fmt.Println(checkSumV2("131"))
	fmt.Println(superpalindromesInRange("100000000", "12345654321"))
}

// found, every num's sqrt should be in [0, 1, 2, 3]
func superpalindromesInRange(L string, R string) int {
	cnt := 0
	l, _ := strconv.Atoi(L)
	r, _ := strconv.Atoi(R)
	llen, rlen := len(L), len(R)
	k, w := (llen+1)/2, (llen-1)/4
	for i := 0; i < pow2(w); i++ {
		s := fmt.Sprintf(fmt.Sprintf("%%.%db", w), i)
		s = "1" + mirror(s, k%2 == 1) + "1"
		v, _ := strconv.Atoi(s)
		fmt.Println(s, v, v*v, l, r)
		if v*v >= l && v*v <= r {
			cnt++
		}
	}
	// cacl num whose length is len(L)
	for i := (llen + 3) / 2; i <= rlen/2; i++ {
		if i%2 == 0 {
			cnt++
		} else {
			cnt += 3
		}
		cnt += pow2((i - 1) / 2)
	}
	return cnt
}

// eg: abc, false will be abccbe, true will be abcba
func mirror(s string, flag bool) string {
	l := 2 * len(s)
	if flag {
		l--
	}
	res := make([]byte, l)
	for i, j := 0, l-1; i <= j; {
		res[i], res[j] = s[i], s[i]
		i++
		j--
	}
	return string(res)
}

func pow2(n int) int {
	x := 1
	for i := 0; i < n; i++ {
		x *= 2
	}
	return x
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

func checkSumV2(s string) bool {
	sum := 0
	for i := 0; i < len(s); i++ {
		i := int(s[i]) - 48
		sum += i * i
	}
	return sum < 10
}
