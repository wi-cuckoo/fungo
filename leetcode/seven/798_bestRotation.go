package seven

/*
给你一个数组 nums，我们可以将它按一个非负整数 k 进行轮调，这样可以使数组变为 [nums[k], nums[k + 1], ... nums[nums.length - 1], nums[0], nums[1], ..., nums[k-1]] 的形式。
此后，任何值小于或等于其索引的项都可以记作一分。

例如，数组为 nums = [2,4,1,3,0]，我们按 k = 2 进行轮调后，它将变成 [1,3,0,2,4]。这将记为 3 分，因为 1 > 0 [不计分]、3 > 1 [不计分]、0 <= 2 [计 1 分]、2 <= 3 [计 1 分]，4 <= 4 [计 1 分]。
在所有可能的轮调中，返回我们所能得到的最高分数对应的轮调下标 k 。如果有多个答案，返回满足条件的最小的下标 k 。

示例 1：
输入：nums = [2,3,1,4,0]
输出：3
解释：
下面列出了每个 k 的得分：
k = 0,  nums = [2,3,1,4,0],    score 2
k = 1,  nums = [3,1,4,0,2],    score 3
k = 2,  nums = [1,4,0,2,3],    score 3
k = 3,  nums = [4,0,2,3,1],    score 4
k = 4,  nums = [0,2,3,1,4],    score 3
所以我们应当选择 k = 3，得分最高。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/smallest-rotation-with-highest-score
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// 1. 动态规划思路
// 每一轮后，首位变末位，其他位置元素的索引减 1. 此时分四种情况
// * 特殊情况是首位变末位，固定加一分
// * 本轮之前索引大于值的元素仍然记一分
// * 本轮之前索引等于值的元素减一分
// * 本轮之前索引小于值的元素仍然不得分
// 那么状态转移函数可定义 f(k) = f(k-1)-(k-1轮时索引等于值的元素个数)+1
// 而任意k伦时索引等于值的个数，可由每个元素需要经过多少轮走到等于值的索引点统计得出
func bestRotation(nums []int) int {
	n := len(nums)
	steps := make([]int, n)
	score := 0 // k = 0 时得分
	for i := 0; i < n; i++ {
		steps[(i-nums[i]+n)%n]++
		if i >= nums[i] {
			score++
		}
	}
	maxScore, k := score, 0
	for i := 1; i < n; i++ {
		score = score - steps[i-1] + 1
		if score > maxScore {
			maxScore, k = score, i
		}
	}
	return k
}
