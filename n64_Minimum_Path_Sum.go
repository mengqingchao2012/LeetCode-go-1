package main

import "LeetCode-go/utils"

//Time：O(m*n)，Space:O(m*n)
func minPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])

	dp := make([][]int, m) //注意这里初始化的方法
	for i := range dp {
		dp[i] = make([]int, n)
	}

	dp[0][0] = grid[0][0]

	for i := 1; i < m; i++ {
		dp[i][0] = dp[i - 1][0] + grid[i][0]
	}

	for j := 1; j < n; j++ {
		dp[0][j] = dp[0][j - 1] + grid[0][j]
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = utils.Min(dp[i - 1][j], dp[i][j - 1]) + grid[i][j]
		}
	}
	return dp[m - 1][n - 1]
}
