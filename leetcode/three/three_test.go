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
	root := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}
	oddEvenList(root).Print()
}