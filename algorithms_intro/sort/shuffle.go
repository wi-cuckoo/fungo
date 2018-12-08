package sort

import (
	"math/rand"
)

func shuffle(nums []int) {
	rand.Shuffle(len(nums), func(i, j int) {
		nums[i], nums[j] = nums[j], nums[i]
	})
}
