package sort

// SelectSort sort nums with selection [3, 2, 4, 5, 1, 9]
func SelectSort(nums []int) {
	l := len(nums)
	for i := 0; i < l; i++ {
		if i == l-1 {
			break
		}
		n, min := i, nums[i] // index of max num in nums[i:]
		for j := i + 1; j < l; j++ {
			if nums[j] < min {
				min = nums[j]
				n = j
			}
		}
		nums[i], nums[n] = nums[n], nums[i]
	}
}

// QuickSort use quick sort
func QuickSort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	pivot := nums[0]
	less := make([]int, 0)
	greater := make([]int, 0)
	for _, i := range nums[1:] {
		if i > pivot {
			greater = append(greater, i)
		} else {
			less = append(less, i)
		}
	}
	result := append(QuickSort(less), pivot)
	return append(result, QuickSort(greater)...)
}
