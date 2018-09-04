package main

import (
	"fmt"
	"math"
)

func main() {
	num := -1234233
	fmt.Println(reverseV2(num))
}

func reverse(num int) int {
	var max = 1<<31 - 1
	var min = -1 << 31
	cache := make([]int, 0)
	for {
		n := num % 10
		num /= 10
		cache = append(cache, n)
		if num == 0 {
			break
		}
	}
	res := 0
	for i, n := range cache {
		res += n * int(math.Pow10(len(cache)-i-1))
		if res > max || res < min {
			return 0
		}
	}
	return res
}

func reverseV2(num int) int {
	var max = 1<<31 - 1
	var min = -1 << 31
	var res int
	for {
		n := num % 10
		num /= 10
		if res > max/10 || (res == max/10 && n > 7) {
			return 0
		}
		if res < min/10 || (res == min/10 && n < -8) {
			return 0
		}
		res = res*10 + n
		if num == 0 {
			break
		}
	}
	return res
}
