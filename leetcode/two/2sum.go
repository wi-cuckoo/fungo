package two

// @n^2
func twoSum(nums []int, target int) []int {
	for i, a := range nums {
		for j, b := range nums[i+1:] {
			if a+b == target {
				return []int{i, i + 1 + j}
			}
		}
	}
	return []int{}
}

// @N*2
func twoSumV2(nums []int, target int) []int {
	cache := map[int]int{}
	for i, val := range nums {
		cache[val] = i
	}
	for a, val := range nums {
		another := target - val
		b, ok := cache[another]
		if ok && a != b {
			return []int{a, b}
		}
	}
	return []int{}
}

// @N
func twoSumV3(nums []int, target int) []int {
	cache := map[int]int{}
	for a, val := range nums {
		another := target - val
		b, ok := cache[another]
		if ok && a != b {
			return []int{b, a}
		}
		cache[val] = a
	}
	return []int{}
}
