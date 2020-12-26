package main

import ."LeetCode-go/utils"

func maxProfit(prices []int, fee int) int {
	n := len(prices)
	if n < 1 {
		return 0
	}

	// dp[i][0] 表示第 i 天手上没有有股票时的最大利润
	// dp[i][1] 表示第 i 天手上持有股票时的最大利润
	dp := make([][2]int, n)
	for i := 0; i < n; i++ {
		dp[i] = [2]int{}
	}
	dp[0][0] = 0
	dp[0][1] = -prices[0]

	for i := 1; i < n; i++ {
		dp[0][0] = Max(dp[0][0], dp[0][1] + prices[i] - fee)
		dp[0][1] = Max(dp[0][1], dp[0][0] - prices[i])
	}
	return dp[0][0]
}

func maxProfit1(prices []int, fee int) int {
	n := len(prices)
	if n < 1 {
		return 0
	}

	d0 := 0
	d1 := -prices[0]

	for i := 1; i < n; i++ {
		d0 = Max(d0, d1 + prices[i] - fee)
		d1 = Max(d1, d0 - prices[i])
	}
	return d0
}