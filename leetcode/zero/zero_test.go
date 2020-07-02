package zero

import (
	"fmt"
	"testing"

	"github.com/wi-cuckoo/fungo/model"
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
	intervals := []model.Interval{
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
	expect := []model.Interval{
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

func TestSpiralOrder(t *testing.T) {
	matrix := [][]int{
		{1, 3, 5},
		// {10, 11, 16},
		// {23, 30, 34},
		// {7, 20, 50},
	}
	expect := []int{1, 3, 5} //, 16, 34, 50, 20, 7, 23, 10, 11, 30}
	res := spiralOrder(matrix)
	if len(res) != len(expect) {
		t.Fatalf("length of result should be %d, not %d", len(expect), len(res))
	}
	for i, n := range expect {
		if res[i] != n {
			t.Errorf("res[%d] should be %d, not %d", i, n, res[i])
		}
	}
}

func TestFirstMissingPositive(t *testing.T) {
	nums := []int{-1, 1, 1}
	x := firstMissingPositive(nums)
	if x != 3 {
		t.Error("should be 3, not", x)
	}
}
func TestMergeTwoLists(t *testing.T) {
	l1 := model.NewListNode([]int{1, 2, 4, 5})
	l2 := model.NewListNode([]int{1, 3, 4})
	// mergeTwoLists(l1, l2).Print()
	mergeTwoListsV2(l1, l2).Print()
}

func TestMergeKLists(t *testing.T) {
	lists := []*model.ListNode{
		model.NewListNode([]int{1, 2, 4, 5}),
		model.NewListNode([]int{2, 3, 9}),
		model.NewListNode([]int{6, 8, 9}),
	}
	mergeKLists(lists).Print()
}

func TestMinWindow(t *testing.T) {
	// s := minWindow("KBMADBOBECODEBACVV", "ABC")
	// t.Log(s)

	t.Log(minWindow("cabwefgewcwaefgcf", "cae"))
}

func TestRestoreIPAddresses(t *testing.T) {
	t.Log(restoreIPAddresses("172162541"))
}

func TestMultiply(t *testing.T) {
	if s := multiply("99", "99"); s != "56088" {
		t.Error("123*456 should be 56088, not", s)
	}

}

func TestGetPermutation(t *testing.T) {
	t.Log(getPermutation(3, 6))
	t.Log(getPermutation(4, 9))
}

func TestSearch(t *testing.T) {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	if i := search(nums, 3); i != -1 {
		t.Errorf("target 3 should be at %d, not %d\n", -1, i)
	}
	t.Log(search([]int{5, 1, 3, 4}, 3))
}

func TestTrap(t *testing.T) {
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	if w := trap(height); w != 6 {
		t.Errorf("water should be %d, not %d", 6, w)
	}
	t.Log(trap([]int{2, 0, 2}))
}

func TestSqrt(t *testing.T) {
	pairs := map[int]int{
		1:   1,
		3:   1,
		4:   2,
		101: 10,
	}
	for k, v := range pairs {
		if m := sqrt2(k); m != v {
			t.Errorf("sqrt(%d) should be %d, not %d\n", k, v, m)
		}
	}
}

func TestFullJustify(t *testing.T) {
	// words := []string{"This", "is", "an", "example", "of", "text", "justification."}
	words := []string{"What", "must", "be", "acknowledgment", "shall", "be"}
	// words := []string{"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain", "to", "a", "computer.", "Art", "is", "everything", "else", "we", "do"}
	for _, s := range fullJustify(words, 16) {
		fmt.Println(s)
	}
}

func TestMaxSubArray(t *testing.T) {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	// if n := maxSubArray(nums); n != 6 {
	// 	t.Error("max sub array sum should be 6, not", n)
	// }
	if n := maxSubArrayV2(nums); n != 6 {
		t.Error("max sub array sum should be 6, not", n)
	}
}

func TestPermuteUnique(t *testing.T) {
	nums := []int{1, 2, 1}
	res := permuteUnique(nums)
	if len(res) != 3 {
		t.Error("length of result should 6, not", len(res))
	}
	t.Log(res)
}

func TestIsMatch(t *testing.T) {
	egs := map[[2]string]bool{
		{"aa", "*"}:        true,
		{"adceb", "a*b"}:   true,
		{"acdcb", "a*c?b"}: false,
		{"cb", "?a"}:       false,
		{"a", "a*"}:        true,
	}
	for sp, b := range egs {
		if b != isMatch(sp[0], sp[1]) {
			t.Errorf("if %s match %s %t\n", sp[0], sp[1], b)
		}
	}
}

func TestLengthOfLongestSubstring(t *testing.T) {
	egs := map[string]int{
		"aabaab!bb": 3,
		"":          0,
		" ":         1,
		"abcabcbb":  3,
		"a":         1,
	}
	for e, t2 := range egs {
		if t1 := lengthOfLongestSubstring(e); t1 != t2 {
			t.Fatalf("expect %d for %s, got %d", t2, e, t1)
		}
	}
}
