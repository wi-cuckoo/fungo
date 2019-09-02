package main

func main() {
	l, w := 1680, 640
	divide(l, w)

	println(sum([]int{2, 4, 6}))~
}

func divide(l, w int) {
	if l == 2*w {
		println(w)
		return
	}
	nl := 0
	if l > w {
		nl = w
		w = l % w
	} else {
		nl = l
		w = w % l
	}
	divide(nl, w)
}

func sum(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	return nums[0] + sum(nums[1:])
}
