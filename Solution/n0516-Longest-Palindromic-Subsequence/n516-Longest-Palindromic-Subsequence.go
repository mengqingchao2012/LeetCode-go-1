package main

import ."LeetCode-go/utils"

func longestPalindromeSubseq(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}

	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
		dp[i][i] = 1
	}

	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			if s[i] == s[j] {
				dp[i][j] = 2 + dp[i + 1][j - 1]
			} else {
				dp[i][j] = Max(dp[i + 1][j], dp[i][j - 1])
			}
		}
	}
	return dp[0][n - 1]
}