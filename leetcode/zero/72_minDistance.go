/*
  给定两个单词 word1 和 word2，计算出将 word1 转换成 word2 所使用的最少操作数 。

  你可以对一个单词进行如下三种操作：

  插入一个字符
  删除一个字符
  替换一个字符
  示例 1:

  输入: word1 = "horse", word2 = "ros"
  输出: 3
  解释:
  horse -> rorse (将 'h' 替换为 'r')
  rorse -> rose (删除 'r')
  rose -> ros (删除 'e')
  示例 2:

  输入: word1 = "intention", word2 = "execution"
  输出: 5
  解释:
  intention -> inention (删除 't')
  inention -> enention (将 'i' 替换为 'e')
  enention -> exention (将 'n' 替换为 'x')
  exention -> exection (将 'n' 替换为 'c')
  exection -> execution (插入 'u')
*/

package zero

import "math"

//  ## 分析
//  考虑word1[0:i]到word2[0:j]的编辑距离 i<len(word1),j<len(word2)
//  当word1[i] == word2[j] 时，则只需要计算 word1[i:] 到 Word2[j:]的编辑距离
//  如果不相等，则需要操作一下，分三种情况
//  1. 将word[i]改成word2[j]，则之后只需计算 word1[i+1:]到word2[j+1:]的编辑距离
//  2. 将word1[i]删除，之后计算 word1[i+1:] 到 word2[j:]的编辑距离
//  3. 将 word2[j] 插入到 word1[i] 之前，之后计算 word1[i:]到word2[j+1:]的编辑距离

func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	dp := make([][]int, m+1)
	// 0 <= i <= m; j = 0 时，编辑距离都为 i
	// 0 <= j <= n; i = 0 时，编辑距离都为 j
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
		dp[i][0] = i
	}
	for j := 0; j <= n; j++ {
		dp[0][j] = j
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			// 如果字母相等，不需要操作，跳过
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
				continue
			}
			// 如果字母不相等，问题就大了，需要比较三种情况下最优方案
			d1, d2, d3 := dp[i-1][j]+1, dp[i][j-1]+1, dp[i-1][j-1]+1
			// dp[i-1][j]+1 表示把Word2的最后的字符插入到word1末尾，操作分加一
			// dp[i][j-1]+1 表示把word2最后一个字符删掉，操作分加一
			// dp[i-1][j-1] 表示将word2最后一个字符替换成word1最后一个字符，操作分加一
			// 与上面的分析一致，但理解起来不太直观
			dp[i][j] = minimal(d1, d2, d3)
		}
	}
	return dp[m][n]
}

func minimal(x ...int) int {
	m := math.MaxInt32
	for _, v := range x {
		if v < m {
			m = v
		}
	}
	return m
}
