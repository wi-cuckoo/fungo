package eight

import "testing"

func TestLemonadeChange(t *testing.T) {
	if lemonadeChange([]int{10, 10}) {
		t.Error("[10,10] should be false")
	}
	if lemonadeChange([]int{5, 5, 10, 10, 20}) {
		t.Error("[5,5,10,10,20] should be false")
	}
	if !lemonadeChange([]int{5, 5, 5, 10, 20}) {
		t.Error("[5,5,5,10,20] should be false")
	}
}

func TestStrWithout3a3b(t *testing.T) {
	t.Log(strWithout3a3b(2, 5))
}

func TestLongestMountain(t *testing.T) {
	t.Log(longestMountain([]int{2, 2, 2}))
	t.Log(longestMountain([]int{2, 1, 4, 7, 3, 2, 5}))
}
