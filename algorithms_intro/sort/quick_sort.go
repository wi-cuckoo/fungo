package sort

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
