package dp

// CoinInLine check if one guy can win with first operator
/*
Method: CoinInLine
Rule:There are n coins in a line. Two players take turns to take one or two coins from right side until there are no more coins left. The player who take the last coin wins.
could you please decide the first player will win or lose?
*/
func CoinInLine(i int) bool {
	dp := make([]bool, i+1)
	dp[1] = true // when just one coin, first win
	dp[2] = true // when just two, first win
	for k := 3; k <= i; k++ {
		dp[k] = !(dp[k-1] && dp[k-2])
	}
	return dp[i]
}
