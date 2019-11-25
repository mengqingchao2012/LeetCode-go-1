package main

//Time：O(m*n)，Space:O(m*n)
func uniquePaths(m int, n int) int {
	if m < 1 || n < 1 {
		return 0
	}

	var dp [100][100]int //dp[i][j]代表到达第(i,j)个格子的所有路径的总数

	for i := 0; i < m; i++ { //第一列的值都为1
		dp[i][0] = 1
	}
	for j := 0; j < n; j++ { //第一行的值都为1
		dp[0][j] = 1
	}

	for i := 1; i < m; i++ { //其余位置可以从左边或是上边过来，所以值等于左边的值加上上边的值
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}
