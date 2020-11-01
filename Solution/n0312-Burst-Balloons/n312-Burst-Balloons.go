package main

import ."LeetCode-go/utils"

func maxCoins(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	// 在两端加上 1，处理边界情况
	ballons := make([]int, n + 2)
	for i := 1; i < n + 1; i++ {
		ballons[i] = nums[i - 1]
	}
	ballons[0], ballons[n + 1] = 1, 1

	dp := make([][]int, n + 2)
	for i := 0; i < n + 2; i++ {
		dp[i] = make([]int, n + 2)
	}

	// lens 表示当前枚举区间的长度，长度最小为 3
	for lens := 3; lens <= n + 2; lens++ {
		// 设定左端点
		for i := 0; i + lens - 1 <= n + 1; i++ {
			// 设定右端点
			j := i + lens - 1
			// 枚举最后一个戳破的气球 k
			for k := i + 1; k < j; k++ {
				dp[i][j] = Max(dp[i][j], dp[i][k] + dp[k][j] + ballons[i] * ballons[k] * ballons[j])
			}
		}
	}
	return dp[0][n + 1]
}