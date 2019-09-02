/*
给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。

示例:

输入: [0,1,0,3,12]
输出: [1,3,12,0,0]
说明:

必须在原数组上操作，不能拷贝额外的数组。
尽量减少操作次数。
*/

package two

func moveZeroes(nums []int) {
	for {
		again, firstZero := false, 0
		for i := firstZero; i < len(nums)-1; i++ {
			if nums[i] == 0 && nums[i+1] != 0 {
				nums[i], nums[i+1] = nums[i+1], nums[i]
				again = true
			}
			if nums[i] == 0 && firstZero <= 0 {
				firstZero = i
			}
		}
		if again == false {
			return
		}
	}
}

func moveZeroesV2(nums []int) {
	i, j := 0, 0
	for ; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[j] = nums[i]
			j++
		}
	}
	for ; j < len(nums); j++ {
		nums[j] = 0
	}
}
