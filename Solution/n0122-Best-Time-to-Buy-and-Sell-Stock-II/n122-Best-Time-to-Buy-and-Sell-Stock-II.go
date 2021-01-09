package main

import ."LeetCode-go/utils"

// 动态规划
func maxProfit(prices []int) int {
	n := len(prices)
	if n < 2 {
		return 0
	}

	// dp[i][0] 表示第 i 天手上不持有股票时的最大利润
	// dp[i][1] 表示第 i 天手上持有股票时的最大利润
	dp := make([][2]int, n)
	for i := 0; i < n; i++ {
		dp[i] = [2]int{}
	}
	dp[0][0] = 0
	dp[0][1] = -prices[0]

	for i := 1; i < n; i++ {
		dp[i][0] = Max(dp[i - 1][0], dp[i - 1][1] + prices[i])
		dp[i][1] = Max(dp[i - 1][1], dp[i - 1][0] - prices[i])
	}
	return dp[n - 1][0]
}

// 贪心：只要后一天的价格高于当天价格，就执行一次买入卖出
func maxProfit1(prices []int) int {
	n := len(prices)
	if n < 2 {
		return 0
	}

	res := 0
	for i := 1; i < n; i++ {
		if prices[i] > prices[i - 1] {
			res += prices[i] - prices[i - 1]
		}
	}
	return res
}