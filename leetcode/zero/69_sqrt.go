/*
实现 int sqrt(int x) 函数。

计算并返回 x 的平方根，其中 x 是非负整数。

由于返回类型是整数，结果只保留整数的部分，小数部分将被舍去。

示例 1:

输入: 4
输出: 2
示例 2:

输入: 8
输出: 2
说明: 8 的平方根是 2.82842...,
     由于返回类型是整数，小数部分将被舍去。
*/

package zero

func mySqrt(x int) int {
	if x > 0 && x < 4 {
		return 1
	}
	low, high := 0, x
	for i := x / 2; i > 0; i /= 2 {
		if i*i == x {
			return i
		}
		if i*i < x {
			low = i
			break
		}
		high = i
	}
	for i := low + 1; i <= high; i++ {
		if i*i > x {
			return i - 1
		}
	}
	return 0
}

func sqrt(x int) int {
	n := x
	for n*n > x {
		n = (n + x/n) / 2
	}
	return n
}

// 二分查找
func sqrt2(x int) int {
	if x <= 1 {
		return x
	}
	l, r := 1, x
	for l <= r {
		mid := (l + r) / 2
		m := mid * mid
		if m == x {
			return mid
		}
		if m > x {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return r
}
