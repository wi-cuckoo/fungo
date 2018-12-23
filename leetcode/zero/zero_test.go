package zero

import (
	"testing"
)

// TestIsValidSudoku ..
func TestIsValidSudoku(t *testing.T) {
	board := [][]byte{
		{'8', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	isValidSudoku(board)
}

func TestJump(t *testing.T) {
	nums := []int{4, 1, 1, 3, 1, 1, 1}
	step := jump(nums)
	if step != 2 {
		t.Error("step should be 2")
	}
	t.Log(step)

	// 2
	nums = []int{2, 3, 1, 1, 4}
	step = jump(nums)
	if step != 2 {
		t.Error("step should be 2")
	}
	t.Log(step)
}

func TestSetZeroes(t *testing.T) {
	matrix := [][]int{
		{0, 1, 2, 0},
		{3, 4, 5, 2},
		{1, 3, 1, 5},
	}
	setZeroesV2(matrix)
	expected := [][]int{
		{0, 0, 0, 0},
		{0, 4, 5, 0},
		{0, 3, 1, 0},
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] != expected[i][j] {
				t.Error(matrix[i][j], "should be", expected[i][j])
			}
		}
	}
	t.Log(matrix)
}

func TestThreeSum(t *testing.T) {
	nums := []int{-1, 0, 1, 2, -1, -4}
	set := threeSum(nums)
	t.Log(set)

	nums = []int{0, 0, 0, 0}
	set = threeSum(nums)
	t.Log(set)

	nums = []int{-4, -2, -2, -2, 0, 1, 2, 2, 2, 3, 3, 4, 4, 6, 6}
	set = threeSum(nums)
	t.Log(set)
}

func TestLongestPalindrome(t *testing.T) {
	s := "babad"
	r := longestPalindrome(s)
	if r != "bab" {
		t.Error("r should be aba or bab, not", r)
	}

	s = "abbc"
	r = longestPalindrome(s)
	if r != "bb" {
		t.Error("r should be bb, not", r)
	}
}

func TestGroupAnagrams(t *testing.T) {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	res := groupAnagrams(strs)
	t.Log(res)
}

func TestPermute(t *testing.T) {
	nums := []int{1, 2, 3}
	res := permute(nums)
	if len(res) != 6 {
		t.Error("length of result should 6, not", len(res))
	}
	t.Log(res)
}

func TestSubsets(t *testing.T) {
	nums := []int{1, 2, 3}
	res := subsets(nums)
	if len(res) != 8 {
		t.Error("length of result should 8, not", len(res))
	}
	t.Log(res)
}

func TestExist(t *testing.T) {
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	if exist(board, "ABCCED") == false {
		t.Error("ABCCED should be exist")
	}

	board = [][]byte{
		{'a'},
	}
	if exist(board, "ab") {
		t.Error("a should be not exist")
	}
}

func TestSortColors(t *testing.T) {
	nums := []int{2, 0, 2, 1, 1, 0}
	sortColorsV2(nums)
	for i, v := range []int{0, 0, 1, 1, 2, 2} {
		if nums[i] != v {
			t.Errorf("nums[%d] should be %d, not %d", i, v, nums[i])
		}
	}
}

func TestSearchMatrix(t *testing.T) {
	matrix := [][]int{
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 50},
	}
	if !searchMatrix(matrix, 3) {
		t.Error("3 should exist")
	}

	// if searchMatrix(matrix, 14) {
	// 	t.Error("14 should not exist")
	// }

	// if !searchMatrix(matrix, 1) {
	// 	t.Error("1 should exist")
	// }

	// if !searchMatrix(matrix, 10) {
	// 	t.Error("14 should exist")
	// }

	matrix = [][]int{{1}}
	if !searchMatrix(matrix, 1) {
		t.Error("1 should  exist")
	}
}

func TestMergeIntervals(t *testing.T) {
	intervals := []Interval{
		{1, 3},
		{2, 6},
		{8, 10},
		{13, 18},
	}
	result := merge(intervals)
	if len(result) != 3 {
		t.Log(result)
		t.Fatal("length of result should be 3, not", len(result))
	}
	expect := []Interval{
		{1, 6},
		{8, 10},
		{13, 18},
	}
	for i, v := range expect {
		if result[i] != v {
			t.Errorf("result[%d] should be %+v, not %+v", i, v, result[i])
		}
	}
}

func TestSearchRotatedArray(t *testing.T) {
	nums := []int{4, 5, 6, 7, 8, 1, 2, 3}
	if search(nums, 4) != 0 {
		t.Errorf("4 should be index with 0")
	}
}
