package fourteen

/*
给你一个整数 n ，请你返回所有 0 到 1 之间（不包括 0 和 1）满足分母小于等于  n 的 最简 分数 。分数可以以 任意 顺序返回。

 

示例 1：

输入：n = 2
输出：["1/2"]
解释："1/2" 是唯一一个分母小于等于 2 的最简分数。
示例 2：

输入：n = 3
输出：["1/2","1/3","2/3"]
示例 3：

输入：n = 4
输出：["1/2","1/3","1/4","2/3","3/4"]
解释："2/4" 不是最简分数，因为它可以化简为 "1/2" 。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/simplified-fractions
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// no better idea, just mark the invalid fraction
func simplifiedFractions(n int) []string {
	if n == 1 {
		return []string{}
	}
	matrix := make([][]int, n+1)
	for i := 2; i <= n; i++ {
		matrix[i] = make([]int, i)
	}
	for i := 2; i <= n; i++ { // 分母
		for j := 1; j < i; j++ { // 分子
			if matrix[i][j] == 0 {
				matrix[i][j] = 1
			}
			for m := 2; m*i <= n; m++ {
				matrix[m*i][m*j] = -1
			}
		}
	}
	ans := make([]string, 0, n/2)
	for i := 2; i <= n; i++ { // 分母
		for j := 1; j < i; j++ { // 分子
			if matrix[i][j] == 1 {
				ans = append(ans, fmt.Sprintf("%d/%d", j, i))
			}
		}
	}
	return ans
}
