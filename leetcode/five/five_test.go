package five

import "testing"

func TestSubarraySum(t *testing.T) {
	if s := subarraySum([]int{1, 2, 1, 2, 1}, 3); s != 4 {
		t.Fatalf("got %d, not pass", s)
	}
	if s := subarraySum([]int{1}, 0); s != 0 {
		t.Fatalf("got %d, not pass", s)
	}
}

func TestCheckSubarraySum(t *testing.T) {
	if !checkSubarraySum([]int{23, 2, 4, 6, 7}, 6) {
		t.Fatal("not pass")
	}
}
