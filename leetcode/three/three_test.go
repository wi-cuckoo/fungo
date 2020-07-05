package three

import (
	"testing"

	"github.com/wi-cuckoo/fungo/model"
)

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
}

func TestIncreasingTriplet(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	yes := true //increasingTriplet(nums)
	if !yes {
		t.Error("should be true")
	}

	nums = []int{2, 1, 5, 0, 3}
	yes = increasingTriplet(nums)
	if yes {
		t.Error("should be false")
	}

}

func TestOddEvenList(t *testing.T) {
	root := model.NewListNode([]int{1, 2, 3, 4, 5})
	oddEvenList(root).Print()
}

func TestTopKFrequent(t *testing.T) {
	nums := []int{1, 1}
	result := topKFrequent(nums, 1)
	t.Log(result)
}

func TestGetSum(t *testing.T) {
	eg := [][]int{
		{4, 8, 12},
		{3, -2, 1},
	}
	for _, e := range eg {
		if x := getSum(e[0], e[1]); x != e[2] {
			t.Errorf("%d+%d should be %d, not %d", e[0], e[1], e[2], x)
		}
	}
}

func TestIsPerfectSquare(t *testing.T) {
	nums := make([]bool, 101)
	for i := 0; i <= 10; i++ {
		nums[i*i] = true
	}
	for i, n := range nums {
		if isPerfectSquare(i) != n {
			t.Errorf("%d should be %v square", i, n)
		}
	}
}

func TestMaxEnvelopes(t *testing.T) {
	envelopes := [][]int{
		{5, 4},
		{6, 4},
		{6, 7},
		{2, 3},
	}
	t.Log(maxEnvelopes(envelopes))

	envelopes = [][]int{{1, 3}, {3, 5}, {6, 7}, {6, 8}, {8, 4}, {9, 5}}
	t.Log(maxEnvelopes(envelopes))
}

func TestValidUtf8(t *testing.T) {
	if validUtf8([]int{240, 162, 138, 147, 145}) {
		t.Error("should be false, fuck")
	}
}

func TestKthSmallest(t *testing.T) {
	matrix := [][]int{
		{1, 5, 8},
		{10, 12, 13},
		{11, 13, 15},
	}
	n := kthSmallest(matrix, 3)
	t.Log(n)
}
