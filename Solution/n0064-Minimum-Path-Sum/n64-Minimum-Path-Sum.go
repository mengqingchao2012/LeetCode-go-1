package main

import ."LeetCode-go/utils"

// 方法一：Time：O(m*n)，Space:O(m*n)
func minPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dp := make([][]int, m)
	dp[0] = make([]int, n)

	dp[0][0] = grid[0][0]
	for i := 1; i < m; i++ {
		dp[i] = make([]int, n)
		dp[i][0] = dp[i - 1][0] + grid[i][0]
	}

	for i := 1; i < n; i++ {
		dp[0][i] = dp[0][i - 1] + grid[0][i]
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = Min(dp[i - 1][j], dp[i][j - 1]) + grid[i][j]
		}
	}

	return dp[m - 1][n - 1]
}

// 方法二：Time：O(m*n)，Space:O(n)，n 是内部数组的长度，即列数
func minPathSum1(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dp := make([]int, n)
	dp[0] = grid[0][0]

	for i := 1; i < n; i++ {
		dp[i] = dp[i - 1] + grid[0][i]
	}

	for i := 1; i < m; i++ {
		dp[0] = dp[0] + grid[i][0]
		for j := 1; j < n; j++ {
			dp[j] = Min(dp[j], dp[j - 1]) + grid[i][j]
		}
	}
	return dp[n - 1]
}