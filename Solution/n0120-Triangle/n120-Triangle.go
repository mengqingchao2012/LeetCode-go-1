package main

import ."LeetCode-go/utils"

// 方法一：自顶向下
func minimumTotal(triangle [][]int) int {
	m := len(triangle)
	dp := make([][]int, m)
	dp[0] = make([]int, m)
	dp[0][0] = triangle[0][0]

	for i := 1; i < m; i++ {
		dp[i] = make([]int, m)
		dp[i][0] = dp[i - 1][0] + triangle[i][0]
		dp[i][i] = dp[i - 1][i - 1] + triangle[i][i]
		for j := 1; j < i; j++ {
			dp[i][j] = Min(dp[i - 1][j - 1], dp[i - 1][j]) + triangle[i][j]
		}
	}

	res := dp[m - 1][0]
	for i := 1; i < m; i++ {
		if dp[m - 1][i] < res {
			res = dp[m - 1][i]
		}
	}
	return res
}

// 方法二：自底向上
func minimumTotal1(triangle [][]int) int {
	m := len(triangle)
	dp := make([][]int, m)
	dp[m - 1] = make([]int, m)
	copy(dp[m - 1], triangle[m - 1])

	for i := m - 2; i >= 0; i-- {
		dp[i] = make([]int, m)
		for j := 0; j <= i; j++ {
			dp[i][j] = Min(dp[i + 1][j + 1], dp[i + 1][j]) + triangle[i][j]
		}
	}
	return dp[0][0]
}

// 方法三：优化方法二
func minimumTotalOn(triangle [][]int) int {
	m := len(triangle)
	dp := make([]int, m)
	copy(dp, triangle[m - 1])

	for i := m - 2; i >= 0; i-- {
		for j := 0; j <= i; j++ {
			dp[j] = Min(dp[j + 1], dp[j]) + triangle[i][j]
		}
	}
	return dp[0]
}