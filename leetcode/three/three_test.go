package three

import "testing"

func TestCoinChange(t *testing.T) {
	coins := []int{1, 2, 5}
	count := coinChange(coins, 11)
	if count != 3 {
		t.Error("count should be 3")
	}
	t.Log(count)

	coins = []int{186, 419, 83, 408}
	count = coinChange(coins, 6249)
	if count != 20 {
		t.Error("count should be 20")
	}
}

func TestLengthOfLIS(t *testing.T) {
	nums := []int{10, 9, 2, 5, 3, 7, 101, 18}
	l := lengthOfLIS(nums)
	if l != 4 {
		t.Error("length should be 4")
	}
	t.Log(l)

	nums = []int{3, 5, 6, 2, 5, 4, 19, 5, 6, 7, 12}
	l = lengthOfLISV2(nums)
	if l != 6 {
		t.Error("length should be 6")
	}
	t.Log(l)
}
