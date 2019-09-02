package sort

import "math/rand"

func shuffle(nums []int) {
	// 使用官方的shuffle
	/*
		rand.Shuffle(len(nums), func(i, j int) {
			nums[i], nums[j] = nums[j], nums[i]
		})
	*/

	// 仿官方实现
	for i := len(nums) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		nums[i], nums[j] = nums[j], nums[i]
	}
}
