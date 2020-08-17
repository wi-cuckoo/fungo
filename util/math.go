package util

// Max for many nums
func Max(a int, nums ...int) int {
	for _, n := range nums {
		if a < n {
			a = n
		}
	}
	return a
}

// Min for many nums
func Min(a int, nums ...int) int {
	for _, n := range nums {
		if a > n {
			a = n
		}
	}
	return a
}

// Abs return a positive num
func Abs(a int) int {
	if a < 0 {
		return 0 - a
	}
	return a
}
