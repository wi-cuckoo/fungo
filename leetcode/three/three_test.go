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
